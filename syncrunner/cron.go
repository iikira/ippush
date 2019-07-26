package syncrunner

import (
	"github.com/iikira/BaiduPCS-Go/pcsutil/pcstime"
	"github.com/robfig/cron/v3"
)

type (
	// ErrCallback 错误处理回调, err 为 nil 时代表成功
	ErrCallback func(err error)
)

// PrepareCron 准备 cron 执行
func (r *Runner) PrepareCron(spec string, callback ErrCallback) *cron.Cron {
	c := cron.New(cron.WithLocation(pcstime.CSTLocation)) // 东八区
	c.AddFunc(spec, func() {
		err := r.Run()
		if callback != nil {
			callback(err)
		}
	})
	return c
}
