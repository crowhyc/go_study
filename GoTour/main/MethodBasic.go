package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Hypot(v.X, v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v)

}
