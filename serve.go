package finger

import (
	"fmt"
	"net"
	"io"
)

// Serve accepts incoming connections from 'listener',
// deals with parsing and validating the finger-protocol request,
// and hands off the connection to 'handler' to handle.
//
//	var handler finger.Handler
//	
//	// ...
//	
//	err := finger.Serve(listener, handler)
//	if nil != err {
//		return err
//	}
func Serve(listener net.Listener, handler Handler) error {

	if nil == listener {
		return errNilListener
	}
	defer listener.Close()

	if nil == handler {
		return errNilHandler
	}

	for {
		conn, err := listener.Accept()
		if nil != err {
			return fmt.Errorf("finger-protocol server had problem accepting incoming connection: %w", err)
		}

		go handle(conn, handler)
	}

	return nil
}

func handle(conn net.Conn, handler Handler) error {

	if nil == conn {
		return errNilConnection
	}
	defer conn.Close()

	var request Request
	{
		requestline, err := ReadRequestLine(conn)
		if nil != err && io.EOF != err {
			return fmt.Errorf("problem reading request-line when trying to handle new finger-request: %w", err)
		}

		request, err = ParseRequest(requestline)
		if nil != err {
			return fmt.Errorf("problem parsing request-line when trying to handle new finger-request: %w", err)
		}
	}

	var responseWriter ResponseWriter = NewResponseWriter(conn)

	handler.HandleFinger(responseWriter, request)
	return nil
}
