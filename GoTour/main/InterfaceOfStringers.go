package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Liu yuanyuan ", 18}
	b := Person{"Shi tt", 250}
	fmt.Println(a, b)

}
