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
    "errors"
    "runtime"
    "strings"
)

// Here defines a set of standard errors
var (
    ErrTooManyArguments  = errors.New("too many arguments")
    ErrMissingGoFile     = errors.New("missing target .go file")
    ErrUnexpectedParams  = errors.New("unexpected list of parameters")
    ErrNoMatchedPlatform = errors.New("no matched platform")
)

// Options defines the structure of compilation options
type Options struct {
    // StandardGO is the standard go build flags and Arguments,
    // gos only forwards them without parsing and processing.

    // * Cross will parse an -e identifier to determine
    // * if a compilation error message needs to be printed. The default is not to print.
    StandardGO []string

    Output  string
    Package string
    Platform
    // * ShowErr will be true when the -e identifier is present
    ShowErr bool

    raw []string
}

// NewOptions is used to parse cmd args
func NewOptions(args []string) (*Options, error) {
    // set default
    options := &Options{
        Platform: Platform{
            OS:   runtime.GOOS,
            Arch: runtime.GOARCH,
        },
        raw: args,
    }

    var pos uint8
    for i := 0; i < len(args); i++ {
        arg := args[i]

        if pos == 0 {
            // * -e hook
            // parse show error identifier
            if arg == "-e" {
                options.ShowErr = true
                continue
            }

            // * -o hook
            if arg == "-o" {
                i++
                if i < len(args) {
                    options.Output = args[i]
                    continue
                }
                return nil, ErrUnexpectedParams
            }

            if strings.HasSuffix(arg, ".go") {
                pos++
                options.Package = arg
                continue
            }
            options.StandardGO = append(options.StandardGO, arg)
        } else {
            switch pos++; pos {
            case 2:
                options.OS = arg
            case 3:
                options.Arch = arg
            default:
                return nil, ErrTooManyArguments
            }
        }

    }

    if options.Package == "" {
        return nil, ErrMissingGoFile
    }

    return options, nil
}

// GetCompileWorkhorse is used to get all compilation tasks
func (o *Options) GetCompileWorkhorse() ([]*Workhorse, error) {
    platforms := allPlatforms
    if o.OS != "all" {
        platforms = platforms.FilterByOS(o.OS)
    }
    if o.Arch != "all" {
        platforms = platforms.FilterByArch(o.Arch)
    }

    var horses []*Workhorse

    if len(platforms) == 0 {
        return horses, ErrNoMatchedPlatform
    }

    for _, platform := range platforms {
        horses = append(horses, &Workhorse{
            Package:    o.Package,
            StandardGO: o.StandardGO,
            Output:     o.Output,
            Platform:   platform,
        })
    }
    return horses, nil
}
