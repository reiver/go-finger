package finger

import (
	"strings"
)

func (receiver Request) Sentence() string {
	var buffer strings.Builder

	receiver.writeTo(&buffer)

	return buffer.String()
}
