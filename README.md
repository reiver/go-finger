# go-finger

Package **finger** implements the **finger-protocol**, for the Go programming language.

I.e., generally, what is defined in:

* IETF RFC-742 — https://datatracker.ietf.org/doc/html/rfc742
* IETF RFC-1288 — https://datatracker.ietf.org/doc/html/rfc1288

Although this package add some modernizations that are compatible with IETF RFC-742,
and in the spirt of the **finger-protocol**,
_that you can choose whether to use or not use_.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-finger

[![GoDoc](https://godoc.org/github.com/reiver/go-finger?status.svg)](https://godoc.org/github.com/reiver/go-finger)

## Simple Finger Client

This is an example of how to create very very very simple **finger-protocol** client using this package.

It is meant for educational purposes.
A more useful **finger-protocol** client would include a lot more features.
But anyway....

We first need to decide what the **finger-protocol** target will be.

What is a **finger-protocol** "target"‽

Well, for example, if we were to run the **finger** program as:

```
finger joeblow@example.com
```

Then out target would be:

```go
"joeblow"
```

So, in our Go code we would have something like:

```go
package main

import (
	"github.com/reiver/go-finger"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	//@TODO
}
```

Now we need to put the `target` inside of a **finger-protocol** request.

We do that with the following code:

```go
package main

import (
	"github.com/reiver/go-finger"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	//@TODO
}
```

Next, we need to figure out what is the TCP-address of the **finger-protocol server** our **finger-protocol client** will connect to.

TCP-addresses (as a string) look something like:

	"example.com:79"

	"reiver.link:1971"

	"181.70.250.13:1079"

	"12.23.34.45:7979"

I.e., it is a 'host' and a TCP-port (with a colon ":" between them).

We are using the following as our reference example:

```
finger joeblow@example.com
```

So what is our "address"‽

It cannot be just this:
```go
"example.com"
```

Since it doesn't include a TCP-port.

But, as it turns out, this package can handle this situation.
When the TCP-port (in the address) is missing, it defaults it to the default **finger-protocol** TCP-port.
(Which is TCP-port 79.)

So, now our code becomes:

```go
package main

import (
	"github.com/reiver/go-finger"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	var address finger.Address = finger.CreateAddressHost("example.com")

	//@TODO
}
```

Now we need to connect to the server.
We will do that using `net.Dial()`.

So....

```go
package main

import (
	"github.com/reiver/go-finger"

	"fmt"
	"net"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	var address finger.Address = finger.CreateAddressHost("example.com")

	conn, err := net.Dial("tcp", address.Resolve())
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	defer conn.Close()

	//@TODO
}
```

And now we need to create client that will use this TCP connection to make a **finger-protocol request**.

So, this will create the client:

```go
package main

import (
	"github.com/reiver/go-finger"

	"fmt"
	"net"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	var address finger.Address = finger.CreateAddressHost("example.com")

	conn, err := net.Dial("tcp", address.Resolve())
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	defer conn.Close()

	var client finger.Client = finger.AssembleClient(conn)

	//@TODO
}
```

And then this will send the **finger-protocol request**.

```go
package main

import (
	"github.com/reiver/go-finger"

	"fmt"
	"net"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	var address finger.Address = finger.CreateAddressHost("example.com")

	conn, err := net.Dial("tcp", address.Resolve())
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	defer conn.Close()

	var client finger.Client = finger.AssembleClient(conn)

	responseReader, err := client.Do(request)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}

	//@TODO
}
```

And finally, this will send the **finger-protocol response** this program received from the **finger-protocol server** to STDOUT (so that we can see it):

```go
package main

import (
	"github.com/reiver/go-finger"

	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	var target finger.Target = finger.CreateTarget("joeblow")

	var request finger.Request = finger.AssembleRequestTarget(target)

	var address finger.Address = finger.CreateAddressHost("example.com")

	conn, err := net.Dial("tcp", address.Resolve())
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}
	defer conn.Close()

	var client finger.Client = finger.AssembleClient(conn)

	responseReader, err := client.Do(request)
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}

	io.Copy(os.Stdout, responseReader)
}
```

## Unicde UTF-8

**This package supports Unicode UTF-8 text encoding for the finger-protocol.**

The finger-protocol existed **before** any IETF RFC was written about it, but —

There are 2 IETF RFCs that are relavent for the finger-protocol:

* IETF RFC-742 (released in 1977) — https://datatracker.ietf.org/doc/html/rfc742
* IETF RFC-1288 (released in 1991) — https://datatracker.ietf.org/doc/html/rfc1288

IETF RFC-742 does _not_ directly or indirectly forbid the use of Unicode UTF-8 — because IETF RFC-742 does _not_ seem to explicitly specify what character-set should be used.
IETF RFC-1288 indirectly forbids the use of Unicode UTF-8.

**This package sides with the original IETF RFC-742, and supports Unicode UTF-8 text.**

This package does this (supports the Unicode UTF-8 encoding of text) to provide a _more modern_ implementation of the finger-protocol.

Someone in the future can write a new IETF RFC for the finger-protocol that gives permission to use the Unicode UTF-8 text encoding.

## Import

To import package **finger** use import code like the follownig:

```
import "github.com/reiver/go-finger"
```

## Installation

To install package **finger** do the following:

```
GOPROXY=direct go get https://github.com/reiver/go-finger
```

## Author

Package **finger** was written by [Charles Iliya Krempeaux](http://reiver.link/)
