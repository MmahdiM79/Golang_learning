package main

import (
	"fmt"
	"net"
	"os"
)

const BUFFER_SIZE = 8192

func main() {
	args := os.Args[1:]

	con, err := net.Listen("tcp", args[0]+":"+args[1])
	if err != nil {
		panic(err)
	}
	defer con.Close()
	fmt.Printf("Server started on %s ...\n\n", con.Addr())

	for {
		conn, err := con.Accept()
		if err != nil {
			panic(err)
		}

		go clientHandler(conn)
	}
}

func clientHandler(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("\nClient connected from: %s\n\n", conn.RemoteAddr(), )

	for {
		// read message from connection
		buf := make([]byte, BUFFER_SIZE)
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		msg := string(buf[:n])
		if msg == "bye" {
			break
		}

		fmt.Printf("\n[Client %s]: %s\n", conn.RemoteAddr(), msg)
	}

	fmt.Println("Client", conn.RemoteAddr(), "disconnected")
}
