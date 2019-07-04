//worker pool is a collection of goroutines working concurrently to perform a job

package main

import (
	"fmt"
	"time"
)

// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}

/* wydrukuje:
[main] main() started
[main] Wrote 5 tasks
[worker 0] Sending result by worker 0
[worker 1] Sending result by worker 1
[main] Result 0 : 4
[main] Result 1 : 16
[worker 2] Sending result by worker 2
[main] Result 2 : 0
[worker 0] Sending result by worker 0
[main] Result 3 : 36
[worker 1] Sending result by worker 1
[main] Result 4 : 64
[main] main() stopped
*/
