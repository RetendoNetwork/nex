package nex

import (
	"net"
	"strconv"
)

type Server struct {
	socket              *net.UDPConn
	accessKey           string
	prudpVersion        int
	nexVersion          int
	keySize             int
	fragmentSize        int16
	signatureKey        int
	genericEventHandles map[string][]func(PacketInterface)
	prudpV0EventHandles map[string][]func(*PacketV0)
	prudpV1EventHandles map[string][]func(*PacketV1)
}

func (server *Server) Listen(port int) {
	server.ListenUDP(port)
}

func (server *Server) ListenUDP(port int) {
	protocol := "udp"

	udpAddress, err := net.ResolveUDPAddr(protocol, ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	socket, err := net.ListenUDP(protocol, udpAddress)
	if err != nil {
		panic(err)
	}

	quit := make(chan struct{})

	server.SetSocket(socket)

	<-quit
}

// TODO: WebSocket Listen Port

func (server *Server) GetSocket() *net.UDPConn {
	return server.socket
}

func (server *Server) SetSocket(socket *net.UDPConn) {
	server.socket = socket
}

func (server *Server) GetPrudpVersion() int {
	return server.prudpVersion
}

func (server *Server) SetPrudpVersion(prudpVersion int) {
	server.prudpVersion = prudpVersion
}

func (server *Server) GetNexVersion() int {
	return server.nexVersion
}

func (server *Server) SetNexVersion(nexVersion int) {
	server.nexVersion = nexVersion
}

func (server *Server) GetSignatureKey() int {
	return server.signatureKey
}

func (server *Server) SetSignatureKey(SignatureKey int) {
	server.signatureKey = SignatureKey
}

func (server *Server) GetFragmentSize() int16 {
	return server.fragmentSize
}

func (server *Server) SetFragmentSize(FragmentSize int16) {
	server.fragmentSize = FragmentSize
}

func (server *Server) GetAccessKey() string {
	return server.accessKey
}

func (server *Server) SetAccessKey(accessKey string) {
	server.accessKey = accessKey
}

func (server *Server) GetKeySize() int {
	return server.keySize
}

func (server *Server) SetKeySize(KeySize int) {
	server.keySize = KeySize
}

func (server *Server) OnData(event string, handler interface{}) {
	switch handler.(type) {
	case func(PacketInterface):
		server.genericEventHandles[event] = append(server.genericEventHandles[event], handler.(func(PacketInterface)))
	case func(*PacketV0):
		server.prudpV0EventHandles[event] = append(server.prudpV0EventHandles[event], handler.(func(*PacketV0)))
	case func(*PacketV1):
		server.prudpV1EventHandles[event] = append(server.prudpV1EventHandles[event], handler.(func(*PacketV1)))
	}
}

func NewServer() *Server {
	server := &Server{
		prudpVersion: 1,
		keySize:      32,
		fragmentSize: 1300,
		signatureKey: 1,
	}

	return server
}
