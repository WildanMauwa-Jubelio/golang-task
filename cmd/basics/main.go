package main

import "fmt"

// TASK 4: Loops, Conditionals, and Functions
// TASK 4A: Loops - The for statement
// Task A: FizzBuzz
func runFizzBuzz() {
	fmt.Println("\nFizzBuzz (1-100):")
	fmt.Println("===================")

	for i := 1; i <= 100; i++ {
		if i%15 == 0 { // Multiple of both 3 and 5
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 { // Multiple of 3
			fmt.Print("Fizz ")
		} else if i%5 == 0 { // Multiple of 5
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}

		// Add line break every 10 items for better readability
		if i%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Task B: Slice Summation
func sumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func testSliceSummation() {
	fmt.Println("\nSlice Summation Test:")
	fmt.Println("========================")

	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{10, 20, 30},
		{-5, -10, 15, 20},
		{100},
		{},
	}

	for i, slice := range testCases {
		sum := sumSlice(slice)
		fmt.Printf("Test %d: %v -> Sum: %d\n", i+1, slice, sum)
	}
}

// TASK 4B: Conditional Logic - The switch Statement
// Task A: Day of the Week
func getDayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func testDayOfWeek() {
	fmt.Println("\nDay of the Week Test:")
	fmt.Println("========================")

	// Test all days plus invalid cases
	testDays := []int{1, 2, 3, 4, 5, 6, 7, 0, 8, -1}

	for _, day := range testDays {
		result := getDayOfWeek(day)
		fmt.Printf("Day %d: %s\n", day, result)
	}
}

// Task B: Type Inspector
func inspectType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("The type is int (value: %d)\n", v)
	case string:
		fmt.Printf("The type is string (value: %s)\n", v)
	case bool:
		fmt.Printf("The type is bool (value: %t)\n", v)
	default:
		fmt.Printf("The type is unknown (value: %v, type: %T)\n", v, v)
	}
}

func testTypeInspector() {
	fmt.Println("\nType Inspector Test:")
	fmt.Println("=======================")

	// Test different types
	testValues := []interface{}{
		42,
		"Hello, World!",
		true,
		false,
		3.14,
		[]int{1, 2, 3},
		map[string]int{"key": 123},
	}

	for i, value := range testValues {
		fmt.Printf("Test %d: ", i+1)
		inspectType(value)
	}
}

// TASK 4C: Functions
// Task A: Multiple Return Values
func calculate(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func testCalculate() {
	fmt.Println("\nCalculate Function Test:")
	fmt.Println("===========================")

	// Test cases
	testCases := [][2]int{
		{5, 3},
		{10, 7},
		{-4, 6},
		{0, 15},
		{12, 12},
	}

	for _, testCase := range testCases {
		a, b := testCase[0], testCase[1]
		sum, product := calculate(a, b)
		fmt.Printf("calculate(%d, %d) -> Sum: %d, Product: %d\n", a, b, sum, product)
	}
}

// Task B: Variadic Sum Function
func sumAll(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func testSumAll() {
	fmt.Println("\nVariadic Sum Function Test:")
	fmt.Println("==============================")

	// Test with different numbers of arguments
	fmt.Printf("sumAll(1, 2) = %d\n", sumAll(1, 2))
	fmt.Printf("sumAll(5, 10, 15, 20) = %d\n", sumAll(5, 10, 15, 20))
	fmt.Printf("sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) = %d\n", sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Printf("sumAll() = %d\n", sumAll())
	fmt.Printf("sumAll(100) = %d\n", sumAll(100))
	fmt.Printf("sumAll(-5, -10, 15, 20, -3) = %d\n", sumAll(-5, -10, 15, 20, -3))
}

// TASK 4: Main Demo Function
func runTask4Demo() {
	fmt.Println("\nStarting Task 4 Demo: Loops, Conditionals, and Functions")
	fmt.Println("=============================================================")

	// Loops: The for statement
	runFizzBuzz()
	testSliceSummation()

	// Conditional Logic: The switch Statement
	testDayOfWeek()
	testTypeInspector()

	// Functions
	testCalculate()
	testSumAll()

	fmt.Println("\nTask 4 Demo completed!")
	fmt.Println("=========================")
}

func main() {
	fmt.Println("TASK 4: Loops, Conditionals, and Functions")
	fmt.Println("==========================================")

	runTask4Demo()
}
