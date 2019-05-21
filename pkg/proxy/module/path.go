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
	"errors"
	"fmt"
	"path"
	"strings"
	"unicode/utf8"
)

// PathType defines the path type
type PathType uint8

// Defined a common set of module path types for Module Path
const (
	TypePathUnknown PathType = iota
	TypePathList
	TypePathLatest
	TypePathInfo
	TypePathMod
	TypePathZip
)

// Defined a common set of error types for Module Path
var (
	ErrUnknownPathType = errors.New("unknown path type")
	ErrInvalidPath     = errors.New("invalid path")
	ErrInvalidPathType = errors.New("invalid path type")
	ErrUnknownVersion  = errors.New("unknown version")
)

// Path is used to parse the request path and generate the module info.
type Path struct {
	Module
	pType PathType

	raw      string
	segments []string
}

// NewPath is used to create a new path instance
func NewPath(raw string) (*Path, error) {
	p := &Path{
		raw: cleanPath(raw),
	}
	if err := p.prepare(); err != nil {
		return nil, err
	}
	return p, nil
}

func (m *Path) prepare() error {
	switch m.GetType() {
	case TypePathUnknown:
		return ErrUnknownPathType
	case TypePathLatest:
		return nil
	default:
		if segments := m.getSegments(); len(segments) < 2 {
			return ErrInvalidPath
		}
	}
	m.Module = *NewModule(m.GetModAddr(), m.GetModVersion())
	return nil
}

// GetType is used to get the path type of specified module path
func (m *Path) GetType() PathType {
	if m.pType != 0 {
		return m.pType
	}
	switch ext := path.Ext(m.raw); ext {
	case ".info":
		m.pType = TypePathInfo
	case ".mod":
		m.pType = TypePathMod
	case ".zip":
		m.pType = TypePathZip
	default:
		m.pType = TypePathUnknown
		if strings.HasSuffix(m.raw, "/@v/list") {
			m.pType = TypePathList
		}
		if strings.HasSuffix(m.raw, "@latest") {
			m.pType = TypePathLatest
		}
	}
	return m.pType
}

// GetModVersion is used to get the module version of specified module path
func (m *Path) GetModVersion() string {
	if m.version != "" {
		return m.version
	}
	segment := m.getLastSegment()
	length := len(segment)
	switch pathType := m.GetType(); pathType {
	case TypePathList:
		break
	case TypePathInfo:
		m.version = segment[:length-5]
	case TypePathMod, TypePathZip:
		m.version = segment[:length-4]
	case TypePathLatest:
		m.version = "latest"
	default:
		break
	}
	return m.version
}

// GetModAddr is used to get the module address of specified module path
func (m *Path) GetModAddr() string {
	if m.addr == "" {
		i := strings.Index(m.raw, "/@v/")
		if i == -1 {
			i = strings.Index(m.raw, "/@latest")
			if i == -1 {
				return ""
			}
		}
		m.addr = m.raw[:i]
	}
	return m.addr
}

func (m *Path) getSegments() []string {
	if len(m.segments) == 0 {
		m.segments = strings.Split(m.raw, "/@v/")
	}
	return m.segments
}

func (m *Path) getLastSegment() string {
	segments := m.getSegments()
	return segments[len(segments)-1]
}

func cleanPath(addr string) string {
	return strings.Trim(addr, "/")
}

func joinPath(addrs ...string) string {
	var r []string
	for i, addr := range addrs {
		if addr == "" {
			continue
		}
		if i != 0 {
			addr = strings.TrimLeft(addr, "/")
		}
		r = append(r, strings.TrimRight(addr, "/"))
	}
	return strings.Join(r, "/")
}

func encodePath(addr string) (string, error) {
	haveUpper := false
	for _, r := range addr {
		if r == '!' || r >= utf8.RuneSelf {
			// This should be disallowed by CheckPath, but diagnose anyway.
			// The correctness of the encoding loop below depends on it.
			return "", fmt.Errorf("internal error: inconsistency in EncodePath")
		}
		if 'A' <= r && r <= 'Z' {
			haveUpper = true
		}
	}

	if !haveUpper {
		return addr, nil
	}

	var buf []byte
	for _, r := range addr {
		if 'A' <= r && r <= 'Z' {
			buf = append(buf, '!', byte(r+'a'-'A'))
		} else {
			buf = append(buf, byte(r))
		}
	}
	return string(buf), nil
}

func decodePath(addr string) string {
	var buf []byte

	bang := false
	for _, r := range addr {
		if r >= utf8.RuneSelf {
			return ""
		}
		if bang {
			bang = false
			if r < 'a' || 'z' < r {
				return ""
			}
			buf = append(buf, byte(r+'A'-'a'))
			continue
		}
		if r == '!' {
			bang = true
			continue
		}
		if 'A' <= r && r <= 'Z' {
			return ""
		}
		buf = append(buf, byte(r))
	}
	if bang {
		return ""
	}
	return string(buf)
}
