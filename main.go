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

/*
Pacakge main is the starting point of gos program

from now on, use gos instead of go:

	go get => gos get
	go build => gos build
	go run => gos run
	go ... => gos ...

gos is compatible with all go commands and has go mod/get equipped with smart GOPROXY, it automatically distinguishes between private and public repositories and uses GOPROXY to download your lost package when appropriate.

gos has a few extra commands to enhance your development experience:

	cross      agile and fast cross compiling
	proto      quick and easy compilation of proto files
	...

You can use -h on these sub commands to get more information.

Project address: https://github.com/storyicon/gos
*/
package main

import (
	"github.com/storyicon/gos/cmd"
	"github.com/storyicon/gos/pkg/proxy"
)

func main() {
	go proxy.Default().Run()
	cmd.CmdRoot.Execute()
}
