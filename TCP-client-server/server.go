package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]

	con, err := net.Listen("tcp", args[0]+":"+args[1])
	if err != nil {
		panic(err)
	}
	defer con.Close()

	conn, err := con.Accept()

	// read message from connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}

