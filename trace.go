package trace

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	enabled = true
	output  = io.Writer(os.Stdout)
	lock    = sync.Mutex{}
)

// Enable enables tracing.
func Enable() {
	lock.Lock()
	defer lock.Unlock()

	enabled = true
}

// Disable disables tracing.
func Disable() {
	lock.Lock()
	defer lock.Unlock()

	enabled = false
}

// SetWriter sets the output of trace lines
// to an io.Writer.
func SetWriter(writer io.Writer) {
	lock.Lock()
	defer lock.Unlock()

	output = writer
}

// SetOutputFile creates a file at filename and sets it as the output destination.
// If filename points to an existing file, it will be truncated. An error is returned
// if the file could not be opened.
func SetOutputFile(filename string) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	output = f

	return nil
}

// goroutineNum returns the goroutine number
// the caller is running on.
func goroutineNum() int {
	b := make([]byte, 20)
	runtime.Stack(b, false)
	var goroutineNum int

	fmt.Sscanf(string(b), "goroutine %d ", &goroutineNum)
	return goroutineNum
}

func getTraceLine() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	caller := runtime.FuncForPC(pc)
	goroutine := goroutineNum()

	return fmt.Sprintf("[goroutine %d] %s:%d %s",
		goroutine, filepath.Base(file), line, caller.Name())
}

// Trace prints the goroutine number, file, line number, and name
// of the calling function, as well as any optional arguments.
func Trace(args ...interface{}) {
	if !enabled {
		return
	}

	traceLine := getTraceLine()
	if traceLine == "" {
		return
	}

	message := fmt.Sprint(args...)

	if message != "" {
		message = "[" + message + "]"
	}

	fmt.Fprintln(output, traceLine, message)
}

// Tracef prints the goroutine number, file, line number, and name
// of the calling function, as well as any optional arguments printed
// with a specific format.
func Tracef(format string, args ...interface{}) {
	if !enabled {
		return
	}

	traceLine := getTraceLine()
	if traceLine == "" {
		return
	}

	message := fmt.Sprintf(format, args...)

	if message != "" {
		message = "[" + message + "]"
	}

	fmt.Fprintln(output, traceLine, message)
}
