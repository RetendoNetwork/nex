package nex

type ConnState uint8

const (
	NotConnected ConnState = iota
	Connecting
	Connected
	Disconnecting
	Faulty
)
