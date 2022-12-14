package magicfinger

import (
	"github.com/reiver/go-fck"
)

const (
	errBadFieldName           = fck.Error("bad field name")
	errBadFieldBody           = fck.Error("bad field body")
	errNilFile                = fck.Error("nil file")
	errNilMagicResponseWriter = fck.Error("nil magic response writer")
	errNilReceiver            = fck.Error("nil receiver")
	errNilResponseWriter      = fck.Error("nil response writer")
	errNilWriter              = fck.Error("nil writer")
)
