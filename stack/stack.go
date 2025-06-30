package stack

import (
	"runtime"
)

const (
	defaultCallStackPosition = 3
	stacktraceMaxDepth       = 32
)

// stack represents a stack of program counters.
type Stack []uintptr

// Callers Возвращает массив указателей из стека вызовов (переходов)
// PC - program counter
func Callers(stackPosition int) *Stack {
	const depth = stacktraceMaxDepth
	var pcs [depth]uintptr

	// https://pkg.go.dev/runtime#Callers
	// Пропускаем три вызова, которые является pc самого caller
	n := runtime.Callers(stackPosition, pcs[:])
	var st Stack = pcs[0:n]
	return &st
}

func (s *Stack) StackTrace() Trace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}
