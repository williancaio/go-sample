package main

import (
	"fmt"
	"math"
)

func main() {
	r := retangle{1, 2}
	writeArea(r)

}

type retangle struct {
	h float64
	l float64
}

type cicle struct {
	r float64
}

type form interface {
	area() float64
}

func writeArea(f form) {
	fmt.Println(f.area())
}

func (r retangle) area() float64 {
	return r.h * r.l
}

func (c cicle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
