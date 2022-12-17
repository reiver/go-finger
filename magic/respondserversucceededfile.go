package magicfinger

import (
	"github.com/reiver/go-finger"

	"fmt"
	"io"
	"io/fs"
	"strconv"
)

func RespondServerSucceededFile(rw finger.ResponseWriter, request finger.Request, file fs.File) (e error) {
	if nil == rw {
		return errNilResponseWriter
	}

	const punctuation string = "!"
	const verb        string = "SERVER-SUCCEEDED"
	var   object      string = QuoteSentence(request.Sentence())

	var mrw MagicResponseWriter = NewMagicResponseWriter(rw, punctuation, verb, object)
	if nil == mrw {
		return RespondServerErred(rw, request)
		e = fmt.Errorf("problem creating magic-finger response-writer: %w", errInternalError)
	}

	defer func(){
		if err := mrw.Close(); nil != err {
			if nil == e {
				e = fmt.Errorf("problem closing magic-finger \"%s%s %s\" connection to client: %w", punctuation, verb, object, err)
			}
		}
	}()

	if nil == file {
		return errNilFile
	}


	var fileinfo fs.FileInfo
	{
		var err error

		fileinfo, err = file.Stat()
		if nil != err {
			return RespondServerErred(rw, request)
		}
		if nil == fileinfo {
			return RespondServerErred(rw, request)
		}
	}

	var filemode fs.FileMode
	{
		filemode = fileinfo.Mode()
	}

	{
		if !filemode.IsRegular() {
			return RespondServerFailed(rw, request)
		}
	}

	var contentLength int64
	{
		contentLength = fileinfo.Size()
	}

	{
		mrw.AddField("Content-Length", strconv.FormatInt(contentLength, 10))
//@TODO: Content-Digest
	}

	{
		io.Copy(mrw, file)
	}

	return nil
}

