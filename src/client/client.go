package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/IO008/go_chat_test/scheme"
)

func StartClient(ip string, port int16) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	protocol := scheme.Protocol{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		fmt.Fprint(conn, text)
		protocol.Pack([]byte(text))
		/*message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message reply: " + message)*/
	}
}
