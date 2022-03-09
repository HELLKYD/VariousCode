package gomodule

import (
	"fmt"
	"net"
)

func StartTestServer() {
	dStream, err := net.Listen("tcp", ":9998")

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
		go doSomething(connection)
	}
}

func doSomething(con net.Conn) {
	fmt.Println("User connected")
	con.Close()
}
