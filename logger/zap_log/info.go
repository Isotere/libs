package zap_log

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}
