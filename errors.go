package nex

import "fmt"

type Errors struct {
	RCode   uint32
	Message string
}

// Return the error name from the error code
func (s Errors) Error() string {
	// See the code here https://github.com/PretendoNetwork/nex-go/blob/4885f237400082a8528743dc9499218c7d19c50c/errors.go
	rcode := s.RCode
	if rcode&errorMask != 0 {
		rcode = rcode & ^errorMask
	}
	return fmt.Sprintf("[%s] %s", ErrorNameFromCode(rcode), s.Message)
}

// Create a new error
func NewError(rcode uint32, message string) *Errors {
	if rcode&errorMask == 0 {
		rcode = rcode | errorMask
	}

	return &Errors{
		RCode:   rcode,
		Message: message,
	}
}
