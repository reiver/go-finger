package finger

import (
	"fmt"
	"strings"
)

type Query struct {
	UserName UserName
	Addresses []Address
}

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
				q.UserName = SomeUserName(query)
			default:
				q.UserName = SomeUserName(query[:index])
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

func (receiver Query) String() string {

	var buffer strings.Builder

	{
		username, usernameIsSomething := receiver.UserName.Unwrap()
		if usernameIsSomething {
			buffer.WriteString(username)
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
