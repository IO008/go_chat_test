package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	description string
}

const (
	login       = "login"
	register    = "register"
	friendList  = "friend list"
	sendMessage = "send message"
	sendMedia   = "send media"
	exit        = "exit"
)

var commands = []Command{
	{description: login},
	{description: register},
	{description: friendList},
	{description: sendMessage},
	{description: sendMedia},
	{description: exit},
}

func (c Command) ShowAllCommand() {
	for key, value := range commands {
		fmt.Println(key, " ", value.description)
	}
}

func (c Command) ReadCommand(commandChan chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		command := scanner.Text()
		commandChan <- command
		if isExitCommand(command) {
			close(commandChan)
			return
		}
	}
}

func (c Command) HandleCommand(commandChan <-chan string) {
	for cmd := range commandChan {
		num, err := getCommandNumber(cmd)
		if err != nil {
			fmt.Println("handle command error", err)
			continue
		}
		switch commands[num].description {
		case login:
			fmt.Println("login")
		case register:
			fmt.Println("register")
		case friendList:
			fmt.Println("friend list")
		case sendMessage:
			fmt.Println("send message")
		case sendMedia:
			fmt.Println("send media")
		case exit:
			fmt.Println("exit")
			return
		}
	}
}

func getCommandNumber(command string) (int, error) {
	num, err := strconv.Atoi(command)
	if err != nil {
		return -1, err
	}
	if num < 0 || num >= len(commands) {
		return -1, fmt.Errorf("index out of range %d", num)
	}
	return num, nil
}

func isExitCommand(command string) bool {
	num, err := getCommandNumber(command)
	if err != nil {
		return false
	}
	return commands[num].description == exit
}
