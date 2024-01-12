package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := "127.0.0.1:9000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError3(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError3(err)
	for {
		handleClient2(conn)
	}
}
func handleClient2(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
func checkError3(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
