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

func TestModule_GetListAddr(t *testing.T) {
	type fields struct {
		version string
		addr    string
	}
	type args struct {
		base       string
		isUpstream bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/graphquery/@v/list",
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery",
			},
			args: args{
				base:       "/mnt/f/golang",
				isUpstream: false,
			},
			want:    "/mnt/f/golang/" + downloadPath + "/github.com/storyicon/graphquery/@v/list",
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				version: "v1.0.0",
				addr:    "/github.com/storyicon/graphquery",
			},
			args: args{
				base:       "https://goproxy.io/",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/graphquery/@v/list",
			wantErr: false,
		},
		{
			name: "test3",
			fields: fields{
				version: "v1.0.0",
				addr:    "/github.com/storyicon/graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/graphquery/@v/list",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		m := &Module{
			version: tt.fields.version,
			addr:    tt.fields.addr,
		}
		got, err := m.GetListAddr(tt.args.base, tt.args.isUpstream)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestModule_GetInfoAddr(t *testing.T) {
	type fields struct {
		version string
		addr    string
	}
	type args struct {
		base       string
		isUpstream bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/graphquery/@v/v1.0.0.info",
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery/",
			},
			args: args{
				base:       "/mnt/f/golang",
				isUpstream: false,
			},
			want:    "/mnt/f/golang/" + downloadPath + "/github.com/storyicon/graphquery/@v/v1.0.0.info",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		m := &Module{
			version: tt.fields.version,
			addr:    tt.fields.addr,
		}
		got, err := m.GetInfoAddr(tt.args.base, tt.args.isUpstream)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestModule_GetLatestAddr(t *testing.T) {
	type fields struct {
		version string
		addr    string
	}
	type args struct {
		base       string
		isUpstream bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/graphquery/@latest",
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/graphquery/",
			},
			args: args{
				base:       "/mnt/f/golang",
				isUpstream: false,
			},
			want:    "/mnt/f/golang/" + downloadPath + "/github.com/storyicon/graphquery/@latest",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		m := &Module{
			version: tt.fields.version,
			addr:    tt.fields.addr,
		}
		got, err := m.GetLatestAddr(tt.args.base, tt.args.isUpstream)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestModule_GetModAddr(t *testing.T) {
	type fields struct {
		version string
		addr    string
	}
	type args struct {
		base       string
		isUpstream bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/Graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/!graphquery/@v/v1.0.0.mod",
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/Graphquery/",
			},
			args: args{
				base:       "/mnt/f/golang",
				isUpstream: false,
			},
			want:    "/mnt/f/golang/" + downloadPath + "/github.com/storyicon/!graphquery/@v/v1.0.0.mod",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		m := &Module{
			version: tt.fields.version,
			addr:    tt.fields.addr,
		}
		got, err := m.GetModAddr(tt.args.base, tt.args.isUpstream)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestModule_GetZipAddr(t *testing.T) {
	type fields struct {
		version string
		addr    string
	}
	type args struct {
		base       string
		isUpstream bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/Graphquery",
			},
			args: args{
				base:       "https://goproxy.io",
				isUpstream: true,
			},
			want:    "https://goproxy.io/github.com/storyicon/!graphquery/@v/v1.0.0.zip",
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				version: "v1.0.0",
				addr:    "github.com/storyicon/Graphquery/",
			},
			args: args{
				base:       "/mnt/f/golang",
				isUpstream: false,
			},
			want:    "/mnt/f/golang/" + downloadPath + "/github.com/storyicon/!graphquery/@v/v1.0.0.zip",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		m := &Module{
			version: tt.fields.version,
			addr:    tt.fields.addr,
		}
		got, err := m.GetZipAddr(tt.args.base, tt.args.isUpstream)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
