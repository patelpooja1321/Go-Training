package main

import "fmt"

type Circle struct {
	radious float32
	pi      float32
}

func (c Circle) Area_cal() float32 {
	return c.pi * c.radious * c.radious

}

type Rectangle struct {
	Length, Width float32
}

func (r Rectangle) Area_cal() float32 {
	return r.Length * r.Width

}

type shape interface {
	Area_cal() float32
}

func main() {
	cirl := Circle{3.14, 10}
	rct := Rectangle{15, 25}

	result := []shape{cirl, rct}

	for _, shape := range result {
		fmt.Println(shape.Area_cal())
	}

}
