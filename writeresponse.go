package finger

import (
	"fmt"
	"io"
)

const responseMagic = "\xEF\xBB\xBF"+ "# FINGER"+"\r\n"

const clientErredString   = responseMagic+ "!ERRED"   +" "+ "client" +"\r\n"+"\r\n"
const serverErredString   = responseMagic+ "!ERRED"   +" "+ "server" +"\r\n"+"\r\n"
const serverFailedString  = responseMagic+ "!FAILED"  +" "+ "server" +"\r\n"+"\r\n"
const serverRefusedString = responseMagic+ "!REFUSED" +" "+ "server" +"\r\n"+"\r\n"

// WriteResponseClientErred is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol client erred with its request somehow.
//
// This should only be used when the request is an invalid finger-protocol request.
//
// Not when the server doesn't have or doesn't want to give a response to something.
// In those cases, finger.WriteResponseServerRefused() should instead be used.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseClientErred(rw)
//		
//		// ...
//		
//	}
func WriteResponseClientErred(writer io.Writer) error {
	if nil == writer {
		return errNilResponseWriter
	}

	const s string = clientErredString

	n, err := io.WriteString(writer, s)
	if nil != err {
		return fmt.Errorf("problem writing cilent-error finger-protocol response: %w", err)
	}
	if expected, actual := len(s), n; expected != actual {
		return fmt.Errorf("problem writing cilent-error finger-protocol response: actually wrote %d bytes but expected to write %d bytes", actual, expected)
	}

	return nil
}

// WriteResponseServerErred is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server erred somehow when trying
// to handle the finger-protocol client's request.
//
// For example, maybe the server tried to read a file, but the hard drive is corrupted, and the
// operating system returned an error.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseServerErred(rw)
//		
//		// ...
//		
//	}
func WriteResponseServerErred(writer io.Writer) error {
	if nil == writer {
		return errNilResponseWriter
	}

	const s string = serverErredString

	n, err := io.WriteString(writer, s)
	if nil != err {
		return fmt.Errorf("problem writing server-error finger-protocol response: %w", err)
	}
	if expected, actual := len(s), n; expected != actual {
		return fmt.Errorf("problem writing server-error finger-protocol response: actually wrote %d bytes but expected to write %d bytes", actual, expected)
	}

	return nil
}

// WriteResponseServerFailed is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server failed somehow when trying
// to handle the finger-protocol client's request.
//
// For example, maybe the user the client is interested doesn't exist on the server.
//
//	func handleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseServerFailed(rw)
//		
//		// ...
//		
//	}
func WriteResponseServerFailed(writer io.Writer) error {
	if nil == writer {
		return errNilResponseWriter
	}

	const s string = serverFailedString

	n, err := io.WriteString(writer, s)
	if nil != err {
		return fmt.Errorf("problem writing server-error finger-protocol response: %w", err)
	}
	if expected, actual := len(s), n; expected != actual {
		return fmt.Errorf("problem writing server-error finger-protocol response: actually wrote %d bytes but expected to write %d bytes", actual, expected)
	}

	return nil
}

// WriteResponseServerRefused is a helper function that can be used by a finger-protocol server
// to tell a finger-protocol client that the finger-protocol server refused to handle the
// finger-protocol client's request.
//
// For example, maybe server the server won't handle the "/PULL" switch for a user.
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
func WriteResponseServerRefused(writer io.Writer) error {
	if nil == writer {
		return errNilResponseWriter
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

