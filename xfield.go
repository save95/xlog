package xlog

// 属性集合
type Fields map[string]interface{}

// 日志前置属性接口
// 前置属性应该伴随整个日志生命周期，即每次日志的输出都应该在内容的前面输出这些属性
type XPreField interface {
	// 设置属性
	Set(key string, value interface{})

	// 将所有属性转成字符串，用于日志输出
	String() string
}

// 日志属性接口
type XField interface {
	// 将前置属性设置到 XLog 中
	// 该前置属性会伴随整个日志生命周期，但是其值会被后续同名的属性替代
	WithPreField(field XPreField) XLog

	// 将属性设置到 XLog 中
	// 该属性只会被第一个日志输出
	// 接受可选参数，如果可选参数实现了 XFieldFormatter 接口，日志中属性的输出格式由其处理
	// 如果需要设置多个，请使用 `WithFields`
	WithField(key string, value interface{}, options ...interface{}) XLog

	// 将属性集合设置到 XLog 中
	// 该属性集合只会被第一个日志输出
	// 接受可选参数，如果可选参数实现了 XFieldFormatter 接口，日志中属性的输出格式由其处理
	WithFields(fields Fields, options ...interface{}) XLog
}

// XFiled 格式化接口
type XFieldFormatter interface {
	Format(fields Fields) string
}
