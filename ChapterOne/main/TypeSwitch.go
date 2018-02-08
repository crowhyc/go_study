package main

import (
	"math"
	"fmt"
)

func main() {
	x, y := 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = int(f)
	fmt.Print(x, y, z)

}
