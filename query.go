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
	User User
	Addresses []Address
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
				q.User = SomeUser(query)
			default:
				q.User = SomeUser(query[:index])
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

			q.Addresses = append(q.Addresses, address)
		}
	}

	return q, nil

}

// ClientParameters returns the information need for a finger-protocol client
// to make a finger-protocol request.
func (receiver Query) ClientParameters() (Address, Query) {
	var addresses []Address = receiver.Addresses

	length := len(addresses)

	if length < 1 {
		return DefaultAddress(), receiver
	}

	return addresses[length-1], Query{
		User: receiver.User,
		Addresses: addresses[:length-1],
	}
}

func (receiver Query) String() string {

	var buffer strings.Builder

	{
		user, userIsSomething := receiver.User.Unwrap()
		if userIsSomething {
			buffer.WriteString(user)
		}
	}

	{
		for _, address := range receiver.Addresses {
			buffer.WriteRune('@')

			buffer.WriteString(address.String())
		}
	}

	return buffer.String()
}

// Targets returns the equivalent finger.Target to finger.Query.
func (receiver Query) Target() Target {
	if NoUser() == receiver.User && len(receiver.Addresses) < 1 {
		return NoTarget()
	}

	return SomeTarget(receiver.String())
}
