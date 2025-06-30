package zap_log

func (l *Logger) Close() {
	_ = l.logger.Sync()
}
