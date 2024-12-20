package nex

type PacketInterface interface {
	Payload() []byte
	SetPayload(payload []byte)
}
