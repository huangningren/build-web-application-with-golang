package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:8989"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)
	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
