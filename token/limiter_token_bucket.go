package token

import (
	// "fmt"
	"fmt"
	"github/edte/elimit/config"
	"sync"
	"time"
)

const (
	defaultBucketSize = 50
)

// 令牌桶限流算法
type LimitTokenBucket struct {
	bucketSize int64 // 令牌桶容量大小
	rate       int64
	circle     time.Duration
	once       sync.Once
	bucket     chan struct{} // 令牌桶

	tokenRate time.Duration // 产生令牌的速率
	cur       int64         // 当前令牌个数
	lasttime  time.Time     // 上次放入令牌时间
}

func NewTokenBucketLimit(c *config.Config) *LimitTokenBucket {
	l := &LimitTokenBucket{
		bucketSize: defaultBucketSize,
		rate:       c.Rate,
		circle:     c.Circle,
		once:       sync.Once{},
		bucket:     make(chan struct{}, defaultBucketSize),
		tokenRate:  c.Circle / time.Duration(c.Rate),
		lasttime:   time.Now(),
		cur:        defaultBucketSize, // 默认令牌是满的
	}

	return l
}

// 同步限流
func (li *LimitTokenBucket) Wait() {
	// [step 1] 开协程走定时器、定时生成令牌
	li.once.Do(func() {
		go func() {
			t := time.NewTicker(li.tokenRate)

			for {
				select {
				case <-t.C:
					li.bucket <- struct{}{}
				}
			}
		}()
	})

	// [step 2] 取一个令牌
	<-li.bucket
}

// 异步限流
// 先更新令牌情况，然后尝试取令牌
func (li *LimitTokenBucket) Allow() bool {
	fmt.Println(li.cur)

	// [step 0] 备份当前令牌数
	old := li.cur

	// [step 1] 获取当前时间
	now := time.Now()

	fmt.Println(now.Sub(li.lasttime), li.tokenRate, now.Sub(li.lasttime)/li.tokenRate)

	// [step 2] 更新当前令牌数
	// 计算公式：当前令牌数 + 生成时间/生成速率
	li.cur = li.cur + int64(now.Sub(li.lasttime)/li.tokenRate)

	// [step 3] 如果超出桶容量，那么更新容量为满的
	if li.cur > li.bucketSize {
		li.cur = li.bucketSize
	}

	// [step 4] 更新生成令牌时间
	if old != li.cur {
		li.lasttime = now
	}

	// [step 5] 如果令牌为空，则取令牌失败
	if li.cur == 0 {
		return false
	}

	// [step 6] 取令牌成功
	li.cur--
	return true
}
