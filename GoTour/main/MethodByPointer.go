package main

import (
	"math"
	"fmt"
)

type VertexPointer struct {
	X, Y float64
}

func (v *VertexPointer) Scale(f float64) {
	v.Y = v.Y * f
	v.X = v.X * f
}
func (v VertexPointer) Abs() float64 {
	return math.Hypot(v.X, v.Y)
}
func main() {
	v := &VertexPointer{3, 4}
	v.Scale(3)
	fmt.Println(v, v.Abs())

}
