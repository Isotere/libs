package zap_log

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}
