package nex

import "fmt"

type Errors struct {
	RCode   uint32
	Message string
}

func (s Errors) Error() string {
	rcode := s.RCode
	if int(rcode)&errorMask != 0 {
		rcode = rcode & ^uint32(errorMask)
	}
	return fmt.Sprintf("[%s] %s", ErrorNameFromCode(rcode), s.Message)
}

func NewError(rcode uint32, message string) *Errors {
	if int(rcode)&errorMask == 0 {
		rcode = uint32(int(rcode) | errorMask)
	}

	return &Errors{
		RCode:   rcode,
		Message: message,
	}
}
