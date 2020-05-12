# Sudo

<p>
  <a href="https://pkg.go.dev/github.com/zeebo/sudo"><img src="https://img.shields.io/badge/doc-reference-007d9b?logo=go&style=flat-square" alt="go.dev" /></a>
  <a href="https://goreportcard.com/report/github.com/zeebo/sudo"><img src="https://goreportcard.com/badge/github.com/zeebo/sudo?style=flat-square" alt="Go Report Card" /></a>
  <a href="https://sourcegraph.com/github.com/zeebo/sudo?badge"><img src="https://sourcegraph.com/github.com/zeebo/sudo/-/badge.svg?style=flat-square" alt="SourceGraph" /></a>
</p>

sudo is a package to make reflect more powerful (and dangerous).

It exports a single function,  `Sudo`, which when passed a `reflect.Value`, will return a new `reflect.Value` with the read-only restrictions removed.
