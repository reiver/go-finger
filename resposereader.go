package finger

import (
	"net"
)

// A ResponseReader interface is used by a finger-protocol client to read a finger-protocol response.
type ResponseReader interface {
	Close() error
	LocalAddr() net.Addr
	Read([]byte) (int, error)
	ReadByte() (byte, error)
	ReadRune() (rune, int, error)
	RemoteAddr() net.Addr
}
