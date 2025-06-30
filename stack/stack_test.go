package stack

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	stackTraceTest := Callers(defaultCallStackPosition)

	t.Run("stack depth", func(t *testing.T) {
		assert.Equal(t, 2, len(stackTraceTest.StackTrace()))
	})
}

func testCallers() *Stack {
	// для тестирования формата нам нужна только одна запись в стеке
	const depth = 1
	var pcs [depth]uintptr

	// для работы тестов пропускаем не три, а два вызова
	n := runtime.Callers(2, pcs[:])
	var st Stack = pcs[0:n]
	return &st
}

func Test_Stack_Formatters(t *testing.T) {
	stackTraceTest := testCallers()

	t.Run("stack trace as S", func(t *testing.T) {
		formatted := fmt.Sprintf("%s\n", stackTraceTest.StackTrace())
		assert.Contains(t, formatted, "stack_test.go")
	})

	t.Run("stack trace as V", func(t *testing.T) {
		formatted := fmt.Sprintf("%v\n", stackTraceTest.StackTrace())
		assert.Contains(t, formatted, "stack_test.go:31")
	})

	t.Run("stack trace as +V", func(t *testing.T) {
		formatted := fmt.Sprintf("%+v\n", stackTraceTest.StackTrace())
		assert.Contains(t, formatted, "stack/stack_test.go:31")
	})
}
