package finger

import (
	"strings"
)

type Query struct {
	UserName UserName
	Hosts []Host
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

			var host Host
			{
				switch {
				case index < 0:
					host = SomeHost(query)
				default:
					host = SomeHost(query[:index])
					query = query[index:]
				}
			}

			q.Hosts = append(q.Hosts, host)
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
		for _, host := range receiver.Hosts {
			buffer.WriteRune('@')

			s, _ := host.Unwrap()
			buffer.WriteString(s)
		}
	}

	return buffer.String()
}
