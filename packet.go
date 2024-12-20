package nex

type Packet struct {
	request RMCRequest
}

func (packet *Packet) RMCRequest() RMCRequest {
	return packet.request
}
