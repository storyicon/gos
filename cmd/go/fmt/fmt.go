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

package fmt

import (
    "os"

    "github.com/spf13/cobra"
    "github.com/storyicon/gos/pkg/util"
)

// CmdFmt runs the command 'gofmt -l -w' on the packages named by the import paths.
var CmdFmt = &cobra.Command{
    Use:   "fmt [-n] [-x] [packages]",
    Short: "gofmt (reformat) package sources",
    Long: `
Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths. It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
`,
    DisableFlagParsing: true,
}

func init() {
    CmdFmt.Run = func(cmd *cobra.Command, args []string) {
        fd := util.GetGoBinaryCMD("fmt", args)
        fd.Env = util.GetEnvWithLocalProxy()
        fd.Stdout = os.Stdout
        fd.Stderr = os.Stderr
        util.RunCMDWithExit(fd)
    }
}
