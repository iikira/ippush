package syncrunner

import (
	"context"
	"github.com/iikira/iikira-go-utils/utils/getip"
	"github.com/iikira/ippush/syncip"
	"log"
	"os"
)

const (
	// EnvCFAPIKey api key 环境变量 key
	EnvCFAPIKey = "CF_API_KEY"
	// EnvCFAPIEmail api email 环境变量 key
	EnvCFAPIEmail = "CF_API_EMAIL"
)

type (
	// Runner 执行者定义
	Runner struct {
		APIKey   string
		APIEmail string
		task     RunTask
		s        *syncip.SyncIP
		relateIP string // 最近的ip
	}
)

// NewRunner 初始化 Runner, 从环境变量读取配置
func NewRunner() (*Runner, error) {
	r := &Runner{}
	err := r.initInfo()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewRunnerWithInfo 初始化 Runner
func NewRunnerWithInfo(apiKey, apiEmail string) *Runner {
	return &Runner{
		APIKey:   apiKey,
		APIEmail: apiEmail,
	}
}

// initInfo 从环境变量初始化 apiKey 和 apiEmail
func (r *Runner) initInfo() (err error) {
	var ok bool
	if r.APIKey == "" {
		r.APIKey, ok = os.LookupEnv(EnvCFAPIKey)
		if !ok {
			return ErrNotSetEnv
		}
	}
	if r.APIEmail == "" {
		r.APIEmail, ok = os.LookupEnv(EnvCFAPIEmail)
		if !ok {
			return ErrNotSetEnv
		}
	}
	return nil
}

func (r *Runner) lazyInit(ctx context.Context) (err error) {
	err = r.initInfo()
	if err != nil {
		return
	}

	if r.task == nil {
		return ErrNotSetTask
	}

	if r.s == nil {
		r.s, err = syncip.NewSyncIP(r.APIKey, r.APIEmail)
		if err != nil {
			return
		}

	}

	if !r.s.IsSetZone() {
		// set zone
		err = r.s.SetZone(ctx, r.task.ZoneName())
		if err != nil {
			return
		}
	}

	return nil
}

// SetTask 设置工作模式
func (r *Runner) SetTask(task RunTask) {
	r.task = task
}

// Run 执行一次
func (r *Runner) Run(ctx context.Context) error {
	err := r.lazyInit(ctx)
	if err != nil {
		return err
	}

	ipAddr, err := getip.IPInfoFromNetease()
	if err != nil {
		return err
	}
	log.Println(ipAddr)

	if ipAddr == r.relateIP {
		// ip未变化, 跳过更新
		return nil
	}

	err = r.s.SetARecord(ctx, r.task.AName(), ipAddr)
	if err != nil {
		return err
	}

	r.relateIP = ipAddr // set relate ip
	return nil
}
