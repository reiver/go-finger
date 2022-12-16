package finger

import (
	"github.com/reiver/go-fck"
)

const (
	errInternalError  = fck.Error("internal error")
	errNilConnection  = fck.Error("nil connection")
	errNilHandler     = fck.Error("nil handler")
	errNilListener    = fck.Error("nil listener")
	errNilReceiver    = fck.Error("nil receiver")
	errNilWriter      = fck.Error("nil writer")
	errSlashNotFound  = fck.Error("slash not found")
)
