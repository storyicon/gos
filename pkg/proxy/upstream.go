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
	"io"
	"net/http"

	"errors"

	"github.com/storyicon/gos/pkg/proxy/module"
)

type upstreamFetcher struct {
	UpsteamAddr string
}

func newUpstreamFetcher(upstream string) *upstreamFetcher {
	return &upstreamFetcher{
		UpsteamAddr: upstream,
	}
}

// List is used to list all versions of the specified package
// It is one of the standard interfaces specified by GOPROXY
func (c *upstreamFetcher) List(mod *module.Module) (io.ReadCloser, error) {
	return c.fetch(mod.GetListAddr)
}

// Info is used to return information about the specified version of the specified package
// It is one of the standard interfaces specified by GOPROXY
func (c *upstreamFetcher) Info(mod *module.Module) (io.ReadCloser, error) {
	return c.fetch(mod.GetInfoAddr)
}

// Latest is used to return the latest version of the specified package
// It is one of the standard interfaces specified by GOPROXY
func (c *upstreamFetcher) Latest(mod *module.Module) (io.ReadCloser, error) {
	return c.fetch(mod.GetLatestAddr)
}

// Mod is used to return module info about the specified version of the specified package
// It is one of the standard interfaces specified by GOPROXY
func (c *upstreamFetcher) Mod(mod *module.Module) (io.ReadCloser, error) {
	return c.fetch(mod.GetModAddr)
}

// Zip is used to return zip file about the specified version of the specified package
// It is one of the standard interfaces specified by GOPROXY
func (c *upstreamFetcher) Zip(mod *module.Module) (io.ReadCloser, error) {
	return c.fetch(mod.GetZipAddr)
}

func (c *upstreamFetcher) fetch(addrFunc func(string, bool) (string, error)) (io.ReadCloser, error) {
	addr, err := addrFunc(c.UpsteamAddr, true)
	if err != nil {
		return nil, err
	}

	r, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 || r.ContentLength == 0 {
		return nil, errors.New("not found")
	}
	return r.Body, nil
}
