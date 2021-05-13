package geometry

import (
	"encoding/json"
	"fmt"
	"math"
)

type Triangle struct {
	size int
}

func (t *Triangle) doubleSize() {
	t.size *= 2
}
func (t *Triangle) SetSize(size int) {
	t.size = size
}
func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}
func (t *Triangle) name() int {
	return t.size * 2
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	Size float64
}

func (s Square) Perimeter() float64 {
	return s.Size * s.Size
}

func (s Square) Area() float64 {
	return s.Size * 4
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.Radius
}
func PrintInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

type CustomWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (c CustomWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, s := range resp {
		fmt.Println(s.FullName)
	}
	return len(p), nil
}
