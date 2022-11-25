package finger

import (
	"strconv"
	"strings"
)

type Address struct {
	Host Host
	Port Port
}

func (receiver Address) String() string {

	host, hostIsSomething := receiver.Host.Unwrap()
	if !hostIsSomething {
		return ""
	}

	port, portIsSomething := receiver.Port.Unwrap()

	var buffer strings.Builder
	{
		buffer.WriteString(host)

		if portIsSomething {
			buffer.WriteRune(':')
			buffer.WriteString(strconv.FormatUint(uint64(port), 10))
		}
	}

	return buffer.String()
}
