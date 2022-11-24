package finger

// A Handler responds to a finger-protocol request.
type Handler interface {
	HandleFinger(ResponseWriter, Request)
}
