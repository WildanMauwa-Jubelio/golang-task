package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Global variables (scope global)
var (
	appName = "Golang Task Manager"
	version = "1.0.0"
)

// Task represents a task with its details
type Task struct {
	ID          int
	Name        string
	Description string
	Command     string
	Directory   string
}

// Available tasks
var tasks = []Task{
	{
		ID:          1,
		Name:        "Web Server (Fiber)",
		Description: "Start Fiber web server with health endpoint on port 9200",
		Command:     "go run main.go",
		Directory:   "cmd/webserver",
	},
	{
		ID:          2,
		Name:        "Cron Job Scheduler",
		Description: "Run background cron job that executes every 10 seconds",
		Command:     "go run main.go",
		Directory:   "cmd/cronjob",
	},
	{
		ID:          3,
		Name:        "Inventory Management System",
		Description: "Demo inventory management with structs and methods",
		Command:     "go run main.go",
		Directory:   "cmd/inventory",
	},
	{
		ID:          4,
		Name:        "Programming Basics",
		Description: "Demonstrate loops, conditionals, and functions (FizzBuzz, switch, etc.)",
		Command:     "go run main.go",
		Directory:   "cmd/basics",
	},
	{
		ID:          5,
		Name:        "Concurrent Product Aggregator",
		Description: "Fetch product data concurrently using goroutines and channels",
		Command:     "go run main.go",
		Directory:   "cmd/aggregator",
	},
}

func displayWelcome() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Printf("║                    %s v%s                    ║\n", appName, version)
	fmt.Println("╠════════════════════════════════════════════════════════════════╣")
	fmt.Println("║                        Task Selection Menu                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func displayTasks() {
	fmt.Println("Available Tasks:")
	fmt.Println("================")

	for _, task := range tasks {
		fmt.Printf("[%d] %s\n", task.ID, task.Name)
		fmt.Printf("    %s\n", task.Description)
		fmt.Printf("    Directory: %s\n", task.Directory)
		fmt.Println()
	}

	fmt.Println("[0] Exit")
	fmt.Println("========================================")
}

func getUserChoice() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your choice (0-5): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Remove whitespace and newlines
		input = strings.TrimSpace(input)

		// Convert to integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0-5.")
			continue
		}

		// Validate range
		if choice < 0 || choice > len(tasks) {
			fmt.Printf("Invalid choice. Please enter a number between 0-%d.\n", len(tasks))
			continue
		}

		return choice
	}
}

func runTask(taskID int) {
	// Find task by ID
	var selectedTask Task
	found := false
	for _, task := range tasks {
		if task.ID == taskID {
			selectedTask = task
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		return
	}

	fmt.Printf("\nStarting Task %d: %s\n", taskID, selectedTask.Name)
	fmt.Println("=" + strings.Repeat("=", len(selectedTask.Name)+20))
	fmt.Printf("Directory: %s\n", selectedTask.Directory)
	fmt.Printf("Command: %s\n\n", selectedTask.Command)

	// Prepare command
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "cd "+selectedTask.Directory+" && "+selectedTask.Command)
	} else {
		cmd = exec.Command("bash", "-c", "cd "+selectedTask.Directory+" && "+selectedTask.Command)
	}

	// Set output to current process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\nError running task: %v\n", err)
	} else {
		fmt.Printf("\nTask %d completed successfully!\n", taskID)
	}

	fmt.Println("\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func showProjectStructure() {
	fmt.Println("\nProject Structure:")
	fmt.Println("======================")
	fmt.Println("golang-task/")
	fmt.Println("├── cmd/                     # Main applications")
	fmt.Println("│   ├── webserver/           # Task 1: Fiber web server")
	fmt.Println("│   │   └── main.go")
	fmt.Println("│   ├── cronjob/             # Task 2: Cron job scheduler")
	fmt.Println("│   │   └── main.go")
	fmt.Println("│   ├── inventory/           # Task 3: Inventory management")
	fmt.Println("│   │   └── main.go")
	fmt.Println("│   ├── basics/              # Task 4: Programming basics")
	fmt.Println("│   │   └── main.go")
	fmt.Println("│   └── aggregator/          # Task 5: Concurrent aggregator")
	fmt.Println("│       └── main.go")
	fmt.Println("├── pkg/                     # Library packages")
	fmt.Println("│   ├── inventory/           # Inventory management library")
	fmt.Println("│   │   └── inventory.go")
	fmt.Println("│   └── fetcher/             # Product fetcher library")
	fmt.Println("│       └── fetcher.go")
	fmt.Println("├── docs/                    # Documentation")
	fmt.Println("├── internal/                # Private application code")
	fmt.Println("├── main.go                  # Task launcher/menu")
	fmt.Println("├── go.mod                   # Go module declaration")
	fmt.Println("├── go.sum                   # Go dependencies")
	fmt.Println("└── README.md                # Project documentation")
	fmt.Println()
}

func main() {
	for {
		// Clear screen (works on both Windows and Unix)
		if runtime.GOOS == "windows" {
			exec.Command("cmd", "/c", "cls").Run()
		} else {
			exec.Command("clear").Run()
		}

		displayWelcome()
		showProjectStructure()
		displayTasks()

		choice := getUserChoice()

		if choice == 0 {
			fmt.Println("\nThank you for using Golang Task Manager!")
			fmt.Println("Goodbye!")
			break
		}

		runTask(choice)
	}
}
