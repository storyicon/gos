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

package vet

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
)

// CmdVet runs the Go vet command on the packages named by the import paths
var CmdVet = &cobra.Command{
	Use:   "vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]",
	Short: "report likely mistakes in packages",
	Long: `
Vet runs the Go vet command on the packages named by the import paths.

For more about vet and its flags, see 'go doc cmd/vet'.
For more about specifying packages, see 'go help packages'.
For a list of checkers and their flags, see 'go tool vet help'.
For details of a specific checker such as 'printf', see 'go tool vet help printf'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

The -vettool=prog flag selects a different analysis tool with alternative
or additional checks.
For example, the 'shadow' analyzer can be built and run using these commands:

  go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
  go vet -vettool=$(which shadow)

The build flags supported by go vet are those that control package resolution
and execution, such as -n, -x, -v, -tags, and -toolexec.
For more about these flags, see 'go help build'.

See also: go fmt, go fix.
`,
	DisableFlagParsing: true,
}

func init() {
	CmdVet.Run = func(cmd *cobra.Command, args []string) {
		fd := util.GetGoBinaryCMD("vet", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}
