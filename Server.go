package gomodule

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func StartServer() {

	dStream, err := net.Listen("tcp", ":9998")
	fmt.Println("Server started")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer dStream.Close()

	for {
		connection, err := dStream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection)
	}
}
func handleConnection(connection net.Conn) {
	fmt.Println("User connected")
	sendMessage(connection, "Hello")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(input)
	}
	connection.Close()
}

func sendMessage(con net.Conn, text string) {
	io.WriteString(con, fmt.Sprint(text, "\n"))
}
