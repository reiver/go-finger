package finger

// A Request is used to hold a finger-protocol request.
//
// Likely, most people would only use a finger.Request when writing
// a finger.Handler. For example:
//
//	func (receiver *MyFingerHandler) HandleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//	}
//
// Over the networks a raw finger-protocol request can look like these:
//
//	"\r\n"
//
//	"joeblow\r\n"
//
//	"janedoe\r\n"
//
//	"/W joeblow\r\n"
//
//	"/W janedoe\r\n"
//
//	"/W\r\n"
//
// In Go code, these would become:
//
//	// "\r\n"
//	var request finger.Request = finger.EmptyRequest()
//
//	// "joeblow\r\n"
//	var request finger.Request = finger.CreateRequestTarget("joeblow")
//
//	// "janedoe\r\n"
//	var request finger.Request = finger.CreateRequestTarget("janedoe")
//
//	// "/W joeblow\r\n"
//	var request finger.Request = finger.CreateRequest("W", "joeblow")
//
//	// "/W janedoe\r\n"
//	var request finger.Request = finger.CreateRequest("W", "janedoe")
//
//	// "/W\r\n"
//	var request finger.Request = finger.CreateRequestSwitch("W")
type Request struct {
	swtch Switch
	target Target
}

func AssembleRequest(swtch Switch, target Target) Request {
	return Request{
		swtch: swtch,
		target: target,
	}
}

func AssembleRequestSwitch(swtch Switch) Request {
	return Request{
		swtch: swtch,
	}
}

func AssembleRequestTarget(target Target) Request {
	return Request{
		target: target,
	}
}

// EmptyRequest is used to create an empty finger.Request.
// I.e., with no finger-protocol request user or finger-protocol request target.
func EmptyRequest() Request {
	return Request{}
}

// CreateRequest is used to create a finger.Request with a finger-protocol request switch, and
// a finger-protocol request target.
//
// For example, a call like this:
//
//	var request finger.Request = finger.CreateRequest("W", "dariush")
//
// Is equivalent to the (raw) finger-protocol request:
//
//	"/W dariush\r\n"
func CreateRequest(swtch string, target string) Request {
	return Request{
		swtch: SomeSwitch(swtch),
		target: SomeTarget(target),
	}
}

// CreateRequestSwitch is used to create a finger.Request with just a finger-protocol request switch
// (but no finger-protocol request target).
//
// For example, a call like this:
//
//	var request finger.Request = finger.CreateRequestSwitch("W")
//
// Is equivalent to the (raw) finger-protocol request:
//
//	"/W\r\n"
func CreateRequestSwitch(swtch string) Request {
	return Request{
		swtch: SomeSwitch(swtch),
	}
}

// CreateRequestTarget is used to create a finger.Request with just a finger-protocol request target
// (but no finger-protocol request switch).
//
// For example, a call like this:
//
//	var request finger.Request = finger.CreateRequestTarget("dariush")
//
// Is equivalent to the (raw) finger-protocol request:
//
//	"dariush\r\n"
func CreateRequestTarget(target string) Request {
	return Request{
		target: SomeTarget(target),
	}
}
