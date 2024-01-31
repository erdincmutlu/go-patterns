package main

import "fmt"

// Liskov Substitution Principle
// If you have an API works with base class, it should also work with
// derived class

// Dealing with geometrics shapes

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// We are setting width and height together. This is violating the
// Liskov Substitution Principle
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("Expected an area of %d, but got %d\n", expectedArea, actualArea)
}

type Square2 struct {
	size int // width, height
}

func (s Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	// Everything in UseIt should work with Square, but it doesn't work
	sq := NewSquare(5)
	UseIt(sq)
}
