package window

import (
	"github/edte/elimit/config"
	"time"
)

// 滑动窗口算法
type LimitWindow struct {
	queue  []*time.Time
	rate   int64
	circle time.Duration
}

func NewWindowLimit(c *config.Config) *LimitWindow {
	l := &LimitWindow{
		rate:   c.Rate,
		circle: c.Circle,
		queue:  make([]*time.Time, 0),
	}

	return l
}

func (li *LimitWindow) Wait() {
	panic("not implemented") // TODO: Implement
}

//  滑动窗口限流算法
func (li *LimitWindow) Allow() bool {
	// [step 1] 获取当前时间
	now := time.Now()

	// [step 2] 如果队列没满，直接放入并且成功
	if len(li.queue) < int(li.rate) {
		li.queue = append(li.queue, &now)
		return true
	}

	// [step 3] 取出最早的请求成功时间
	first := li.queue[0]

	// [step 4] 如果最早的请求成功时间到现在在周期内，说明周期内的请求量超过了队列长度即限制，故应该限流
	if now.Sub(*first) <= li.circle {
		return false
	}

	// [step 5] 如果请求不在周期内，说明可以继续请求，并且队列出队，然后入队
	li.queue = li.queue[1:]
	li.queue = append(li.queue, &now)

	return true
}
