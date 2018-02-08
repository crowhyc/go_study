package main

import (
	"math"
	"fmt"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Hypot(x, y)

	}
	fmt.Println(hypot(3, 4))
}
