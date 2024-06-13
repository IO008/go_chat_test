package scheme

import "testing"

func TestPack(t *testing.T) {
	var err error
	protocol := new(Protocol)
	var content = "complete package"
	var data []byte
	data = []byte(content)
	var packed []byte
	packed, err = protocol.Pack(data)
	if err != nil {
		t.Error("Error packing:", err)
		return
	}

	var unpackBytes []byte
	unpackBytes, err = protocol.Unpack(packed)
	if err != nil {
		t.Error("Error unpacking:", err)
		return
	}

	if string(unpackBytes) != content {
		t.Errorf("unpck complte package error except %s actual %s", content, string(unpackBytes))
		return
	}
}
