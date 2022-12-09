package finger

import (
	"fmt"
	"strings"
)

const (
	bufferLimit = 4194304 // 4 MB
)

// ParseRequest parses a string for a finger-protocol request.
func ParseRequest(s string) (Request, error) {

	if "" == s {
		return EmptyRequest(), nil
	}

	// Unicode text encode as UTF-8 can have a magic-bytes prefix.
	//
	//	0xEF 0xBB 0xBF
	//
	// Here we check for it.
	// And if it is there, remove it.
	if 3 <= len(s) {
		s0 := s[0]
		s1 := s[1]
		s2 := s[2]

		if 0xEF == s0 && 0xBB == s1 && 0xBF == s2 {
			s  = s[3:]
		}
	}

	var switchString string
	{
		if 0 < len(s) && '/' == s[0] {
			const SP rune = ' ' // SPACE

			var index int = strings.IndexRune(s, SP)


			if index < 0 {
				switchString = s
				s = ""
			} else {
				switchString = s[:index]
				s = s[1+index:]

				loop: for 0 < len(s) {
					s0 := s[0]

					switch s0 {
					case ' ':
						s = s[1:]
					default:
						break loop
					}
				}
			}
		}
	}

	var swtch Switch
	{
		if "" != switchString {
			var err error

			swtch, err = ParseSwitch(switchString)
			if nil != err {
				return EmptyRequest(), fmt.Errorf("problem parsing switch in request-line when parsing finger-protocol request: %w", err)
			}
		}
	}

	var target Target
	{
		if "" != s {
			target = CreateTarget(s)
		}
	}

	return Request{
		swtch:  swtch,
		target: target,
	}, nil
}
