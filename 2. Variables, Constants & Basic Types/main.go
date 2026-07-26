package main

import "fmt"

const Pi = 3.14159

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {

	fmt.Println("Pi: ", Pi)
	fmt.Println("Wednesday is day number:", Wednesday)
	fmt.Print("\nKB: ", KB, "\nMB: ", MB, "\nGB: ", GB)
}
