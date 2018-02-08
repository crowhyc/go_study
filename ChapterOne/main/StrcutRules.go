package main

import "fmt"

type VertexD struct {
	X, Y int
}

var (
	v1 = VertexD{1, 2}
	v2 = VertexD{X: 1}
	v3 = VertexD{}
	p  = &VertexD{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
