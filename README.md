# go-finger

Package **finger** implements the **finger-protocol**, for the Go programming language.

I.e., generally, what is defined in:

* IETF RFC-742 — https://datatracker.ietf.org/doc/html/rfc742
* IETF RFC-1288 — https://datatracker.ietf.org/doc/html/rfc1288

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-finger

[![GoDoc](https://godoc.org/github.com/reiver/go-finger?status.svg)](https://godoc.org/github.com/reiver/go-finger)

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
