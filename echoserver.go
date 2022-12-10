package finger

// EchoServer returns the the finger-client's request as a response.
//
// EchoServer is useful for debugging purposes.
//
//	finger.Serve(listener, finger.EchoServer())
func EchoServer() Handler {
	return echoServer(0)
}

var _ Handler = echoServer(0)

type echoServer int

func (echoServer) HandleFinger(responsewriter ResponseWriter, request Request) {

	defer responsewriter.Close()

	request.WriteTo(responsewriter)
}
