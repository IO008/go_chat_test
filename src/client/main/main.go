package main

import (
	"fmt"

	"github.com/IO008/go_chat_test/client"
)

func main() {
	fmt.Println("Start client")

	client.StartClient("localhost", 8080)

	fmt.Println("End client")
}
