package finger

import (
	"fmt"
	"strings"
)

// A Query represents a parsed finger-protocol target.
//
// For example, consider the following finger-protocol request:
//
//	"/W dariush@changelog.ca@example.com\r\n"
//
// When left as just an uninterpreted string, the finger-protocol target is:
//
//	"dariush@changelog.ca@example.com"
//
// Interpretting this finger-protocol target string, and separating out the:
//
// • user (i.e,. "dariush")
//
// • addresses (i.e., "changelog.ca" and "example.com")
//
// Is what makes it a finger-protocol query.
type Query struct {
	user User
	addresses []Address
}

func NoQuery() Query {
	return Query{}
}

func SomeQueryAddresses(addresses ...Address) Query {
	return Query{
		addresses: addresses,
	}
}

func SomeQueryHost(host string) Query {
	return Query{
		addresses: []Address{
			SomeAddressHost(host),
		},
	}
}

func SomeQueryHostPort(host string, port uint16) Query {
	return Query{
		addresses: []Address{
			SomeAddress(host, port),
		},
	}
}

func SomeQueryHosts(hosts ...string) Query {
	var addresses []Address

	for _, hostString := range hosts {

		var address Address

		if "" != hostString {
			address = SomeAddressHost(hostString)
		}

		addresses = append(addresses, address)
	}

	return Query{
		addresses: addresses,
	}
}

func SomeQueryUser(user string) Query {
	return Query{
		user: CreateUser(user),
	}
}

func SomeQueryUserAddresses(user string, addresses ...Address) Query {
	return Query{
		user: CreateUser(user),
		addresses: addresses,
	}
}

func SomeQueryUserHost(user string, host string) Query {
	return Query{
		user: CreateUser(user),
		addresses: []Address{
			SomeAddressHost(host),
		},
	}
}

func SomeQueryUserHosts(user string, hosts ...string) Query {
	var addresses []Address

	for _, hostString := range hosts {

		var address Address

		if "" != hostString {
			address = SomeAddressHost(hostString)
		}

		addresses = append(addresses, address)
	}

	return Query{
		user: CreateUser(user),
		addresses: addresses,
	}
}

func SomeQueryUserHostPort(user string, host string, port uint16) Query {
	return Query{
		user: CreateUser(user),
		addresses: []Address{
			SomeAddress(host, port),
		},
	}
}

// ParseQuery parses a (target) string for a finger-protocol query.
func ParseQuery(query string) (Query, error) {

	if "" == query {
		return Query{}, nil
	}

	var q Query

	{
		if '@' != query[0] {

			var index int = strings.IndexRune(query, '@')

			switch {
			case index < 0 && "" != query:
				q.user = CreateUser(query)
			default:
				q.user = CreateUser(query[:index])
				query = query[index:]
			}
		}
	}

	{
		for {
			if "" == query {
				break
			}
			if '@' != query[0] {
				break
			}

			query = query[1:]

			var index int = strings.IndexRune(query, '@')

			var address Address
			{
				var s string

				switch {
				case index < 0:
					s = query
					query = ""
				default:
					s = query[:index]
					query = query[index:]
				}

				var err error

				address, err = ParseAddress(s)
				if nil != err {
					return Query{}, fmt.Errorf("problem parsing finger-protocol query: %w", err)
				}
			}

			q.addresses = append(q.addresses, address)
		}
	}

	return q, nil
}

// ClientParameters returns the information need for a finger-protocol client
// to make a finger-protocol request.
func (receiver Query) ClientParameters() (Address, Query) {
	var addresses []Address = receiver.addresses

	length := len(addresses)

	if length < 1 {
		return DefaultAddress(), receiver
	}

	return addresses[length-1], Query{
		user: receiver.user,
		addresses: addresses[:length-1],
	}
}

func (receiver Query) String() string {

	var buffer strings.Builder

	{
		user, userIsSomething := receiver.user.Unwrap()
		if userIsSomething {
			buffer.WriteString(user)
		}
	}

	{
		for _, address := range receiver.addresses {
			buffer.WriteRune('@')

			buffer.WriteString(address.String())
		}
	}

	return buffer.String()
}

// Targets returns the equivalent finger.Target to finger.Query.
func (receiver Query) Target() Target {
	if EmptyUser() == receiver.user && len(receiver.addresses) < 1 {
		return EmptyTarget()
	}

	return CreateTarget(receiver.String())
}
