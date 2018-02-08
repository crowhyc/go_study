package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().MarshalBinary())
}
