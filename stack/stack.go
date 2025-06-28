package stack

import (
	"runtime"
)

const stacktraceMaxDepth = 32

// stack represents a stack of program counters.
type Stack []uintptr

// Callers Возвращает массив указателей из стека вызовов (переходов)
// PC - program counter
func Callers() *Stack {
	const depth = stacktraceMaxDepth
	var pcs [depth]uintptr

	// https://pkg.go.dev/runtime#Callers
	// Пропускаем три вызова, которые является pc самого caller
	n := runtime.Callers(3, pcs[:])
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
