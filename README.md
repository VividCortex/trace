trace [![GoDoc](https://godoc.org/github.com/VividCortex/trace?status.png)](https://godoc.org/github.com/VividCortex/trace)
=====
Trace is a simple package that helps with debugging.

> The most effective debugging tool is still careful thought, coupled with judiciously placed print statements.
> — Brian Kernighan

`trace.Trace()` prints the goroutine number, file, line number, and name
of the calling function.

Example
---
Here is Go by Example's [Worker Pools example](https://gobyexample.com/worker-pools) with a
call to trace:

```go
package main

import (
	"fmt"
	"time"

	"github.com/VividCortex/trace"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		trace.Trace(j)

		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}

```

Running the program prints the following:

```
[goroutine 20] traceTest.go:12 main.worker [1]
worker 1 processing job 1
[goroutine 21] traceTest.go:12 main.worker [2]
worker 2 processing job 2
[goroutine 22] traceTest.go:12 main.worker [3]
worker 3 processing job 3
[goroutine 20] traceTest.go:12 main.worker [4]
worker 1 processing job 4
[goroutine 21] traceTest.go:12 main.worker [5]
worker 2 processing job 5
[goroutine 22] traceTest.go:12 main.worker [6]
worker 3 processing job 6
[goroutine 20] traceTest.go:12 main.worker [7]
worker 1 processing job 7
[goroutine 21] traceTest.go:12 main.worker [8]
worker 2 processing job 8
[goroutine 22] traceTest.go:12 main.worker [9]
worker 3 processing job 9
```
