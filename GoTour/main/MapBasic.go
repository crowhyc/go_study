package main

import "fmt"

type MapStruct struct {
	Lat, Long float64
}

var m map[string]MapStruct

func main() {
	m = make(map[string]MapStruct)
	m["Bell Labs"] = MapStruct{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
