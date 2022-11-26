package finger

import (
	"fmt"
	"strconv"
	"strings"
)

// An Address represents a finger-protocol address that is part of a finger-protocol query.
//
// Here are some example finger-protocol queries and a list of each of their finger-protocol addresses:
//
//	"joeblow@example.com"
//	// finger-protocol address -> "example.com"
//
//	"dariush@changelog.ca"
//	// finger-protocol address -> "changelog.ca"
//
//	"joeblow@example.com:1971"
//	// finger-protocol address -> "example.com:1971"
//
//	"dariush@changelog.ca:12345"
//	// finger-protocol address -> "changelog.ca:12345"
//
//	"joeblow@example.com@something.social"
//	// finger-protocol addresses -> "example.com", "something.social"
//
//	"dariush@changelog.ca@once.com@twice.net@thrice.org@fource.dev"
//	// finger-protocol addresses -> "changelog.ca", "once.com", "twice.net", "thrice.org", "fource.dev"
//
//	"joeblow@example.com:1971@something.social"
//	// finger-protocol addresses -> "example.com", "something.social"
//
//	"dariush@changelog.ca@once.com:54321@twice.net@thrice.org:1212@fource.dev"
//	// finger-protocol addresses -> "changelog.ca", "once.com:54321", "twice.net", "thrice.org:1212", "fource.dev"
//
// Note that not all of these have a TCP-port.
//
// With the finger-protocol, if a TCP-port isn't specified, then it defaults to TCP-port 79.
//
// To have the TCP-port explicitly added in, use the Resolve method.
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
func EmptyAddress() Address {
	return Address{}
}

// CreateAddress is used to create a finger.Address with something in it.
func CreateAddress(host string, port uint16) Address {
	return Address {
		host: CreateHost(host),
		port: CreatePort(port),
	}
}

// CreateAddressHost is used to create a finger.Address with something in it.
func CreateAddressHost(host string) Address {
	return Address {
		host: CreateHost(host),
	}
}

// CreateAddressPort is used to create a finger.Address with something in it.
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
func ParseAddress(s string) (Address, error) {

	if "" == s {
		return Address{}, nil
	}

	index := strings.IndexRune(s, ':')

	if index < 0 {
		return Address{
			host: CreateHost(s),
		}, nil
	}

	var address Address
	{
		var host string = s[:index]

		address.host = CreateHost(host)

		var err error
		address.port, err = ParsePort(s[1+index:])
		if nil != err {
			return address, fmt.Errorf("problem parsing finger-protocol port: %w", err)
		}
	}

	return address, nil
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
func (receiver Address) Resolve() string {
	return fmt.Sprintf("%s:%d", receiver.host.Resolve(), receiver.port.Resolve())
}

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
