package xlog

// 日志等级接口
type XLevel interface {
	// 设置日志等级
	// 日志等级按 Panic, Fatal, Error, Warn, Info, Debug, Trace 从零开始依次递增
	SetLevel(level int)

	// 通过字符串的方式设置日志等级
	// 支持以下日志级别：trace, debug, info, warning, error, fatal, panic
	SetLevelByString(level string)

	// 获得当前的日志等级值
	GetLevel() int
}
