package finger

import (
	"io"
	"net"
	"time"
)

// ConnectedReadCloser represents the reading part of a net.Conn.
type ConnectedReadCloser interface {
	io.ReadCloser
	Connection
	SetReadDeadline(time.Time) error
}

var _ ConnectedReadCloser = net.Conn(nil)

// ConnectedWriteCloser represents the writing part of a net.Conn.
type ConnectedWriteCloser interface {
	io.WriteCloser
	Connection
	SetWriteDeadline(time.Time) error
}

var _ ConnectedWriteCloser = net.Conn(nil)

type Connection interface {
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

var _ Connection = net.Conn(nil)



