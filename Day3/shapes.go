package main

import "math"

type Shape interface{
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}
// method for the struct
func (r Rectangle) Area() float64  {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64  {
    return c.Radius* c.Radius * math.Pi
}


type Triangle struct{
    Base float64
    Height float64
}
func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}
// func Perimeter(length  , breadth float64) float64{
// 	return 2 *( length+breadth)
// }

func Perimeter(rectangle Rectangle) float64{
	return 2 *( rectangle.Width+rectangle.Height)
}


