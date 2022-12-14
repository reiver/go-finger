package magicfinger

import (
	"github.com/reiver/go-finger"

	"github.com/reiver/go-utf8"

	"fmt"
	"io"
	"net"
	"strings"
)

var _ MagicResponseWriter = &internalMagicResponseWriter{}

type internalMagicResponseWriter struct {
	rw finger.ResponseWriter
	buffer strings.Builder
	magicBuffered bool
	headerWritten bool
}

// NewMagicResponseWriter creates a new magicfinger.MagicResponseWriter.
func NewMagicResponseWriter(rw finger.ResponseWriter) MagicResponseWriter {
	return &internalMagicResponseWriter{
		rw:rw,
	}
}

func (receiver *internalMagicResponseWriter) AddField(name string, body string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if strings.ContainsAny(name, "\r\n:") {
		return errBadFieldName
	}
	if strings.ContainsAny(body, "\r\n") {
		return errBadFieldBody
	}

	receiver.bufferOnceMagic()

	receiver.buffer.WriteString(name)
	receiver.buffer.WriteString(": ")
	receiver.buffer.WriteString(body)
	receiver.buffer.WriteString("\r\n")

	return nil
}

func (receiver *internalMagicResponseWriter) bufferOnceMagic() {
	if nil == receiver {
		return
	}

	if !receiver.magicBuffered {
		receiver.buffer.WriteString(magic)
		receiver.magicBuffered = true
	}
}

func (receiver *internalMagicResponseWriter) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	var rw finger.ResponseWriter
	{
		rw = receiver.rw
		if nil == rw {
			return errNilResponseWriter
		}
	}

	return rw.Close()
}

func (receiver internalMagicResponseWriter) LocalAddr() net.Addr {

	var rw finger.ResponseWriter
	{
		rw = receiver.rw
		if nil == rw {
			return nil
		}
	}

	return rw.LocalAddr()
}

func (receiver internalMagicResponseWriter) RemoteAddr() net.Addr {

	var rw finger.ResponseWriter
	{
		rw = receiver.rw
		if nil == rw {
			return nil
		}
	}

	return rw.RemoteAddr()
}

func (receiver *internalMagicResponseWriter) Write(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	{
		err := receiver.writeOnceHeader()
		if  nil != err {
			return 0, err
		}
	}

	var rw finger.ResponseWriter
	{
		rw = receiver.rw
		if nil == rw {
			return 0, errNilResponseWriter
		}
	}

	return rw.Write(p)
}

func (receiver *internalMagicResponseWriter) WriteByte(b byte) error {
	if nil == receiver {
		return errNilReceiver
	}

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

func (receiver *internalMagicResponseWriter) writeOnceHeader() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.bufferOnceMagic()

	var rw finger.ResponseWriter
	{
		rw = receiver.rw
		if nil == rw {
			return errNilResponseWriter
		}
	}

	if !receiver.headerWritten {
		receiver.buffer.WriteString("\r\n")

		_, err := io.WriteString(rw, receiver.buffer.String())
		if nil != err {
			return fmt.Errorf("problem writnig header for finger-protocol response: %w", err)
		}
		receiver.headerWritten = true
	}

	return nil
}

func (receiver *internalMagicResponseWriter) WriteRune(r rune) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var writer io.Writer = receiver

	var wrapped utf8.RuneWriter = utf8.RuneWriterWrap(writer)
	var runewriter runeWriter = &wrapped

	return runewriter.WriteRune(r)
}

func (receiver *internalMagicResponseWriter) WriteString(s string) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var writer io.Writer = receiver

	return io.WriteString(writer, s)
}
