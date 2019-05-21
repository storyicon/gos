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

package concurrent

import "sync"

// Concurrent for simple concurrency control
type Concurrent struct {
	channel chan bool
	wait    *sync.WaitGroup
}

// New is used to initial a concurrent control object
func New(limit int) *Concurrent {
	return &Concurrent{
		wait:    &sync.WaitGroup{},
		channel: make(chan bool, limit),
	}
}

// Add is used to add a task
func (c *Concurrent) Add(n int) {
	c.wait.Add(n)
	for n > 0 {
		n--
		c.channel <- true
	}
}

// Done is used to accomplish a task
func (c *Concurrent) Done() {
	c.wait.Done()
	<-c.channel
}

// Wait is used to wait for all tasks to be completed
func (c *Concurrent) Wait() {
	c.wait.Wait()
}
