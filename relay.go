package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	var echoServerConn net.Conn
	var clientConn net.Conn

	fmt.Println("listening on port 8081")

	echoServerlistener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer echoServerlistener.Close()

	for {
		echoServerConn, err = echoServerlistener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		if echoServerConn != nil {

			fmt.Println("established relay address: localhost:8082")

			clientlistener, err := net.Listen("tcp", ":8082")

			if err != nil {
				log.Fatal(err)
			}
			defer clientlistener.Close()

			clientConn, err = clientlistener.Accept()
			if err != nil {
				log.Fatal(err)
			}

		}

		if echoServerConn != nil && clientConn != nil {
			go forwardData(echoServerConn, clientConn)
		}

		fmt.Println("established relay address: localhost:8081")

	}
}

func forwardData(echoServerConn net.Conn, clientConn net.Conn) {

	for {
		bufferMem := make([]byte, 1400)

		// read data from client
		data, err := clientConn.Read(bufferMem)
		if err != nil {
			log.Println("connection error in client side")
		}
		rdata := bufferMem[:data]

		fmt.Println("message recived : ", string(rdata))

		// send client data to echo server
		p, err := echoServerConn.Write(rdata)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(p)

		// read bounced data from echo server
		bouncedData, err := echoServerConn.Read(bufferMem)
		if err != nil {
			log.Println("connection error in echo server side")
		}
		bdata := bufferMem[:bouncedData]

		// write bounced data back to client
		q, err := clientConn.Write(bdata)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(q)

	}

}
