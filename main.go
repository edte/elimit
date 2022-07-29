package main

import (
	// "fmt"
	// "time"

	"fmt"

	"go.uber.org/ratelimit"
	rate2 "golang.org/x/time/rate"
)

func main() {
	l := rate2.NewLimiter(20, 5)
	l.Allow()

	l1 := ratelimit.New(1999999)
	// prev := time.Now()

	for i := 0; i < 10; i++ {
		l1.Take()
		fmt.Println(i)
		// fmt.Println(i, now.Sub(prev))
		// prev = now
	}
}
