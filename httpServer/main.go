package main

import (
	"fmt"
	"net/http"
	"os"
	"errors"
	"io"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello, HTTP!\n")
}