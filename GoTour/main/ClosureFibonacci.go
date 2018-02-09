package main

import "fmt"

func fibonacci() func() int64 {
	a := int64(0)
	b := int64(1)
	return func() int64 {
		result := a + b
		a = b
		b = result
		return result
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
