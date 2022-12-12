package finger

import (
	"io"
	"os"
	osuser "os/user"
	"path/filepath"
	"strings"
)

// ClassicServer works similar to classic finger servers.
// And supports user .plan files, etc.
// As well as adding some modernizations.
//
//	finger.Serve(listener, finger.ClassicServer())
func ClassicServer() Handler {
	return classicServer(0)
}

var _ Handler = classicServer(0)

type classicServer int

func (classicServer) HandleFinger(responsewriter ResponseWriter, request Request) {

	defer responsewriter.Close()

	target, targetIsSomething := request.Target().Unwrap()

	if !targetIsSomething {
		WriteResponseServerRefused(responsewriter)
/////////////// RETURN
		return
	}

	query, err := ParseQuery(target)
	if nil != err {
		WriteResponseClientErred(responsewriter)
/////////////// RETURN
		return
	}

	if 0 < query.LenAddresses() {
		WriteResponseServerRefused(responsewriter)
/////////////// RETURN
		return
	}

	var user string
	{
		var something bool

		user, something = query.User().Unwrap()
		if !something {
			WriteResponseServerRefused(responsewriter)
/////////////////////// RETURN
			return
		}
	}

	var homepath string
	var realname string
	{
		u, err := osuser.Lookup(user)
		if nil != err {
			switch err.(type) {
			case osuser.UnknownUserError:
				WriteResponseServerFailed(responsewriter)
/////////////////////////////// RETURN
				return
			default:
				WriteResponseServerErred(responsewriter)
/////////////////////////////// RETURN
				return
			}
		}
		if nil == u {
			WriteResponseServerErred(responsewriter)
/////////////////////// RETURN
			return
		}

		homepath = u.HomeDir
		if "" == homepath {
			WriteResponseServerFailed(responsewriter)
/////////////////////// RETURN
			return
		}

		realname = u.Name
	}


	var reader io.Reader
	{
		var planpath string = filepath.Join(homepath, ".plan")

		planfile, err := os.Open(planpath)
		if nil != err {
			switch {
			case os.IsNotExist(err):
				WriteResponseServerFailed(responsewriter)
			default:
				WriteResponseServerErred(responsewriter)
			}
/////////////////////// RETURN
			return

		}
		defer planfile.Close()

		planinfo, err := planfile.Stat()
		if nil != err {
			WriteResponseServerErred(responsewriter)
/////////////////////// RETURN
			return
		}
		if planinfo.IsDir() {
			WriteResponseServerFailed(responsewriter)
/////////////////////// RETURN
			return
		}

		reader = planfile
	}

	{
		var header strings.Builder

		header.WriteString("Name: ")
		header.WriteString(realname)
		header.WriteString("\r\n")
		header.WriteString("\r\n")

		io.WriteString(responsewriter, header.String())
		io.Copy(responsewriter, reader)
/////////////// RETURN
		return
	}
}
