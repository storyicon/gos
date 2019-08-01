// Copyright 2019 storyicon@foxmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cross

import (
    "log"
    "runtime"
    "sync"

    "github.com/hashicorp/go-multierror"
    "github.com/spf13/cobra"
)

// CmdCross is the command line for cross compilation
var CmdCross = &cobra.Command{
    Use:   "cross [common go build flags] [file] [darwin|linux|windows|freebsd|netbsd|all] [amd64|386|arm|s390x|mips|mipsle|mips64|mips64le|all]",
    Short: "agile and fast cross compiling",
    Long: `
Usage:
    gos cross [common go build flags] [-e] [go file] [os] [arch]

    [common go build flags] you can use any standard flag for go build here, such as /-tags="..."/-a/-o/..., as you would with native go build
    [-e] when you specify -e, compilation errors will be printed
    [file] the go file you wanna to build
    [os] the OS such as linux/darwin/windows/freebsd/netbsd/openbsd/android/dragonfly/nacl/solaris/plan9, you can also use "all" to compile all OS
    [arch] the Arch such as amd64/386/arm/arm64/s390x/mips/mipsle/mips64/mips64le, you can also use "all" to compile all Arch

    - Compile all platform
    gos cross main.go all all

    - Compile all linux arch
    gos cross main.go linux all

    - Compile windows/amd64
    gos cross main.go windows amd64

    - Compile with standard go build flags
    gos cross -tags="prod" -ldflags="-s -w" -a main.go all all

    - Compile with error info printed
    gos cross -e main.go all all
    `,
    DisableFlagParsing: true,
}

func init() {
    CmdCross.Run = Run
}

// Run command
func Run(cmd *cobra.Command, args []string) {
    if len(args) == 1 {
        if arg := args[0]; arg == "-h" || arg == "help" {
            printUsage()
            return
        }
    }

    options, err := NewOptions(args)
    if err != nil {
        log.Println(err)
        return
    }

    horses, err := options.GetCompileWorkhorse()
    if err != nil {
        log.Println(err)
        return
    }

    var wg sync.WaitGroup
    var lock sync.Mutex
    var errs error
    threads := make(chan bool, runtime.NumCPU())
    for _, horse := range horses {
        wg.Add(1)
        threads <- true
        go func(horse *Workhorse) {
            defer func() {
                wg.Done()
                <-threads
            }()
            err := horse.Compile()
            platform := horse.Platform.String()
            if err != nil {
                log.Printf("* %s: failed", platform)
                lock.Lock()
                defer lock.Unlock()
                errs = multierror.Append(errs, err)
                return
            }
            log.Printf("* %s: successed", platform)
        }(horse)
    }
    wg.Wait()
    if errs != nil && options.ShowErr {
        log.Println(errs)
    }
}

func printUsage() {
    log.Println(CmdCross.Long)
}
