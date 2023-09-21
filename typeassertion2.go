package main

import (
	"fmt"
	"math"
)

// Shape is an interface for shapes that can calculate their area.
type Shape interface {
	Area() float64
}

// Circle is a type representing a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle is a type representing a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Create a slice of shapes
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Width: 4.0, Height: 6.0},
	}
//we are making it as a slice of iniializing both the structs
	// Calculate and print the area of each shape
	for _, shape := range shapes {
		// Use type assertion to access the Area() method of each shape
		area := shape.Area()
		fmt.Printf("Area: %.2f\n", area)
	}
}