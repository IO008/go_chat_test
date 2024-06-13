package scheme

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type protocolSerialize interface {
	Serialize() ([]byte, error)

	Deserialize([]byte) error

	ProtocolId() int32
}

type Protocol struct {
	buffer     bytes.Buffer
	dataLength uint32
}

func (Protocol) Pack(data []byte) ([]byte, error) {
	l := len(data)
	buffer := bytes.Buffer{}
	var err error
	err = binary.Write(&buffer, binary.LittleEndian, uint32(l))
	if err != nil {
		fmt.Println("Error writing data length:", err)
		return nil, err
	}
	err = binary.Write(&buffer, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (p *Protocol) Unpack(buffer []byte) ([]byte, error) {
	binary.Write(&p.buffer, binary.LittleEndian, buffer)

	var err error
	if p.dataLength == 0 && p.buffer.Len() >= 4 { // read head
		err = binary.Read(&p.buffer, binary.LittleEndian, &p.dataLength)
		if err != nil {
			return nil, err
		}
	}

	if p.dataLength > 0 && uint32(p.buffer.Len()) >= p.dataLength { // read data
		result := make([]byte, p.dataLength)
		binary.Read(&p.buffer, binary.LittleEndian, &result)
		p.dataLength = 0
		return result, nil
	}
	return nil, nil
}
