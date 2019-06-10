package syncip

import (
	"errors"
)

var (
	// ErrNotSetAPI 未初始化 api
	ErrNotSetAPI = errors.New("api not set")
	// ErrNotSetZone 未初始化 zone
	ErrNotSetZone = errors.New("zone not set")
	// ErrNoSuchRecord 未找到任何记录
	ErrNoSuchRecord = errors.New("no such records")
)
