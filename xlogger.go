package xlog

// 日志记录器接口
type XLogger interface {
	XLog
	XLevel
	XField

	// 设置是否同时在控制台输出日志
	SetStdPrint(b bool) bool

	// 设置当前日志记录器的 Field 格式化方式
	// 一般实现需要遵循以下优先级：当前日志记录器的 Formatter > 临时 Formatter > 默认 Formatter
	SetFieldFormatter(f XFieldFormatter)

	// 获得日志存储方式
	GetStack() Stack
}
