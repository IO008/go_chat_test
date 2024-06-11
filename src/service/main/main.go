package main

import (
	"fmt"

	"github.com/IO008/go_chat_test/service"
)

const listenPort = 8080

func main() {
	fmt.Println("Start service")

	service.StartChatService(listenPort)

	fmt.Println("End main")
}
