package scheme

type protocolSerialize interface {
	Serialize() ([]byte, error)

	ProtocolId() int32
}
