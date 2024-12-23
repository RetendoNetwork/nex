package nex

type PacketV0 struct {
	Packet
	checksum uint8
}

func (pkt *PacketV0) SetChecksum(checksum uint8) {
	pkt.checksum = checksum
}

// Returns the checksum of the packet
func (pkt *PacketV0) GetChecksum() uint8 {
	return pkt.checksum
}

// Version returns the version of the packet
func (p *PacketV0) Version() int {
	return int(p.version)
}

func NewPacketV0(server *Server, data []byte) *PacketV0 {
	pkt := NewPacket(server, data)
	pktv0 := &PacketV0{Packet: *pkt}

	return pktv0
}
