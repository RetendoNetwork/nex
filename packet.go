package nex

type Packet struct {
	request  RMCRequest
	response RMCResponse
}

func (pkt *Packet) RMCRequest() RMCRequest {
	return pkt.request
}

func (pkt *Packet) RMCResponse() RMCResponse {
	return pkt.response
}
