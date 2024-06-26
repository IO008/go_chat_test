package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IO008/go_chat_test/client/cli"
)

func main() {
	fmt.Println("Start up client")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	commandChan := make(chan string)

	command := cli.Command{}
	command.ShowAllCommand()

	go command.ReadCommand(commandChan)
	command.HandleCommand(commandChan)

	//client.StartClient("localhost", 8080)
	sig := <-sigChan
	fmt.Printf("exit for receive sig %s \n", sig)
}
