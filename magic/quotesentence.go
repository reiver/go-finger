package magicfinger

import (
	"strings"
)

func QuoteSentence(sentence string) string {

	var s string = sentence
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `{`, `\{`)
	s = strings.ReplaceAll(s, `}`, `\}`)

	var buffer strings.Builder

	buffer.WriteRune('{')
	buffer.WriteString(s)
	buffer.WriteRune('}')

	return buffer.String()
}
