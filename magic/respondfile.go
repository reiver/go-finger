package magicfinger

import (
	"io"
	"io/fs"
	"strconv"
)

func RespondFile(mrw MagicResponseWriter, file fs.File) error {
	if nil == mrw {
		return errNilMagicResponseWriter
	}
	if nil == file {
		return errNilFile
	}

	var fileinfo fs.FileInfo
	{
		var err error

		fileinfo, err = file.Stat()
		if nil != err {
			return RespondServerErred(mrw)
		}
		if nil == fileinfo {
			return RespondServerErred(mrw)
		}
	}

	var filemode fs.FileMode
	{
		filemode = fileinfo.Mode()
	}

	{
		if !filemode.IsRegular() {
			return RespondServerFailed(mrw)
		}
	}

	var contentLength int64
	{
		contentLength = fileinfo.Size()
	}

	{
		mrw.AddField("Content-Length", strconv.FormatInt(contentLength, 10))
	}

	{
		io.Copy(mrw, file)
	}

	return nil
}

