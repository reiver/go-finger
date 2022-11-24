package finger

import (
	"strings"
)

// String requests the raw finger-protocol request.
func (receiver Request) String() string {
	var buffer strings.Builder

	receiver.WriteTo(&buffer)

	return buffer.String()
}
