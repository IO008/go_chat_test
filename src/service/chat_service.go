package service

import (
	"bufio"
	"fmt"
	"net"
)

func StartChatService(port int16) {
	fmt.Println("Chat service started")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Listening on port", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Print("Message Received:", string(message))

		conn.Write([]byte("Message Received\n"))
	}
}
