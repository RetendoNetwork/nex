package nex

type PacketV1 struct {
	Packet
}

func (p *PacketV1) Version() int {
	return 1
}
