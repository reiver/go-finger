package finger

import (
	"fmt"
)

// User represents a finger-protocol user.
//
// For example, in this finger-protocol request:
//
//	"dariush@example.com@something.social\r\n"
//
// “dariush” is the user.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a User with code like
// the following:
//
//	var user finger.User
//	
//	// ...
//	
//	fmt.Printf("finger user: %#v", user)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a User, do something like the following:
//
//	var user finger.User
//	
//	// ...
//	
//	value, something := user.Unwrap()
//	if something {
//		// a value was set for the finger.User
//	} else {
//		// no value was set for the finger.User
//	}
//
// Note that this is unwrapping the finger.User optional-type.
type User struct {
	value string
	something bool
}

// EmptyUser is used to create a finger.User with nothing in it.
func EmptyUser() User {
	return User{}
}

// CreateUser is used to create a finger.User with something in it.
func CreateUser(value string) User {
	return User{
		value:value,
		something:true,
	}
}

// Unwrap is used to unwrap a finger.User.
//
//	var user finger.User
//	
//	// ...
//	
//	value, something := user.Unwrap()
//
// If finger.User is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.User is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver User) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var user finger.User = finger.CreateUser("dariush")
//	
//	// ...
//	
//	fmt.Printf("user = %#v", user)
//
//	// Output:
//	// user = finger.CreateUser("dariush")
//
// Also, for example:
//
//	var user finger.User = finger.EmptyUser()
//	
//	// ...
//	
//	fmt.Printf("user = %#v", user)
//
//	// Output:
//	// user = finger.EmptyUser()
func (receiver User) GoString() string {
	if !receiver.something {
		return "finger.EmptyUser()"
	}

	return fmt.Sprintf("finger.CreateUser(%#v)", receiver.value)
}
