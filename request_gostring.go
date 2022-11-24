package finger

import (
	"strings"
)

// GoString returns the Go code equivalent of the construction of a finger.Request.
func (receiver Request) GoString() string {
	var buffer strings.Builder

	buffer.WriteString("finger.Request{")
	if NoSwitch() != receiver.Switch {
		buffer.WriteString(" Switch: ")
		buffer.WriteString(receiver.Switch.GoString())
	}
	if NoSwitch() != receiver.Switch && NoTarget() != receiver.Target {
		buffer.WriteRune(',')
	}
	if NoTarget() != receiver.Target {
		buffer.WriteString(" Target: ")
		buffer.WriteString(receiver.Target.GoString())
	}
	if NoSwitch() != receiver.Switch || NoTarget() != receiver.Target {
		buffer.WriteRune(' ')
	}
	buffer.WriteString("}")

	return buffer.String()
}
