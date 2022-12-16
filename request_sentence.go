package finger

import (
	"github.com/reiver/go-utf8"

	"io"
	"strings"
)

func (receiver Request) Sentence() string {
	var buffer strings.Builder

	receiver.writeSentenceTo(&buffer)

	return buffer.String()
}

func (receiver Request) writeSentenceTo(w io.Writer) (int64, error) {
	if nil == w {
		return 0, errNilWriter
	}

	var n64 int64
	{
		swtch, switchIsSomething  := receiver.swtch.Unwrap()
		target, targetIsSomething := receiver.target.Unwrap()

		if switchIsSomething {
			n, err := io.WriteString(w, swtch)
			n64 += int64(n)
			if nil != err {
				return n64, err
			}
		}

		if switchIsSomething && targetIsSomething {
			n, err := utf8.WriteRune(w, ' ')
			n64 += int64(n)
			if nil != err {
				return n64, err
			}
		}

		if targetIsSomething {
			n, err := io.WriteString(w, target)
			n64 += int64(n)
			if nil != err {
				return n64, err
			}
		}
	}

	return n64, nil
}
