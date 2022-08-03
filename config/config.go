package config

import (
	"time"
)

// type
type LimitType int

const (
	TypeWindows      LimitType = iota // 滑动窗口
	TypeCount                         // 计数
	TypeTokenBucket                   // 令牌桶
	TypeLeakyBucket                   // 漏桶
	TypeLimitChannel                  // 巧妙利用 channel 和定时器进行限流
	TypeAdaptive                      // 自适应算法
)

// config
type Config struct {
	Rate      int64         // 请求最多次数
	Circle    time.Duration // 请求周期
	LimitType LimitType     // 实现算法类型,默认 todo
}
