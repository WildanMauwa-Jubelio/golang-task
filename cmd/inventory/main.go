package main

import (
	"fmt"

	"github.com/jubelio/learn_golang/pkg/inventory"
)

// TASK 3: Inventory Management Demo Function
func runInventoryDemo() {
	fmt.Println("\nStarting Inventory Management System Demo...")

	// Create a new inventory
	inv := inventory.NewInventory()

	// Add several different items
	fmt.Println("\nAdding items to inventory:")

	inv.AddItem(inventory.Item{Name: "Laptop", Price: 1200.00, Quantity: 5})
	inv.AddItem(inventory.Item{Name: "Mouse", Price: 25.50, Quantity: 20})
	inv.AddItem(inventory.Item{Name: "Keyboard", Price: 75.00, Quantity: 15})
	inv.AddItem(inventory.Item{Name: "Monitor", Price: 300.00, Quantity: 8})
	inv.AddItem(inventory.Item{Name: "Headphones", Price: 150.00, Quantity: 12})

	// Print initial inventory
	inv.PrintInventory()

	// Demonstrate updating existing items
	fmt.Println("\nUpdating existing items:")
	inv.AddItem(inventory.Item{Name: "Laptop", Price: 1150.00, Quantity: 3}) // Update price and add quantity
	inv.AddItem(inventory.Item{Name: "Mouse", Price: 22.00, Quantity: 10})   // Update price and add quantity

	// Add new item after initial setup
	inv.AddItem(inventory.Item{Name: "Webcam", Price: 80.00, Quantity: 6})

	// Print final inventory state
	fmt.Println("\nFinal inventory state:")
	inv.PrintInventory()

	// Demonstrate additional methods
	fmt.Printf("\nTotal items in inventory: %d\n", inv.GetItemCount())
	fmt.Printf("Total inventory value: $%.2f\n", inv.GetTotalValue())

	// Get specific item
	if item, exists := inv.GetItem("Laptop"); exists {
		fmt.Printf("Found item: %s - $%.2f (Qty: %d)\n", item.Name, item.Price, item.Quantity)
	}
}

func main() {
	fmt.Println("TASK 3: Inventory Management System")
	fmt.Println("==================================")

	runInventoryDemo()

	fmt.Println("\nInventory Management System Demo completed!")
}
