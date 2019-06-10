package syncrunner

import (
	"errors"
)

var (
	// ErrNotSetEnv 未设置环境变量
	ErrNotSetEnv = errors.New("cf env not set")
	// ErrNotSetTask 未设置工作模式
	ErrNotSetTask = errors.New("cf task not set")
)
