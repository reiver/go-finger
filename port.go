package finger

import (
	"fmt"
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

// NoPort is used to create a finger.Port with nothing in it.
func NoPort() Port {
	return Port{}
}

// SomePort is used to create a finger.Port with something in it.
func SomePort(value uint16) Port {
	return Port{
		value:value,
		something:true,
	}
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

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var port finger.Port = finger.SomePort(79)
//	
//	// ...
//	
//	fmt.Printf("port = %#v", port)
//
//	// Output:
//	// port = finger.SomePort(79)
//
// Also, for example:
//
//	var port finger.Port = finger.NoPort()
//	
//	// ...
//	
//	fmt.Printf("port = %#v", port)
//
//	// Output:
//	// port = finger.NoPort()
func (receiver Port) GoString() string {
	if !receiver.something {
		return "finger.NoPort()"
	}

	return fmt.Sprintf("finger.SomePort(%#v)", receiver.value)
}
