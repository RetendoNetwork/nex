package nex

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	socket       *net.UDPConn
	accessKey    string
	prudpVersion int
	KeySize      int
	FragmentSize int
}

func (server *Server) Listen(port int) {
	protocol := "udp"

	udpAddress, _ := net.ResolveUDPAddr(protocol, ":"+strconv.Itoa(port))
	socket, _ := net.ListenUDP(protocol, udpAddress)

	server.SetSocket(socket)

	fmt.Println("Server was listening on port", port)
}

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

func (server *Server) GetFragmentSize() int {
	return server.FragmentSize
}

func (server *Server) SetFragmentSize(FragmentSize int) {
	server.FragmentSize = FragmentSize
}

func (server *Server) GetAccessKey() string {
	return server.accessKey
}

func (server *Server) SetAccessKey(accessKey string) {
	server.accessKey = accessKey
}

func (server *Server) GetKeySize() int {
	return server.KeySize
}

func (server *Server) SetKeySize(KeySize int) {
	server.KeySize = KeySize
}

func NewServer() *Server {
	server := &Server{
		KeySize:      32,
		FragmentSize: 1300,
		prudpVersion: 1,
	}

	return server
}
