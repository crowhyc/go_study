package main

import (
	"math"
	"fmt"
)

type MWPR struct {
	X, Y float64
}

//当方法接收的是值而不是指针时,修改的只是这个值的副本,而不是真正的对象
func (v *MWPR) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *MWPR) Abs() float64 {
	return math.Hypot(v.X, v.Y)
}

func main() {
	v := &MWPR{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
