package main

import "fmt"

func  swapStr(x,y string) (string, string)  {
	return y, x


}
func main() {
	fmt.Print(swapStr("x", "y"))

}