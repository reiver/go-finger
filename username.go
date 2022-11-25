package finger

import (
	"fmt"
)

// UserName represents a finger-protocol username.
//
// For example, in this finger-protocol request:
//
//	"dariush@example.com@something.social\r\n"
//
// “dariush” is the username.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a UserName with code like
// the following:
//
//	var username finger.UserName
//	
//	// ...
//	
//	fmt.Printf("finger username: %#v", username)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a UserName, do something like the following:
//
//	var username finger.UserName
//	
//	// ...
//	
//	value, something := username.Unwrap()
//	if something {
//		// a value was set for the finger.UserName
//	} else {
//		// no value was set for the finger.UserName
//	}
//
// Note that this is unwrapping the finger.UserName optional-type.
type UserName struct {
	value string
	something bool
}

// NoUserName is used to create a finger.UserName with nothing in it.
func NoUserName() UserName {
	return UserName{}
}

// SomeUserName is used to create a finger.UserName with something in it.
func SomeUserName(value string) UserName {
	return UserName{
		value:value,
		something:true,
	}
}

// Unwrap is used to unwrap a finger.UserName.
//
//	var username finger.UserName
//	
//	// ...
//	
//	value, something := username.Unwrap()
//
// If finger.UserName is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.UserName is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver UserName) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var username finger.UserName = finger.SomeUserName("dariush")
//	
//	// ...
//	
//	fmt.Printf("username = %#v", username)
//
//	// Output:
//	// username = finger.SomeUserName("dariush")
//
// Also, for example:
//
//	var username finger.UserName = finger.NoUserName()
//	
//	// ...
//	
//	fmt.Printf("username = %#v", username)
//
//	// Output:
//	// username = finger.NoUserName()
func (receiver UserName) GoString() string {
	if !receiver.something {
		return "finger.NoUserName()"
	}

	return fmt.Sprintf("finger.SomeUserName(%#v)", receiver.value)
}
