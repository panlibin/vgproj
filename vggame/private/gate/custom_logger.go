package gate

import logger "github.com/panlibin/vglog"

type CustomLogger struct {
	*logger.Logger
}

func (l *CustomLogger) Printf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}
