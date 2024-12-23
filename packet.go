package nex

type Packet struct {
	server   *Server
	version  uint8
	data     []byte
	payload  []byte
	request  RMCRequest
	response RMCResponse
}

func (pkt *Packet) RMCRequest() RMCRequest {
	return pkt.request
}

func (pkt *Packet) RMCResponse() RMCResponse {
	return pkt.response
}

func NewPacket(server *Server, data []byte) *Packet {
	pkt := &Packet{
		server:  server,
		data:    data,
		payload: []byte{},
	}

	return pkt
}
