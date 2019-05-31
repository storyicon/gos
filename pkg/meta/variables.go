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
	"os"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

// Key of Environmental Variables
const (
	EnvGosUpstreamAddress = "GOS_UPSTREAM_ADDRESS"
	EnvGosDebug           = "GOS_DEBUG"
)

var v atomic.Value

var defaultSysVar = SystemVar{
	GoBinaryPath:    "go",
	ProxyListenAddr: "",
	UpstreamAddr:    "https://athens.azurefd.net",
}

// SystemVar defines the structure of system variables
type SystemVar struct {
	GoBinaryPath    string
	ProxyListenAddr string
	UpstreamAddr    string
}

// LoadConfig is used to load variables info to system variables
func LoadConfig(variables SystemVar) {
	v.Store(variables)
}

// GetConfig is used to obtain the current config
func GetConfig() SystemVar {
	return v.Load().(SystemVar)
}

// SetUpstreamAddr is used to modify the upstream address of gos
func SetUpstreamAddr(path string) {
	config := GetConfig()
	config.UpstreamAddr = path
	LoadConfig(config)
}

// SetGoBinaryPath is used to modify the location of go binary files
func SetGoBinaryPath(path string) {
	config := GetConfig()
	config.GoBinaryPath = path
	LoadConfig(config)
}

// GetGoBinaryPath is used to get the GoBinaryPath currently configured
func GetGoBinaryPath() string {
	config := GetConfig()
	return config.GoBinaryPath
}

// GetLocalProxyListenAddr is used to get the ProxyListenAddr currently configured
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

	upstream := os.Getenv(EnvGosUpstreamAddress)
	if upstream != "" {
		defaultSysVar.UpstreamAddr = upstream
	}

	debug := os.Getenv(EnvGosDebug)
	if debug != "" {
		log.SetLevel(log.DebugLevel)
		log.Debugln("debug mode is on")
	}

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
