package stack

import (
	"runtime"
)

const stacktraceMaxDepth = 32

// stack represents a stack of program counters.
type stack []uintptr

// Callers Возвращает массив указателей из стека вызовов (переходов)
// PC - program counter
func Callers() *stack {
	const depth = stacktraceMaxDepth
	var pcs [depth]uintptr

	// https://pkg.go.dev/runtime#Callers
	// Пропускаем три вызова, которые является pc самого caller
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func (s *stack) StackTrace() Trace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}
