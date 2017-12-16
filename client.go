package main

import (
	"bufio"
	//"os"
	"fmt"
	"net"
	"strings"
)



func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8082")
	for {
		// read in input from stdin
		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Text to send: ")
		//text, _ := reader.ReadString('\n')
		// send to socket
		//fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("Message from server: "+message)

		newmessage := strings.ToUpper(message)

		conn.Write([]byte(newmessage + "\n"))
	}
}

