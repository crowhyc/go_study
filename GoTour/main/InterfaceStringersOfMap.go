package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	strArr := make([]string, len(ip))
	for i := range ip {
		strArr[i] = string(ip[i])
	}
	fmt.Println(string(ip[:]))
	fmt.Printf("%v,length %d \n", strArr, len(strArr))
	return strings.Join(strArr, ".")
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for i, v := range addrs {
		fmt.Printf("%v : %v\n", i, v)
	}
}
