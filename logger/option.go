package logger

import (
	"io"
	"os"
)

type options struct {
	w         io.Writer
	logLevel  Level
	traceIDFn TraceIDFn
	events    Events
}

type Option func(o *options)

func WithWriter(w io.Writer) Option {
	return func(o *options) {
		o.w = w
	}
}

func WithConsoleWriter(w io.Writer) Option {
	return WithWriter(os.Stdout)
}

func WithTraceIDFn(fn TraceIDFn) Option {
	return func(o *options) {
		o.traceIDFn = fn
	}
}

func WithLogLevel(lvl Level) Option {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func WithDebugEvent(fn EventFn) Option {
	return func(o *options) {
		o.events.Debug = fn
	}
}

func WithInfoEvent(fn EventFn) Option {
	return func(o *options) {
		o.events.Info = fn
	}
}

func WithWarnEvent(fn EventFn) Option {
	return func(o *options) {
		o.events.Warn = fn
	}
}

func WithErrorEvent(fn EventFn) Option {
	return func(o *options) {
		o.events.Error = fn
	}
}

func compileOptions(in *options, opts ...Option) error {
	for _, opt := range opts {
		opt(in)
	}

	return nil
}
