package magicfinger_test

import (
	"github.com/reiver/go-finger"
	"github.com/reiver/go-finger/magic"
	"github.com/reiver/go-finger/test"

	"github.com/reiver/go-strfs"

	"io/fs"
	"strconv"
	"time"

	"testing"
)

func TestResponseFile(t *testing.T) {

	const filename = "something.txt"

	const s string = "Hello world!"+"\r\n"+"\r\n"+"0123456789"+"\r\n"
	var content strfs.Content = strfs.CreateContent(s)
	var regularfile strfs.RegularFile = strfs.RegularFile{
		FileContent: content,
		FileName:    filename,
		FileModTime: time.Now(),
	}
	var file fs.File = &regularfile

	const swtch  string = "/GET"
	const target string = "joeblow/something.txt"

	var request finger.Request = finger.CreateRequest(swtch, target)

	var conn testfinger.TestConnectedWriteCloser
	var rw finger.ResponseWriter = finger.NewResponseWriter(&conn)

	magicfinger.RespondServerSucceededFile(rw, request, file)

	var actual string = conn.String()
	var expected string =
		"\xEF\xBB\xBF"+
		"Magic-Finger"                                          +"\r\n"+
		"!SERVER-SUCCEEDED {/GET joeblow/something.txt}"        +"\r\n"+
		"Content-Length: "+strconv.FormatInt(int64(len(s)), 10) +"\r\n"+
		""                                                      +"\r\n"+
		s

	if expected != actual {
		t.Errorf("The actual magic-finger file response was not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
