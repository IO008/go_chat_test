package connection

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/IO008/go_chat_test/scheme"
)

var isStartNextPackage bool = true

func Unpackage(buffer []byte) error {

	if isStartNextPackage {
		isStartNextPackage = false
		var id int32
		reader := bytes.NewReader(buffer)
		err := binary.Read(reader, binary.LittleEndian, &id)
		if err != nil {
			fmt.Printf("Error reading id(% X): %v\n", buffer, err)
			return err
		}
		fmt.Printf("Deserializing id: %x\n", id)

		if id == scheme.MessageId {
			var message scheme.Message
			err := message.Deserialize(buffer[4:])
			if err != nil {
				fmt.Printf("Error deserializing message: %v\n", err)
				return err
			}
			fmt.Printf("Message content: %s\n", message.Content)
		}
	}

	fmt.Println("Deserializing buffer")
	return nil
}
