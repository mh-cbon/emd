# emd

[![travis Status](https://travis-ci.org/mh-cbon/emd.svg?branch=master)](https://travis-ci.org/mh-cbon/emd)[![appveyor Status](https://ci.appveyor.com/api/projects/status/github/mh-cbon/emd?branch=master&svg=true)](https://ci.appveyor.com/project/mh-cbon/emd)
[![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/emd)](https://goreportcard.com/report/github.com/mh-cbon/emd)

[![GoDoc](https://godoc.org/github.com/mh-cbon/emd?status.svg)](http://godoc.org/github.com/mh-cbon/emd)


Enhanced Markdown template processor.


See [emd README file](https://raw.githubusercontent.com/mh-cbon/emd/master/README.e.md)

# Install

Check the [release page](https://github.com/mh-cbon/emd/releases)!

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


###### $ emd -help
```sh
emd - 0.0.0

Usage
  -h	Show help
  -help
    	Show help
  -version
    	Show version

Commands
	gen	Process an emd file.
```


###### $ emd gen -help
```sh
emd - 0.0.0

Command "gen": Process an emd file.
  -data string
    	JSON map of data
  -h	Show help
  -help
    	Show help
  -in string
    	Input src file (default "README.e.md")
  -out string
    	Output destination, defaults to stdout (default "-")
```

# Cli examples

```sh

# Reads content of README.e.md, outputs to README.md
emd gen -out README.md

# same with data injections,
emd gen -out README.md --data='{"CsUUID":"xxxx"}'

```

# Templates helper

#### Define data

It is possible to define data directly into the `README` file,

insert a prelude such as

```
---
title: "Easygen - Easy to use universal code/text generator"
date: "2016-01-01T22:13:12-05:00"
categories: ["Tech"]
tags: ["go","programming","easygen","CLI"]
---
{ {.categories} }
```

directly followed by yout content.

The keys are injected into the template `dot`, the value are `json` decoded.

#### Data

| Key | Description |
| --- | --- |
| Name | Project name based on the cwd (example: emd). |
| User | User name based on the cwd (example: mh-cbon). |
| ProviderURL | The vcs provider url (example: github.com). |
| ProviderName | The vcs provider name (example: github). |
| URL | Project url as determined by the cwd (example: github.com/mh-cbon/emd). |
| Branch | Current vcs branch name (defaults to master). |

#### Function

| Name | Description |
| --- | --- |
| color(color string, contet string]) | Embed given content wiht triple backquote syntax colorizer support. |
| cat(f string[, colorizer string]) | Displays a file header. Reads and returns the file body. |
| exec(bin string, args ...string) | Displays a command line header. Executes and returns the command line response. |
| pkgdoc(files ...string) | Reads the first of the files, or `main.go`, lookup for its package comment and returns it as plain text. |
| gotest(run string, args ...string) | Executes `go test -v -run <run> <args>` and returns its output. |
| render(name string, data interface{}, keyValues ...interface{}) | Renders given template name, using data as its data. Additionnal data values can be declared using `keyValues ...interface{}` signature, such as `render("x", data, "key1", "val1", "key2", "val2")`. |

__deprecated helpers__

| Name | Description |
| --- | --- |
| file(f string[, colorizer string]) | read and display a file enclosed with triples backquotes. If `colorizer` is empty, it defaults to the file extension. |
| cli(bin string, args ...string) | execute and display a command line enclosed with triples backquotes. The highlight defaults to `sh`. |

#### Templates

##### std

| Name | Description | Params |
| --- | --- | --- |
| gh/releases | Show a text to link the release page. | |
| badge/travis | Show a travis badge. | |
| badge/appveyor | Show an appveyor badge. | |
| badge/codeship | Show a codeship badge. | __CsUUID__: the codeship project UUID. Within your `e.md` file use the `render` function, `{render "badge/codeship" . "CsUUID" "xxxxxx"}`. Via cli, add it with `--data '{"CsUUID": "xxxxxx"}'`. |
| choco/install | Show an sh snippet to install the package with chocolatey. | |
| linux/gh_src_repo | Show an sh snippet to install the package via linux repositories (deb/rpm). | |
| linux/gh_pkg | Show an sh snippet to install the package via linux packages (deb/rpm). | |

##### go

| Name | Description | Params |
| --- | --- | --- |
| go/install | Show an sh snippet to install the package via `go get`. | |
| badge/godoc | Show a godoc badge. | |
| badge/goreport | Show a goreport badge. | |

##### go-nonstd

| Name | Description | Params |
| --- | --- | --- |
| glide/install | Show an sh snippet to install the package via `glide`. | |


# API example


###### > main_test.go
```go
package main_test

import (
	"os"

	"github.com/mh-cbon/emd/emd"
	"github.com/mh-cbon/emd/std"
)

var projectName = "dummy"

// ExampleGenerate demonstrates the generation
// of the given README.e.md source file
// to os.Stdout.
func Example() {

	// make a new instance of emd.Generator.
	gen := emd.NewGenerator()

	// set the main template.
	gen.AddTemplate("{{.Name}}")

	// set the data available in templates.
	gen.SetDataMap(map[string]interface{}{"Name": projectName})

	// register a plugin
	if err := std.Register(gen); err != nil {
		panic(err)
	}

	// process the template.
	if err := gen.Execute(os.Stdout); err != nil {
		panic(err)
	}
	// Output: dummy
}
```

# Recipes

#### Generate HTML content

To directly generate HTML content out of `emd` output, for example, with `gh-markdown-cli`,

```sh
npm install gh-markdown-cli -g
emd gen | mdown
```

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

<no value>

