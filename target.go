package finger

import (
	"fmt"
)

// Target is used to hold a finger-protocol request ‘target’.
//
// Note that Target is implemented as an optional-type.
// (Which in other programming languages is called a option-type or a maybe-type.)
//
// ⁂
//
// Most people are probably only going to use Target as part
// of finger.Request.
//
// I.e,.
//
//	func (receiver *MyFingerHandler) HandleFinger(rw finger.ResponseWriter, request finger.Request) {
//		
//		// ...
//		
//		fmt.Printf("finger request target: %#v", request.Target)
//		
//		// ...
//		
//		target, something := request.Target.Unwrap()
//		
//		// ...
//		
//	}
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Target with code like
// the following:
//
//	fmt.Printf("finger request target: %#v", target)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Target, do something like the following:
//
//	var target finger.Target
//	
//	// ...
//	
//	value, something := target.Unwrap()
//	if something {
//		// a value was set for the finger.Target
//	} else {
//		// no value was set for the finger.Target
//	}
//
// Note that this is unwrapping the Target optional-type.
//
// ⁂
//
// One type of finger-protocol request looks like:
//
//	"joeblow\r\n"
//
// Another type of finger-protocol request looks like:
//
//	"/W joeblow\r\n"
//
// (There are other types of finger-protocol requests, but they aren't relevant here.)
//
// For each of these finger-protocol requests, what would be stored in the code (for finger.Target) is:
//
//	var target finger.Target = finger.CreateTarget("joeblow")
type Target struct {
	value string
	something bool
}

// EmptyTarget is used to create a finger.Target with nothing in it.
func EmptyTarget() Target {
	return Target{}
}

// CreateTarget is used to create a finger.Target with something in it.
func CreateTarget(value string) Target {
	return Target{
		value:value,
		something:true,
	}
}

// Unwrap is used to unwrap a finger.Target.
//
//	value, something := target.Unwrap()
//
// If finger.Target is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Target is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Target) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var target finger.Target = finger.CreateTarget("dariush")
//	
//	// ...
//	
//	fmt.Printf("target = %#v", target)
//
//	// Output:
//	// target = finger.CreateTarget("dariush")
//
// Also, for example:
//
//	var target finger.Target = finger.EmptyTarget()
//	
//	// ...
//	
//	fmt.Printf("target = %#v", target)
//
//	// Output:
//	// target = finger.EmptyTarget()
func (receiver Target) GoString() string {
	if !receiver.something {
		return "finger.EmptyTarget()"
	}

	return fmt.Sprintf("finger.CreateTarget(%#v)", receiver.value)
}
