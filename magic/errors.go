package magicfinger

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errBadFieldName      = erorr.Error("bad field name")
	errBadFieldBody      = erorr.Error("bad field body")
	errNilFile           = erorr.Error("nil file")
	errInternalError     = erorr.Error("internal error")
	errNilReceiver       = erorr.Error("nil receiver")
	errNilResponseWriter = erorr.Error("nil response writer")
	errNilWriter         = erorr.Error("nil writer")
)
