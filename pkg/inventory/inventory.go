package inventory

import "fmt"

// Item struct represents an inventory item
type Item struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Inventory struct manages the collection of items
type Inventory struct {
	Items map[string]Item `json:"items"`
}

// NewInventory creates a new inventory instance
func NewInventory() *Inventory {
	return &Inventory{
		Items: make(map[string]Item),
	}
}

// AddItem adds a new item to the inventory or updates existing item
func (inv *Inventory) AddItem(item Item) {
	if existingItem, exists := inv.Items[item.Name]; exists {
		// Update existing item: update price and add to quantity
		existingItem.Price = item.Price
		existingItem.Quantity += item.Quantity
		inv.Items[item.Name] = existingItem
		fmt.Printf("Updated item: %s (New Quantity: %d, Price: %.2f)\n",
			item.Name, existingItem.Quantity, existingItem.Price)
	} else {
		// Add new item
		inv.Items[item.Name] = item
		fmt.Printf("Added new item: %s (Quantity: %d, Price: %.2f)\n",
			item.Name, item.Quantity, item.Price)
	}
}

// PrintInventory prints the details of all items in the inventory
func (inv *Inventory) PrintInventory() {
	fmt.Println("\n========================================")
	fmt.Println("INVENTORY REPORT")
	fmt.Println("========================================")

	if len(inv.Items) == 0 {
		fmt.Println("No items in inventory.")
		return
	}

	fmt.Printf("%-20s %-10s %-10s\n", "Name", "Price", "Quantity")
	fmt.Println("----------------------------------------")

	totalValue := 0.0
	for _, item := range inv.Items {
		itemValue := item.Price * float64(item.Quantity)
		totalValue += itemValue
		fmt.Printf("%-20s $%-9.2f %-10d\n", item.Name, item.Price, item.Quantity)
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("Total Inventory Value: $%.2f\n", totalValue)
	fmt.Println("========================================")
}

// GetTotalValue calculates the total value of all items in inventory
func (inv *Inventory) GetTotalValue() float64 {
	totalValue := 0.0
	for _, item := range inv.Items {
		totalValue += item.Price * float64(item.Quantity)
	}
	return totalValue
}

// GetItemCount returns the total number of different items
func (inv *Inventory) GetItemCount() int {
	return len(inv.Items)
}

// GetItem retrieves a specific item by name
func (inv *Inventory) GetItem(name string) (Item, bool) {
	item, exists := inv.Items[name]
	return item, exists
}
