package finger

import (
	"net"
)

// A Client is a finger-protocol client.
//
// It handles sending a finger-protocol client request and receiving a finger-protocol server response.
//
//	var address finger.Address = finger.CreateAddress("example.com", 79)
//	
//	// ...
//	
//	var conn net.Conn
//	
//	conn, err = net.Dial("tcp", address.Resolve())
//	
//	// ...
//	
//	var client finger.Client = finger.AssembleClient(conn)
//	
//	// ...
//	
//	var request finger.Request = finger.SomeRequestTarget("joeblow")
//	
//	// ...
//	
//	responseReader, err := client.Do(request)
//	
//	// ...
//	
//	io.Copy(os.Stdout, response)
type Client struct {
	conn net.Conn
}

// AssebleClient returns a client that communicates over a net.Conn.
//
// Example
//
//	var conn net.Conn
//	
//	// ...
//	
//	var fingerClient finger.Client = finger.AssembleClient(conn)
func AssembleClient(conn net.Conn) Client {
	return Client{
		conn:conn,
	}
}

// Close closes the wrapped net.Conn.
func (receiver *Client) Close() error {
	if nil == receiver {
		return nil
	}

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.Close()
}

// LocalAddr returns the local-address of the connection, if known.
func (receiver Client) LocalAddr() net.Addr {
	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.LocalAddr()
}

// RemoteAddr returns the remote-address of the connection, if known.
func (receiver Client) RemoteAddr() net.Addr {
	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.RemoteAddr()
}
