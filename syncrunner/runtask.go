package syncrunner

type (
	// RunTask 工作模式
	RunTask interface {
		ZoneName() string // 域名
		AName() string    // a记录名字
	}
)
