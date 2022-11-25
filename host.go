package finger

import (
	"fmt"
)

const (
	defaultHost = "127.0.0.1"
)

// Host represents a finger-protocol host.
//
// For example, in this finger-protocol request:
//
//	"dariush@example.com@something.social\r\n"
//
// “example.com” and “something.social” are hosts.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Host with code like
// the following:
//
//	var host finger.Host
//	
//	// ...
//	
//	fmt.Printf("finger host: %#v", host)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Host, do something like the following:
//
//	var host finger.Host
//	
//	// ...
//	
//	value, something := host.Unwrap()
//	if something {
//		// a value was set for the finger.Host
//	} else {
//		// no value was set for the finger.Host
//	}
//
// Note that this is unwrapping the finger.Host optional-type.
type Host struct {
	value string
	something bool
}

// NoHost is used to create a finger.Host with nothing in it.
func NoHost() Host {
	return Host{}
}

// SomeHost is used to create a finger.Host with something in it.
func SomeHost(value string) Host {
	return Host{
		value:value,
		something:true,
	}
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var host finger.Host = finger.SomeHost("dariush")
//	
//	// ...
//	
//	fmt.Printf("host = %#v", host)
//
//	// Output:
//	// host = finger.SomeHost("dariush")
//
// Also, for example:
//
//	var host finger.Host = finger.NoHost()
//	
//	// ...
//	
//	fmt.Printf("host = %#v", host)
//
//	// Output:
//	// host = finger.NoHost()
func (receiver Host) GoString() string {
	if !receiver.something {
		return "finger.NoHost()"
	}

	return fmt.Sprintf("finger.SomeHost(%#v)", receiver.value)
}

func (receiver Host) Resolve() string {
	if !receiver.something {
		return defaultHost
	}

	return receiver.value
}

// Unwrap is used to unwrap a finger.Host.
//
//	var host finger.Host
//	
//	// ...
//	
//	value, something := host.Unwrap()
//
// If finger.Host is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Host is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Host) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}
