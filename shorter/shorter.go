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

package shorter

// Shorter is a inerface to Make and Load a shorten url
type Shorter interface {
	// ID return a short id
	ID() string

	// Gen to generate a shorten url from a long one
	//
	// url is the long url, shorten if we offered
	// we can set to it.
	Gen(url string, shorten ...string) (string, error)

	// Get is a method to take a original url from shorten one.
	//
	// shorten is the only one param
	Get(shorten string) (string, error)
}
