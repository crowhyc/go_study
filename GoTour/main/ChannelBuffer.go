package main

import "fmt"

func main() {
	c := make(chan int, 9)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
