package magicfinger_test

import (
	"github.com/reiver/go-finger"
	"github.com/reiver/go-finger/magic"
	"github.com/reiver/go-finger/test"

	"testing"
)

func TestRespondServerErred(t *testing.T) {

	const object string = "joeblow/something.txt"

	var buffer testfinger.TestConnectedWriteCloser
	var rw finger.ResponseWriter = finger.NewResponseWriter(&buffer)

	magicfinger.RespondServerErred(object, rw)

	var actual   string = buffer.String()
	var expected string =
		"\xEF\xBB\xBF"+
		"Magic-Finger"    +"\r\n"+
		"!ERRED "+ object +"\r\n"

	if expected != actual {
		t.Errorf("The actual magic-finger server-erred response is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
