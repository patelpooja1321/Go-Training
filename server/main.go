package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server is listening")
	listener, err := net.Listen("tcp", "localhost:3000")

	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("New Connection")

		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
		datasize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:datasize]
		fmt.Println("received message:", string(data))

		_, err = conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message Sent:", string(data))

	}

}
