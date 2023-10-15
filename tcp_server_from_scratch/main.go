package main

import (
	"log"
	"net"
	"time"
)

func makeConnectionInAction(conn net.Conn) {
	/* connection have 2 properties: 
		1.) Read the request came on conn 
	 	2.) Write the response after processing
	*/

	// read from the connection
	buffer := make([]byte, 1024);
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("request in processing")
	time.Sleep(10 * time.Second)

	// write the response on the connection (means returning the response to the user)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!!\r\n"))

	// close the connection
	conn.Close()
}

func listenMultipleConnections(listener net.Listener) {
	for {
		log.Println("waiting for client to connect")

		// accept the connections on the listener (blocking call)
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		log.Println("client connected")
		go makeConnectionInAction(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	listenMultipleConnections(listener)
}