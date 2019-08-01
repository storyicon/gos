/*
 * Copyright 2019 storyicon@foxmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package clean

import (
    "os"

    "github.com/spf13/cobra"
    "github.com/storyicon/gos/pkg/util"
)

// CmdClean removes object files from package source directories
var CmdClean = &cobra.Command{
    Use:   "clean [clean flags] [build flags] [packages]",
    Short: "remove object files and cached files",
    Long: `
Clean removes object files from package source directories.
The go command builds most objects in a temporary directory,
so go clean is mainly concerned with object files left by other
tools or by manual invocations of go build.

Specifically, clean removes the following files from each of the
source directories corresponding to the import paths:

    _obj/            old object directory, left from Makefiles
    _test/           old test directory, left from Makefiles
    _testmain.go     old gotest file, left from Makefiles
    test.out         old test log, left from Makefiles
    build.out        old test log, left from Makefiles
    *.[568ao]        object files, left from Makefiles

    DIR(.exe)        from go build
    DIR.test(.exe)   from go test -c
    MAINFILE(.exe)   from go build MAINFILE.go
    *.so             from SWIG

In the list, DIR represents the final path element of the
directory, and MAINFILE is the base name of any Go source
file in the directory that is not included when building
the package.

The -i flag causes clean to remove the corresponding installed
archive or binary (what 'go install' would create).

The -n flag causes clean to print the remove commands it would execute,
but not run them.

The -r flag causes clean to be applied recursively to all the
dependencies of the packages named by the import paths.

The -x flag causes clean to print remove commands as it executes them.

The -cache flag causes clean to remove the entire go build cache.

The -testcache flag causes clean to expire all test results in the
go build cache.

The -modcache flag causes clean to remove the entire module
download cache, including unpacked source code of versioned
dependencies.

For more about build flags, see 'go help build'.

For more about specifying packages, see 'go help packages'.
`,
    DisableFlagParsing: true,
}

func init() {
    CmdClean.Run = func(cmd *cobra.Command, args []string) {
        fd := util.GetGoBinaryCMD("clean", args)
        fd.Env = util.GetEnvWithLocalProxy()
        fd.Stdout = os.Stdout
        fd.Stderr = os.Stderr
        util.RunCMDWithExit(fd)
    }
}
