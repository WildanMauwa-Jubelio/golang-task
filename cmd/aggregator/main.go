package main

import (
	"fmt"
	"time"

	"github.com/jubelio/learn_golang/pkg/fetcher"
)

// runConcurrentProductAggregator demonstrates concurrent data fetching
func runConcurrentProductAggregator() {
	fmt.Println("\nStarting Concurrent Product Data Aggregator")
	fmt.Println("===============================================")

	// Predefined slice of product IDs
	productIDs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Fetching data for product IDs: %v\n\n", productIDs)

	// Create product fetcher
	pf := fetcher.NewProductFetcher()

	// Fetch all products concurrently
	productTitles, fetchErrors, elapsedTime := pf.FetchAllProducts(productIDs)

	// Print results
	fmt.Println("\nAGGREGATION RESULTS")
	fmt.Println("======================")

	fmt.Printf("Total execution time: %v\n", elapsedTime)
	fmt.Printf("Total products requested: %d\n", len(productIDs))
	fmt.Printf("Successfully fetched: %d\n", len(productTitles))
	fmt.Printf("Errors encountered: %d\n\n", len(fetchErrors))

	// Print successful product titles
	if len(productTitles) > 0 {
		fmt.Println("SUCCESSFULLY FETCHED PRODUCTS:")
		fmt.Println("-----------------------------------")
		for i, title := range productTitles {
			fmt.Printf("%d. %s\n", i+1, title)
		}
		fmt.Println()
	}

	// Print errors if any
	if len(fetchErrors) > 0 {
		fmt.Println("ERRORS ENCOUNTERED:")
		fmt.Println("----------------------")
		for i, err := range fetchErrors {
			fmt.Printf("%d. %s\n", i+1, err.Error())
		}
		fmt.Println()
	}

	// Performance summary
	fmt.Println("PERFORMANCE SUMMARY:")
	fmt.Println("-----------------------")
	fmt.Printf("• Concurrent execution: %d goroutines\n", len(productIDs))
	fmt.Printf("• Average time per request: %v\n", elapsedTime/time.Duration(len(productIDs)))
	fmt.Printf("• Success rate: %.1f%%\n", float64(len(productTitles))/float64(len(productIDs))*100)

	fmt.Println("\nConcurrent Product Data Aggregator completed!")
	fmt.Println("=================================================")
}

// TASK 5: Demonstration of Key Primitives
func demonstrateGoroutinePrimitives() {
	fmt.Println("\nDemonstrating Key Go Concurrency Primitives")
	fmt.Println("===============================================")

	fmt.Println("\n1. GOROUTINES:")
	fmt.Println("   - Used for loop to launch worker goroutines")
	fmt.Println("   - Each product ID gets its own goroutine")
	fmt.Println("   - Proper variable passing to avoid capture issues")

	fmt.Println("\n2. sync.WaitGroup:")
	fmt.Println("   - wg.Add(1) called for each worker goroutine")
	fmt.Println("   - wg.Done() called when each worker completes")
	fmt.Println("   - wg.Wait() blocks main until all workers finish")

	fmt.Println("\n3. CHANNELS:")
	fmt.Println("   - Buffered results channel: make(chan string, len(productIDs))")
	fmt.Println("   - Buffered errors channel: make(chan error, len(productIDs))")
	fmt.Println("   - Workers send outcomes to appropriate channels")

	fmt.Println("\n4. FAN-OUT/FAN-IN PATTERN:")
	fmt.Println("   - Fan-out: Launch multiple goroutines from main")
	fmt.Println("   - Fan-in: Collect all results back to main function")

	fmt.Println("\n5. CONCURRENCY BENEFITS:")
	fmt.Println("   - Multiple API calls execute simultaneously")
	fmt.Println("   - Reduced total execution time")
	fmt.Println("   - Better resource utilization")
}

func main() {
	fmt.Println("TASK 5: Concurrent Product Data Aggregator")
	fmt.Println("=========================================")

	// Run concurrent product aggregator demo
	runConcurrentProductAggregator()

	// Demonstrate goroutine primitives
	demonstrateGoroutinePrimitives()
}
