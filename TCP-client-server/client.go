package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]

	conn, err := net.Dial("tcp", args[0]+":"+args[1])
	if err != nil {
		panic(err)
	}
	defer conn.Close()


	var msg string
	fmt.Scanf("%s", &msg)

	conn.Write([]byte(msg))
}