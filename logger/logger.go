package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Isotere/libs/stack"
)

const defaultCallStackPosition = 3

// TraceIDFn represents a function that can return the trace id from
// the specified context.
type TraceIDFn func(ctx context.Context) string

// Logger represents a logger for logging information.
type Logger struct {
	handler   slog.Handler
	traceIDFn TraceIDFn
}

// New constructs a newLogger log for application use.
func New(serviceName string, opts ...Option) *Logger {
	o := defaultOptions()
	_ = compileOptions(o, opts...)

	return newLogger(o.w, o.logLevel, serviceName, o.traceIDFn, o.events)
}

// Debug logs at LevelDebug with the given context.
func (log *Logger) Debug(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelDebug, defaultCallStackPosition, msg, true, args...)
}

// Info logs at LevelInfo with the given context.
func (log *Logger) Info(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelInfo, defaultCallStackPosition, msg, false, args...)
}

// Warn logs at LevelWarn with the given context.
func (log *Logger) Warn(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelWarn, defaultCallStackPosition, msg, false, args...)
}

// Error logs at LevelError with the given context.
func (log *Logger) Error(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelError, defaultCallStackPosition, msg, true, args...)
}

func (log *Logger) write(ctx context.Context, level Level, caller int, msg string, withStack bool, args ...any) {
	slogLevel := slog.Level(level)

	if !log.handler.Enabled(ctx, slogLevel) {
		return
	}

	stackTrace := stack.Callers(caller)

	r := slog.NewRecord(time.Now(), slogLevel, msg, (*stackTrace)[0])

	if log.traceIDFn != nil {
		args = append(args, "trace_id", log.traceIDFn(ctx))
	}

	if withStack {
		args = append(args, "stacktrace", stackTrace.StackTrace())
	}

	r.Add(args...)

	_ = log.handler.Handle(ctx, r)
}

func newLogger(w io.Writer, minLevel Level, serviceName string, traceIDFn TraceIDFn, events Events) *Logger {
	// Convert the file name to just the name.ext when this key/value will be logged.
	f := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			if source, ok := a.Value.Any().(*slog.Source); ok {
				v := fmt.Sprintf("%s:%d", filepath.Base(source.File), source.Line)
				return slog.Attr{Key: "file", Value: slog.StringValue(v)}
			}
		}

		return a
	}

	// Construct the slog JSON handler for use.
	handler := slog.Handler(slog.NewJSONHandler(w, &slog.HandlerOptions{AddSource: true, Level: slog.Level(minLevel), ReplaceAttr: f}))

	// If events are to be processed, wrap the JSON handler around the custom
	// log handler.
	if events.Debug != nil || events.Info != nil || events.Warn != nil || events.Error != nil {
		handler = newLogHandler(handler, events)
	}

	// Attributes to add to every log.
	attrs := []slog.Attr{
		{Key: "service", Value: slog.StringValue(serviceName)},
	}

	// Add those attributes and capture the final handler.
	handler = handler.WithAttrs(attrs)

	return &Logger{
		handler:   handler,
		traceIDFn: traceIDFn,
	}
}

func defaultOptions() *options {
	return &options{
		w:         os.Stdout,
		logLevel:  LevelInfo,
		traceIDFn: nil,
	}
}
