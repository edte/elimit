package leaky

import (
	"github/edte/elimit/config"
	"sync"
	"time"
)

const (
	//  todo: 这里桶大小需要改成可配置的, 一般线上写成配置文件
	defaultBucketSize = 10
)

// 漏桶限流算法
type LimitLeakyBucket struct {
	bucket chan struct{} // 桶
	ok     chan struct{} // 桶流出控制
	rate   int64         // 速率上限
	circle time.Duration // 周期
	once   sync.Once

	lasttime   int64   // 上次更新时间
	bucketSize float64 // 桶大小
	cur        float64 // 当前水位
	waterWate  float64 // 水流出速率
}

func NewLeakyBucketLimit(c *config.Config) *LimitLeakyBucket {
	l := &LimitLeakyBucket{
		bucketSize: defaultBucketSize, // 默认桶大小
		bucket:     make(chan struct{}, defaultBucketSize),
		ok:         make(chan struct{}, 1),
		rate:       c.Rate,
		circle:     c.Circle,
		once:       sync.Once{},
		lasttime:   time.Now().UnixNano(),
		cur:        0,
		waterWate:  float64(c.Rate) / float64(c.Circle),
	}

	return l
}

// 同步限流
func (li *LimitLeakyBucket) Wait() {
	// [step 1] 流量入桶
	li.bucket <- struct{}{}

	// [step 2] 开一个协程，然后以固定的速度流出，固定速度的实现很简单，直接计算出相邻流出时间，然后 sleep 即可
	li.once.Do(func() {
		go func() {
			// [step 2.1] 开计时器
			t := time.NewTicker(li.circle / time.Duration(li.rate))

			// [step 2.2] 桶流出
			for {
				// [step 2.3] 但时间到则桶流出，然后控制当前流出
				select {
				case <-t.C:
					<-li.bucket
					li.ok <- struct{}{}
				}
			}
		}()
	})

	<-li.ok
}

// 异步限流
// 先更新漏水情况，然后尝试加水
func (li *LimitLeakyBucket) Allow() bool {
	// [step 1] 获取当前时间
	now := time.Now().UnixNano()

	// [step 2] 更新当前水位
	// 计算公式：当前水位 - 流水时间*流水速率
	// 最后注意要大于 0
	li.cur = max(0, li.cur-float64(now-li.lasttime)*li.waterWate)

	// [step 3] 更新漏水时间
	li.lasttime = now

	// [step 4] 尝试加水到桶中
	if li.cur+1 < li.bucketSize {
		li.cur++
		return true
	}

	// [step 5] 水桶已满
	return false
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
