package gooselogger

import (
	"strings"

	"go.uber.org/zap"
)

type GooseLogger struct {
	logger *zap.SugaredLogger
}

// NewGooseLogger returns a new GooseLogger from a zap logger
func NewGooseLogger(logger *zap.SugaredLogger) *GooseLogger {
	return &GooseLogger{logger: logger}
}

func (l *GooseLogger) clean(str string) string {
	// str = strings.TrimPrefix(str, "goose: ")
	return strings.TrimSuffix(str, "\n")
}

func (l *GooseLogger) Fatal(v ...interface{})                 { l.logger.Fatal(v...) }                   //nolint:revive
func (l *GooseLogger) Fatalf(format string, v ...interface{}) { l.logger.Fatalf(l.clean(format), v...) } //nolint:revive
func (l *GooseLogger) Print(v ...interface{})                 { l.logger.Info(v...) }                    //nolint:revive
func (l *GooseLogger) Println(v ...interface{})               { l.logger.Infoln(v...) }                  //nolint:revive
func (l *GooseLogger) Printf(format string, v ...interface{}) { l.logger.Infof(l.clean(format), v...) }  //nolint:revive
