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

package proxy

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/json-iterator/go"
	"github.com/storyicon/gos/pkg/proxy/module"
	"github.com/storyicon/gos/pkg/util"
)

// Fetcher defined the interface of GORPOXY api
type Fetcher interface {
	List(*module.Module) (io.ReadCloser, error)
	Info(*module.Module) (io.ReadCloser, error)
	Latest(*module.Module) (io.ReadCloser, error)
	Mod(*module.Module) (io.ReadCloser, error)
	Zip(*module.Module) (io.ReadCloser, error)
}

// localFetcher uses a local puller for pulling private repositories
type localFetcher struct {
	GoPath string
	Env    []string
	Config
}

func newLocalFetcher(c Config) (*localFetcher, error) {
	fetcher := &localFetcher{}
	dir, err := fetcher.allocateTempDir()
	if err != nil {
		return nil, err
	}
	fetcher.Config = c
	fetcher.GoPath = dir
	fetcher.prepareEnv()
	return fetcher, nil
}

func (c *localFetcher) list(mod *module.Module) (*module.List, error) {
	fd := c.executeGo("list", []string{
		"-m", "-versions", "-json", mod.GetAddr(),
	})
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	fd.Stdout = stdout
	fd.Stderr = stderr
	err := fd.Run()
	if err != nil {
		return nil, errors.New(stderr.String())
	}

	moduleList := &module.List{}
	err = jsoniter.NewDecoder(stdout).Decode(moduleList)
	return moduleList, err
}

func (c *localFetcher) List(mod *module.Module) (io.ReadCloser, error) {
	moduleList, err := c.list(mod)
	if err != nil {
		return nil, err
	}
	reader := ioutil.NopCloser(strings.NewReader(moduleList.GetVersions().String()))
	return reader, nil
}

func (c *localFetcher) Info(mod *module.Module) (io.ReadCloser, error) {
	info, err := mod.GetStorageInfo(c.GoPath)
	if err == nil {
		return info, nil
	}
	err = c.fetch(mod)
	if err != nil {
		return nil, err
	}
	return mod.GetStorageInfo(c.GoPath)
}

func (c *localFetcher) Latest(mod *module.Module) (io.ReadCloser, error) {
	list, err := c.list(mod)
	if err != nil {
		return nil, err
	}
	info, err := jsoniter.Marshal(&module.Info{
		Version: list.Version,
		Time:    list.Time,
	})
	return ioutil.NopCloser(bytes.NewReader(info)), err
}

func (c *localFetcher) Mod(mod *module.Module) (io.ReadCloser, error) {
	moduleMod, err := mod.GetStorageMod(c.GoPath)
	if err == nil {
		return moduleMod, nil
	}
	err = c.fetch(mod)
	if err != nil {
		return nil, err
	}

	return mod.GetStorageMod(c.GoPath)
}

func (c *localFetcher) Zip(mod *module.Module) (io.ReadCloser, error) {
	moduleZip, err := mod.GetStorageZip(c.GoPath)
	if err == nil {
		return moduleZip, nil
	}
	err = c.fetch(mod)
	if err != nil {
		return nil, err
	}

	return mod.GetStorageZip(c.GoPath)
}

func (c *localFetcher) prepareEnv() {
	if len(c.Env) != 0 {
		return
	}
	c.Env = append(os.Environ(), []string{
		"GOPATH=" + c.GoPath,
		"GOPROXY=direct",
		"GOCACHE=" + filepath.Join(c.GoPath, "cache"),
		// ? Is it necessary to force GO111MODULE=on
		"GO111MODULE=", 
	}...)
}

func (c *localFetcher) fetch(mod *module.Module) error {
	fd := c.executeGo("mod", []string{
		"download", mod.GetAddrWithVersion(),
	})
	stderr := &bytes.Buffer{}
	fd.Stderr = stderr
	fd.Stdout = ioutil.Discard
	err := fd.Run()
	if err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

func (c *localFetcher) executeGo(subcmd string, args []string) *exec.Cmd {
	fd := exec.Command(c.GoBinaryPath, util.Prepend(args, subcmd)...)
	fd.Env = c.Env
	fd.Dir = c.GoPath
	return fd
}

func (c *localFetcher) allocateTempDir() (string, error) {
	path := filepath.Join(os.TempDir(), ".gos", "storage")
	return path, os.MkdirAll(path, os.ModePerm)
}
