# Sudo

[![GoDoc](https://godoc.org/github.com/zeebo/sudo?status.svg)](https://godoc.org/github.com/zeebo/sudo)
[![Sourcegraph](https://sourcegraph.com/github.com/zeebo/sudo/-/badge.svg)](https://sourcegraph.com/github.com/zeebo/sudo?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeebo/sudo)](https://goreportcard.com/report/github.com/zeebo/sudo)

sudo is a package to make reflect more powerful (and dangerous).

It exports a single function,  `Sudo`, which when passed a `reflect.Value`, will return a new `reflect.Value` with the read-only restrictions removed.
