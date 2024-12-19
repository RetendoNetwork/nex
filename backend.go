package nex

// TODO: Continue this.

type LoginResult struct {
	Pid           int64
	Ticket        []byte
	SourceKey     []byte
	SecureStation *SecureStation
}

type SecureStation struct {
	PID     int64
	CID     int64
	Address string
	Port    int
	Sid     int64
}

type BackEndClient struct {
	AuthHost string
	AuthPort int
}
