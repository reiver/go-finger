# testfinger

Package `testfinger` is provide tools to help test code that uses the `github.com/reiver/go-finger` package.

## Example

For example, imagine that you have finger handler function like the following:
```go
import "github.com/reiver/go-finger"

// ...

func handlerFunc(rw finger.ResponseWriter, request finger.Request) {

	// ...

}
```

To help you write tests for a handler function like this, you can use `testfinger`'s `testfinger.TestConnectedWriteCloser` type.
