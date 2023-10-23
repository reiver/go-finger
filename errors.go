package finger

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	ErrClientErred   = erorr.Error("client erred")
	ErrServerErred   = erorr.Error("server erred")
	ErrServerFailed  = erorr.Error("server failed")
	ErrServerRefused = erorr.Error("server refused")
)

const (
	errInternalError  = erorr.Error("internal error")
	errNilConnection  = erorr.Error("nil connection")
	errNilHandler     = erorr.Error("nil handler")
	errNilListener    = erorr.Error("nil listener")
	errNilReceiver    = erorr.Error("nil receiver")
	errNilWriter      = erorr.Error("nil writer")
	errSlashNotFound  = erorr.Error("slash not found")
)
