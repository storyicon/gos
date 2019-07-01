# Gos: Armed Golang 💪 `v1.1`
[![CircleCI](https://circleci.com/gh/storyicon/gos/tree/master.svg?style=svg)](https://circleci.com/gh/storyicon/gos/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/storyicon/gos)](https://goreportcard.com/report/github.com/storyicon/gos)  [![Build Status](https://travis-ci.org/storyicon/gos.svg?branch=master)](https://travis-ci.org/storyicon/gos) [![GoDoc](https://godoc.org/github.com/storyicon/gos?status.svg)](https://godoc.org/github.com/storyicon/gos) [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/storyicon/Lobby)

![gos](https://raw.githubusercontent.com/storyicon/gos/master/docs/screenshot/gos.png)

The current gos is still an alpha version, welcome more heroes to comment and improve it 🍓, you can add more commands to it, or modify something to make it perform better.

You can download the compiled binary program here: [Release Page](https://github.com/storyicon/gos/releases/)

  - [Brief introduction](#brief-introduction)
  - [How to start](#how-to-start)
  - [What GOS can do:](#what-gos-can-do)           
          - [1. Fully compatible with Go native commands](#1-fully-compatible-with-go-native-commands)           
          - [2. Simpler Cross-Compilation](#2-simpler-cross-compilation)           
          - [3. Rapid generation of .proto](#3-rapid-generation-of-proto)           
          - [4. Go proxy solution](#4-go-proxy-solution)           

## :beer: News
>  Please use golang with version >= `1.12`. [About `go: cannot find main module`](https://github.com/storyicon/gos/blob/master/docs/go-version.md)
> :astonished: Is there still a **dial-tcp-timeout**? Starting with v1.1, you can use GOS_UPSTREAM_ADDRESS to specify the upstream proxy address: [What's new in v1.1](https://github.com/storyicon/gos/blob/master/docs/something-new-in-v1.1.md) (2019-5-30)

>  🌷 Here's something about "GOPROXY supports comma-separated list of URLs": [On GOPROXY](https://github.com/storyicon/gos/blob/master/docs/on-goproxy.md)  (2019-5-23)


## 🦄 Brief introduction

from now on, use gos instead of go:

```bash
go get => gos get
go build => gos build
go run => gos run
go ... => gos ...
```

gos is compatible with all go commands and has go mod/get equipped with smart `GOPROXY`, 
it automatically distinguishes between private and public repositories 
and uses `GOPROXY` to download your lost package when appropriate.

gos has a few extra commands to enhance your development experience:

```bash
  cross      agile and fast cross compiling
  proto      quick and easy compilation of proto files
```

You can use `-h` on these sub commands to get more information.              

## 🐋 How to start

This can't be simpler.              
According to your system type, download the zip file from the [release page](https://github.com/storyicon/gos/releases/), unzip, rename the binaries to `gos` and put it in your `$PATH`. Then use `gos` as if you were using the `go` command.               
You can also download the source code and compile it using `go build -o gos main.go`                

Note: The prerequisite for gos to work properly is that the [go binary](https://golang.org/dl/) is in your $PATH. If you need to use the `gos proto` command, you need the [protoc binary](https://github.com/protocolbuffers/protobuf/releases) too.              

## :tangerine: What GOS can do: 

### 1. Fully compatible with Go native commands 

You can use `gos` just like you would with the `go` command. Compatible with all flags and arguments, such as the following:

```bash
go get -u -v github.com/xxxx/xxxx
=>
gos get -u -v github.com/xxxx/xxxx
```

### 2. Simpler Cross-Compilation

You can use `gos cross` command for simpler cross-compilation:

```bash
# Compile Linux platform binaries for the current system architecture
# For example, if your computer are amd64, it will compile main.go into the binary of linux/amd64 architecture.
gos cross main.go linux

# Specify the build platform and architecture
gos cross main.go linux amd64
gos cross main.go linux arm
gos cross main.go linux 386
gos cross main.go windows amd64
gos cross main.go darwin 386

# Compiling binary files for all architectures on the specified platform
gos cross main.go linux all
gos cross main.go windows all

# Compiling binary files for all platforms on the specified architecture
gos cross main.go all amd64

# Trying to compile binary files for all platforms and architectures
gos cross all all

# Compile with standard go build flags
gos cross -tags="prod" -ldflags="-s -w" -a main.go all all

# Compile with error info printed
gos cross -e main.go all all

# Compile with CGO enabled
CGO_ENABLED=1 gos cross all all
```

Gos uses parallel compilation, very fast 🚀, but still depends on the configuration of your operating system.

more information: `gos cross -h`

### 3. Rapid generation of .proto

This feature may only be useful to RPC developers. You can compile proto files more easily, as follows:

```bash
# Compile a single file
gos proto helloworld.proto

# Compile all proto files under the current folder (excluding subfolders)
gos proto all

# Compile all proto files in the current directory and all subdirectories
gos proto all/all
```

Of course, the precondition is that you have a [protoc binary](https://github.com/protocolbuffers/protobuf/releases) in your $PATH.

more information: `gos proto -h`

### 4. Go proxy solution

There is a dilemma here. If you don't use `GOPROXY`, there may be a large number of Package pull timeouts (network reasons) or non-existence (repository rename, delete or migrate), like the following:
```bash
unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: 
dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
```
```
go: google.golang.org/grpc@v1.21.0: unrecognized import path "google.golang.org/grpc" (https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
```

If use `GOPROXY`, you will not be able to pull the private repositories (github, gitlab, etc) properly, like that:
```
go get github.com/your_private_repo: unexpected status (https://athens.azurefd.net/github.com/your_private_repo/@v/list): 500 Internal Server Error
```

GOS strengthens all of GO's native commands, no matter it's go mod/get/build/run/....Any situation that might cause a package pull, gos will intelligently determine whether the current repository to be pulled needs to use `GOPROXY`.


**Now, live your thug life 😎**
