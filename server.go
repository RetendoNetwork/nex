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
}

func (srv *Server) Listen(port int) {
	srv.ListenUDP(port)
}

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

func (srv *Server) GetSocket() *net.UDPConn {
	return srv.socket
}

func (srv *Server) SetSocket(socket *net.UDPConn) {
	srv.socket = socket
}

func (srv *Server) GetPrudpVersion() int {
	return srv.prudpVersion
}

func (srv *Server) SetPrudpVersion(prudpVersion int) {
	srv.prudpVersion = prudpVersion
}

func (srv *Server) GetNexVersion() int {
	return srv.nexVersion
}

func (srv *Server) SetNexVersion(nexVersion int) {
	srv.nexVersion = nexVersion
}

func (srv *Server) GetSignatureKey() int {
	return srv.signatureKey
}

func (srv *Server) SetSignatureKey(SignatureKey int) {
	srv.signatureKey = SignatureKey
}

func (srv *Server) GetFragmentSize() int {
	return srv.fragmentSize
}

func (srv *Server) SetFragmentSize(FragmentSize int) {
	srv.fragmentSize = FragmentSize
}

func (srv *Server) GetAccessKey() string {
	return srv.accessKey
}

func (srv *Server) SetAccessKey(accessKey string) {
	srv.accessKey = accessKey
}

func (srv *Server) GetKeySize() int {
	return srv.keySize
}

func (srv *Server) SetKeySize(KeySize int) {
	srv.keySize = KeySize
}

func (srv *Server) OnData(event string, handler interface{}) {
	switch handler.(type) {
	case func(PacketInterface):
		srv.genericEventHandles[event] = append(srv.genericEventHandles[event], handler.(func(PacketInterface)))
	case func(*PacketV0):
		srv.prudpV0EventHandles[event] = append(srv.prudpV0EventHandles[event], handler.(func(*PacketV0)))
	case func(*PacketV1):
		srv.prudpV1EventHandles[event] = append(srv.prudpV1EventHandles[event], handler.(func(*PacketV1)))
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
	}

	return srv
}
