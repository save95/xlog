package xlog

// 日志记录器集合接口
// 用于一个项目需要多个记录器的场景
type XLoggers interface {
	XLevel

	// 设置是否同时在控制台输出日志
	SetStdPrint(b bool)
	// 从日志记录器集合中获取指定的记录器
	GetLogger(category string) XLogger
}
