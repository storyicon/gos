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

package tool

import (
    "os"

    "github.com/spf13/cobra"
    "github.com/storyicon/gos/pkg/util"
)

// CmdTool runs the go tool command identified by the arguments.
var CmdTool = &cobra.Command{
    Use:   "tool [-n] command [args...]",
    Short: "run specified go tool",
    Long: `
Tool runs the go tool command identified by the arguments.
With no arguments it prints the list of known tools.

The -n flag causes tool to print the command that would be
executed but not execute it.

For more about each tool command, see 'go doc cmd/<command>'.
`,
    DisableFlagParsing: true,
}

func init() {
    CmdTool.Run = func(cmd *cobra.Command, args []string) {
        fd := util.GetGoBinaryCMD("tool", args)
        fd.Env = util.GetEnvWithLocalProxy()
        fd.Stdout = os.Stdout
        fd.Stderr = os.Stderr
        util.RunCMDWithExit(fd)
    }
}
