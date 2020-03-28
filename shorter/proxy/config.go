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

package proxy

// Config is object to contians configuration of proxy
//
// define the request flow to proxy a shorten service
type Config struct {
	Host        string            `json:"host,omitempty" yaml:"host"`
	Headers     map[string]string `json:"headers,omitempty" yaml:"headers"`
	Create, Get Request           `json:"create,omitempty" yaml:"create"`
}

// Request define a request to handle with http request
type Request struct {
	Method  string            `json:"method,omitempty" yaml:"method"`
	Path    string            `json:"path,omitempty" yaml:"path"`
	Headers map[string]string `json:"headers,omitempty" yaml:"headers"`
	Body    string            `json:"body,omitempty" yaml:"body"` // send the data with body
	Data    string            `json:"data,omitempty" yaml:"data"` // parse the data from reponse, if data from header?
}
