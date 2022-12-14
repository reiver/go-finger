package magicfinger

type runeWriter interface {
	WriteRune(rune) (int, error)
}
