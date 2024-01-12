package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.ParseIP("127.0.0.1")
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}
