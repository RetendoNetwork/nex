package nex

type RMCRequest struct {
	protocol int
	callID   int
	method   int
}

func (rmc *RMCRequest) GetProtocol() int {
	return rmc.protocol
}

func (rmc *RMCRequest) SetProtocol(protocol int) {
	rmc.protocol = protocol
}

func (rmc *RMCRequest) GetCallID() int {
	return rmc.callID
}

func (rmc *RMCRequest) SetCallID(callID int) {
	rmc.callID = callID
}

func (rmc *RMCRequest) GetMethod() int {
	return rmc.method
}

func (rmc *RMCRequest) SetMethod(method int) {
	rmc.method = method
}

func NewRMCRequest() RMCRequest {
	return RMCRequest{}
}
