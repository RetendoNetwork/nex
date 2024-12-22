package nex

type User struct {
	pid      uint64
	username string
	password string
}

func NewUser(pid uint64, username string, password string) *User {
	return &User{
		pid:      pid,
		username: username,
		password: password,
	}
}
