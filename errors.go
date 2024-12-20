package nex

import "fmt"

type Errors struct {
	RCode   uint32
	Message string
}

func (s Errors) Error() string {
	rcode := s.RCode
	if rcode&errorMask != 0 {
		rcode = rcode & ^errorMask
	}
	return fmt.Sprintf("[%s] %s", ErrorNameFromCode(rcode), s.Message)
}

func NewError(rcode uint32, message string) *Errors {
	if rcode&errorMask == 0 {
		rcode = rcode | errorMask
	}

	return &Errors{
		RCode:   rcode,
		Message: message,
	}
}
