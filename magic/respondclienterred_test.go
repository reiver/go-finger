package magicfinger_test

import (
	"github.com/reiver/go-finger"
	"github.com/reiver/go-finger/magic"
	"github.com/reiver/go-finger/test"

	"testing"
)

func TestRespondClientErred(t *testing.T) {

	const swtch  string = "/GET"
	const target string = "joeblow/something.txt"

	var request finger.Request = finger.CreateRequest(swtch, target)

	var conn testfinger.TestConnectedWriteCloser
	var rw finger.ResponseWriter = finger.NewResponseWriter(&conn)

	magicfinger.RespondClientErred(rw, request)

	var actual   string = conn.String()
	var expected string =
		"\xEF\xBB\xBF"+
		"Magic-Finger"                               +"\r\n"+
		"!CLIENT-ERRED {/GET joeblow/something.txt}" +"\r\n"+
		""                                           +"\r\n"

	if expected != actual {
		t.Errorf("The actual magic-finger cilent-erred response is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
