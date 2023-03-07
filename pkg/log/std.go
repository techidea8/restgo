package log

var std Logger // 标准输出

func WithField(key string, value interface{}) Logger {
	return std.WithField(key, value)
}
func WithFields(fields Fields) Logger {
	return std.WithFields(fields)
}
func Trace(args ...interface{}) {
	std.Trace(args...)
}
func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}
func Debug(args ...interface{}) {
	std.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}
func Print(args ...interface{}) {
	std.Print(args...)
}
func Println(args ...interface{}) {
	std.Println(args...)
}
func Info(args ...interface{}) {
	std.Info(args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Warn(args ...interface{}) {
	std.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Error(args ...interface{}) {
	std.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
func Panic(args ...interface{}) {
	std.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// 初始化全局日志
func SetDefaultLogger(s Logger) {
	std = s
}
