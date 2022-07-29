package count

import (
	"github/edte/elimit/config"
	"sync"
	"sync/atomic"
	"time"
)

// limitCount 计数器限流算法实现
type limitCount struct {
	rate   int64         // 计数上限
	begin  time.Time     // 开始计数时间
	circle time.Duration // 计数周期
	count  int64         // 当前计数

	once sync.Once
}

func NewCountLimit(c *config.Config) *limitCount {
	l := &limitCount{
		rate:   c.Rate,
		begin:  time.Now(),
		circle: c.Circle,
		count:  0,
	}

	return l
}

// 同步限流，相对于异步，只是在需要限流时，剩余周期时间内直接睡眠
func (li *limitCount) Wait() {
	// [step 1] 计数器自增
	atomic.AddInt64(&li.count, 1)

	// [step 2] 如果没有达到速率上限，则不限速, 这里与是否超过计数周期限制内无关，没超过可以，超过则更说明速率慢，不用限速
	if li.count <= li.rate {
		return
	}

	// [step 3] 获取当前时间
	t := time.Now()

	// [step 4] 如果计数时间小于周期，则说明在周期时间内计数达到了上限，则需要限速,直接睡眠剩余周期内的时间
	if t.Sub(li.begin) < li.circle {
		m := li.begin.Add(li.circle).Sub(t)
		// fmt.Printf("sleep, %v\n", m)
		time.Sleep(m)
		return
	}

	// [step 5] 如果计数时间大于周期,则虽然超出了计数上限，但是立刻就进入了下一个计数周期，也不会达到新的上限，所以重置开始时间和计数，并且不限速
	atomic.StoreInt64(&li.count, 0)
	li.begin = t

	return
}

// 异步限流，直接返回限流结果，不阻塞
// 使用原子的方式计数和赋值
func (li *limitCount) Allow() bool {
	// [step 1] 计数器自增
	atomic.AddInt64(&li.count, 1)

	// [step 2] 如果没有达到速率上限，则不限速, 这里与是否超过计数周期限制内无关，没超过可以，超过则更说明速率慢，不用限速
	if li.count <= li.rate {
		return true
	}

	// [step 3] 获取当前时间
	t := time.Now()

	// [step 4] 如果计数时间小于周期，则说明在周期时间内计数达到了上限，则需要限速
	if t.Sub(li.begin) < li.circle {
		return false
	}

	// [step 5] 如果计数时间大于周期,则虽然超出了计数上限，但是立刻就进入了下一个计数周期，也不会达到新的上限，所以重置开始时间和计数，并且不限速
	atomic.StoreInt64(&li.count, 0)
	li.begin = t

	return true
}

func (li *limitCount) Acquire() {
	li.once.Do(func() {

	})

}
