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

package cross

import (
	"io/ioutil"
	"runtime"

	"bytes"

	"fmt"

	"path"
	"strings"

	"os"

	"github.com/storyicon/gos/pkg/util"
)

// Workhorse is a compilation unit
type Workhorse struct {
	Package    string
	StandardGO []string
	Output     string
	Platform
}

// Compile is the highlight
func (h *Workhorse) Compile() error {
	env := append(util.GetEnvWithLocalProxy(),
		"GOOS="+h.OS,
		"GOARCH="+h.Arch,
		h.GetCGOEnv(),
	)
	// In order to avoid duplicate names
	output, err := h.GetRealOutput()
	if err != nil {
		return err
	}

	stderr := &bytes.Buffer{}
	args := append(h.StandardGO, "-o", output, h.Package)

	cmd := util.GetGoBinaryCMD("build", args)
	cmd.Env = env
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		err = fmt.Errorf("%s: %s", h.Platform.String(), stderr.String())
		return err
	}
	return nil
}

// GetCGOEnv is used to determine whether to enable CGO
// when there is no corresponding environment variable.
func (h *Workhorse) GetCGOEnv() string {
	Cgo := os.Getenv("CGO_ENABLED")
	if Cgo == "" {
		if runtime.GOOS == h.OS && runtime.GOARCH == h.Arch {
			Cgo = "1"
		} else {
			Cgo = "0"
		}
	}
	return fmt.Sprintf("CGO_ENABLED=%s", Cgo)
}

// GetRealOutput is used to generate the real output address
func (h *Workhorse) GetRealOutput() (string, error) {
	if h.Output == "" {
		ext := path.Ext(h.Package)
		if ext != ".go" {
			return "", ErrMissingGoFile
		}
		h.Output = strings.TrimRight(h.Package, ext)
	}

	output := strings.Join([]string{
		h.Output,
		h.OS,
		h.Arch,
	}, "_")

	if h.OS == "windows" {
		output += ".exe"
	}

	return output, nil
}
