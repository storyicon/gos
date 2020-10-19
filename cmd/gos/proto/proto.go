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

package proto

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "sync"

    "github.com/storyicon/gos/pkg/concurrent"

    "github.com/hashicorp/go-multierror"

    "github.com/spf13/cobra"
)

// CmdProto is the command line for cross compilation
var CmdProto = &cobra.Command{
    Use:   "proto [file]",
    Short: "quick and easy compilation of proto files",
    Long: `
Usage:
    gos proto [file]

    [file] the proto file you wanna to compile, gos supports two additional wildcards: 
    "all" means compiling all proto files under the current folder (excluding subfolders);
    "all/all" means compiling all proto files in the current directory and all subdirectories
    
    - Compile a single file
    gos proto helloworld.proto

    - Compile all proto files under the current folder (excluding subfolders)
    gos proto all

    - Compile all proto files in the current directory and all subdirectories
    gos proto all/all
`,
    Args:               cobra.MinimumNArgs(1),
    DisableFlagParsing: true,
}

func init() {
    CmdProto.Run = Run
}

// GoPath is the go path
var GoPath = os.Getenv("GOPATH")

// GoPathSrc is the address of $GOPATH/src
var GoPathSrc = filepath.Join(GoPath, "src")

// Run command
func Run(cmd *cobra.Command, in []string) {
    args, flags := filterArgs(in)
    if len(args) == 0 {
        log.Println("no args provided")
        os.Exit(1)
        return
    }

    proto := args[0]

    var err error
    switch proto {
    case "all":
        err = WalkCurrent(flags)
    case "all/all":
        err = WalkIter(flags)
    default:
        err = Generate(proto, flags)
    }

    if err != nil {
        log.Println(err)
        os.Exit(1)
        return
    }
    log.Println("Good job, you are ready to go :)")
}

func filterArgs(in []string) (args []string, flags []string) {
    for _, arg := range in {
        if strings.HasPrefix("--", arg) {
            flags = append(flags, arg)
        } else {
            args = append(args, arg)
        }
    }
    return args, flags
}

// WalkIter is used to traverse the current folder and its subfolders,
// find all proto files and execute the generate command
func WalkIter(flags []string) error {
    dir, err := os.Getwd()
    if err != nil {
        return err
    }

    var errs error
    var lock sync.Mutex
    var c = concurrent.New(runtime.NumCPU())

    fatalErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return nil
        }
        c.Add(1)
        go func() {
            defer c.Done()
            ext := filepath.Ext(path)
            if ext != ".proto" {
                return
            }
            err = Generate(path, flags)
            if err != nil {
                lock.Lock()
                errs = multierror.Append(errs, err)
                lock.Unlock()
            }
        }()
        return nil
    })
    c.Wait()

    if fatalErr != nil {
        errs = multierror.Append(errs, fatalErr)
    }

    return errs
}

// WalkCurrent is used to traverse all proto files
// in the current folder and execute the generate command
func WalkCurrent(flags []string) error {
    dir, err := os.Getwd()
    if err != nil {
        return err
    }
    elements, err := ioutil.ReadDir(dir)
    if err != nil {
        return nil
    }

    var lock sync.Mutex
    var errs error
    var c = concurrent.New(runtime.NumCPU())
    for _, ele := range elements {
        if ele.IsDir() {
            continue
        }
        name := ele.Name()
        ext := filepath.Ext(name)

        if ext != ".proto" {
            continue
        }

        c.Add(1)
        go func(name string) {
            defer c.Done()
            err := Generate(name, flags)
            if err != nil {
                lock.Lock()
                defer lock.Unlock()
                errs = multierror.Append(errs, err)
            }
        }(name)
    }
    c.Wait()
    return errs
}

// Generate is used to execute the generate command for the specified proto file
func Generate(proto string, flags []string) error {
    path, name := filepath.Split(proto)
    fd := exec.Command("protoc", append([]string{
        "--proto_path=" + GoPath,
        "--proto_path=" + GoPathSrc,
        "--go_out=plugins=grpc:.",
        "--proto_path=.",
        name,
    }, flags...)...)
    stderr := &bytes.Buffer{}
    fd.Stdout = stderr
    fd.Stderr = stderr
    fd.Dir = path
    if err := fd.Run(); err != nil {
        err = fmt.Errorf("%s: %s", proto, stderr.String())
        return err
    }
    return nil
}
