package scheme

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const MessageId int32 = 0x00001

type Message struct {
	Content string
}

func (m Message) Serialize() ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, MessageId)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.LittleEndian, int32(len(m.Content)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.LittleEndian, []byte(m.Content))
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (m *Message) Deserialize(data []byte) error {
	reader := bytes.NewReader(data)

	var length int32
	err := binary.Read(reader, binary.LittleEndian, &length)
	if err != nil {
		fmt.Printf("Error reading length(%d): %v\n", length, err)
		return err
	}
	var contentBytes = make([]byte, length)
	err = binary.Read(reader, binary.LittleEndian, &contentBytes)
	if err != nil {
		fmt.Printf("Error reading content(%s): %v\n", string(contentBytes), err)
		return err
	}
	m.Content = string(contentBytes)
	return nil
}

func (m Message) ProtocolId() int32 {
	return MessageId
}
