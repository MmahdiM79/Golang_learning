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
	for {
		fmt.Scanf("%s", &msg)
		if msg == "bye" {
			break
		}

		conn.Write([]byte(msg))
	}
}