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

package util

import (
	"net"
	"os"
	"os/exec"

	"strings"

	"github.com/storyicon/gos/pkg/meta"
)

// Prepend is used to prepend a string type element into []string
func Prepend(arr []string, ele string) []string {
	return append([]string{ele}, arr...)
}

func GetGoBinaryCMD(subcmd string, args []string) *exec.Cmd {
	binary := meta.GetGoBinaryPath()
	return exec.Command(binary, Prepend(args, subcmd)...)
}

func GetEnvWithLocalProxy() []string {
	proxy := meta.GetLocalProxyListenAddr()
	_, port, _ := net.SplitHostPort(proxy)
	return append(os.Environ(), "GOPROXY=http://127.0.0.1:"+port)
}

func GetEnvWithoutGoProxy() []string {
	return append(os.Environ(), "GOPROXY=")
}

func RestoreCMDArgs(args []string) []string {
	var r []string
	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
			if i := strings.Index(arg, "="); i != -1 {
				symbol := `"`
				if strings.Contains(arg[i:], `"`) {
					symbol = `'`
				}
				arg = arg[:i+1] + symbol + arg[i+1:] + symbol
			}
		}
		r = append(r, arg)
	}
	return r
}
