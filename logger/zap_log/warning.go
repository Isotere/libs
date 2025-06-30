package zap_log

func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}
