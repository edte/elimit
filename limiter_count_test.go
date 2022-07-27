package main

import (
	"fmt"
	"github/edte/elimit/tools"
	"log"
	"testing"
	"time"
)

func TestNewCount(t *testing.T) {
}

func Test_limitCount_Wait(t *testing.T) {
	l := NewCountLimit(&config{
		rate:   10,
		circle: time.Microsecond * 1000,
	})

	pre := time.Now()
	for i := 0; i < 1000; i++ {
		l.Wait()
		t := time.Now()
		log.Println(t.Sub(pre))
		pre = t
	}
}

func Test_limitCount_Allow(t *testing.T) {
	l := NewCountLimit(&config{
		rate:   10,
		circle: time.Microsecond * 10,
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
