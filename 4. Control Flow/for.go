package main

import "fmt"

func main() {
	// classic 3-part for loop
	for i := 0; i < 3; i++ {
		fmt.Println("classic: ", i)
	}

	// "while" style - just the condition
	n := 0
	for n < 3 {
		fmt.Println("while-style:", n)
		n++
	}

	// infinite loop with a break
	counter := 0
	for {
		if counter >= 3 {
			break
		}
		fmt.Println("infinite+break: ", counter)
		counter++
	}

	// range over a slice
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	// labeled loops for breaking out of nested loops
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer // skips to next i, not just next j
			}
			fmt.Println("i,j: ", i, j)
		}
	}
}
