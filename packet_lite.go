package nex

type PacketLite struct {
	Packet
}

func (p *PacketLite) Version() int {
	return 2
}
