package nex

type PacketV0 struct {
	Packet
}

func (p *PacketV0) Version() int {
	return int(p.version)
}
