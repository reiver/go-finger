package finger

import (
	"sourcecode.social/reiver/go-netln"

	"fmt"
	"io"
	"strings"
)

// ReadRequestLine reads in a "\r\n" terminated line from 'reader',
// and returns it.
//
// The returned value does not include the terminating "\r\n".
//
// So if what is read from 'reader' (logically) is:
//
//	"/W joeblow\r\n"
//
// Then what is returned is"
//
//	"/W joeblow"
//
// Note that typically the value returned from ReadRequestLine would be given to ParseRequest.
// For example:
//
//	var reader io.Reader
//	
//	// ...
//	
//	requestline, err := finger.ReadRequestLine(reader)
//	if nil != err {
//		return err
//	}
//	
//	request, err := finger.ParseRequest(requestline)
//	if nil != err {
//		return err
//	}
func ReadRequestLine(reader io.Reader) (string, error) {

	var buffer strings.Builder

	{
		_, err := netln.CopyLine(&buffer, reader)
		if io.EOF == err {
			return buffer.String(), io.EOF
		}
		if nil != err {
			return "", fmt.Errorf("problem reading request-line: %w", err)
		}
	}

	return buffer.String(), nil
}
