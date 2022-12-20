package testfinger

import (
	"github.com/reiver/go-arbitrary"

	"net"
	"time"
)

// TestConnectedWriteCloser is a finger.ConnectedWritCloser that can be used for testing purposes.
// You can think of it as a "dummy" net.Conn when the net.Conn is only written to.
//
// What has been written into it can be retreived with its String method.
//
// Example usage:
//
//	var buffer testfinger.TestConnectedWriteCloser
//	
//	var responsewriter finger.ResponseWriter = finger.NewResponseWriter(&buffer)
//	
//	// ...
//
//	var writtenData string = buffer.String()
//
// You can give the responsewriter you make out of one of these to a handler fuction.
// And check to see what the handler function actually wrote.
type TestConnectedWriteCloser struct {
	storage []byte
	closed bool
	localAddr net.Addr
	remoteAddr net.Addr
}

func (receiver *TestConnectedWriteCloser) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.closed = true

	return nil
}

func (receiver *TestConnectedWriteCloser) LocalAddr() net.Addr {
	if nil == receiver {
		return nil
	}

	if receiver.closed {
		return nil
	}

	if nil == receiver.localAddr {
		receiver.localAddr = arbitrary.NetAddr()
	}

	return receiver.localAddr
}

func (receiver *TestConnectedWriteCloser) RemoteAddr() net.Addr {
	if nil == receiver {
		return nil
	}

	if receiver.closed {
		return nil
	}

	if nil == receiver.remoteAddr {
		receiver.remoteAddr = arbitrary.NetAddr()
	}

	return receiver.remoteAddr
}

func (receiver *TestConnectedWriteCloser) SetWriteDeadline(t time.Time) error {
	if nil == receiver {
		return errNilReceiver
	}

//@TODO
	return nil
}

func (receiver *TestConnectedWriteCloser) String() string {
	if nil == receiver {
		return ""
	}

	return string(receiver.storage)
}

func (receiver *TestConnectedWriteCloser) Write(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	if receiver.closed {
		return 0, errClosed
	}

	if len(p) <= 0 {
		return 0, nil
	}

	receiver.storage = append(receiver.storage, p...)
	return len(p), nil
}
