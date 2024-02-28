package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	// "time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8000")

	for {
		fmt.Println("I am here+++++++++++++++++++++++++++++++++++")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error occured: %s\n", err)
			continue
		}
		// conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		fmt.Println("I am hereee----------------------")
		n, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\nHAHAHA\r\n")
				return
			}
			fmt.Printf("error: %s\n", err)
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\nYOOOOOO\r\n")
			return
		}

		fmt.Printf("Receivedddddddddddd %s\n", buffer[:n])
		fmt.Fprintf(conn, "Server sent %s", buffer[:n])
		fmt.Println("I am RECUVEDDDDDD ENDDDDDDD+_+_+_+_+_+_+_+_+_+-")

	}
}