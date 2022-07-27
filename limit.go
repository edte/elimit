package main

import "time"

// type
type LimitType int

const (
	TypeWindows     LimitType = iota // 滑动窗口
	TypeCount                        // 计数
	TypeTokenBucket                  // 令牌桶
	TypeLeakyBucket                  // 漏桶
)

// config
type config struct {
	rate      int64         // 请求最多次数
	circle    time.Duration // 请求周期
	limitType LimitType     // 实现算法类型,默认 todo
}

type Option func(*config)

func WithRate(rate int64) Option {
	return func(c *config) {
		c.rate = rate
	}
}

func WithLimitType(t LimitType) Option {
	return func(c *config) {
		c.limitType = t
	}
}

func WithCircle(t time.Duration) Option {
	return func(c *config) {
		c.circle = t
	}
}

// Limiter
type Limiter interface {
	Wait()       // 同步睡眠
	Allow() bool // 异步返回 bool
}

func New(rate int64, opts ...Option) (l Limiter) {
	c := &config{
		rate:      rate,
		circle:    time.Second, // 默认一秒钟
		limitType: TypeWindows, // 默认 xx 算法实现
	}

	for _, opt := range opts {
		opt(c)
	}

	switch c.limitType {
	case TypeWindows:

	case TypeCount:
		l = NewCountLimit(c)
	case TypeTokenBucket:

	case TypeLeakyBucket:

	default:
	}

	return
}

func NewFunc(rate int, take func()) (l Limiter) {
	return
}
