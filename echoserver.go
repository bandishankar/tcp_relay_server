package main

import (
	"bufio"
	"fmt"
	"strings"
	"net"
)

// simple echo server

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("Message from server: "+message)

		newmessage := strings.ToUpper(message)

		conn.Write([]byte(newmessage + "\n"))
	}
}

