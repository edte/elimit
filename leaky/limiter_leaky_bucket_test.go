package leaky

import (
	"fmt"
	"github/edte/elimit/config"
	"github/edte/elimit/tools"
	"testing"
	"time"
)

func TestNewLeakyBucketLimit(t *testing.T) {
}

func TestLimitLeakyBucket_Wait(t *testing.T) {
	l := NewLeakyBucketLimit(&config.Config{
		Rate:   1,
		Circle: time.Second,
	})

	pre := time.Now()
	for i := 0; i < 10000; i++ {
		l.Wait()
		// fmt.Println("##")
		t := time.Now()
		fmt.Println(t.Sub(pre))
		pre = t
	}
}

func TestLimitLeakyBucket_Allow(t *testing.T) {
	l := NewLeakyBucketLimit(&config.Config{
		Rate:   1,
		Circle: time.Nanosecond * 10000,
	})

	data := make([]tools.Item, 0)

	for i := 0; i < 1000; i++ {
		data = append(data, tools.Item{
			Time: time.Now(),
			Data: l.Allow(),
		})
	}

	tools.Listen(8899, data)
}
