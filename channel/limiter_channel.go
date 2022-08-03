package channel

import (
	// "fmt"
	"fmt"
	"github/edte/elimit/config"
	"sync"
	"time"
)

// limitChannel 使用 channel 进行限流
type limitChannel struct {
	begin  time.Time     // 开始时间
	circle time.Duration // 计数周期
	ch     chan struct{} // 通道
	mutex  sync.Mutex
}

func NewChannelLimit(c *config.Config) *limitChannel {
	// [step 1] 初始化 channel
	l := &limitChannel{
		begin:  time.Now(),
		circle: c.Circle,
		ch:     make(chan struct{}, c.Rate),
		mutex:  sync.Mutex{},
	}

	fmt.Println(c.Circle)

	// 开个 channel，然后定时清空 channel(容量不变)
	go func() {
		t := time.NewTicker(c.Circle)
		for {
			select {
			case <-t.C:
				fmt.Println("clear channel ")

				l.mutex.Lock()
				for i := 0; i < len(l.ch); i++ {
					<-l.ch
				}
				l.mutex.Unlock()
			}
		}
	}()

	return l
}

//  同步阻塞，直接尝试往 channel 里放，放成功就不限流
func (li *limitChannel) Wait() {
	li.ch <- struct{}{}
}

// 异步阻塞，暂时只实现接口，无法异步
func (li *limitChannel) Allow() bool {
	panic("not implemented") // TODO: Implement
}
