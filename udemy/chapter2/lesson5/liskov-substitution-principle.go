package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}

type Rectangle struct {
	Width  int
	Height int
}

func(r *Rectangle) GetWidth() int {
	return r.Width
}
func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	excpectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()
	fmt.Print("Expected area: ", excpectedArea, " but got actual area: ", actualArea)

}

type Square struct {
	Rectangle
}

func NewSquare(width int) *Square {
	return &Square{
		Rectangle: Rectangle{
			Width: width,
			Height: width,
		},
	}
}

func (s *Square) SetWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height int) {
	s.Height = height
	s.Width = height
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc:= &Rectangle{2, 3}
	UseIt(rc)
	sq := NewSquare(5)
	UseIt(sq)
	
 }