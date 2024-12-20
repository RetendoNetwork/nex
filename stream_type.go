package nex

type StreamType uint8

func (st StreamType) EnumIndex() uint8 {
	return uint8(st)
}

const (
	DO StreamType = iota
	RV
	OldRVSec
	SBMGMT
	NAT
	SessionDiscovery
	NATEcho
	Routing
	Game
	TVSecure
	Relay
)
