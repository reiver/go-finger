package classicfinger

import (
	"github.com/reiver/go-finger"

	"fmt"
	"io"
)

func RespondServerErred(rw finger.ResponseWriter, request finger.Request) (e error) {
	if nil == rw {
		return errNilResponseWriter
	}

	defer func(){
		if err := rw.Close(); nil != err {
			if nil == e {
				e = fmt.Errorf("problem closing classic-finger connection to client: %w", err)
			}
		}
	}()

	var msg string = fmt.Sprint("server erred! - %s", request.Sentence())

	{
		n, err := io.WriteString(rw, msg)
		if nil != err {
			e = fmt.Errorf("problem sending server resopnse to client: %w", err)
			return
		}
		if expected, actual := len(msg), n; expected != actual {
			e = fmt.Errorf("problem sending server response to client: actual number of bytes written is %d but expected it to be %d", actual, expected)
			return
		}
	}

	return
}
