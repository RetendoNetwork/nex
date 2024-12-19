package nex

import (
	"fmt"
	"net"
)

type Server struct {
	socket    *net.UDPConn
	accessKey string
}

func (server *Server) Listen(address string) {

	protocol := "udp"

	udpAddress, _ := net.ResolveUDPAddr(protocol, address)
	socket, _ := net.ListenUDP(protocol, udpAddress)

	server.SetSocket(socket)

	fmt.Println("Server was listening on ", udpAddress)
}

func (server *Server) GetSocket() *net.UDPConn {
	return server.socket
}

func (server *Server) SetSocket(socket *net.UDPConn) {
	server.socket = socket
}

func (server *Server) GetAccessKey() string {
	return server.accessKey
}

func (server *Server) SetAccessKey(accessKey string) {
	server.accessKey = accessKey
}

func NewServer() *Server {
	server := &Server{}

	return server
}
