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
	"io/ioutil"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/storyicon/gos/pkg/meta"
	"github.com/storyicon/gos/pkg/proxy/module"
)

var discard = func(*Context) {}

type Engine struct {
	Config
	backend Backend
	pool    sync.Pool
	s       *gin.Engine
}

func New(config *Config) *Engine {
	engine := &Engine{}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}

	s := gin.New()
	s.Use(engine.Interceptor())
	err := config.fix()
	if err != nil {
		panic(err)
	}

	engine.s = s
	engine.Config = *config
	return engine
}

func Default() *Engine {
	c := meta.GetConfig()
	return New(&Config{
		GoBinaryPath: c.GoBinaryPath,
		UpstreamAddr: c.UpstreamAddr,
		ListenAddr:   c.ProxyListenAddr,
	})
}

func (engine *Engine) SetBackend(backend Backend) {
	// to protect
	if engine.backend == nil {
		engine.backend = backend
	}
}

func (engine *Engine) GetBackend() Backend {
	// Use Default Backend
	if engine.backend == nil {
		engine.backend = newGosBackend(engine.Config)
	}
	return engine.backend
}

func (engine *Engine) Run() error {
	logrus.Debugln("Run on:", engine.ListenAddr)
	return engine.s.Run(engine.ListenAddr)
}

func (engine *Engine) Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		path, err := module.NewPath(c.Request.URL.Path)
		if err != nil {
			return
		}
		ctx := engine.createContext(path, c)
		defer func() {
			ctx.reset()
			engine.pool.Put(ctx)
		}()
		engine.GetHandler(path)(ctx)
	}
}

func (engine *Engine) GetHandler(p *module.Path) func(*Context) {
	backend := engine.GetBackend()
	switch r := p.GetType(); r {
	case module.TypePathList:
		return backend.List
	case module.TypePathInfo:
		return backend.Info
	case module.TypePathMod:
		return backend.Mod
	case module.TypePathZip:
		return backend.Zip
	case module.TypePathLatest:
		return backend.Latest
	default:
		return discard
	}
}

func (engine *Engine) createContext(path *module.Path, c *gin.Context) *Context {
	ctx := engine.pool.Get().(*Context)
	ctx.reset()
	ctx.Path = path
	ctx.Context = c
	return ctx
}

func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
}
