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

package env

import (
	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
	"os"
)

var CmdEnv = &cobra.Command{
	Use:   "env [-json] [var ...]",
	Short: "print Go environment information",
	Long:  `
Env prints Go environment information.

By default env prints information as a shell script
(on Windows, a batch file). If one or more variable
names is given as arguments, env prints the value of
each named variable on its own line.

The -json flag prints the environment in JSON format
instead of as a shell script.

For more about environment variables, see 'go help environment'.
`,
	DisableFlagParsing: true,
}

func init(){
	CmdEnv.Run = func(cmd *cobra.Command, args []string){
		fd := util.GetGoBinaryCMD("env", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}
