package nex

type StreamType uint8

func (st StreamType) EnumIndex() uint8 {
	return uint8(st)
}

const (
	DO StreamType = iota + 1
	RV
	OldRVSec
	SBMGMT
	NAT
	SessionDiscovery
	NATEcho
	Routing
	Game
	RVSecure
	Relay
)
