package finger

import (
	"fmt"
)

// Actor represents a finger-protocol actor.
//
// For example, in this finger-protocol request:
//
//	"dariush@example.com@something.social\r\n"
//
// “dariush” is the actor.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Actor with code like
// the following:
//
//	var actor finger.Actor
//	
//	// ...
//	
//	fmt.Printf("finger actor: %#v", actor)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Actor, do something like the following:
//
//	var actor finger.Actor
//	
//	// ...
//	
//	value, something := actor.Unwrap()
//	if something {
//		// a value was set for the finger.Actor
//	} else {
//		// no value was set for the finger.Actor
//	}
//
// Note that this is unwrapping the finger.Actor optional-type.
type Actor struct {
	value string
	something bool
}

// EmptyActor is used to create a finger.Actor with nothing in it.
func EmptyActor() Actor {
	return Actor{}
}

// CreateActor is used to create a finger.Actor with something in it.
func CreateActor(value string) Actor {
	return Actor{
		value:value,
		something:true,
	}
}

// Unwrap is used to unwrap a finger.Actor.
//
//	var actor finger.Actor
//	
//	// ...
//	
//	value, something := actor.Unwrap()
//
// If finger.Actor is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Actor is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Actor) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var actor finger.Actor = finger.CreateActor("dariush")
//	
//	// ...
//	
//	fmt.Printf("actor = %#v", actor)
//
//	// Output:
//	// actor = finger.CreateActor("dariush")
//
// Also, for example:
//
//	var actor finger.Actor = finger.EmptyActor()
//	
//	// ...
//	
//	fmt.Printf("actor = %#v", actor)
//
//	// Output:
//	// actor = finger.EmptyActor()
func (receiver Actor) GoString() string {
	if !receiver.something {
		return "finger.EmptyActor()"
	}

	return fmt.Sprintf("finger.CreateActor(%#v)", receiver.value)
}
