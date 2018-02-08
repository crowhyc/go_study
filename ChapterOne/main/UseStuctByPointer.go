package main

import "fmt"

type VertexC struct {
	X int
	Y int
}

func main() {
	v := VertexC{12, 3}
	p := &v
	p.X = 1e9
	fmt.Print(v)

}
