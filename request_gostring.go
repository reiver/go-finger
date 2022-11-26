package finger

import (
	"fmt"
)

// GoString returns the Go code equivalent of the construction of a finger.Request.
func (receiver Request) GoString() string {


	swtch,  swtchIsSomething  := receiver.swtch.Unwrap()
	target, targetIsSomething := receiver.target.Unwrap()

	switch {
	case swtchIsSomething && targetIsSomething:
		return fmt.Sprintf("finger.CreateRequest(%q, %q)", swtch, target)
	case swtchIsSomething && !targetIsSomething:
		return fmt.Sprintf("finger.CreateRequestSwitch(%q)", swtch)
	case !swtchIsSomething && targetIsSomething:
		return fmt.Sprintf("finger.CreateRequestTarget(%q)", target)
	default:
		return "finger.EmptyRequest()"
	}
}
