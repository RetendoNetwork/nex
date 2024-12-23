package nex

type RMCRequest struct {
	protocol uint8
	callID   uint32
	method   uint32
}

func (rmc *RMCRequest) GetProtocol() uint8 {
	return rmc.protocol
}

func (rmc *RMCRequest) SetProtocol(protocol uint8) {
	rmc.protocol = protocol
}

func (rmc *RMCRequest) GetCallID() uint32 {
	return rmc.callID
}

func (rmc *RMCRequest) SetCallID(callID uint32) {
	rmc.callID = callID
}

func (rmc *RMCRequest) GetMethod() uint32 {
	return rmc.method
}

func (rmc *RMCRequest) SetMethod(method uint32) {
	rmc.method = method
}

func NewRMCRequest() RMCRequest {
	return RMCRequest{}
}

type RMCResponse struct {
	protocol  uint8
	customID  uint16
	callID    uint32
	method    uint32
	success   int
	data      []byte
	errorCode uint32
}

// Returns the protocol of the response
func (rmc *RMCResponse) CustomID() uint16 {
	return rmc.customID
}

// Returns the protocol of the response
func (rmc *RMCResponse) SetCustomID(customID uint16) {
	rmc.customID = customID
}

// Returns the protocol of the response
func (rmc *RMCResponse) SetSuccess(method uint32, data []byte) {
	rmc.success = 1
	rmc.method = method
	rmc.data = data
}

// Sets the error code for the response
func (rmc *RMCResponse) SetError(RCode uint32) {
	if RCode&errorMask == 0 {
		RCode = uint32(RCode | errorMask)
	}
	rmc.success = 0
	rmc.errorCode = RCode
}

// Creates a new RMCResponse object
func NewRMCResponse(protocol uint8, callID uint32) RMCResponse {
	rmc := RMCResponse{
		protocol: protocol,
		callID:   callID,
	}

	return rmc
}
