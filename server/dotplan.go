package finger

import (
	"io"
	"os"
	osuser "os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// WriteResopnseDotPlan is a helper function that can be used by a finger-protocol server
// to response with something similar to the classic .plan file response for a user.
//
// For example:
//
//	func handleFinger(responsewriter finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		err := finger.WriteResponseDotPlan(resopnsewriter, username)
//		
//		// ...
//		
//	}
func WriteResponseDotPlan(writer io.Writer, user string) error {

	var homepath string
	var realname string
	{
		u, err := osuser.Lookup(user)
		if nil != err {
			switch err.(type) {
			case osuser.UnknownUserError:
				WriteResponseServerFailed(writer)
/////////////////////////////// RETURN
				return nil
			default:
				WriteResponseServerErred(writer)
/////////////////////////////// RETURN
				return nil
			}
		}
		if nil == u {
			WriteResponseServerErred(writer)
/////////////////////// RETURN
			return nil
		}

		homepath = u.HomeDir
		if "" == homepath {
			WriteResponseServerFailed(writer)
/////////////////////// RETURN
			return nil
		}

		realname = u.Name
	}

	var reader io.Reader
	var contentLength int64
	{
		var planpath string = filepath.Join(homepath, ".plan")

		planfile, err := os.Open(planpath)
		if nil != err {
			switch {
			case os.IsNotExist(err):
				WriteResponseServerFailed(writer)
			default:
				WriteResponseServerErred(writer)
			}
/////////////////////// RETURN
			return nil

		}
		defer planfile.Close()

		planinfo, err := planfile.Stat()
		if nil != err {
			WriteResponseServerErred(writer)
/////////////////////// RETURN
			return nil
		}
		if planinfo.IsDir() {
			WriteResponseServerFailed(writer)
/////////////////////// RETURN
			return nil
		}

		reader = planfile
		contentLength = planinfo.Size()
	}

	var header strings.Builder
	{
		header.WriteString(magic)

		header.WriteString("Name: ")
		header.WriteString(realname)
		header.WriteString("\r\n")

		header.WriteString("User-Name: ")
		header.WriteString(user)
		header.WriteString("\r\n")

		header.WriteString("Content-Length: ")
		header.WriteString(strconv.FormatInt(contentLength, 10))
		header.WriteString("\r\n")

		header.WriteString("\r\n")
	}

	{
		io.WriteString(writer, header.String())
		io.Copy(writer, reader)
	}

	return nil
}
