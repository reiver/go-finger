package finger

import (
	"github.com/reiver/go-finger"
)

// ClassicServer works similar to classic finger servers.
// And supports user .plan files, etc.
// As well as adding some modernizations.
//
//	finger.Serve(listener, finger.ClassicServer())
func ClassicServer(responsewriter finger.ResponseWriter, request finger.Request) {

	defer responsewriter.Close()

	var target string
	{
		var something bool

		target, something = request.Target().Unwrap()

		if !something {
			RespondServerRefused(responsewriter, request)
/////////////////////// RETURN
			return
		}
	}

	var query finger.Query
	{
		var query finger.Query

		query, err := finger.ParseQuery(target)
		if nil != err {
			RespondClientErred(responsewriter, request)
/////////////////////// RETURN
			return
		}

		if 0 < query.LenAddresses() {
			RespondServerRefused(responsewriter, request)
/////////////////////// RETURN
			return
		}
	}

	var username string
	{
		var something bool

		username, something = query.Actor().Unwrap()
		if !something {
			RespondServerRefused(responsewriter, request)
/////////////////////// RETURN
			return
		}
	}

	RespondServerSucceededDotPlan(responsewriter, username)
}
