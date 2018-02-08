package main

import "fmt"

func sqrt(x float64) float64 {
	var last float64
	result := x/2
	for i:=0;i<10;i++{
		last=result
		result = (result+x/result)/2
		fmt.Println(last)
	}
	return result
}

func main() {
	fmt.Println(sqrt(2))
}
