package finger

import (
	"github.com/reiver/go-utf8"

	"fmt"
	"io"
	"net"
)

var _ ResponseWriter = internalResponseWriter{}

type internalResponseWriter struct {
	conn net.Conn
}

func (receiver internalResponseWriter) Close() error {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return errNilConnection
		}
	}

	return conn.Close()
}

func (receiver internalResponseWriter) LocalAddr() net.Addr {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.LocalAddr()
}

func (receiver internalResponseWriter) RemoteAddr() net.Addr {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return nil
		}
	}

	return conn.RemoteAddr()
}

func (receiver internalResponseWriter) Write(p []byte) (int, error) {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return 0, errNilConnection
		}
	}

	return conn.Write(p)
}

func (receiver internalResponseWriter) WriteByte(b byte) error {

	var buffer [1]byte
	var p []byte = buffer[:]

	n, err := receiver.Write(p)
	if nil != err {
		return fmt.Errorf("problem writing byte for finger-protocol response: %w", err)
	}

	if expected, actual := 1, n; expected != actual {
		return fmt.Errorf("problem writing byte for finger-protocol response: actual number of bytes written is %d but expected it to be %d", actual, expected)
	}

	return nil
}

func (receiver internalResponseWriter) WriteRune(r rune) (int, error) {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return 0, errNilConnection
		}
	}

	var wrapped utf8.RuneWriter = utf8.RuneWriterWrap(conn)
	var runewriter runeWriter = &wrapped

	return runewriter.WriteRune(r)
}

func (receiver internalResponseWriter) WriteString(s string) (int, error) {

	var conn net.Conn
	{
		conn = receiver.conn
		if nil == conn {
			return 0, errNilConnection
		}
	}

	return io.WriteString(conn, s)
}
