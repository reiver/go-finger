package magicfinger

import (
	"github.com/reiver/go-finger"

	"fmt"
)

// RespondServerErred is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server erred when trying
// to handle the finger-protocol client's request.
//
// For example, maybe server's hard disk crashed.
// for some reason.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseServerErred(object, rw)
//		
//		// ...
//		
//	}
func RespondServerErred(object string, rw finger.ResponseWriter) error {
	if nil == rw {
		return errNilResponseWriter
	}

	const punctuation string = "!"
	const verb        string = "ERRED"

	var mrw MagicResponseWriter = NewMagicResponseWriter(punctuation, verb, object, rw)
	if err := mrw.Flush(); nil != err {
		return fmt.Errorf("problem sending magic-finger \"%s%s %s\" response to client: %w", punctuation, verb, object, err)
	}

	return nil
}