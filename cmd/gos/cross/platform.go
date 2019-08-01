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

package cross

import (
    "fmt"
)

// Platform defines the compilation platform
type Platform struct {
    OS   string
    Arch string
}

func (p *Platform) String() string {
    return fmt.Sprintf("%s/%s", p.OS, p.Arch)
}

// Platforms defines a set of compilation platforms
type Platforms []Platform

var allPlatforms = Platforms{
    {"android", "arm"},
    {"android", "386"},

    {"darwin", "386"},
    {"darwin", "amd64"},
    {"darwin", "arm"},
    {"darwin", "arm64"},

    {"dragonfly", "386"},
    {"dragonfly", "amd64"},

    {"freebsd", "386"},
    {"freebsd", "amd64"},
    {"freebsd", "arm"},

    {"linux", "386"},
    {"linux", "s390x"},
    {"linux", "amd64"},
    {"linux", "arm"},
    {"linux", "arm64"},
    {"linux", "ppc64"},
    {"linux", "ppc64le"},
    {"linux", "mips"},
    {"linux", "mipsle"},
    {"linux", "mips64"},
    {"linux", "mips64le"},

    {"nacl", "amd64"},
    {"nacl", "amd64p32"},
    {"nacl", "arm"},

    {"netbsd", "386"},
    {"netbsd", "amd64"},
    {"netbsd", "arm"},

    {"openbsd", "386"},
    {"openbsd", "amd64"},

    {"plan9", "386"},
    {"plan9", "amd64"},
    {"plan9", "arm"},

    {"solaris", "amd64"},

    {"windows", "386"},
    {"windows", "amd64"},

    {"js", "wasm"},
}

// FilterByOS filters the platform by type of OS
func (p Platforms) FilterByOS(os string) (ps Platforms) {
    for _, v := range p {
        if v.OS == os {
            ps = append(ps, v)
        }
    }
    return ps
}

// FilterByArch filters the platform by type of arch
func (p Platforms) FilterByArch(arch string) (ps Platforms) {
    for _, v := range p {
        if v.Arch == arch {
            ps = append(ps, v)
        }
    }
    return ps
}
