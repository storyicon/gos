## What's new in `v1.2`

#### 1. Fixed some errors

`Splitter` is a component of `GOS` that attempts to determine whether a repository is private before a `go get` request is sent (It was just a simple head request), but after testing, it may caused some pull errors, like the following:      

```bash
$ GOS_DEBUG=1 gos get github.com/schollz/progressbar/v2
debug mode is on
local proxy run on: [::]:10736
upstream address: https://athens.azurefd.net
abs local: github.com/schollz/progressbar/v2
```
`abs local` means that `GOS` treats it as a private repository. But obviously, this judgment is wrong, because our `Splitter` model is too simple. So we temporarily removed `Splitter` until it was smart enough.

#### 2. WebAssembly support

Added support for [WebAssembly](https://webassembly.org/) in the `go cross` command, now you can do this:
```bash
gos cross main.go js all
```
or this:
```bash
gos cross main.go js wasm
```
It is roughly equivalent to:
```bash
GOOS=js GOARCH=wasm go build main.go
```
