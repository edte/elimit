package adaptive

import (
	"github/edte/elimit/config"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// 自适应限流算法
// todo: 当前请求量的统计、过去 5min 最大请求量的统计（滑动窗口采集）
type limitAdaptive struct {
	lastLimitTime *time.Time // 上一次限流时间
	requestNum    int        // 当前请求量
	lastLoad      int        // 过去 5min 内允许的最大吞吐量
}

func NewAdaptiveLimit(c *config.Config) *limitAdaptive {
	l := &limitAdaptive{}

	return l
}

// 只是占位
func (li *limitAdaptive) Wait() {
	panic("not implemented") // TODO: Implement
}

func (li *limitAdaptive) Allow() bool {
	// [step 1] 获取 cpu 使用率
	f, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("get cpu use info failed when allow, error:%s", err)
		return true
	}

	// [step 2] 获取当前时间
	t := time.Now()

	// [step 3] 判断 cpu 负载是否小于 80%
	if f[0] < 0.8 {
		// [step 3.1] 判断上一次限流时间是否在 1s 内
		if t.Sub(*li.lastLimitTime) < time.Second {
			// [step 3.1.1] 判断请求负载是否超过,超过就直接拒绝
			if li.requestNum > li.lastLoad {
				return false
			}
		}
		// [step 3.2] 如果上一次限流在 1s 外、或者请求量没有超过负载，则允许
		return true
	}

	// [step 4] 如果 cpu 负载超过了 80%， 则再根据请求负载判断是否要限流
	return li.requestNum > li.lastLoad
}
