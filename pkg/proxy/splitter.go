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
    "net/http"
)

// StreamDestType defines the destination to process the request
type StreamDestType uint8

// Define a set of common stream dest type
const (
    StreamDestTypeUnknown StreamDestType = iota
    StreamDestTypeLocal
    StreamDestTypeUpstream
)

const (
    httpProtocol  = "http://"
    httpsProtocol = "https://"
)

// StreamSplitter is used to separate requests,
// it determines whether a goproxy request should use upstream or local
type StreamSplitter interface {
    Split(c *Context) StreamDestType
}

type gosStreamSplitter struct{}

func newGosStreamSplitter() *gosStreamSplitter {
    return &gosStreamSplitter{}
}

// Split is used to determine whether a goproxy request should use upstream or local
func (s *gosStreamSplitter) Split(c *Context) StreamDestType {
    domain := c.Module.GetDomain()
    switch domain {
    // case for special uses
    default:
        return StreamDestTypeUpstream
    }
}

func (s *gosStreamSplitter) ping(addr string) StreamDestType {
    resp, err := http.Head(addr)
    if err != nil || resp.StatusCode != 200 {
        return StreamDestTypeLocal
    }
    return StreamDestTypeUpstream
}

func (s *gosStreamSplitter) githubSplit(addr string) StreamDestType {
    return s.ping(httpsProtocol + addr)
}
