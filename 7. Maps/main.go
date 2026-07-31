package main

import "fmt"

func main() {
	// map literal
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println(ages)

	// make() also creates maps
	scores := make(map[string]int)
	scores["math"] = 90
	scores["science"] = 85
	fmt.Println(scores)

	// The "comma ok" idiom - safely check if a key exists
	value, ok := scores["history"]
	fmt.Println("history score:", value, "present:", ok)

	// Deleting a key
	delete(scores, "math")
	fmt.Println("after delete:", scores)

	// Iterating - NOTE: map iteration order is randomized by Go on purpose!
	for subject, score := range scores {
		_ = subject
		_ = score // (not printing here since order isn't guaranteed)
	}
	fmt.Println("iterated over", len(scores), "entries")

	// Maps with struct values
	type Point struct{ X, Y int }
	locations := map[string]Point{
		"origin": {0, 0},
		"unit":   {1, 1},
	}
	fmt.Println("origin:", locations["origin"])

	// Zero value of a map is nil - reading is fine, writing panics!
	var nilMap map[string]int
	fmt.Println("read from nil map:", nilMap["anything"]) // safe, returns zero value
	// nilMap["x"] = 1 // this would PANIC: assignment to entry in nil map
}
