package errors

import (
	"fmt"
	"testing"

	stderrors "errors"
)

func stdErrorsF(at, depth int) error {
	if at >= depth {
		return stderrors.New("std error")
	}
	return stdErrorsF(at+1, depth)
}

func myErrors(at, depth int) error {
	if at >= depth {
		return New("my error")
	}
	return myErrors(at+1, depth)
}

func myErrorsWithTrace(at, depth int) error {
	if at >= depth {
		return New("my error").WithStacktrace()
	}
	return myErrors(at+1, depth)
}

// GlobalE is an exported global to store the result of benchmark results,
// preventing the compiler from optimising the benchmark functions away.
var GlobalE interface{}

func BenchmarkErrors(b *testing.B) {
	type run struct {
		stack int
		trace bool
		std   bool
	}
	runs := []run{
		{10, false, false},
		{10, true, false},
		{10, false, true},
		{100, false, false},
		{100, true, false},
		{100, false, true},
		{1000, false, false},
		{1000, true, false},
		{1000, false, true},
	}
	for _, r := range runs {
		part := "myErrors"
		if r.std {
			part = "stdErrors"
		}
		if r.trace {
			part = "myErrorsTrace"
		}

		name := fmt.Sprintf("%s-stack-%d", part, r.stack)
		b.Run(name, func(b *testing.B) {
			var err error
			f := myErrors
			if r.std {
				f = stdErrorsF
			}
			if r.trace {
				f = myErrorsWithTrace
			}
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				err = f(0, r.stack)
			}
			b.StopTimer()
			GlobalE = err
		})
	}
}

func BenchmarkStackFormatting(b *testing.B) {
	type run struct {
		stack  int
		format string
	}
	runs := []run{
		{10, "%s"},
		{10, "%v"},
		{10, "%+v"},
		{30, "%s"},
		{30, "%v"},
		{30, "%+v"},
		{60, "%s"},
		{60, "%v"},
		{60, "%+v"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-stdstack-%d", r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			err := stdErrorsF(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	for _, r := range runs {
		name := fmt.Sprintf("%s-stack-%d", r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			err := myErrors(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	for _, r := range runs {
		name := fmt.Sprintf("%s-stacktrace-%d", r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			st := myErrorsWithTrace(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, st)
			}
			b.StopTimer()
		})
	}
	GlobalE = stackStr
}
