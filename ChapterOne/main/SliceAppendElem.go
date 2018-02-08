package main

import "fmt"

func main() {
	var a []int
	PrintSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	PrintSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	PrintSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	PrintSlice("a", a)
}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
