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
// In Go code, these would become:
//
//	// "\r\n"
//	var request finger.Request = finger.Request{
//		Switch: finger.NoSwitch(),
//		Target: finger.NoTarget(),
//	}
//
//	// "joeblow\r\n"
//	var request finger.Request = finger.Request{
//		Switch: finger.NoSwitch(),
//		Target: finger.SomeTarget("joeblow"),
//	}
//
//	// "janedoe\r\n"
//	var request finger.Request = finger.Request{
//		Switch: finger.NoSwitch(),
//		Target: finger.SomeTarget("janedoe"),
//	}
//
//	// "/W joeblow\r\n"
//	var request finger.Request = finger.Request{
//		Switch: finger.SomeSwitch("W"),
//		Target: finger.SomeTarget("joeblow"),
//	}
//
//	// "/W janedoe\r\n"
//	var request finger.Request = finger.Request{
//		Switch: finger.SomeSwitch("W"),
//		Target: finger.SomeTarget("janedoe"),
//	}
type Request struct {
	Switch Switch
	Target Target
}

