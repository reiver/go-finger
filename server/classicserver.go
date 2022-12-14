package finger

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

	var username string
	{
		var something bool

		username, something = query.User().Unwrap()
		if !something {
			WriteResponseServerRefused(responsewriter)
/////////////////////// RETURN
			return
		}
	}

	WriteResponseDotPlan(responsewriter, username)
}
