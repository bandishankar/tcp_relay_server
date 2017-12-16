package main


import (
	"net"
	"bufio"
	"strconv"
	"fmt"
	"strconv"
)



//var relayServerPort = flag.String("serverport", "8081", " relay server port ")


var firstClient net.Listener ;
var connFirst net.Conn;


func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {

	// Create a waitgroup that will handle
	//wg := sync.WaitGroup{}

	// Add a counter to the waitgroup and then run the
	// http server in a goroutine
	//wg.Add(1)


	//fmt.Println("Launching server...")
	//
	//// listen on all interfaces
	//firstClient, _ = net.Listen("tcp", ":8081")
	//secondClient,_ := net.Listen("tcp", ":8082")
	//
	//// accept connection on port
	//connFirst, _ = firstClient.Accept()
	//connSecond, _ := secondClient.Accept()
	//
	//Pipe(connFirst,connSecond)


	server, _ := net.Listen("tcp", ":8081")
	server1, _ := net.Listen("tcp", ":8082")
	if server == nil {
		panic("couldn't start listening: " )
	}

	if server1 == nil {
		panic("couldn't start listening: " )
	}
	conns := clientConns(server)
	conns1 := clientConns(server1)
	for {
		go handleConn(<-conns)
		go handleConn(<-conns1)
	}

	//// run loop forever (or until ctrl-c)
	//for {
	//	// will listen for message to process ending in newline (\n)
	//	message, _ := bufio.NewReader(connFirst).ReadString('\n')
	//	// output message received
	//	fmt.Print("Message Received:", string(message))
	//	// sample process for string received
	//	//newmessage := strings.ToUpper(message)
	//	// send new string back to client
	//	connSecond.Write([]byte(message + "\n"))
	//
	//	rmessage, _ := bufio.NewReader(connSecond).ReadString('\n')
	//	fmt.Print("Message returned :", string(rmessage))
	//
	//	connFirst.Write([]byte(rmessage + "\n"))
	//}




	//for {
	//	conn, err := firstClient.Accept()
	//	check(err, "Accepted connection.")
	//
	//	go func() {
	//		buf := bufio.NewReader(connFirst)
	//
	//		for {
	//			message, err := buf.ReadString('\n')
	//
	//			if err != nil {
	//				fmt.Printf("Client disconnected.\n")
	//				break
	//			}
	//
	//			newmessage := strings.ToUpper(message)
	//
	//			conn.Write([]byte(newmessage + "\n"))
	//		}
	//	}()
	//}



}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}



func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, _ := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: ")
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(),
				client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

