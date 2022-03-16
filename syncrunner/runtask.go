package syncrunner

type (
	// RunTask 工作模式
	RunTask interface {
		ZoneName() string // 域名
		AName() string    // a记录名字
	}

	staticTask struct {
		zoneName string
		aName    string
	}
)

func NewStaticTask(zoneName, aName string) RunTask {
	return &staticTask{
		zoneName: zoneName,
		aName:    aName,
	}
}

func (s *staticTask) ZoneName() string {
	return s.zoneName
}

func (s *staticTask) AName() string {
	return s.aName
}
