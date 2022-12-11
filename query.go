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
	path Path
	addresses []Address
}

func EmptyQuery() Query {
	return Query{}
}

func AssembleQueryAddresses(addresses ...Address) Query {
	return Query{
		addresses: addresses,
	}
}

func AssembleQueryUser(user User) Query {
	return Query{
		user: user,
	}
}

func AssembleQueryUserPath(user User, path Path) Query {
	return Query{
		user: user,
		path: path,
	}
}

func AssembleQueryUserPathAddresses(user User, path Path, addresses ...Address) Query {
	return Query{
		user: user,
		path: path,
		addresses: addresses,
	}
}

func AssembleQueryUserAddresses(user User, addresses ...Address) Query {
	return Query{
		user: user,
		addresses: addresses,
	}
}

func CreateQueryHost(host string) Query {
	return Query{
		addresses: []Address{
			CreateAddressHost(host),
		},
	}
}

func CreateQueryHostPort(host string, port uint16) Query {
	return Query{
		addresses: []Address{
			CreateAddress(host, port),
		},
	}
}

func CreateQueryHosts(hosts ...string) Query {
	var addresses []Address

	for _, hostString := range hosts {

		var address Address

		if "" != hostString {
			address = CreateAddressHost(hostString)
		}

		addresses = append(addresses, address)
	}

	return Query{
		addresses: addresses,
	}
}

func CreateQueryUser(user string) Query {
	return Query{
		user: CreateUser(user),
	}
}

func CreateQueryUserHost(user string, host string) Query {
	return Query{
		user: CreateUser(user),
		addresses: []Address{
			CreateAddressHost(host),
		},
	}
}

func CreateQueryUserHosts(user string, hosts ...string) Query {
	var addresses []Address

	for _, hostString := range hosts {

		var address Address

		if "" != hostString {
			address = CreateAddressHost(hostString)
		}

		addresses = append(addresses, address)
	}

	return Query{
		user: CreateUser(user),
		addresses: addresses,
	}
}

func CreateQueryUserPath(user string, path string) Query {
	return Query{
		user: CreateUser(user),
		path: CreatePath(path),
	}
}

func CreateQueryUserPathHost(user string, path string, host string) Query {
	return Query{
		user: CreateUser(user),
		path: CreatePath(path),
		addresses: []Address{
			CreateAddressHost(host),
		},
	}
}


func CreateQueryUserPathHosts(user string, path string, hosts ...string) Query {
	var addresses []Address

	for _, hostString := range hosts {

		var address Address

		if "" != hostString {
			address = CreateAddressHost(hostString)
		}

		addresses = append(addresses, address)
	}

	return Query{
		user: CreateUser(user),
		path: CreatePath(path),
		addresses: addresses,
	}
}

func CreateQueryUserHostPort(user string, host string, port uint16) Query {
	return Query{
		user: CreateUser(user),
		addresses: []Address{
			CreateAddress(host, port),
		},
	}
}

func CreateQueryUserPathHostPort(user string, path string, host string, port uint16) Query {
	return Query{
		user: CreateUser(user),
		path: CreatePath(path),
		addresses: []Address{
			CreateAddress(host, port),
		},
	}
}

// ParseQuery parses a (target) string for a finger-protocol query.
func ParseQuery(query string) (Query, error) {

	if "" == query {
		return EmptyQuery(), nil
	}

	var q Query

	{
		var query0 byte = query[0]

		if '/' != query0 && '@' != query0 {

			var indexSolidus int = strings.IndexRune(query, '/')
			var indexAt      int = strings.IndexRune(query, '@')

			// Ex: "joeblow"
			if indexSolidus < 0 && indexAt < 0 {
				q.user = CreateUser(query)
/////////////////////////////// RETURN
				return q, nil
			}

			var index int
			switch {
			// Ex: "joeblow/once/twice/thrice/fource"
			case 0 <= indexSolidus && indexAt < 0 && "" != query:
				index = indexSolidus
			// Ex: "joeblow@example.om"
			case indexSolidus < 0 && 0 <= indexAt && "" != query:
				index = indexAt
			// Ex: "joeblow@example.com"
			case 0 <= indexSolidus && 0 <= indexAt && "" != query:
				index = indexSolidus
				if indexAt < index {
					index = indexAt
				}
			}

			q.user = CreateUser(query[:index])
			query = query[index:]
		}
	}

	if "" == query {
		return q, nil
	}

	{
		var query0 byte = query[0]

		if '/' == query0 {

			var indexAt int = strings.IndexRune(query, '@')

			if indexAt < 0 {
				q.path = CreatePath(query)
/////////////////////////////// RETURN
				return q, nil
			}

			var index int = indexAt

			q.path = CreatePath(query[:index])
			query = query[index:]
		}
	}

	if "" == query {
		return q, nil
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

// ParseRFC1288Query parses a (target) string for an older style finger-protocol query as defined IETF RFC-1288.
func ParseRFC1288Query(query string) (Query, error) {

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
		path: receiver.path,
		addresses: addresses[:length-1],
	}
}

func (receiver Query) isEmpty() bool {
	return EmptyUser() == receiver.user &&
	       EmptyPath() == receiver.path &&
	       len(receiver.addresses) < 1
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
		path, pathIsSomething := receiver.path.Unwrap()
		if pathIsSomething {
			buffer.WriteString(path)
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
	if receiver.isEmpty() {
		return EmptyTarget()
	}

	return CreateTarget(receiver.String())
}
