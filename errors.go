package nex

import "fmt"

type Errors struct {
	Code    uint32
	Message string
}

func (s Errors) Error() string {
	var code = s.Code

	if int(code)&errorMask != 0 {
		code = code & uint32(errorMask)
	}

	return fmt.Sprintf("[%s] %s", ErrorNameFromCode(code), s.Message)
}
