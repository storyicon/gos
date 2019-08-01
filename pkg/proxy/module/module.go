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

package module

import (
    "io"
    "net"
    "os"
    "strings"
)

const downloadPath = "pkg/mod/cache/download"

// Module defines the structure of package module
type Module struct {
    version string
    addr    string
}

// NewModule is used to init a module
func NewModule(addr, version string) *Module {
    return &Module{
        addr:    decodePath(cleanPath(addr)),
        version: version,
    }
}

// GetAddrWithVersion is used to get the module name with version
func (m *Module) GetAddrWithVersion() string {
    return strings.Join([]string{
        m.addr, "@", m.version,
    }, "")
}

// GetVersion is used to get the module version
func (m *Module) GetVersion() string {
    return m.version
}

// GetAddr is used to get the module address
func (m *Module) GetAddr() string {
    return m.addr
}

// GetDomain is used to get the domain of remote repo address
func (m *Module) GetDomain() string {
    i := strings.Index(m.addr, "/")
    if i == -1 {
        return m.addr
    }
    addr := m.addr[:i]
    domain, _, err := net.SplitHostPort(addr)
    if err != nil {
        return addr
    }
    return domain
}

// GetListAddr is used to get list address
func (m *Module) GetListAddr(base string, isUpstream bool) (string, error) {
    return m.getModuleAddr(base, TypePathList, isUpstream)
}

// GetInfoAddr is used to get info address
func (m *Module) GetInfoAddr(base string, isUpstream bool) (string, error) {
    return m.getModuleAddr(base, TypePathInfo, isUpstream)
}

// GetLatestAddr is used to get latest address
func (m *Module) GetLatestAddr(base string, isUpstream bool) (string, error) {
    return m.getModuleAddr(base, TypePathLatest, isUpstream)
}

// GetModAddr is used to get mod address
func (m *Module) GetModAddr(base string, isUpstream bool) (string, error) {
    return m.getModuleAddr(base, TypePathMod, isUpstream)
}

// GetZipAddr is used to get zip address
func (m *Module) GetZipAddr(base string, isUpstream bool) (string, error) {
    return m.getModuleAddr(base, TypePathZip, isUpstream)
}

// GetStorageInfo is used to get the info of module from the local storage
func (m *Module) GetStorageInfo(base string) (io.ReadCloser, error) {
    return m.getStorage(base, m.GetInfoAddr)
}

// GetStorageMod is used to get the mod of module from the local storage
func (m *Module) GetStorageMod(base string) (io.ReadCloser, error) {
    return m.getStorage(base, m.GetModAddr)
}

// GetStorageZip is used to get the zip file of module from the local storage
func (m *Module) GetStorageZip(base string) (io.ReadCloser, error) {
    return m.getStorage(base, m.GetZipAddr)
}

func (m *Module) getStorage(base string, addrFunc func(string, bool) (string, error)) (io.ReadCloser, error) {
    addr, err := addrFunc(base, false)
    if err != nil {
        return nil, err
    }
    return os.Open(addr)
}

func (m *Module) getModuleAddr(base string, t PathType, isUpstream bool) (rel string, err error) {
    var n string
    var sep = "@v"
    switch t {
    case TypePathList:
        n = "list"
    case TypePathLatest:
        n = "@latest"
        sep = ""
    case TypePathInfo, TypePathMod, TypePathZip:
        if m.version == "" {
            return "", ErrUnknownPathType
        }
        switch t {
        case TypePathInfo:
            n = m.version + ".info"
        case TypePathMod:
            n = m.version + ".mod"
        case TypePathZip:
            n = m.version + ".zip"
        }
    default:
        return "", ErrUnknownPathType
    }

    rel = joinPath(m.addr, sep, n)

    rel, err = encodePath(rel)
    if err != nil {
        return "", err
    }

    if !isUpstream {
        base = joinPath(base, downloadPath)
    }

    return joinPath(base, rel), nil
}
