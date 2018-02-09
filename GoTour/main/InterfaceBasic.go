package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

type ExFloat float64

func (f ExFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type ExVertex struct {
	X, Y float64
}

func (v *ExVertex) Abs() float64 {
	return math.Hypot(v.X, v.Y)
}

func main() {
	var a Abser
	f := ExFloat(-math.Sqrt2)
	v := ExVertex{3, 4}
	a = f
	a = &v
	a = &f
	fmt.Println(a.Abs())
}
