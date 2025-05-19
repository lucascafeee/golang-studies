package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Result represents a processed data item with potential error
type Result struct {
	Value int
	Err   error
}

// worker processes items from jobs channel and sends results to results channel
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup, activeWorkers *int64) {
	defer wg.Done()
	
	// Register worker as active
	atomic.AddInt64(activeWorkers, 1)
	defer atomic.AddInt64(activeWorkers, -1)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d shutting down: %v", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simulate work with varying durations
			processTime := time.Duration(50+rand.Intn(200)) * time.Millisecond
			time.Sleep(processTime)

			// Process the job
			value := job * job
			
			// Send the result back
			select {
			case results <- Result{Value: value}:
				// Result sent successfully
			case <-ctx.Done():
				log.Printf("Worker %d abandoning result due to shutdown: %v", id, ctx.Err())
				return
			}
		}
	}
}

// generator yields a sequence of integers
func generator(ctx context.Context, max int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			// Try to send the value, respecting cancellation
			select {
			case out <- i:
				// Value sent successfully
			case <-ctx.Done():
				log.Printf("Generator shutting down: %v", ctx.Err())
				return
			}
			
			// Simulate generation time
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	return out
}

// fanOut distributes work across multiple workers
func fanOut(ctx context.Context, input <-chan int, workerCount int, activeWorkers *int64) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(workerCount)
	
	// Start workers
	for i := 1; i <= workerCount; i++ {
		go worker(ctx, i, input, results, &wg, activeWorkers)
	}
	
	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	return results
}

// fanIn multiplexes multiple input channels into a single output channel
func fanIn(ctx context.Context, channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	multiplexed := make(chan Result)
	
	// Start a goroutine for each input channel
	multiplex := func(ch <-chan Result) {
		defer wg.Done()
		for result := range ch {
			select {
			case multiplexed <- result:
				// Result forwarded successfully
			case <-ctx.Done():
				return
			}
		}
	}
	
	// Wait for all input channels to close
	wg.Add(len(channels))
	for _, ch := range channels {
		go multiplex(ch)
	}
	
	// Close output channel when all input channels are closed
	go func() {
		wg.Wait()
		close(multiplexed)
	}()
	
	return multiplexed
}

// monitorResourceUsage tracks active goroutines
func monitorResourceUsage(ctx context.Context, activeWorkers *int64) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Printf("Active workers: %d", atomic.LoadInt64(activeWorkers))
		}
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// Setup cancellation context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Initialize metrics
	var activeWorkers int64
	go monitorResourceUsage(ctx, &activeWorkers)
	
	// Configure pipeline parameters
	const (
		totalJobs    = 100
		workerCount  = 5
		bufferSize   = 10
	)
	
	log.Printf("Starting pipeline with %d jobs and %d workers", totalJobs, workerCount)
	
	// Stage 1: Generate jobs
	jobs := generator(ctx, totalJobs)
	
	// Stage 2: Process jobs with multiple workers (fan-out)
	processedResults := fanOut(ctx, jobs, workerCount, &activeWorkers)
	
	// Stage 3: Collect and process results
	var totalProcessed int
	var sum int
	
	// Process results with timeout
	resultTimeout := time.After(10 * time.Second)
	
	for {
		select {
		case <-ctx.Done():
			log.Printf("Pipeline cancelled: %v", ctx.Err())
			return
			
		case <-resultTimeout:
			log.Printf("Pipeline timed out after processing %d results", totalProcessed)
			return
			
		case result, ok := <-processedResults:
			if !ok {
				// All results have been processed
				log.Printf("Pipeline complete! Processed %d results with sum: %d", 
					totalProcessed, sum)
				return
			}
			
			if result.Err != nil {
				log.Printf("Error processing job: %v", result.Err)
				continue
			}
			
			// Process the result
			totalProcessed++
			sum += result.Value
			
			// Log progress periodically
			if totalProcessed%10 == 0 {
				log.Printf("Progress: %d/%d jobs processed", totalProcessed, totalJobs)
			}
		}
	}
}
