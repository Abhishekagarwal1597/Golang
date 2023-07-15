package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a parent context
	parentCtx := context.Background()

	// Create a derived context with a cancellation function
	ctx, cancel := context.WithCancel(parentCtx)

	// Start a goroutine to perform some task
	go performTask(ctx)

	// Sleep for 2 seconds to simulate the passage of time
	time.Sleep(2 * time.Second)

	// Cancel the context to signal termination
	cancel()

	// Wait for a moment to allow the goroutine to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Main goroutine terminated.")
}

func performTask(ctx context.Context) {
	// Check if the context is canceled
	select {
	case <-ctx.Done():
		fmt.Println("Task canceled.")
		return
	default:
		// Perform some task
		fmt.Println("Task in progress...")
		time.Sleep(3 * time.Second)
		fmt.Println("Task completed.")
	}
}
