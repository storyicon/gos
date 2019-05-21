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

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/cmd/go/build"
	"github.com/storyicon/gos/cmd/go/clean"
	"github.com/storyicon/gos/cmd/go/doc"
	"github.com/storyicon/gos/cmd/go/env"
	"github.com/storyicon/gos/cmd/go/fix"
	"github.com/storyicon/gos/cmd/go/fmt"
	"github.com/storyicon/gos/cmd/go/generate"
	"github.com/storyicon/gos/cmd/go/get"
	"github.com/storyicon/gos/cmd/go/install"
	"github.com/storyicon/gos/cmd/go/list"
	"github.com/storyicon/gos/cmd/go/mod"
	"github.com/storyicon/gos/cmd/go/run"
	"github.com/storyicon/gos/cmd/go/test"
	"github.com/storyicon/gos/cmd/go/tool"
	"github.com/storyicon/gos/cmd/go/version"
	"github.com/storyicon/gos/cmd/go/vet"
	"github.com/storyicon/gos/cmd/gos/cross"
	"github.com/storyicon/gos/cmd/gos/proto"
)

// CmdRoot is the root command
var CmdRoot = cobra.Command{
	Use:   "ROOT Command",
	Short: "print introduction",
	Long: `
Gos: Armed Golang 💪

from now on, use gos instead of go:

go get => gos get
go build => gos build
go ... => gos ...

gos is compatible with all go commands and has go mod/get equipped with smart GOPROXY, 
it automatically distinguishes between private and public repositories 
and uses GOPROXY to download your lost package when appropriate.

gos has a few extra commands to enhance your development experience:

  cross      agile and fast cross compiling
  proto      quick and easy compilation of proto files

You can use -h on these sub commands to get more information.
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(cmd.Long)
	},
}

func init() {
	CmdRoot.AddCommand(
		// Standard GO
		build.CmdBuild,
		clean.CmdClean,
		doc.CmdDoc,
		env.CmdEnv,
		fix.CmdFix,
		fmt.CmdFmt,
		generate.CmdGenerate,
		get.CmdGet,
		install.CmdInstall,
		list.CmdList,
		mod.CmdMod,
		run.CmdRun,
		test.CmdTest,
		tool.CmdTool,
		version.CmdVersion,
		vet.CmdVet,

		// GOS
		cross.CmdCross,
		proto.CmdProto,
	)
}
