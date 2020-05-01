package main

import "fmt"

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

func main() {
	c1 := Circle{60}
	s1 := Square{20}
	// one way of printing
	fmt.Println(c1.area())
	fmt.Println(s1.area())
	// second way of printng
	PrintShape(c1)
}

func PrintShape(sh Shape) {
	fmt.Println(sh.area())
}

func (s Square) area() float64 {
	return s.Side * 4
}

func (c Circle) area() float64 {
	return c.Radius * 60
}
