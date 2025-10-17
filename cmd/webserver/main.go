package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Global variables (scope global)
var (
	appName = "Golang Task Web Server"
	version = "1.0.0"
)

// Task 1: Health endpoint
func healthEndpoint(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "ok",
		"message":   "Server is running",
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"app":       appName,
		"version":   version,
	})
}

func main() {
	// TASK 1: Fiber Web Framework Setup
	fmt.Println("TASK 1: Fiber Web Server Setup")
	fmt.Println("==============================")
	fmt.Printf("Starting %s v%s\n", appName, version)
	fmt.Println("Initializing Fiber web server on port 9200...")

	app := fiber.New(fiber.Config{
		AppName: appName + " " + version,
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Routes
	app.Get("/health", healthEndpoint)

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "Welcome to " + appName,
			"version":   version,
			"endpoints": []string{"/health"},
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	fmt.Println("Web server is ready!")
	fmt.Println("Available endpoints:")
	fmt.Println("- GET http://localhost:9200/")
	fmt.Println("- GET http://localhost:9200/health")
	fmt.Println("\nStarting server...")

	// Start server
	log.Fatal(app.Listen(":9200"))
}
