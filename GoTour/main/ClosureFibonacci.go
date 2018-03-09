package main

import (
	"fmt"
	"runtime"
)

func fibonacci() func() int64 {
	a, b := 0, 1
	return func() int64 {
		a, b = b, a+b
		return int64(b)
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
