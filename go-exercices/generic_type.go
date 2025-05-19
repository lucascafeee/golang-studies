package main

import (
	"fmt"
	"sync"
)

// ProcessingItem represents an item being processed
type ProcessingItem[T any] struct {
	Value     T
	Timestamp int64
	Status    string
}

// Type alias for a queue of items being processed
type ProcessingQueue[T any] = []ProcessingItem[T]

// Alias for a thread-safe map with strong typing
type ConcurrentMap[K comparable, V any] = sync.Map

// Structure for processing results
type Result[T any] struct {
	Data  T
	Ok    bool
	Errs  []error
}

// Alias for a result channel
type ResultChan[T any] = chan Result[T]

// processBatch processes items concurrently
func processBatch[T any, R any](items ProcessingQueue[T], process func(T) (R, error)) ResultChan[R] {
	results := make(ResultChan[R], len(items))
	
	for _, item := range items {
		go func(i ProcessingItem[T]) {
			data, err := process(i.Value)
			
			if err != nil {
				results <- Result[R]{
					Ok:   false,
					Errs: []error{err},
				}
				return
			}
			
			results <- Result[R]{
				Data: data,
				Ok:   true,
			}
		}(item)
	}
	
	return results
}

func main() {
	// Create a queue of integers to process
	numQueue := ProcessingQueue[int]{
		{Value: 10, Timestamp: 1714300000, Status: "new"},
		{Value: 20, Timestamp: 1714300010, Status: "new"},
		{Value: 30, Timestamp: 1714300020, Status: "new"},
	}
	
	// A queue of tasks as strings
	taskQueue := ProcessingQueue[string]{
		{Value: "send-email", Timestamp: 1714300000, Status: "new"},
		{Value: "validate-data", Timestamp: 1714300015, Status: "new"},
	}
	
	// A thread-safe cache using our alias
	cache := ConcurrentMap[string, int]{}
	cache.Store("req_count", 157)
	cache.Store("err_count", 3)
	
	// Process numbers and convert to formatted strings
	results := processBatch(numQueue, func(n int) (string, error) {
		return fmt.Sprintf("Result: %d", n*n), nil
	})
	
	// Collect and display results
	for i := 0; i < len(numQueue); i++ {
		r := <-results
		if r.Ok {
			fmt.Println(r.Data)
		}
	}
	
	// Process strings to get lengths
	strResults := processBatch(taskQueue, func(s string) (int, error) {
		return len(s), nil
	})
	
	for i := 0; i < len(taskQueue); i++ {
		r := <-strResults
		if r.Ok {
			fmt.Printf("Task length: %d characters\n", r.Data)
		}
	}
	
	// Demonstrate cache access
	if count, ok := cache.Load("req_count"); ok {
		fmt.Printf("Total requests: %d\n", count)
	}
}
