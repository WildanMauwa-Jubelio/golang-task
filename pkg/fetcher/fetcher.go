package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Product struct untuk parsing JSON response dari dummyjson.com API
type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Brand string  `json:"brand"`
}

// ProductFetcher handles concurrent product fetching
type ProductFetcher struct {
	client  *http.Client
	baseURL string
}

// NewProductFetcher creates a new ProductFetcher instance
func NewProductFetcher() *ProductFetcher {
	return &ProductFetcher{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://dummyjson.com/products",
	}
}

// FetchProduct fetches product data dari API secara concurrent
func (pf *ProductFetcher) FetchProduct(id int, resultsChan chan<- string, errorsChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create API URL
	url := fmt.Sprintf("%s/%d", pf.baseURL, id)

	fmt.Printf("Fetching product ID %d...\n", id)

	// Make HTTP GET request
	resp, err := pf.client.Get(url)
	if err != nil {
		errorsChan <- fmt.Errorf("network error for product ID %d: %v", id, err)
		return
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != 200 {
		errorsChan <- fmt.Errorf("HTTP error for product ID %d: status %d", id, resp.StatusCode)
		return
	}

	// Parse JSON response
	var product Product
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&product); err != nil {
		errorsChan <- fmt.Errorf("JSON parse error for product ID %d: %v", id, err)
		return
	}

	// Send successful result
	resultsChan <- product.Title
	fmt.Printf("Successfully fetched: %s (ID: %d)\n", product.Title, id)
}

// FetchAllProducts fetches multiple products concurrently
func (pf *ProductFetcher) FetchAllProducts(productIDs []int) ([]string, []error, time.Duration) {
	// Create buffered channels
	resultsChan := make(chan string, len(productIDs))
	errorsChan := make(chan error, len(productIDs))

	// Create WaitGroup
	var wg sync.WaitGroup

	// Record start time
	startTime := time.Now()

	// Launch goroutines for each product ID (fan-out)
	fmt.Println("Launching goroutines...")
	for _, id := range productIDs {
		wg.Add(1)
		go pf.FetchProduct(id, resultsChan, errorsChan, &wg)
	}

	// Wait for all goroutines to complete (fan-in)
	fmt.Println("\nWaiting for all fetch operations to complete...")
	wg.Wait()

	// Close channels
	close(resultsChan)
	close(errorsChan)

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	// Collect results
	var productTitles []string
	var fetchErrors []error

	// Collect successful results
	for title := range resultsChan {
		productTitles = append(productTitles, title)
	}

	// Collect errors
	for err := range errorsChan {
		fetchErrors = append(fetchErrors, err)
	}

	return productTitles, fetchErrors, elapsedTime
}
