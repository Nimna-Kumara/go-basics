package main

import "fmt"

// Struct definition
type Point struct {
	X, Y int
}

// Struct with embedded struct (composition) and tags-friendly fields
type Rectangle struct {
	TopLeft Point
	Width   int
	Height  int
}

// Value receiver - operates on a COPY, cannot mutate the original
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Pointer receiver - can mutate the original struct
func (p *Point) MoveBy(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Method on Rectangle using another struct
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	p := Point{X: 1, Y: 2}
	fmt.Println("p:", p.String())

	p.MoveBy(3, 4) // Go automatically takes &p here
	fmt.Println("after MoveBy:", p.String())

	r := Rectangle{TopLeft: Point{0, 0}, Width: 10, Height: 5}
	fmt.Println("rectangle area:", r.Area())
	fmt.Println("top-left corner:", r.TopLeft.String())

	// Struct comparison - structs are comparable if all fields are comparable
	a := Point{1, 1}
	b := Point{1, 1}
	fmt.Println("a == b:", a == b)

	// Anonymous structs - handy for quick one-off groupings
	person := struct {
		Name string
		Age  int
	}{Name: "Nimna", Age: 22}
	fmt.Println(person)
}
