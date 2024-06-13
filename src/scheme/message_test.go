package scheme

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestMessageSerialize(t *testing.T) {
	var testContent = "Hello, World!"
	message := Message{Content: testContent}
	result, err := message.Serialize()
	if err != nil {
		t.Error("Error serializing message:", err)
		return
	}

	var buffer bytes.Buffer
	buffer.Write(result)
	var id int32
	binary.Read(&buffer, binary.LittleEndian, &id)
	if id != message.ProtocolId() {
		t.Error("Invalid id:", id)
		return
	}
	var length int32
	binary.Read(&buffer, binary.LittleEndian, &length)
	if length != int32(len(testContent)) {
		t.Error("Invalid length:", length)
		return
	}
	var contentBytes = make([]byte, length)
	binary.Read(&buffer, binary.LittleEndian, &contentBytes)
	content := string(contentBytes)
	if content != testContent {
		t.Error("Invalid content:", content)
		return
	}
	t.Logf("Test result id: %d, length: %d, content: %s", id, length, content)
}
