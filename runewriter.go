package finger

type runeWriter interface {
	WriteRune(rune) (int, error)
}
