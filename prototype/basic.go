package prototype

import (
	"fmt"
)

// Prototype interface that requires a Clone method
type Prototype interface {
	Clone() Prototype
	GetInfo() string
}

// Concrete struct: Circle
type Circle struct {
	Radius int
	Color  string
}

// Clone method for Circle
func (c *Circle) Clone() Prototype {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

// GetInfo for displaying object details
func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Circle: Radius=%d, Color=%s", c.Radius, c.Color)
}

// Concrete struct: Rectangle
type Rectangle struct {
	Width  int
	Height int
	Color  string
}

// Clone method for Rectangle
func (r *Rectangle) Clone() Prototype {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

// GetInfo for displaying object details
func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Rectangle: Width=%d, Height=%d, Color=%s", r.Width, r.Height, r.Color)
}

// This is a struct to be invoked by main.go
type BasicPrototype struct{}

func (b BasicPrototype) ProcessPattern() bool {
	// Create an original circle
	originalCircle := &Circle{Radius: 10, Color: "Red"}
	clonedCircle := originalCircle.Clone().(*Circle) // Clone the circle

	// Modify the cloned object to show they are independent
	clonedCircle.Color = "Blue"

	// Create an original rectangle
	originalRectangle := &Rectangle{Width: 20, Height: 15, Color: "Green"}
	clonedRectangle := originalRectangle.Clone().(*Rectangle) // Clone the rectangle

	// Modify the cloned rectangle
	clonedRectangle.Width = 25

	// Print results
	fmt.Println("Original and Cloned Objects:")
	fmt.Println(originalCircle.GetInfo())    // Circle: Radius=10, Color=Red
	fmt.Println(clonedCircle.GetInfo())      // Circle: Radius=10, Color=Blue
	fmt.Println(originalRectangle.GetInfo()) // Rectangle: Width=20, Height=15, Color=Green
	fmt.Println(clonedRectangle.GetInfo())   // Rectangle: Width=25, Height=15, Color=Green

	return true
}
