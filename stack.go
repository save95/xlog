package xlog

// 日志存储方式
type Stack int32

const (
	// 单一存储，会将所有日志存储在 runtime.log 中
	SingleStack Stack = iota
	// 按日存储，根据自然日创建日志文件，格式 Y-m-d.log
	DailyStack
)
