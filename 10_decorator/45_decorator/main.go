package main

import "fmt"

// We have a bunch of geometric shapes that we have in the system
// and we want to extend the functionality of those geometric shapes
// by giving them additional properties.

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// If having too many classes, this is not good
// type ColoredSquare struct {
//     Square
//     Color string
// }

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

// You can apply decorator to another decorator
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency",
		t.Shape.Render(), t.Transparency*100)
}

func main() {
	circle := Circle{2}
	circle.Resize(2) // Only Circle shape has Resize method
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
	// redCircle.Resize(2)  // Cannot do this, Resize method is not available

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
