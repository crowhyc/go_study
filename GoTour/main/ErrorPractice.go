package main

import (
	"fmt"
	"math"
)

type SqrtError float64

func (e SqrtError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return math.Sqrt(2), nil
	}
	return 0, SqrtError(-2)
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
