trace
=====
Trace is a simple package that helps with debugging.

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
		trace.Trace()

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
[trace - goroutine 20] tracetest.go:12 main.worker
worker 1 processing job 1
[trace - goroutine 21] tracetest.go:12 main.worker
worker 2 processing job 2
[trace - goroutine 22] tracetest.go:12 main.worker
worker 3 processing job 3
[trace - goroutine 20] tracetest.go:12 main.worker
worker 1 processing job 4
[trace - goroutine 21] tracetest.go:12 main.worker
worker 2 processing job 5
[trace - goroutine 22] tracetest.go:12 main.worker
worker 3 processing job 6
[trace - goroutine 20] tracetest.go:12 main.worker
worker 1 processing job 7
[trace - goroutine 21] tracetest.go:12 main.worker
worker 2 processing job 8
[trace - goroutine 22] tracetest.go:12 main.worker
worker 3 processing job 9
```
