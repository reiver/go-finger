package finger

import (
	"github.com/reiver/go-fck"
)

const (
	errNilConnection      = fck.Error("nil connection")
	errNilHandler         = fck.Error("nil handler")
	errNilListener        = fck.Error("nil listener")
	errNilResponseWriter  = fck.Error("nil response writer")
	errNilWriter          = fck.Error("nil writer")
	errSlashNotFound      = fck.Error("slash not found")
)
