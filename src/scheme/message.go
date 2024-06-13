package scheme

import (
	"bytes"
	"encoding/binary"
)

const id int32 = 0x00001

type Message struct {
	Content string
}

func (m Message) Serialize() ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, id)
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

func (m Message) ProtocolId() int32 {
	return id
}
