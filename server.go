package nex

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	socket                      *net.UDPConn
	accessKey                   string
	prudpVersion                int
	nexVersion                  *nexVersion
	pid                         uint32
	password                    []byte
	ticketVersion               int
	keySize                     int
	fragmentSize                int
	signatureKey                int
	version                     *nexVersion
	datastoreProtocolVersion    *nexVersion
	matchMakingProtocolVersion  *nexVersion
	rankingProtocolVersion      *nexVersion
	ranking2ProtocolVersion     *nexVersion
	messagingProtocolVersion    *nexVersion
	utilityProtocolVersion      *nexVersion
	natTraversalProtocolVersion *nexVersion
	genericEventHandles         map[string][]func(PacketInterface)
	prudpV0EventHandles         map[string][]func(*PacketV0)
	prudpV1EventHandles         map[string][]func(*PacketV1)
	prudpLiteEventHandles       map[string][]func(*PacketLite)
	connIncrementer             *Incrementer[int]
	connected                   bool
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

// NEXVersion returns the server NEX version
func (srv *Server) NEXVersion() *nexVersion {
	return srv.nexVersion
}

// SetDefaultNEXVersion sets the default NEX protocol versions
func (srv *Server) SetDefaultNEXVersion(n *nexVersion) {
	srv.nexVersion = n
	// See the code here https://github.com/PretendoNetwork/nex-go/blob/4885f237400082a8528743dc9499218c7d19c50c/prudp_server.go#L891
	srv.datastoreProtocolVersion = n.Copy()
	srv.matchMakingProtocolVersion = n.Copy()
	srv.rankingProtocolVersion = n.Copy()
	srv.ranking2ProtocolVersion = n.Copy()
	srv.messagingProtocolVersion = n.Copy()
	srv.utilityProtocolVersion = n.Copy()
	srv.natTraversalProtocolVersion = n.Copy()
}

// DataStoreProtocolVersion returns the servers DataStore protocol version
func (srv *Server) DataStoreProtocolVersion() *nexVersion {
	return srv.datastoreProtocolVersion
}

// SetDataStoreProtocolVersion sets the servers DataStore protocol version
func (srv *Server) SetDataStoreProtocolVersion(version *nexVersion) {
	srv.datastoreProtocolVersion = version
}

// MatchMakingProtocolVersion returns the servers MatchMaking protocol version
func (srv *Server) MatchMakingProtocolVersion() *nexVersion {
	return srv.matchMakingProtocolVersion
}

// SetMatchMakingProtocolVersion sets the servers MatchMaking protocol version
func (srv *Server) SetMatchMakingProtocolVersion(version *nexVersion) {
	srv.matchMakingProtocolVersion = version
}

// RankingProtocolVersion returns the servers Ranking protocol version
func (srv *Server) RankingProtocolVersion() *nexVersion {
	return srv.rankingProtocolVersion
}

// SetRankingProtocolVersion sets the servers Ranking protocol version
func (srv *Server) SetRankingProtocolVersion(version *nexVersion) {
	srv.rankingProtocolVersion = version
}

// Ranking2ProtocolVersion returns the servers Ranking2 protocol version
func (srv *Server) Ranking2ProtocolVersion() *nexVersion {
	return srv.ranking2ProtocolVersion
}

// SetRanking2ProtocolVersion sets the servers Ranking2 protocol version
func (srv *Server) SetRanking2ProtocolVersion(version *nexVersion) {
	srv.ranking2ProtocolVersion = version
}

// MessagingProtocolVersion returns the servers Messaging protocol version
func (srv *Server) MessagingProtocolVersion() *nexVersion {
	return srv.messagingProtocolVersion
}

// SetMessagingProtocolVersion sets the servers Messaging protocol version
func (srv *Server) SetMessagingProtocolVersion(version *nexVersion) {
	srv.messagingProtocolVersion = version
}

// UtilityProtocolVersion returns the servers Utility protocol version
func (srv *Server) UtilityProtocolVersion() *nexVersion {
	return srv.utilityProtocolVersion
}

// SetUtilityProtocolVersion sets the servers Utility protocol version
func (srv *Server) SetUtilityProtocolVersion(version *nexVersion) {
	srv.utilityProtocolVersion = version
}

// SetNATTraversalProtocolVersion sets the servers NAT Traversal protocol version
func (srv *Server) SetNATTraversalProtocolVersion(version *nexVersion) {
	srv.natTraversalProtocolVersion = version
}

// NATTraversalProtocolVersion returns the servers NAT Traversal protocol version
func (srv *Server) NATTraversalProtocolVersion() *nexVersion {
	return srv.natTraversalProtocolVersion
}

// PID returns the PID
func (srv *Server) GetPID() uint32 {
	return srv.pid
}

// SetPID sets the PID
func (srv *Server) SetPID(pid uint32) {
	srv.pid = pid
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

func (srv *Server) GetPassword() []byte {
	return srv.password
}

func (srv *Server) SetPassword(password []byte) {
	srv.password = password
}

func (srv *Server) SetTicketVersion(ticketVersion int) {
	srv.ticketVersion = ticketVersion
}

// GetKeySize returns the key size
func (srv *Server) GetKeySize() int {
	return srv.keySize
}

// SetKeySize sets the key size
func (srv *Server) SetKeySize(KeySize int) {
	srv.keySize = KeySize
}

func (srv *Server) SetConnected(connected bool) {
	srv.connected = connected
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
	case func(*PacketLite):
		srv.prudpLiteEventHandles[event] = append(srv.prudpLiteEventHandles[event], h)
	}
}

func (srv *Server) Send(pkt PacketInterface) {
	switch srvPkt := pkt.(type) {
	default:
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
	// TODO: Add SendPacket.
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

// Create a new server
func NewServer() *Server {
	srv := &Server{
		genericEventHandles:   make(map[string][]func(PacketInterface)),
		prudpV0EventHandles:   make(map[string][]func(*PacketV0)),
		prudpV1EventHandles:   make(map[string][]func(*PacketV1)),
		prudpLiteEventHandles: make(map[string][]func(*PacketLite)),
		prudpVersion:          1,
		keySize:               32,
		fragmentSize:          1300,
		signatureKey:          1,
		connIncrementer:       NewIncrementer(10),
	}

	srv.SetDefaultNEXVersion(NewNexVersion(0, 0, 0)) // Set the default version of NEX.

	return srv
}
