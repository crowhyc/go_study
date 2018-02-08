package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	strArr := strings.Fields(s)
	var m = make(map[string]int)
	for _, val := range strArr {
		if v, ok := m[val]; ok == false {
			m[val] = 1
		} else {
			m[val] = v + 1
		}
	}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
