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
    "github.com/asaskevich/govalidator"
    "github.com/storyicon/gos/pkg/meta"
)

// Config contains configuration information for the entire proxy project
type Config struct {
    UpstreamAddr string `valid:"url"`
    ListenAddr   string `valid:"url"`
    GoBinaryPath string
}

func (c *Config) fix() error {
    _, err := govalidator.ValidateStruct(c)
    if err != nil {
        return err
    }

    dc := meta.GetConfig()
    if c.UpstreamAddr == "" {
        c.UpstreamAddr = dc.UpstreamAddr
    }
    if c.GoBinaryPath == "" {
        c.GoBinaryPath = dc.GoBinaryPath
    }
    if c.ListenAddr == "" {
        c.ListenAddr = dc.ProxyListenAddr
    }
    return nil
}
