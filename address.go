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
type Address struct {
	Host Host
	Port Port
}

func DefaultAddress() Address {
	return Address{
		Host: DefaultHost(),
		Port: DefaultPort(),
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
			Host: SomeHost(s),
		}, nil
	}

	var address Address
	{
		var host string = s[:index]

		address.Host = SomeHost(host)

		var err error
		address.Port, err = ParsePort(s[1+index:])
		if nil != err {
			return address, fmt.Errorf("problem parsing finger-protocol port: %w", err)
		}
	}

	return address, nil
}

func (receiver Address) Resolve() string {
	return fmt.Sprintf("%s:%d", receiver.Host.Resolve(), receiver.Port.Resolve())
}

func (receiver Address) String() string {
	host, hostIsSomething := receiver.Host.Unwrap()
	port, portIsSomething := receiver.Port.Unwrap()

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
