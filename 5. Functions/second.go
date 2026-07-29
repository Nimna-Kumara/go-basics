package main

import "fmt"

// A closure factory: returns a function that "remembers" its own counter.
func makeCounter() func() int {
    count := 0
    return func() int {
        count++ // captures and mutates the outer `count` variable
        return count
    }
}

// Functions as values: passed as an argument (higher-order function)
func applyTwice(f func(int) int, x int) int {
    return f(f(x))
}

// Recursion: classic factorial
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    counterA := makeCounter()
    counterB := makeCounter()
    fmt.Println(counterA(), counterA(), counterA()) // 1 2 3 - independent state
    fmt.Println(counterB())                          // 1 - a fresh closure, own state

    double := func(x int) int { return x * 2 } // anonymous function literal
    fmt.Println("applyTwice(double, 3):", applyTwice(double, 3))

    fmt.Println("factorial(6):", factorial(6))

    // defer: LIFO - "last deferred, first run" - executes right before the function returns
    fmt.Println("--- defer demo ---")
    demoDefer()
}

func demoDefer() {
    fmt.Println("start")
    defer fmt.Println("deferred 1 (runs last)")
    defer fmt.Println("deferred 2 (runs second)")
    defer fmt.Println("deferred 3 (runs first, since it was deferred last)")
    fmt.Println("end of function body")
}
