package magicfinger

import (
	"fmt"
	"io"
)

const serverRefusedString = magic+ "!REFUSED" +" "+ "server" +"\r\n"+"\r\n"

// RespondServerRefused is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server refused to handle the
// finger-protocol client's request.
//
// For example, maybe server the server won't handle the "/PULL" switch for a user,
// for some reason.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseServerRefused(rw)
//		
//		// ...
//		
//	}
func RespondServerRefused(writer io.Writer) error {
	if nil == writer {
		return errNilWriter
	}

	const s string = serverRefusedString

	n, err := io.WriteString(writer, s)
	if nil != err {
		return fmt.Errorf("problem writing server-error finger-protocol response: %w", err)
	}
	if expected, actual := len(s), n; expected != actual {
		return fmt.Errorf("problem writing server-error finger-protocol response: actually wrote %d bytes but expected to write %d bytes", actual, expected)
	}

	return nil
}
