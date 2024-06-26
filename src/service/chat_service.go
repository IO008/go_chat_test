package service

import (
	"bufio"
	"fmt"
	"net"

	"github.com/IO008/go_chat_test/scheme"
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

	protocal := new(scheme.Protocol)
	reader := bufio.NewReader(conn)
	var buffer = make([]byte, 1024*4)
	for {
		length, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		protocal.Unpack(buffer[:length])

		// TODO handle custom protocol
		fmt.Print("Message Received:", string(buffer[:length]))

		// TODO ACK message
		conn.Write([]byte("Message Received\n"))
	}
}
