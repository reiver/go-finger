package finger

import (
	"fmt"
	"strconv"
)

const (
	defaultPort = 79
)

var (
	defaultPortString = strconv.FormatUint(defaultPort, 10)
)

// Port represents a finger-protocol port.
// A “port” in this context is is also known as a “TCP port”.
//
// For example, in this finger-protocol request:
//
//	"dariush@example.com:1971@something.social:1234\r\n"
//
// “1971” and “1234” are ports.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Port with code like
// the following:
//
//	var port finger.Port
//	
//	// ...
//	
//	fmt.Printf("finger port: %#v", port)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Port, do something like the following:
//
//	var port finger.Port
//	
//	// ...
//	
//	value, something := port.Unwrap()
//	if something {
//		// a value was set for the finger.Port
//	} else {
//		// no value was set for the finger.Port
//	}
//
// Note that this is unwrapping the finger.Port optional-type.
type Port struct {
	value uint16
	something bool
}

// ParsePort parses a string for a (numeric) port.
func ParsePort(s string) (Port, error) {
	if "" == s {
		return EmptyPort(), nil
	}

	const base int = 10
	const bitsize int = 16

	var u16 uint16
	{
		u64, err := strconv.ParseUint(s, base, bitsize)
		if nil != err {
			return EmptyPort(), fmt.Errorf("problem parsing finger-protocol port: %w", err)
		}

		u16 = uint16(u64)
	}

	return CreatePort(u16), nil
}

// EmptyPort is used to create a finger.Port with nothing in it.
func EmptyPort() Port {
	return Port{}
}

// CreatePort is used to create a finger.Port with something in it.
func CreatePort(value uint16) Port {
	return Port{
		value:value,
		something:true,
	}
}

// DefaultPort is used to create a finger.Port with the value of 79 in it.
// 79 is the default finger-protocol port.
func DefaultPort() Port {
	return CreatePort(defaultPort)
}

// AlternativePort is used to create a finger.Port with the value of 1971 in it.
// 1971 is an alternative port that can be used by the finger-protocol.
func AlternativePort() Port {
	return CreatePort(1971)
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var port finger.Port = finger.CreatePort(79)
//	
//	// ...
//	
//	fmt.Printf("port = %#v", port)
//
//	// Output:
//	// port = finger.CreatePort(79)
//
// Also, for example:
//
//	var port finger.Port = finger.EmptyPort()
//	
//	// ...
//	
//	fmt.Printf("port = %#v", port)
//
//	// Output:
//	// port = finger.EmptyPort()
func (receiver Port) GoString() string {
	if !receiver.something {
		return "finger.EmptyPort()"
	}

	return fmt.Sprintf("finger.CreatePort(%d)", receiver.value)
}

// Resolve resolves a finger.Port.
//
// If the finger.Port has a value, then Resolve returns that.
//
// Else if the finger.Port is empty, then Resolve returns the default finger TCP-port value,
// which is 79.
func (receiver Port) Resolve() uint16 {
	if !receiver.something {
		return defaultPort
	}

	return receiver.value
}

// Set sets the value of a finger.Port.
//
// Set mainly exists so that finger.Port can be as a flag.Value, and thus be used with functions
// such as flag.Var().
func (receiver *Port) Set(value string) error {
	if nil == receiver {
		return nil
	}

	port, err := ParsePort(value)
	if nil != err {
		return err
	}
	*receiver = port
	return nil
}

// String returns the value of a finger.Port.
//
// Note that if finger.Port is empty, then it returns the default finger TCP-port,
// which is 79.
//
// With String you cannot tell the difference between a finger.Port with a value of 79
// and an empty finger.Port.
func (receiver Port) String() string {
	if !receiver.something {
		return defaultPortString
	}

	return strconv.FormatUint(uint64(receiver.value), 10)
}

// Unwrap is used to unwrap a finger.Port.
//
//	var port finger.Port
//	
//	// ...
//	
//	value, something := port.Unwrap()
//
// If finger.Port is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Port is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Port) Unwrap() (uint16, bool) {
	return receiver.value, receiver.something
}

