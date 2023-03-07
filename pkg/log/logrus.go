package log

import (
	"github.com/sirupsen/logrus"
)

type logrusAdapt struct {
	l *logrus.Logger
}

func (s logrusAdapt) WithField(key string, value interface{}) Logger {
	return newFieldAdapt(s.l.WithField(key, value))
}

func (s logrusAdapt) Tracef(format string, args ...interface{}) {
	s.l.Tracef(format, args...)
}

func (s logrusAdapt) Printf(format string, args ...interface{}) {
	s.l.Printf(format, args...)
}
func (s logrusAdapt) Println(args ...interface{}) {
	s.l.Println(args...)
}

func (s logrusAdapt) Debugf(format string, args ...interface{}) {
	s.l.Debugf(format, args...)
}

func (s logrusAdapt) Infof(format string, args ...interface{}) {
	s.l.Infof(format, args...)
}

func (s logrusAdapt) Warnf(format string, args ...interface{}) {
	s.l.Warnf(format, args...)
}

func (s logrusAdapt) Errorf(format string, args ...interface{}) {
	s.l.Errorf(format, args...)
}

func (s logrusAdapt) Panicf(format string, args ...interface{}) {
	s.l.Panicf(format, args...)
}

func (s logrusAdapt) Fatalf(format string, args ...interface{}) {
	s.l.Fatalf(format, args...)
}

func (s logrusAdapt) WithFields(fields Fields) Logger {
	return newFieldAdapt(s.l.WithFields(logrus.Fields(fields)))
}

func (s logrusAdapt) Trace(args ...interface{}) {
	s.l.Trace(args...)
}

func (s logrusAdapt) Debug(args ...interface{}) {
	s.l.Debug(args...)
}

func (s logrusAdapt) Print(args ...interface{}) {
	s.l.Print(args...)
}

func (s logrusAdapt) Info(args ...interface{}) {
	s.l.Info(args...)
}

func (s logrusAdapt) Warn(args ...interface{}) {
	s.l.Warn(args...)
}

func (s logrusAdapt) Error(args ...interface{}) {
	s.l.Error(args...)
}

func (s logrusAdapt) Panic(args ...interface{}) {
	s.l.Panic(args...)
}

func (s logrusAdapt) Fatal(args ...interface{}) {
	s.l.Fatal(args...)
}

// 封装logrus.Entry
type fieldAdapt struct {
	e *logrus.Entry
}

func (f fieldAdapt) WithField(key string, value interface{}) Logger {
	return newFieldAdapt(f.e.WithField(key, value))
}

func (f fieldAdapt) WithFields(fields Fields) Logger {
	return newFieldAdapt(f.e.WithFields(logrus.Fields(fields)))
}

func (f fieldAdapt) Tracef(format string, args ...interface{}) {
	panic("implement me")
}

func (f fieldAdapt) WithError(err error) Logger {
	return newFieldAdapt(f.e.WithError(err))
}

func (f fieldAdapt) Debugf(format string, args ...interface{}) {
	f.e.Debugf(format, args...)
}

func (f fieldAdapt) Infof(format string, args ...interface{}) {
	f.e.Infof(format, args...)
}

func (f fieldAdapt) Printf(format string, args ...interface{}) {
	f.e.Printf(format, args...)
}

func (f fieldAdapt) Warnf(format string, args ...interface{}) {
	f.e.Warnf(format, args...)
}

func (f fieldAdapt) Warningf(format string, args ...interface{}) {
	f.e.Warningf(format, args...)
}

func (f fieldAdapt) Errorf(format string, args ...interface{}) {
	f.e.Errorf(format, args...)
}

func (f fieldAdapt) Fatalf(format string, args ...interface{}) {
	f.e.Fatalf(format, args...)
}

func (f fieldAdapt) Panicf(format string, args ...interface{}) {
	f.e.Panicf(format, args...)
}

func (f fieldAdapt) Debug(args ...interface{}) {
	f.e.Debug(args...)
}

func (f fieldAdapt) Info(args ...interface{}) {
	f.e.Info(args...)
}

func (f fieldAdapt) Print(args ...interface{}) {
	f.e.Print(args...)
}

func (f fieldAdapt) Warn(args ...interface{}) {
	f.e.Warn(args...)
}

func (f fieldAdapt) Warning(args ...interface{}) {
	f.e.Warning(args...)
}

func (f fieldAdapt) Error(args ...interface{}) {
	f.e.Error(args...)
}

func (f fieldAdapt) Fatal(args ...interface{}) {
	f.e.Fatal(args...)
}

func (f fieldAdapt) Panic(args ...interface{}) {
	f.e.Panic(args...)
}

func (f fieldAdapt) Debugln(args ...interface{}) {
	f.e.Debugln(args...)
}

func (f fieldAdapt) Infoln(args ...interface{}) {
	f.e.Infoln(args...)
}

func (f fieldAdapt) Println(args ...interface{}) {
	f.e.Println(args...)
}

func (f fieldAdapt) Warnln(args ...interface{}) {
	f.e.Warnln(args...)
}

func (f fieldAdapt) Warningln(args ...interface{}) {
	f.e.Warningln(args...)
}

func (f fieldAdapt) Errorln(args ...interface{}) {
	f.e.Errorln(args...)
}

func (f fieldAdapt) Fatalln(args ...interface{}) {
	f.e.Fatalln(args...)
}

func (f fieldAdapt) Panicln(args ...interface{}) {
	f.e.Panicln(args...)
}

func (f fieldAdapt) Trace(args ...interface{}) {
	f.e.Trace(args...)
}

func newFieldAdapt(e *logrus.Entry) Logger {
	return fieldAdapt{e}
}

func NewLogrusAdapt(l *logrus.Logger) Logger {
	return &logrusAdapt{
		l: l,
	}
}
