package syncrunner

import (
	"github.com/robfig/cron"
)

type (
	// ErrCallback 错误处理回调, err 为 nil 时代表成功
	ErrCallback func(err error)
)

// CronRun cron 执行
func (r *Runner) CronRun(spec string, callback ErrCallback) *cron.Cron {
	c := cron.New()
	c.AddFunc(spec, func() {
		err := r.Run()
		if callback != nil {
			callback(err)
		}
	})
	c.Start()
	return c
}
