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
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/storyicon/gos/pkg/proxy/module"
)

type Backend interface {
	List(ctx *Context)
	Info(ctx *Context)
	Latest(ctx *Context)
	Mod(ctx *Context)
	Zip(ctx *Context)
}

type Worker func(*module.Module) (io.ReadCloser, error)

type gosBackend struct {
	Config
	StreamSplitter
	upstream Fetcher
	storage  Fetcher
}

func newGosBackend(c Config) *gosBackend {
	storage, err := newLocalFetcher(c)
	upstream := newUpstreamFetcher(c.UpstreamAddr)
	splitter := newGosStreamSplitter()
	if err != nil {
		panic(err)
	}
	return &gosBackend{
		Config:         c,
		StreamSplitter: splitter,
		storage:        storage,
		upstream:       upstream,
	}
}

func (b *gosBackend) List(c *Context) {
	b.RunDefaultWorker(c, b.storage.List, b.upstream.List)
}

func (b *gosBackend) Info(c *Context) {
	b.RunDefaultWorker(c, b.storage.Info, b.upstream.Info)
}

func (b *gosBackend) Latest(c *Context) {
	b.RunDefaultWorker(c, b.storage.Latest, b.upstream.Latest)
}

func (b *gosBackend) Mod(c *Context) {
	b.RunDefaultWorker(c, b.storage.Mod, b.upstream.Mod)
}

func (b *gosBackend) Zip(c *Context) {
	b.RunWorker(c, b.storage.Zip, b.upstream.Zip, func(closer io.ReadCloser, c *Context) {
		defer closer.Close()
		c.Status(http.StatusOK)
		c.Header("Content-Type", "application/zip")
		_, err := io.Copy(c.Writer, closer)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})
}

func (b *gosBackend) RunDefaultWorker(c *Context, storageFunc, upstreamFunc Worker) {
	b.RunWorker(c, storageFunc, upstreamFunc, func(closer io.ReadCloser, c *Context) {
		defer closer.Close()
		bytes, _ := ioutil.ReadAll(closer)
		c.String(http.StatusOK, string(bytes))
	})
}

func (b *gosBackend) RunWorker(c *Context, storageFunc, upstreamFunc Worker, callback func(io.ReadCloser, *Context)) {
	mod := &c.Module
	var (
		feed io.ReadCloser
		err  error
	)

	if b.Split(c) == StreamDestTypeLocal {
		logrus.Debugln(">>> ABS LOCAL", c.Module.GetAddr())
		feed, err = storageFunc(mod)
	} else {
		logrus.Debugln(">>> TRY STREAM Start", c.Module.GetAddr())
		feed, err = upstreamFunc(mod)
		logrus.Debugln(">>> TRY STREAM Finish", err)
		if err != nil {
			feed, err = storageFunc(mod)
			logrus.Debugln(">>> TRY STORAGE", c.Module.GetAddr(), err)
		}
	}
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	callback(feed, c)
}
