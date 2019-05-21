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

package run

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
)

// CmdRun run compiles and runs the named main Go package
var CmdRun = &cobra.Command{
	Use:   "run [build flags] [-exec xprog] package [arguments...]",
	Short: "compile and run Go program",
	Long: `
Run compiles and runs the named main Go package.
Typically the package is specified as a list of .go source files,
but it may also be an import path, file system path, or pattern
matching a single known package, as in 'go run .' or 'go run my/cmd'.

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog:
	'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the system
default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_nacl_386_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.

The exit status of Run is not the exit status of the compiled binary.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build.
`,
	DisableFlagParsing: true,
}

func init() {
	CmdRun.Run = func(cmd *cobra.Command, args []string) {
		fd := util.GetGoBinaryCMD("run", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}
