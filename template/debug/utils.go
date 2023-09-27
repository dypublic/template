package debug

import (
	"fmt"
	"runtime"
	"strings"
)

func TraceF(format string, a ...any) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	name := strings.Split(frame.Function, "/")
	format = "%s:%d %-28s: " + format + "\n"
	var params []any
	params = append(params, frame.File)
	params = append(params, frame.Line)
	params = append(params, name[len(name)-1])
	params = append(params, a...)
	fmt.Printf(format, params...)
}
