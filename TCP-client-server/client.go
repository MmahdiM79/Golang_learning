package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

const BUFFER_SIZE = 8192

func main() {
	args := os.Args[1:]

	conn, err := net.Dial("tcp", args[0]+":"+args[1])
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("Connected to server %s\n\n", conn.RemoteAddr())

	go serverMessages(conn)

	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		msg, _ = reader.ReadString('\n')
		msg = msg[:len(msg)-1]
		
		if msg == "bye" {
			break
		}

		conn.Write([]byte(msg))
	}
}


func serverMessages(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, BUFFER_SIZE)
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		msg := string(buf[:n])
		if msg == "bye" {
			fmt.Println("\n\nServer disconnected :( ")
			os.Exit(0)
		}
	}
}

