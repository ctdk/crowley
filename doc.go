/*
 * Copyright (c) 2020-2025, Jeremy Bingham (<jeremy@goiardi.gl>)
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

/*
`crowley` is a linter that can scan source code files written in any programming language that will make recommendations based on the coding practices it finds in the scanned files. Its output is concise, easy to apply, and would be a valuable addtion to any workflow or CI/CD pipeline.

# Building

Building `crowley` is simple. Cloning the repository and running `go build` and `go install` will do the trick, but you could also copy the `crowley` binary to `/usr/local/bin` or somewhere similar. There are no runtime dependencies.

# Usage

Using `crowley` is straightfoward. Just point it at a source tree directory, run the utility, and it will come back with a list of recommendations of changes. The options are simple and listed below.

	Usage:
	  crowley [OPTIONS]

	Application Options:
	  -v, --version   Print version info.
	  -p, --path=     Root path to scan for linter errors and naughty coding
			  practices. Required.
	  -V, --verbose   Verbose output.
	  -c, --no-color  Disable colorized linter output.

	Help Options:
	  -h, --help      Show this help message

# Bugs

Unlikely.

# Copyright

Copyright (c) Jeremy Bingham, 2019-2025.

# License

`crowley` is licensed under the terms of the Apache License, version 2.0. See the LICENSE file for details.

*/

package main
