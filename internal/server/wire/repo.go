package wire

type MonitorSrc struct {
	Src      string
	Identity string
	Used     uint64
	MaxSrc   uint64
}

type MonitorRepo interface {
	GetDBMonitorData() ([]MonitorSrc, error)
	GetCacheMonitorData() ([]MonitorSrc, error)
}
