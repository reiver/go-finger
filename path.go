package finger

import (
	"fmt"
)

// Path represents a finger-protocol path.
//
// For example, in this finger-protocol request:
//
//	"dariush/img/ani.gif@example.com@something.social\r\n"
//
// “/img/ani.gif” is a path.
//
// ⁂
//
// For debugging, one can see the value (or lack of a value) of a Path with code like
// the following:
//
//	var path finger.Path
//	
//	// ...
//	
//	fmt.Printf("finger path: %#v", path)
//
// Notice that the verb being used is “%#v” (and not just “%v”).
//
// ⁂
//
// To get a value out of a Path, do something like the following:
//
//	var path finger.Path
//	
//	// ...
//	
//	value, something := path.Unwrap()
//	if something {
//		// a value was set for the finger.Path
//	} else {
//		// no value was set for the finger.Path
//	}
//
// Note that this is unwrapping the finger.Path optional-type.
type Path struct {
	value string
	something bool
}

// EmptyPath is used to create a finger.Path with nothing in it.
func EmptyPath() Path {
	return Path{}
}

// CreatePath is used to create a finger.Path with something in it.
func CreatePath(value string) Path {
	return Path{
		value:value,
		something:true,
	}
}

// GoString makes it so that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var path finger.Path = finger.CreatePath("/path/to/file.ext")
//	
//	// ...
//	
//	fmt.Printf("path = %#v", path)
//
//	// Output:
//	// path = finger.CreatePath("/path/to/file.ext")
//
// Also, for example:
//
//	var path finger.Path = finger.EmptyPath()
//	
//	// ...
//	
//	fmt.Printf("path = %#v", path)
//
//	// Output:
//	// path = finger.EmptyPath()
func (receiver Path) GoString() string {
	if !receiver.something {
		return "finger.EmptyPath()"
	}

	return fmt.Sprintf("finger.CreatePath(%#v)", receiver.value)
}

// Set sets the value of a finger.Path.
//
// Set mainly exists so that finger.Path can be as a flag.Value, and thus be used with functions
// such as flag.Var().
func (receiver *Path) Set(value string) error {
	if nil == receiver {
		return nil
	}

	*receiver = CreatePath(value)
	return nil
}

// String returns the value of a finger.Path.
//
// Note that if finger.Path is empty, then it returns the default finger Path,
// which is ""..
//
// With String you cannot tell the difference between a finger.Path with a value of "",
// and an empty finger.Path.
func (receiver Path) String() string {
	if !receiver.something {
		return ""
	}

	return receiver.value
}

// Unwrap is used to unwrap a finger.Path.
//
//	var path finger.Path
//	
//	// ...
//	
//	value, something := path.Unwrap()
//
// If finger.Path is holding something, then ‘something’ (in the code above) is ‘true’.
//
// If finger.Path is holding nothing, then ‘something’ (in the code above) is ‘false’.
func (receiver Path) Unwrap() (string, bool) {
	return receiver.value, receiver.something
}
