package finger

import (
	"sourcecode.social/reiver/go-erorr"

	"fmt"
	"net"
	"strconv"
	"strings"
)

// An Address represents a finger-protocol address that is part of a finger-protocol query.
//
// Address is important because it is used to create a TCP connection using the net.Dial() function.
//
// I.e.,:
//
//	var address finger.Address
//	
//	// ...
//	
//	conn, err := net.Dial("tcp", address.Resolve())
//
// Note that we call the Resolve() method on finger.Address to give us a proper address-string that the net.Dial() can use.
//
// ⁂
//
// A finger-protocol client will likely get a finger.Address (that it can use to make a TCP-connection with the net.Dial() function)
// from the finger.Query.ClientParameters() method.
//
// For example:
//
//	var query finger.Query
//
//	// ...
//
//	address, query := query.ClientParameters()
//
//	// ...
//
//	conn, err := net.Dial("tcp", address.Resolve())
//
// Again notice that we called the Resolve() method on finger.Address to give us a proper address-string that the net.Dial() can use.
//
// ⁂
//
// Here are some example finger-protocol queries with just a finger-protocol actor and a single finger-protocol host:
//
//	"joeblow@example.com"
//	// finger-protocol address -> "example.com"
//	//
//	// TCP-port not explicitly provided, but defaults to 79 — i.e., is equivalent of "example.com:79"
//
//	"dariush@reiver.link"
//	// finger-protocol address -> "reiver.link"
//	//
//	// TCP-port not explicitly provided, but defaults to 79 — i.e., is equivalent of "reiver.link:79"
//
//	"someone@127.0.0.1"
//	// finger-protocol address -> "127.0.0.1"
//	//
//	// TCP-port not explicitly provided, but defaults to 79 — i.e., is equivalent of "127.0.0.1:79"
//
//	"janedoe@12.23.34.45"
//	// finger-protocol address -> "12.23.34.45"
//	//
//	// TCP-port not explicitly provided, but defaults to 79 — i.e., is equivalent of "12.23.34.45:79"
//
// Here are some example finger-protocol queries with just a finger-protocol actor, a single finger-protocol host, and a TCP-port:
//
//	"joeblow@example.com:1971"
//	// finger-protocol address -> "example.com:1971"
//
//	"dariush@reiver.link:12345"
//	// finger-protocol address -> "reiver.link:12345"
//
//	"someone@127.0.0.1:7979"
//	// finger-protocol address -> "127.0.0.1:7979"
//
//	"janedoe@12.23.34.45:79"
//	// finger-protocol address -> "12.23.34.45:79"
//
// Here are some example finger-protocol queries with just a finger-protocol actor, and two finger-protocol host:
//
//	"joeblow@example.com@something.social"
//	// finger-protocol addresses -> "example.com", "something.social"
//
//	"dariush@reiver.link@example.com
//	// finger-protocol addresses -> "reiver.link", "example.com"
//
//	"janedoe@12.23.34.45:79:111.222.3.4:7979"
//	// finger-protocol address -> "12.23.34.45:79", "111.222.3.4:7979"
//
// Here are more example finger-protocol queries with various forms:
//
//	"dariush@reiver.link@once.com@twice.net@thrice.org@fource.dev"
//	// finger-protocol addresses -> "reiver.link", "once.com", "twice.net", "thrice.org", "fource.dev"
//
//	"joeblow@example.com:1971@something.social"
//	// finger-protocol addresses -> "example.com:1971", "something.social"
//
//	"dariush@reiver.link@once.com:54321@twice.net@thrice.org:1212@fource.dev"
//	// finger-protocol addresses -> "reiver.link", "once.com:54321", "twice.net", "thrice.org:1212", "fource.dev"
//
// Note that not all of these have a TCP-port.
//
// With the finger-protocol, if a TCP-port isn't specified, then it defaults to TCP-port 79.
//
// To have the TCP-port explicitly added in (which is needed for net.Dial()), use the Resolve method.
// For example:
//
//	address, err := finger.ParseAddress("example.com")
//	
//	// ...
//	
//	resolvedAddress := address.Resolve()
//	// resolvedAddress -> "example.com:79"
//	
//	conn, err := net.Dial("tcp", resolvedAddress)
type Address struct {
	host Host
	port Port
}

// EmptyAddress is used to create a finger.Address with nothing in it.
//
// For example:
//
//	var address finger.Address = finger.EmptyAddress()
func EmptyAddress() Address {
	return Address{}
}

// CreateAddress is used to create a finger.Address with something in it.
//
// For example:
//
//	var address finger.Address = finger.CreateAddress("example.com", 1971)
func CreateAddress(host string, port uint16) Address {
	return Address {
		host: CreateHost(host),
		port: CreatePort(port),
	}
}

// CreateAddressHost is used to create a finger.Address with something in it.
//
// For example:
//
//	var address finger.Address = finger.CreateAddressHost("example.com")
func CreateAddressHost(host string) Address {
	return Address {
		host: CreateHost(host),
	}
}

// CreateAddressPort is used to create a finger.Address with something in it.
//
// For example:
//
//	var address finger.Address = finger.CreateAddressPort(1971)
func CreateAddressPort(port uint16) Address {
	return Address {
		port: CreatePort(port),
	}
}

// DefaultAddress is used to create a finger.Port with the logical value of 127.0.0.1:79 in it.
func DefaultAddress() Address {
	return Address{
		host: DefaultHost(),
		port: DefaultPort(),
	}
}

// ParseAddress parses a finger-protocol address (as a string).
//
// Some example addresses include:
//
//	""
//
//	"example.com"
//
//	":1971"
//
//	"example.com:1971"
//
// Here are some more random examples:
//
//	"reiver.link"
//
//	"reiver.link:79"
//
//	"reiver.link:1971"
//
//	"reiver.link:7979"
//
//	"12.23.34.45"
//
//	"12.23.34.45:79"
//
//	"12.23.34.45:1971"
//
//	"12.23.34.45:7979"
func ParseAddress(s string) (Address, error) {

	if "" == s {
		return EmptyAddress(), nil
	}

	index := strings.IndexRune(s, ':')

	if index < 0 {
		return CreateAddressHost(s), nil
	}

	var host string = s[:index]

	var port uint16
	{
		p, err := ParsePort(s[1+index:])
		if nil != err {
			return EmptyAddress(), fmt.Errorf("problem parsing finger-protocol port when parsing finger-protocol address: %w", err)
		}

		var something bool
		port, something = p.Unwrap()
		if !something {
			return EmptyAddress(), erorr.Error("unexpected problem when parsing finger-protocol address: empty port")
		}
	}

	if "" == host {
		return CreateAddressPort(port), nil
	}

	return CreateAddress(host, port), nil
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var addressString string = "12.23.34.45:1971"
//	
//	address, err := finger.ParseAddress(addressString)
//	if nil != err {
//		return err
//	}
//	
//	// ...
//	
//	fmt.Printf("address = %#v", address)
//	
//	// Output:
//	// address = finger.CreateAddress("12.23.34.45", 1971)
//
// And, for example:
//
//	var address finger.Address = finger.CreateAddress("example.com", 79)
//	
//	// ...
//	
//	fmt.Printf("address = %#v", address)
//	
//	// Output:
//	// address = finger.CreateAddress("example.com", 79)
//
// Also, for example:
//
//	var address finger.Address = finger.EmptyAddress()
//	
//	// ...
//	
//	fmt.Printf("address = %#v", address)
//	
//	// Output:
//	// address = finger.EmptyAddress()
//
// And, for example:
//
//	var address finger.Address = finger.CreateAddressHost("example.com")
//	
//	// ...
//	
//	fmt.Printf("address = %#v", address)
//	
//	// Output:
//	// address = finger.CreateAddressHost("example.com")
//
// And again, for example:
//
//	var address finger.Address = finger.CreateAddressPort(79)
//	
//	// ...
//	
//	fmt.Printf("address = %#v", address)
//	
//	// Output:
//	// address = finger.CreateAddressPort(79)
func (receiver Address) GoString() string {

	host, hostIsSomething := receiver.host.Unwrap()
	port, portIsSomething := receiver.port.Unwrap()

	switch {
	case hostIsSomething && portIsSomething:
		return fmt.Sprintf("finger.CreateAddress(%q, %d)", host, port)
	case hostIsSomething && !portIsSomething:
		return fmt.Sprintf("finger.CreateAddressHost(%q)", host)
	case !hostIsSomething && portIsSomething:
		return fmt.Sprintf("finger.CreateAddressPort(%d)", port)
	default:
		return "finger.EmptyAddress()"
	}
}

func (receiver Address) Host() Host {
	return receiver.host
}

func (receiver Address) Port() Port {
	return receiver.port
}

// Use what is returned from the Resolve method, to pass to net.Dial().
//
// For example:
//
//	var address finger.Address
//	
//	// ...
//	
//	conn, err := net.Dial("tcp", address.Resolve())
//
// ⁂
//
// Here are some examples:
//
//	var address finger.Address = finger.CreateAddress("example.com", 79)
//	
//	var s string = address.String()
//	// s = "example.com:79"
//	
//	var r string = address.Resolve()
//	// r = "example.com:79"
//
// .
//
//	var address finger.Address = finger.CreateAddressHost("example.com")
//	
//	var s string = address.String()
//	// s = "example.com"
//	
//	var r string = address.Resolve()
//	// r = "example.com:79"
//
// .
//
//	var address finger.Address = finger.CreateAddressPort(79)
//	
//	var s string = address.String()
//	// s = ":79"
//	
//	var r string = address.Resolve()
//	// r = "127.0.0.1:79"
//
// .
//
//	var address finger.Address = finger.EmptyAddress()
//	
//	var s string = address.String()
//	// s = ""
//	
//	var r string = address.Resolve()
//	// r = "127.0.0.1:79"
//
// Notice that when the host was not specified, when it was resolved it was defaulted to "127.0.0.1".
//
// And also notice that when the port was not specified, when it was resolved it was defaulted to "79".
func (receiver Address) Resolve() string {
	var host string = receiver.host.Resolve()
	var port string = strconv.FormatUint(uint64(receiver.port.Resolve()), 10)

	return net.JoinHostPort(host, port)
}

// String returns the address in the format as it was created.
//
// NOTE THAT THE OUTPUT OF finger.Address.String() SHOULD NOT BE USED WITH net.Dial().
// INSTEAD USE THE OUTPUT OF finger.AddressResolve() WITH net.Dial().
//
// ⁂
//
// Here are some examples:
//
//	var address finger.Address = finger.CreateAddress("example.com", 79)
//	
//	var s string = address.String()
//	// s = "example.com:79"
//
// .
//
//	var address finger.Address = finger.CreateAddressHost("example.com")
//	
//	var s string = address.String()
//	// s = "example.com"
//
// .
//
//	var address finger.Address = finger.CreateAddressPort(79)
//	
//	var s string = address.String()
//	// s = ":79"
//
// .
//
//	var address finger.Address = finger.EmptyAddress()
//	
//	var s string = address.String()
//	// s = ""
func (receiver Address) String() string {
	host, hostIsSomething := receiver.host.Unwrap()
	port, portIsSomething := receiver.port.Unwrap()

	if !hostIsSomething && !portIsSomething {
		return ""
	}

	var buffer strings.Builder
	{
		if hostIsSomething {
			buffer.WriteString(host)
		}

		if portIsSomething {
			buffer.WriteRune(':')

			const base int = 10

			buffer.WriteString(strconv.FormatUint(uint64(port), base))
		}
	}

	return buffer.String()
}
