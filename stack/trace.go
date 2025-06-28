package stack

import (
	"fmt"
	"io"
)

// Trace stack (массив) Frames от самых "поздних" к самым "ранним".
type Trace []Frame

// Format реализация интерфейса fmt.Formatter.
//
//	%s	список файлов для каждого фрейма в стеке
//	%v	список файлов с номерами строк, из которых был произведен вызов для каждого фрейма в стеке
//
// Format принимает дополнительные флаги для некоторых "глаголов":
//
//	%+v печатает название файла, название функции и строку вызова для каждого фрейма в стеке
func (st Trace) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			for _, f := range st {
				_, _ = io.WriteString(s, "\n")
				f.Format(s, verb)
			}
		case s.Flag('#'):
			_, _ = fmt.Fprintf(s, "%#v", []Frame(st))
		default:
			st.formatSlice(s, verb)
		}
	case 's':
		st.formatSlice(s, verb)
	}
}

// formatSlice will format this Trace into the given buffer as a slice of
// Frame, only valid when called with '%s' or '%v'.
func (st Trace) formatSlice(s fmt.State, verb rune) {
	_, _ = io.WriteString(s, "[")
	for i, f := range st {
		if i > 0 {
			_, _ = io.WriteString(s, " ")
		}
		f.Format(s, verb)
	}
	_, _ = io.WriteString(s, "]")
}
