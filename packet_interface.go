package nex

type PacketInterface interface {
	Payload() []byte
	SetPayload(payload []byte)
	getFragmentID() uint8
	setFragmentID(fragmentID uint8)
}
