package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:9000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	CheckError2(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	CheckError2(err)
	_, err = conn.Write([]byte("anything"))
	CheckError2(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	CheckError2(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
func CheckError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
