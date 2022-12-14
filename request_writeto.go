package finger

import (
	"io"
	"strings"
)

var _ io.WriterTo = Request{}

// WriteTo writers a finger-protocol request to an io.Writer.
//
// WriteTo returns the number of bytes written, and any error that occurs.
func (receiver Request) WriteTo(writer io.Writer) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

//@TODO: replace with a limited buffer.
	var buffer strings.Builder
	{
		_, err := receiver.writeSentenceTo(&buffer)
		if nil != err {
			return 0, errInternalError
		}

		buffer.WriteString("\r\n")
	}

	var n64 int64
	var err error
	{
		var n int

		n, err = io.WriteString(writer, buffer.String())
		n64 = int64(n)
	}

	return n64, err
}
