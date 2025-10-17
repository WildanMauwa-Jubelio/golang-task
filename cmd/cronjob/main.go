package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

// Global variables (scope global)
var (
	appName  = "Golang Task Cron Job"
	version  = "1.0.0"
	cronJobs *cron.Cron
)

// Task 2: Scheduled task (cron job) - runs every 10 seconds
func scheduledTask() {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] CRON JOB EXECUTED - Background task running every 10 seconds\n", timestamp)
}

func main() {
	fmt.Println("TASK 2: Cron Job Scheduler")
	fmt.Println("==========================")
	fmt.Printf("Starting %s v%s\n", appName, version)
	fmt.Println("Initializing cron scheduler...")

	// Initialize cron scheduler
	cronJobs = cron.New(cron.WithSeconds())

	// Add scheduled job (every 10 seconds)
	_, err := cronJobs.AddFunc("*/10 * * * * *", scheduledTask)
	if err != nil {
		log.Fatal("Failed to add cron job:", err)
	}

	// Start cron scheduler
	cronJobs.Start()
	fmt.Println("Cron scheduler started!")
	fmt.Println("Background task will run every 10 seconds...")
	fmt.Println("Press Ctrl+C to stop")

	// Run initial task
	fmt.Println("\nRunning initial background task...")
	scheduledTask()

	// Wait for interrupt signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal
	<-c

	// Stop cron scheduler
	fmt.Println("\nShutting down cron scheduler...")
	cronJobs.Stop()
	fmt.Println("Cron job stopped. Goodbye!")
}
