package testfinger

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errClosed      = erorr.Error("closed")
	errNilReceiver = erorr.Error("nil receiver")
)
