package finger

import (
	"github.com/reiver/go-fck"
)

const (
	errNilListener    = fck.Error("nil listener")
	errNilReceiver    = fck.Error("nil receiver")
	errNilWriter      = fck.Error("nil writer")
	errSlashNotFound  = fck.Error("slash not found")
)
