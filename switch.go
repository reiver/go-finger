package finger

import (
	"fmt"
)

// Switch is used to hold a finger-protocol request ‘switch’.
//
// Note that Switch is implemented as an optional-type.
// (Which in other programming languages is called a option-type or a maybe-type.)
//
// ⁂
//
// Most people are probably only going to use Switch as part
// of finger.Request.
//
// I.e,.
//
//	func (receiver *MyFingerHandler) HandleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		fmt.Printf("finger request switch: %#v", request.Switch)
//		
//		// ...
//		
//		swtch, something := request.Switch.Unwrap()
//		
//		// ...
//		
//	}
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Switch with code like
// the following:
//
//	fmt.Printf("finger request switch: %#v", swtch)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Switch, do something like the following:
//
//	var swtch finger.Switch
//	
//	// ...
//	
//	value, something := swtch.Unwrap()
//	if something {
//		// a value was set for the finger.Switch
//	} else {
//		// no value was set for the finger.Switch
//	}
//
// Note that this is unwrapping the Switch optional-type.
//
// ⁂
//
// One type of finger-protocol request looks like:
//
//	"/W joeblow\r\n"
//
// Another type of finger-protocol request looks like:
//
//	"/W\r\n"
//
// (There are other types of finger-protocol requests, but they aren't relevant here.)
//
// The "/W" is a finger-protocol request ‘switch’.
// (It is actually the only finger-protocol request ‘switch’ mentioned in the
// IETF RFC-742 & IETF RFC-1288 specifications.)
//
// For each of these finger-protocol requests, what would be stored in the code (for finger.Switch) is:
//
//	var swtch finger.Switch = finger.SomeSwitch("W")
//
// Notice that only "W" is stored, and not the "/".
//
// ⁂
//
// Although the IETF RFCs for finger only mentions the “W” finger-protocol
// request ‘switch’ — if an incoming finger-protocol request has a different
// finger-protocol request ‘switch’, Switch can store it.
//
// For example, if an incoming finger-protocol request was:
//
//	"/PULL joeblow\r\n"
//
// Then Switch would logically be:
//
//	var swtch finger.Switch = finger.SomeSwitch("PULL")
//
// ⁂
//
// IETF RFCs say the following about finger-protocol request ‘switches’ —
//
// The older IETF RFC-742 says:
//
// “if "/W" (called the "Whois switch") also appears on the line given to an
//  ITS server, much fuller descriptions are returned.”
//
// And the newer IETF RFC-1288 says:
//
// “The token /W in the {Q1} or {Q2} query types SHOULD at best be interpreted
//  at the last RUIP to signify a higher level of verbosity in the user
//  information output, or at worst be ignored.”
type Switch struct {
	value string
	something bool
}

// NoSwitch is used to create a finger.Switch with nothing in it.
func NoSwitch() Switch {
	return Switch{}
}

// SomeSwitch is used to create a finger.Switch with something in it.
func SomeSwitch(value string) Switch {
	return Switch{
		value:value,
		something:true,
	}
}

// Unwrap is used to unwrap a finger.Switch.
//
//	value, something := swtch.Unwrap()
//
// If finger.Switch is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Switch is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Switch) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var swtch finger.Switch = finger.SomeSwitch("W")
//	
//	// ...
//	
//	fmt.Printf("swtch = %#v", swtch)
//
//	// Output:
//	// swtch = finger.SomeSwitch("W")
//
// Also, for example:
//
//	var swtch finger.Switch = finger.NoSwitch()
//	
//	// ...
//	
//	fmt.Printf("swtch = %#v", swtch)
//
//	// Output:
//	// swtch = finger.NoSwitch()
func (receiver Switch) GoString() string {
	if !receiver.something {
		return "finger.NoSwitch()"
	}

	return fmt.Sprintf("finger.SomeSwitch(%#v)", receiver.value)
}
