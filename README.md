crowley
=======

`crowley` is a linting utility that can scan source code files written in any programming language that will make recommendations based on the coding practices it finds in the scanned files. Its output is concise, easy to apply, and would be a valuable addtion to any workflow or CI/CD pipeline.

BUILDING
--------

Building `crowley` is simple. Cloning the repository and running `go build` and `go install` will do the trick, but you could also copy the `crowley` binary to `/usr/local/bin` or somewhere similar. There are no runtime dependencies.

USAGE
-----

Using `crowley` is straightfoward. Just point it at a source tree directory, run the utility, and it will come back with a list of recommendations of changes. The options are simple and listed below.

```
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
```

BUGS
----

Unlikely.

COPYRIGHT
---------

Copyright (c) Jeremy Bingham, 2019-2025.

LICENSE
-------

`crowley` is licensed under the terms of the Apache License, version 2.0. See the LICENSE file for details.
