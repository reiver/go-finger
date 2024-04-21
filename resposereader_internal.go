package finger

import (
	"sourcecode.social/reiver/go-utf8"

	"fmt"
	"io"
	"net"
)

var _ ResponseReader = internalResponseReader{}

type internalResponseReader struct {
	conn ConnectedReadCloser
}

func (receiver internalResponseReader) Close() error {

	var conn ConnectedReadCloser
	{
		conn = receiver.conn
		if nil == conn {
			return errNilConnection
		}
	}

	return conn.Close()
}

func (receiver internalResponseReader) LocalAddr() net.Addr {

	var conn ConnectedReadCloser
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.LocalAddr()
}

func (receiver internalResponseReader) Read(p []byte) (int, error) {

	var conn ConnectedReadCloser
	{
		conn = receiver.conn
		if nil == conn {
			return 0, errNilConnection
		}
	}

	return conn.Read(p)
}

func (receiver internalResponseReader) ReadByte() (byte, error) {

	var buffer [1]byte
	var p []byte = buffer[:]

	n, err := receiver.Read(p)
	if nil != err {
		return 0, fmt.Errorf("problem reading byte from finger-protocol response: %w", err)
	}
	if expected, actual := 1, n; expected != actual {
		return 0, fmt.Errorf("problem reading byte from finger-protocol response: actual number of bytes read is %d but expected it to be %d", actual, expected)
	}

	return buffer[0], nil
}

func (receiver internalResponseReader) ReadRune() (rune, int, error) {

	var conn ConnectedReadCloser
	{
		conn = receiver.conn
		if nil == conn {
			return utf8.RuneError, 0, errNilConnection
		}
	}

	var wrapped utf8.RuneReader = utf8.WrapRuneReader(conn)
	var runereader io.RuneReader = &wrapped

	return runereader.ReadRune()
}

func (receiver internalResponseReader) RemoteAddr() net.Addr {

	var conn ConnectedReadCloser
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.RemoteAddr()
}
