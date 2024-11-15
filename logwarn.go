package logwarn

import (
	"io"
)

type LogSink interface {
	// DISABLED
}

func F(format string, args ...any) {
	// DISABLED
}

func FF(out io.Writer, format string, args ...any) {
	// DISABLED
}

func SF(sink LogSink, format string, args ...any) {
	// DISABLED
}
