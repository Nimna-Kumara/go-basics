package main

import "fmt"

// Multiple return values - the (result, error) idiom is everywhere in Go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

// Named return values - "sum" and "product" are pre-declared as zero values
// and returned automatically by a bare `return`.
func sumAndProduct(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // "naked" return - returns sum and product automatically
}

// Variadic parameters - accept zero or more ints
func total(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func main1() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("10 / 2 = ", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	s, p := sumAndProduct(3, 4)
	fmt.Println("sum:", s, "product:", p)

	fmt.Println("total:", total(1, 2, 3, 4))
	fmt.Println("total of nothing:", total())

	// Spreading a slice into a variadic call with ...
	nums := []int{10, 20, 30}
	fmt.Println("total of slice:", total(nums...))
	// The ... operator tells Go to unpack (expand) the slice into individual arguments.

}
