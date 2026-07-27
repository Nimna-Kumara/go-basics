package main

import "fmt"

// "", "\n"

func classify(n int) string {
	// if with an init statement: `remainder` is scoped to this if/else chain
	if remainder := n % 2; remainder == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main1() {
	fmt.Println(7, "is", classify(7))
	fmt.Println(10, "is", classify(10))

	// switch - no fallthrough by default
	day := 3
	switch day {
	case 1, 7:
		fmt.Println("Weekend")
	case 2, 3, 4, 5, 6:
		fmt.Println("Invalid day")
	}

	// switch with no expression == clean if/else chain
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C or below")
	}

	// explicit fallthrough (rare, but available)
	switch 1 {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("also two, because of fallthrough")
	}
}
