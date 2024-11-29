package logger

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Level int8

type Fields map[string]interface{}

type Logger struct {
	*logrus.Entry
}

func NewLogger(level logrus.Level, filePath string) (*Logger, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(absPath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}
	log := logrus.New()
	file, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.Out = file

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(level)
	log.SetFormatter(&MyLoggerFormatter{})
	return &Logger{Entry: log.WithFields(logrus.Fields{})}, nil
}

func (l *Logger) WithFields(fields Fields) *Logger {
	return &Logger{Entry: l.Entry.WithFields(logrus.Fields(fields))}
}

func (l *Logger) WithTrace(ctx context.Context) *Logger {
	ginCtx, ok := ctx.(*gin.Context)

	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.GetString("X-Trace-ID"),
			"span_id":  ginCtx.GetString("X-Span-ID"),
		})
	}
	return l
}

func (l *Logger) logWithContext(ctx context.Context, level logrus.Level, msg string, args ...interface{}) {
	entry := l.WithTrace(ctx).Entry
	switch level {
	case logrus.DebugLevel:
		entry.Debugf(msg, args...)
	case logrus.InfoLevel:
		entry.Infof(msg, args...)
	case logrus.WarnLevel:
		entry.Warnf(msg, args...)
	case logrus.ErrorLevel:
		entry.Errorf(msg, args...)
	case logrus.FatalLevel:
		entry.Fatalf(msg, args...)
	case logrus.PanicLevel:
		entry.Panicf(msg, args...)
	}
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.DebugLevel, msg)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.DebugLevel, format, args...)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.InfoLevel, msg)
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.InfoLevel, format, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.WarnLevel, msg)
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.WarnLevel, format, args...)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.ErrorLevel, msg)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.ErrorLevel, format, args...)
}

func (l *Logger) Fatal(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.FatalLevel, msg)
}

func (l *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.FatalLevel, format, args...)
}

func (l *Logger) Panic(ctx context.Context, msg string) {
	l.logWithContext(ctx, logrus.PanicLevel, msg)
}

func (l *Logger) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.logWithContext(ctx, logrus.PanicLevel, format, args...)
}
