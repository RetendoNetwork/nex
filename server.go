package nex

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	socket              *net.UDPConn
	accessKey           string
	prudpVersion        int
	nexVersion          int
	keySize             int
	fragmentSize        int
	signatureKey        int
	genericEventHandles map[string][]func(PacketInterface)
	prudpV0EventHandles map[string][]func(*PacketV0)
	prudpV1EventHandles map[string][]func(*PacketV1)
	connIncrementer     *Incrementer[int]
}

func (srv *Server) Listen(port int) {
	srv.ListenUDP(port)
}

// ListenUDP starts the UDP server on the specified port
func (srv *Server) ListenUDP(port int) {
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

	srv.SetSocket(socket)

	<-quit
}

// GetSocket returns the current UDP socket
func (srv *Server) GetSocket() *net.UDPConn {
	return srv.socket
}

// SetSocket sets the UDP socket
func (srv *Server) SetSocket(socket *net.UDPConn) {
	srv.socket = socket
}

// GetPRUDPVersion returns the PRUDP version
func (srv *Server) GetPRUDPVersion() int {
	return srv.prudpVersion
}

// SetPRUDPVersion sets the PRUDP version
func (srv *Server) SetPRUDPVersion(prudpVersion int) {
	srv.prudpVersion = prudpVersion
}

// GetNexVersion returns the Nex version
func (srv *Server) GetNexVersion() int {
	return srv.nexVersion
}

// SetNexVersion sets the Nex version
func (srv *Server) SetNexVersion(nexVersion int) {
	srv.nexVersion = nexVersion
}

// GetSignatureKey returns the signature key
func (srv *Server) GetSignatureKey() int {
	return srv.signatureKey
}

// SetSignatureKey sets the signature key
func (srv *Server) SetSignatureKey(SignatureKey int) {
	srv.signatureKey = SignatureKey
}

// GetFragmentSize returns the fragment size
func (srv *Server) GetFragmentSize() int {
	return srv.fragmentSize
}

// SetFragmentSize sets the fragment size
func (srv *Server) SetFragmentSize(FragmentSize int) {
	srv.fragmentSize = FragmentSize
}

// GetAccessKey returns the access key
func (srv *Server) GetAccessKey() string {
	return srv.accessKey
}

// SetAccessKey sets the access key
func (srv *Server) SetAccessKey(accessKey string) {
	srv.accessKey = accessKey
}

// GetKeySize returns the key size
func (srv *Server) GetKeySize() int {
	return srv.keySize
}

// SetKeySize sets the key size
func (srv *Server) SetKeySize(KeySize int) {
	srv.keySize = KeySize
}

// ConnIncrementer returns the connection incrementer
func (srv *Server) ConnIncrementer() *Incrementer[int] {
	return srv.connIncrementer
}

func (srv *Server) OnData(event string, handler interface{}) {
	switch h := handler.(type) {
	case func(PacketInterface):
		srv.genericEventHandles[event] = append(srv.genericEventHandles[event], h)
	case func(*PacketV0):
		srv.prudpV0EventHandles[event] = append(srv.prudpV0EventHandles[event], h)
	case func(*PacketV1):
		srv.prudpV1EventHandles[event] = append(srv.prudpV1EventHandles[event], h)
	}
}

type ServerPacketInterface interface {
	Payload() []byte
	SetPayload(payload []byte)
	getFragmentID() uint8
	setFragmentID(fragmentID uint8)
}

func (srv *Server) Send(pkt PacketInterface) {
	if srvPkt, valid := pkt.(ServerPacketInterface); valid {
		payload := srvPkt.Payload()
		chunkSize := srv.fragmentSize
		totalChunks := len(payload) / chunkSize

		fragID := uint8(1)
		for offset := 0; offset <= totalChunks; offset++ {
			remaining := len(payload)
			if remaining <= chunkSize {
				srvPkt.SetPayload(payload)
				srvPkt.setFragmentID(0)
			} else {
				srvPkt.SetPayload(payload[:chunkSize])
				srvPkt.setFragmentID(fragID)

				payload = payload[chunkSize:]
				fragID++
			}

			srv.SendPacket(srvPkt)
		}
	}
}

func (srv *Server) SendPacket(pkt PacketInterface) {
	// TODO: Do SendPacket func.
}

func (srv *Server) SendRaw(conn *net.UDPAddr, data []byte) error {
	if srv == nil || srv.GetSocket() == nil {
		return fmt.Errorf("server or socket is nil")
	}
	_, err := srv.GetSocket().WriteToUDP(data, conn)
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}
	return nil
}

func NewServer() *Server {
	srv := &Server{
		genericEventHandles: make(map[string][]func(PacketInterface)),
		prudpV0EventHandles: make(map[string][]func(*PacketV0)),
		prudpV1EventHandles: make(map[string][]func(*PacketV1)),
		prudpVersion:        1,
		keySize:             32,
		fragmentSize:        1300,
		signatureKey:        1,
		connIncrementer:     NewIncrementer(10),
	}

	return srv
}
