package magicfinger

import (
	"github.com/reiver/go-finger"
)

// MagicResponseWriter is used to create "Magic-Finger" responses.
//
// A "Magic-Finger" rsponse has a magic-bytes prefix, and header fields.
//
// Note that this shoud not to be confused with the finger.ResponseWriter interface.
//
// One can promote a finger.ResponseWriter to become a magicfinger.MagicResponseWriter.
//
// For example:
//
//	var rw finger.ResponseWriter
//	
//	// ...
//	
//	var mrw magicfinger.MagicResponseWriter = magicfinger.NewMagicResponseWriter(rw)
type MagicResponseWriter interface {
	finger.ResponseWriter
	AddField(name string, body string) error
}
