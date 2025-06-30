package zap_log

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
