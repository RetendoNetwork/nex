package nex

type PacketLite struct {
	Packet
}

// Version returns the version of the packet
func (p *PacketLite) Version() int {
	return 2
}

func NewPacketLite(server *Server, data []byte) *PacketLite {
	pkt := NewPacket(server, data)
	pktlite := &PacketLite{Packet: *pkt}

	return pktlite
}
