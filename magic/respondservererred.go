package magicfinger

import (
	"github.com/reiver/go-finger"

	"fmt"
)

// RespondServerErred is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server erred when trying
// to handle the finger-protocol client's request.
//
// For example, maybe server's hard disk crashed. for some reason.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseServerErred(rw, request)
//		
//		// ...
//		
//	}
func RespondServerErred(rw finger.ResponseWriter, request finger.Request) (e error) {
	if nil == rw {
		return errNilResponseWriter
	}

	const punctuation string = "!"
	const verb        string = "SERVER-ERRED"
	var   object      string = QuoteSentence(request.Sentence())

	var mrw MagicResponseWriter = NewMagicResponseWriter(rw, punctuation, verb, object)

	defer func(){
		if err := mrw.Close(); nil != err {
			if nil == e {
				e = fmt.Errorf("problem closing magic-finger \"%s%s %s\" connection to client: %w", punctuation, verb, object, err)
			}
		}
	}()

	if err := mrw.Flush(); nil != err {
		e = fmt.Errorf("problem sending magic-finger \"%s%s %s\" response to client: %w", punctuation, verb, object, err)
		return
	}

	return nil
}
