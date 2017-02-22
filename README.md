# emd

[![Build Status](https://travis-ci.org/mh-cbon/emd.svg?branch=master)](https://travis-ci.org/mh-cbon/emd)
[![GoDoc](https://godoc.org/github.com/mh-cbon/emd?status.svg)](http://godoc.org/github.com/mh-cbon/emd)


Enhanced Markdown template processor.


# Install

#### Go

```sh
go get github.com/mh-cbon/emd
```


#### Chocolatey

```sh
choco install emd
```

#### linux rpm/deb repository

```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/emd sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/emd sh -xe
```

#### linux rpm/deb standalone package

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/emd sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/emd sh -xe
```

# Usage


__$ emd -help__
```sh
emd - 0.0.0

Usage
  -help
    	Show help
  -version
    	Show version

Commands
	gen	Process an emd file.
```


__$ emd gen -help__
```sh
emd - 0.0.0

Command "gen": Process an emd file.
  -data string
    	JSON map of data
  -help
    	Show help
  -in string
    	Input src file (default "README.e.md")
  -out string
    	Output destination, defaults to stdout (default "-")
success
```

# Cli examples

to generate a README file,
```sh
emd gen -out README.md
```

# API example


__> main_example.go__
```go
// Demonstrates the generation
// of the given README.e.md source file
// to os.Stdout.

package main

import (
	"os"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/std"
)

func genExample() {

	// make a new instance of emd.Generator.
	gen := emd.NewGenerator()

	// set the main template.
	if err := gen.AddFileTemplate("README.e.md"); err != nil {
		panic(err)
	}

	// set the data available in templates.
	gen.SetDataMap(map[string]interface{}{"Name": "projectname"})

	// register a plugin
	if err := std.Register(gen); err != nil {
		panic(err)
	}

	// process the template.
	if err := gen.Execute(os.Stdout); err != nil {
		panic(err)
	}
}
```
