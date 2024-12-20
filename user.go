package nex

type User struct {
	pid      int
	username string
	password string
}

func NewUser(pid int, username string, password string) *User {
	return &User{
		pid:      pid,
		username: username,
		password: password,
	}
}
