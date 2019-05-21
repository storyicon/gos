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

package meta

import (
	"net"
	"sync/atomic"
)

var v atomic.Value

var defaultSysVar = SystemVar{
	GoBinaryPath:    "go",
	ProxyListenAddr: "",
	UpstreamAddr:    "https://athens.azurefd.net",
}

type SystemVar struct {
	GoBinaryPath    string
	ProxyListenAddr string
	UpstreamAddr    string
}

func LoadConfig(variables SystemVar) {
	v.Store(variables)
}

func GetConfig() SystemVar {
	return v.Load().(SystemVar)
}

func SetGoBinaryPath(path string) {
	config := GetConfig()
	config.GoBinaryPath = path
	LoadConfig(config)
}

func GetGoBinaryPath() string {
	config := GetConfig()
	return config.GoBinaryPath
}

func GetLocalProxyListenAddr() string {
	config := GetConfig()
	return config.ProxyListenAddr
}

func init() {
	address, err := allocateAddr()
	if err != nil {
		panic(err)
	}
	defaultSysVar.ProxyListenAddr = address
	LoadConfig(defaultSysVar)
}

func allocateAddr() (string, error) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer ln.Close()
	return ln.Addr().String(), nil
}
