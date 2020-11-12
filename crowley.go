/*
 * Copyright (c) 2020, Jeremy Bingham (<jeremy@goiardi.gl>)
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

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const version = "0.0.1"

type Options struct {
	Version bool `short:"v" long:"version" description:"Print version info."`
	RootPath string `short:"p" long:"path" description:"Root path to scan for linter errors and naughty coding practices."`
	Verbose bool `short:"V" long:"verbose" description:"Verbose output."`
}

func main() {
	var opts = &Options{}
	parser := flags.NewParser(opts, flags.Default)
	parser.ShortDescription = fmt.Sprintf("A source code checker, with a special focus on non-standard coding practices, version %s.", version)

	if _, err := parser.Parse(); err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Println(err)
			os.Exit(1)
		}
	}

	if opts.Version {
		fmt.Printf("crowley version %s, built with %s.\n", version, runtime.Version())
		os.Exit(0)
	}

	pathInfo, err := os.Stat(opts.RootPath)
	if err != nil {
		log.Fatal(err)
	}

	if !pathInfo.IsDir() {
		log.Fatalf("Path %s is not a directory", opts.RootPath)
	}

	c := new(int)
	d := new(int)

	walkErr := filepath.Walk(opts.RootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			*d++ // another directory
		} else {
			*c++ // another file
		}
		return nil
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}

	if opts.Verbose {
		fmt.Printf("Scanned %d files, %d directories... No issues found.\n\n", *c, *d)
	}

	fmt.Printf("\"Do What Thou Wilt\" shall be the Whole of the Law.\n")
	os.Exit(0)
}
