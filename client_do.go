package finger

import (
	"fmt"
	"net"
)

// Do does a finger-orotocol request.
func (receiver Client) Do(request Request) (ResponseReader, error) {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil, errNilConnection
		}
	}

	{
//@TODO: check the number-of-bytes-written
		_, err := request.WriteTo(conn)
		if nil != err {
			return nil, fmt.Errorf("problem writing finger-protocol request to network-connection: %w", err)
		}
	}

	return internalResponseReader{conn}, nil
}

