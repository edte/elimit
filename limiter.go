package main

import (
	"github/edte/elimit/channel"
	"github/edte/elimit/config"
	"github/edte/elimit/count"
	"github/edte/elimit/leaky"
	"time"
)

type Option func(*config.Config)

func WithRate(rate int64) Option {
	return func(c *config.Config) {
		c.Rate = rate
	}
}

func WithLimitType(t config.LimitType) Option {
	return func(c *config.Config) {
		c.LimitType = t
	}
}

func WithCircle(t time.Duration) Option {
	return func(c *config.Config) {
		c.Circle = t
	}
}

// Limiter
type Limiter interface {
	Wait()       // 同步睡眠
	Allow() bool // 异步返回 bool
}

func New(rate int64, opts ...Option) (l Limiter) {
	c := &config.Config{
		Rate:      rate,
		Circle:    time.Second,        // 默认一秒钟
		LimitType: config.TypeWindows, // 默认 xx 算法实现
	}

	for _, opt := range opts {
		opt(c)
	}

	switch c.LimitType {
	case config.TypeWindows:

	case config.TypeCount:
		l = count.NewCountLimit(c)
	case config.TypeTokenBucket:

	case config.TypeLeakyBucket:
		l = leaky.NewLeakyBucketLimit(c)
	case config.TypeLimitChannel:
		l = channel.NewChannelLimit(c)
	default:

	}

	return
}

func NewFunc(rate int, take func()) (l Limiter) {
	return
}
