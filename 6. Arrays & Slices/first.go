package main

import "fmt"

func main() {
	// Array: fixed size, part of the type
	var arr [5]int
	arr[0] = 10
	arr[4] = 50
	fmt.Println("array:", arr, "len:", len(arr))

	// Array literal
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes)

	// Slice literal - most common way to create a slice
	nums := []int{1, 2, 3}
	fmt.Println("slice:", nums, "len:", len(nums), "cap:", cap(nums))

	// append grows the slice (may reallocate)
	nums = append(nums, 4, 5)
	fmt.Println("after append:", nums, "len:", len(nums), "cap:", cap(nums))

	// make() creates a slice with explicit length and capacity
	buffer := make([]int, 3, 10) // length 3, capacity 10
	fmt.Println("buffer:", buffer, "len:", len(buffer), "cap:", cap(buffer))

	// slicing: s[low:high] - half-open range [low, high)
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("letters[1:3]:", letters[1:3]) // ["b" "c"]
	fmt.Println("letters[:2]:", letters[:2])   // ["a" "b"]
	fmt.Println("letters[3:]:", letters[3:])   // ["d" "e"]

	// IMPORTANT: slices sharing an underlying array
	original := []int{1, 2, 3, 4, 5}
	view := original[1:4]
	view[0] = 999 // this mutates `original` too!
	fmt.Println("original after mutating view:", original)

	// 2D slices (slice of slices)
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("grid[1][2]:", grid[1][2])

	// copy() makes an independent duplicate
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 100
	fmt.Println("src:", src, "dst:", dst) // src is untouched
}
