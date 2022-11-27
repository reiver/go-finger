package finger

import (
	"github.com/reiver/go-fck"
)

const (
	errEmptySwitch   = fck.Error("empty switch")
	errNilConnection = fck.Error("nil connection")
	errNilWriter     = fck.Error("nil writer")
	errSlashNotFound = fck.Error("slash not found")
)
