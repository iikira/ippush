package syncrunner

import (
	"context"
	"github.com/iikira/iikira-go-utils/utils/csttime"
	"github.com/robfig/cron/v3"
)

type (
	// ErrCallback 错误处理回调, err 为 nil 时代表成功
	ErrCallback func(err error)
)

// PrepareCron 准备 cron 执行
func (r *Runner) PrepareCron(spec string, first bool, callback ErrCallback) *cron.Cron {
	ctx := context.Background()
	if first {
		// 第一次立即执行
		err := r.Run(ctx)
		if callback != nil {
			callback(err)
		}
	}
	c := cron.New(cron.WithLocation(csttime.CSTLocation)) // 东八区
	c.AddFunc(spec, func() {
		err := r.Run(ctx)
		if callback != nil {
			callback(err)
		}
	})
	return c
}
