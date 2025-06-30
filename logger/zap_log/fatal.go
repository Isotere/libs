package zap_log

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
