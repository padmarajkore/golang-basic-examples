package main

import (
	"fmt"
	"sync"
	"time"
)

// Example 9: Concurrency with Goroutines and Channels
// Demonstrates Go's concurrency features

// Simple goroutine example
func sayHello(id int, wg *sync.WaitGroup) {
	// Defer the WaitGroup.Done() to ensure it's called when the function exits
	defer wg.Done()

	fmt.Printf("Hello from goroutine %d\n", id)
	time.Sleep(time.Millisecond * time.Duration(100*id))
	fmt.Printf("Goodbye from goroutine %d\n", id)
}

// Function that sends data to a channel
func generateNumbers(count int, ch chan<- int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Generating number: %d\n", i)
		ch <- i // Send i to the channel
		time.Sleep(time.Millisecond * 100)
	}
	close(ch) // Close the channel when done
}

// Function that receives from a channel and sends to another channel
func squareNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		result := num * num
		fmt.Printf("Squaring %d: %d\n", num, result)
		out <- result
	}
	close(out) // Close the output channel when done
}

// Worker pool pattern
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * 150) // Simulate work
		results <- job * 2                 // Send result back
	}
}

// Function to demonstrate select statement
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send values on each channel after different delays
	go func() {
		time.Sleep(time.Millisecond * 20)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		ch2 <- "Channel 2"
	}()

	// Use select to wait on both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-time.After(time.Millisecond * 30):
			fmt.Println("Timeout!")
		}
	}
}

// Function that demonstrates a mutex
func mutexExample() {
	var counter = 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Launch 10 goroutines that increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// Lock the mutex to safely update the counter
				mutex.Lock()
				counter++
				mutex.Unlock()
			}

			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

func main() {
	fmt.Println("Concurrency in Go:")

	// Basic goroutines with WaitGroup
	fmt.Println("\n=== Basic Goroutines ===")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sayHello(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines completed")

	// Channels example
	fmt.Println("\n=== Channels Example ===")
	numberChan := make(chan int)
	squareChan := make(chan int)

	// Start the generator and squarer goroutines
	go generateNumbers(5, numberChan)
	go squareNumbers(numberChan, squareChan)

	// Receive and print the results
	for squaredNum := range squareChan {
		fmt.Printf("Received squared result: %d\n", squaredNum)
	}

	// Buffered channel
	fmt.Println("\n=== Buffered Channel ===")
	bufferedChan := make(chan string, 3)

	// Send without a receiver
	bufferedChan <- "First"
	bufferedChan <- "Second"
	bufferedChan <- "Third"

	// Now receive
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	// Worker pool pattern
	fmt.Println("\n=== Worker Pool Pattern ===")
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	numWorkers := 3
	var workerWg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go worker(w, jobs, results, &workerWg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Close the results channel when all workers are done
	go func() {
		workerWg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	// Select statement
	fmt.Println("\n=== Select Statement ===")
	selectExample()

	// Mutex example
	fmt.Println("\n=== Mutex Example ===")
	mutexExample()

	fmt.Println("\nConcurrency examples completed")
}