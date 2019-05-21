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

type Module struct {
	version string
	addr    string
}

func NewModule(addr, version string) *Module {
	return &Module{
		addr:    decodePath(cleanPath(addr)),
		version: version,
	}
}

func (m *Module) GetAddrWithVersion() string {
	return strings.Join([]string{
		m.addr, "@", m.version,
	}, "")
}


func (m *Module) GetVersion() string {
	return m.version
}

func (m *Module) GetAddr() string {
	return m.addr
}

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

func (m *Module) GetListAddr(base string, isUpstream bool) (string, error) {
	return m.getModuleAddr(base, TypePathList, isUpstream)
}

func (m *Module) GetInfoAddr(base string, isUpstream bool) (string, error) {
	return m.getModuleAddr(base, TypePathInfo, isUpstream)
}

func (m *Module) GetLatestAddr(base string, isUpstream bool) (string, error) {
	return m.getModuleAddr(base, TypePathLatest, isUpstream)
}

func (m *Module) GetModAddr(base string, isUpstream bool) (string, error) {
	return m.getModuleAddr(base, TypePathMod, isUpstream)
}

func (m *Module) GetZipAddr(base string, isUpstream bool) (string, error) {
	return m.getModuleAddr(base, TypePathZip, isUpstream)
}

func (m *Module) GetStorageInfo(base string) (io.ReadCloser, error) {
	return m.getStorage(base, m.GetInfoAddr)
}

func (m *Module) GetStorageMod(base string) (io.ReadCloser, error) {
	return m.getStorage(base, m.GetModAddr)
}
func (m *Module) GetStorageZip(base string) (io.ReadCloser, error) {
	return m.getStorage(base, m.GetZipAddr)
}

func (m *Module) getStorage(base string, addrFunc func(string,bool)(string, error)) (io.ReadCloser, error){
	addr, err := addrFunc(base, false)
	if err != nil {
		return nil, err
	}
	return os.Open(addr)
}

func (m *Module) getModuleAddr(base string, t PathType, isUpstream bool) (rel string,err error) {
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