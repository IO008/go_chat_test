package scheme

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestMessageSerialize(t *testing.T) {
	var testContent = "Hello, World!"
	var message Message = Message{Content: testContent}
	var err error
	var result []byte
	result, err = message.Serialize()
	if err != nil {
		t.Error("Error serializing message:", err)
		return
	}

	message = Message{}
	var id int32
	reader := bytes.NewReader(result)
	err = binary.Read(reader, binary.LittleEndian, &id)
	t.Logf("id: %x\n", id)
	if err != nil || id != message.ProtocolId() {
		t.Error("read id error:", id, err)
		return
	}
	result = result[4:]
	err = message.Deserialize(result)
	if err != nil {
		t.Error("Error deserializing message:", err)
		return
	}

	if message.Content != testContent {
		t.Error("Content mismatch after deserialization")
		return
	}
}
