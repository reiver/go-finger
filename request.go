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

// AssembleRequest assembles a finger.Request from a finger.Switch and a finger.Target.
//
// For example:
//
//	var request finger.Request = finger.AssembleRequest("W", "joeblow")
//
// This example is the equivalent of the (raw) finger-protocol request:
//
//	"/W joeblow\r\n"
func AssembleRequest(swtch Switch, target Target) Request {
	return Request{
		swtch: swtch,
		target: target,
	}
}

// AssembleRequest assembles a finger.Request from a finger.Switch.
//
// For example:
//
//	var request finger.Request = finger.AssembleRequest("W")
//
// This example is the equivalent of the (raw) finger-protocol request:
//
//	"/W\r\n"
func AssembleRequestSwitch(swtch Switch) Request {
	return Request{
		swtch: swtch,
	}
}

// AssembleRequest assembles a finger.Request from a finger.Target.
//
// For example:
//
//	var request finger.Request = finger.AssembleRequest("joeblow")
//
// This example is the equivalent of the (raw) finger-protocol request:
//
//	"joeblow\r\n"
func AssembleRequestTarget(target Target) Request {
	return Request{
		target: target,
	}
}

// EmptyRequest is used to create an empty finger.Request.
// I.e., with no finger-protocol request user or finger-protocol request target.
//
// For example:
//
//	var request finger.Request = finger.EmptyRequest()
//
// This example is the equivalent of the (raw) finger-protocol request:
//
//	"\r\n"
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
		target: CreateTarget(target),
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
		target: CreateTarget(target),
	}
}
