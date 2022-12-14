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

	const s string = "Hello world!"+"\r\n"+"\r\n"+"0123456789"+"\r\n"
	var content strfs.Content = strfs.CreateContent(s)
	var regularfile strfs.RegularFile = strfs.RegularFile{
		FileContent: content,
		FileName:    "something.txt",
		FileModTime: time.Now(),
	}
	var file fs.File = &regularfile


	var buffer testfinger.TestConnectedWriteCloser
	var rw finger.ResponseWriter = finger.NewResponseWriter(&buffer)
	var mrw magicfinger.MagicResponseWriter = magicfinger.NewMagicResponseWriter(rw)

	magicfinger.RespondFile(mrw, file)

	{
		var expected string = "\xEF\xBB\xBF"+
		                      "Magic-Finger"                                          +"\r\n"+
		                      ""                                                      +"\r\n"+
		                      "Content-Length: "+strconv.FormatInt(int64(len(s)), 10) +"\r\n"+
		                      ""                                                      +"\r\n"+
		                      s
		var actual   string = buffer.String()

		if expected != actual {
			t.Errorf("The actual magic-finger file response was not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
