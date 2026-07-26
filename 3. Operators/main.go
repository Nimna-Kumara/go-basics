package main

import "fmt"

func main() {
	a, b := 17, 5

	// Arithmatic
	fmt.Print("\nArithmatic")
	fmt.Print("\nadd: ", a+b, "\nsub: ", a-b, "\nmul: ", a*b, "\ndiv: ", a/b, "\nmod: ", a%b)

	// Comparison
	fmt.Print("\n\nComparison")
	fmt.Print("\neq: ", a == b, "\nneq: ", a != b, "\ngt: ", a > b, "\nlte: ", a <= b)

	// Logical
	fmt.Print("\n\nLogical")
	t, f := true, false
	fmt.Print("\nand: ", t && f, "\nor: ", t || f, "\nnot: ", !t)

	// Bitwise
	fmt.Print("\n\nBitwise")
	fmt.Print("\nand: ", a&b, "\nor: ", a|b, "\nxor: ", a^b, "\nandnot: ", a&^b, "\nlshift: ", a<<1, "\nrshift: ", a>>1, "\n")

	// Increment/decrement are STATEMENTS in Go, not expressions
	x := 10
	x++
	x--
	fmt.Println("x after ++ and --: ", x)

	// Go has NO ternary operator (a ? b : c). Use if/else instead.
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("max: ", max)

}
