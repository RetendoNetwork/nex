package nex

import (
	"fmt"
)

var errorMask = 1 << 31

var ErrorNames = map[uint32]string{
	// == Core == //
	0x00010001: "Core::Unknown",
	0x00010002: "Core::NotImplemented",
	0x00010003: "Core::InvalidPointer",
	0x00010004: "Core::OperationAborted",
	0x00010005: "Core::Exception",
	0x00010006: "Core::AccessDenied",
	0x00010007: "Core::InvalidHandle",
	0x00010008: "Core::InvalidIndex",
	0x00010009: "Core::OutOfMemory",
	0x0001000A: "Core::InvalidArgument",
	0x0001000B: "Core::Timeout",
	0x0001000C: "Core::InitializationFailure",
}

func ErrorNameFromCode(errorCode uint32) string {
	if name, exists := ErrorNames[errorCode]; exists {
		return name
	}
	return fmt.Sprintf("Invalid Error Code: %d", errorCode)
}
