package nex

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type NATCheck struct {
	conn         *net.UDPConn
	externalIP   string
	externalPort int
}

func NewNATCheck() (*NATCheck, error) {
	localAddr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	}

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create UDP socket: %v", err)
	}

	client := &NATCheck{
		conn: conn,
	}

	return client, nil
}

func (s *NATCheck) detectExternalAddress(stunAddr string, stunPort int) (string, int, error) {
	message := make([]byte, 16)
	binary.BigEndian.PutUint32(message[0:4], 1) // Type = 1
	binary.BigEndian.PutUint32(message[4:8], 0)
	binary.BigEndian.PutUint32(message[8:12], 0)
	binary.BigEndian.PutUint32(message[12:16], 0)

	serverAddr := fmt.Sprintf("%s:%d", stunAddr, stunPort)
	remoteAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to resolve STUN server: %v", err)
	}

	_, err = s.conn.WriteToUDP(message, remoteAddr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to send message to STUN server: %v", err)
	}

	buf := make([]byte, 16)
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, _, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return "", 0, fmt.Errorf("failed to read from STUN server: %v", err)
	}

	if n == 16 {
		typeField := binary.BigEndian.Uint32(buf[0:4])
		port := binary.BigEndian.Uint32(buf[4:8])
		host := binary.BigEndian.Uint32(buf[8:12])

		if typeField == 1 {
			externalIP := fmt.Sprintf("%d.%d.%d.%d", byte(host>>24), byte(host>>16), byte(host>>8), byte(host))
			return externalIP, int(port), nil
		}
	}

	return "", 0, fmt.Errorf("invalid response from STUN server")
}
