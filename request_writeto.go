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
		swtch, switchIsSomething  := receiver.swtch.Unwrap()
		target, targetIsSomething := receiver.target.Unwrap()

		if switchIsSomething {
			buffer.WriteString(swtch)
		}

		if switchIsSomething && targetIsSomething {
			buffer.WriteRune(' ')
		}

		if targetIsSomething {
			buffer.WriteString(target)
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
