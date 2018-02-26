package main

import (
	"math"
	"fmt"
)

type MethodBasicType struct {
	X, Y float64
}

func (v *MethodBasicType) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := &MethodBasicType{3, 4}
	fmt.Println(v.Abs())
}
