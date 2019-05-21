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

package mod

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
)

// CmdMod provides access to operations on modules.
var CmdMod = &cobra.Command{
	Use:   "mod",
	Short: "module maintenance",
	Long: `
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.
`,
	DisableFlagParsing: true,
}

func init() {
	CmdMod.Run = func(cmd *cobra.Command, args []string) {
		fd := util.GetGoBinaryCMD("mod", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}
