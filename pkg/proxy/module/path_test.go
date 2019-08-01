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
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewPath(t *testing.T) {
    module, err := NewPath("/<module>/@v/<version>.zip")
    assert.Equal(t, nil, err)
    assert.Equal(t, true, module != nil)

    module, err = NewPath("/<module>/<version>.zip")
    assert.Equal(t, ErrInvalidPath, err)
    assert.Equal(t, true, module == nil)
}

func TestModulePath_getSegments(t *testing.T) {
    module, err := NewPath("/<module>/@v/<version>.zip")
    assert.Equal(t, nil, err)
    assert.Equal(t, []string{
        "<module>",
        "<version>.zip",
    }, module.getSegments())
}

func TestModulePath_getLastSegment(t *testing.T) {
    module, err := NewPath("/<module>/@v/<version>.zip")
    assert.Equal(t, nil, err)
    assert.Equal(t, "<version>.zip", module.getLastSegment())
}

func TestModulePath_GetModulePathType(t *testing.T) {
    tests := []struct {
        rawPath string
        want    PathType
        wantErr error
    }{
        {
            rawPath: "/<module>/@v/<version>.zip",
            want:    TypePathZip,
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.mod",
            want:    TypePathMod,
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.info",
            want:    TypePathInfo,
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/list",
            want:    TypePathList,
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.rar",
            want:    TypePathUnknown,
            wantErr: ErrUnknownPathType,
        },
        {
            rawPath: "/<module>/@v/<version>",
            want:    TypePathUnknown,
            wantErr: ErrUnknownPathType,
        },
        {
            rawPath: "/<module>/@v/@latest",
            want:    TypePathLatest,
            wantErr: nil,
        },
    }
    for i, tt := range tests {
        m, err := NewPath(tt.rawPath)
        assert.Equal(t, tt.wantErr, err, i)
        if err == nil {
            got := m.GetType()
            assert.Equal(t, tt.want, got, i)
        }
    }
}

func TestModulePath_GetModuleVersion(t *testing.T) {
    tests := []struct {
        rawPath string
        want    string
        wantErr error
    }{
        {
            rawPath: "/<module>/@v/<version>.zip",
            want:    "<version>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.mod",
            want:    "<version>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.info",
            want:    "<version>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/list",
            want:    "",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.zip",
            want:    "<version>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/@latest",
            want:    "latest",
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        m, err := NewPath(tt.rawPath)
        assert.Equal(t, tt.wantErr, err)
        got := m.GetModVersion()
        assert.Equal(t, tt.want, got)
    }
}

func TestModulePath_GetModuleRemoteAddr(t *testing.T) {
    tests := []struct {
        rawPath string
        want    string
        wantErr error
    }{
        {
            rawPath: "/<vcs>/<author>/<project>/@v/<version>.zip",
            want:    "<vcs>/<author>/<project>",
            wantErr: nil,
        },
        {
            rawPath: "/<vcs>/<group>/<author>/<project>/@v/<version>.mod",
            want:    "<vcs>/<group>/<author>/<project>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.info",
            want:    "<module>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/list",
            want:    "<module>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/<version>.zip",
            want:    "<module>",
            wantErr: nil,
        },
        {
            rawPath: "/<module>/@v/@latest",
            want:    "<module>",
            wantErr: nil,
        },
    }
    for i, tt := range tests {
        m, err := NewPath(tt.rawPath)
        assert.Equal(t, tt.wantErr, err, i)
        got := m.GetModAddr()
        assert.Equal(t, tt.want, got, i)
    }
}
