package magicfinger_test

import (
	"github.com/reiver/go-finger"
	"github.com/reiver/go-finger/magic"
	"github.com/reiver/go-finger/test"

	"testing"
)

func TestRespondServerRefused(t *testing.T) {

	const object string = "joeblow/something.txt"

	var buffer testfinger.TestConnectedWriteCloser
	var rw finger.ResponseWriter = finger.NewResponseWriter(&buffer)

	magicfinger.RespondServerRefused(object, rw)

	var actual   string = buffer.String()
	var expected string =
		"\xEF\xBB\xBF"+
		"Magic-Finger"      +"\r\n"+
		"!REFUSED "+ object +"\r\n"

	if expected != actual {
		t.Errorf("The actual magic-finger server-erred response is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
