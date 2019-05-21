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

package version


import (
	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
	"os"
)

var CmdVersion = &cobra.Command{
	Use:   "version",
	Short: "print Go version",
	Long:  `Version prints the Go version, as reported by runtime.Version.`,
	DisableFlagParsing: true,
}

func init(){
	CmdVersion.Run = func(cmd *cobra.Command, args []string){
		fd := util.GetGoBinaryCMD("version", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}

