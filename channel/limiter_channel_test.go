package channel

import (
	"fmt"
	"github/edte/elimit/config"
	"testing"
	"time"
)

func TestNewChannelLimit(t *testing.T) {
}

func Test_limitChannel_Wait(t *testing.T) {
	fmt.Println("Test_limitChannel_Wait begin...")

	l := NewChannelLimit(&config.Config{
		Rate:   200,
		Circle: time.Second,
	})

	pre := time.Now()
	for i := 0; i < 1000; i++ {
		l.Wait()
		t := time.Now()
		fmt.Println(i, " ", t.Sub(pre))
		pre = t
	}
}

func Test_limitChannel_Allow(t *testing.T) {
}
