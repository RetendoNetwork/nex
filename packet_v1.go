package nex

type PacketV1 struct {
	Packet
}

// Version returns the version of the packet
func (p *PacketV1) Version() int {
	return 1
}

func NewPacketV1(server *Server, data []byte) *PacketV1 {
	pkt := NewPacket(server, data)
	pktv1 := &PacketV1{Packet: *pkt}

	return pktv1
}
