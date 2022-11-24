package finger

import (
	"net"
)

// A ResponseWriter interface is used by a finger-protocol handler to construct a finger-protocol response.
//
// Example
//
// Here is an example:
//
//	func (receiver *MyFingerHandler) HandleFinger(rw finger.ResponseWriter, request finger.Request) {
		
//		// ...
//		
//		defer rw.Close()
//		rw.Write([]byte{0xEF,0xBB,0xBF}) // UTF-8 magic bytes
//		rw.WriteString("Hello world!")
//		
//		// ...
//	}
type ResponseWriter interface {
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Write([]byte) (int, error)
	WriteByte(byte) error
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}
