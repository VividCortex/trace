package trace

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Enabled toggles trace prints to stdout.
var Enabled = true

func goroutineNum() int {
	b := make([]byte, 20)
	runtime.Stack(b, false)
	var goroutineNum int

	fmt.Sscanf(string(b), "goroutine %d ", &goroutineNum)
	return goroutineNum
}

// Trace prints the goroutine number, file, line number, and name
// of the calling function.
func Trace() {
	if !Enabled {
		return
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	caller := runtime.FuncForPC(pc)
	goroutine := goroutineNum()

	fmt.Printf("[trace - goroutine %d] %s:%d %s\n", goroutine, filepath.Base(file), line, caller.Name())
}
