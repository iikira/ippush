package mytask

import (
	"github.com/iikira/ippush/syncrunner"
)

type (
	myRunTask struct{}
)

func NewMyRunTask() syncrunner.RunTask {
	return &myRunTask{}
}

func (*myRunTask) ZoneName() string {
	return "iikira.com"
}

func (*myRunTask) AName() string {
	return "541"
}
