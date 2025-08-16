package main

import "fmt"

// Polymorphism: one interface, many implementations
// Inheritance: Acquiring properties and behaviors of one type into another type
// Encapsulation: Wrapping data (fields) and methods together while restricting direct access to internal details

// write a code to implement Polymorphism

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	Radius float32
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (c *Circle) Area() float32 {
	return 2 * 3.14 * c.Radius
}
func (c *Circle) Perimeter() float32 {
	return 3.14 * c.Radius * c.Radius
}
func (c *Rectangle) Area() float32 {
	return c.Length * c.Breadth
}
func (c *Rectangle) Perimeter() float32 {
	return 2 * (c.Length + c.Breadth)
}

func main() {

	var s Shape

	s = &Circle{Radius: 2}
	fmt.Println("op2:", s.Area())

	s = &Rectangle{Length: 2, Breadth: 2}
	fmt.Println("op4:", s.Area())

}
