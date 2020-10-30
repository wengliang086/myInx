package utils

type GlobalObj struct {
	Host string
	Port int
	Name string

	Version          string
	MaxPackageSize   uint32
	MaxConn          int
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 //SendBuffMsg发送消息的缓冲最大长度

	ConfFilePath string

	LogDir        string
	LogFile       string
	LogDebugClose bool
}
