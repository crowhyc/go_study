package main

import "fmt"

type Element interface {
}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name:%v age%v", p.name, p.age)
}
func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dean", 30}
	for i, v := range list {
		if value, ok := v.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, value)

		} else if value, ok := v.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}
}
