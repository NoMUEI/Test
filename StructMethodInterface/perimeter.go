package StructMethodInterface

import "math"

type Rectangle struct {
	Width float64
	Hight float64
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}


func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Hight + rectangle.Width) * 2

}

func (r Rectangle)Area() float64 {
	return r.Width * r.Hight

}

func (c Circle)Area() float64 {
	return math.Pi*c.Radius*c.Radius
}

