/*
 * Copyright (c) 2020 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License 2.0 (the "License");
 * You may not use this file except in compliance with the License.
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

// Package proxy is a type of shorter proxy to other online shorten service.
package proxy

import (
	"errors"

	"go.zoe.im/shortme/shorter"
)

// proxyShorter is a http proxy to handle the shorter servie
type proxyShorter struct {
	c Config
}

func (p *proxyShorter) do() {}

// Gen method return a new short of long input
func (p *proxyShorter) Gen(url string, shorten ...string) (string, error) {
	// TODO: implement
	return "", errors.New("unimplemented")
}

// Get method return a long url from shorten one
func (p *proxyShorter) Get(shorten string) (string, error) {
	// TODO: implement
	return "", errors.New("unimplemented")
}

// New create a shorter service from configuration
func New(s shorter.Store, c interface{}) (shorter.Shorter, error) {
	// load extends configuration
	// TODO: implement
	return nil, errors.New("unimplemented")
}
