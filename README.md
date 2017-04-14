# emd

[![travis Status](https://travis-ci.org/mh-cbon/emd.svg?branch=master)](https://travis-ci.org/mh-cbon/emd) 
[![appveyor Status](https://ci.appveyor.com/api/projects/status/github/mh-cbon/emd?branch=master&svg=true)](https://ci.appveyor.com/project/mh-cbon/emd) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/emd)](https://goreportcard.com/report/github.com/mh-cbon/emd) [![GoDoc](https://godoc.org/github.com/mh-cbon/emd?status.svg)](http://godoc.org/github.com/mh-cbon/emd) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Enhanced Markdown template processor.


See [emd README file](https://raw.githubusercontent.com/mh-cbon/emd/master/README.e.md)

# TOC
- [Install](#install)
  - [Go](#go)
  - [Chocolatey](#chocolatey)
  - [linux rpm/deb repository](#linux-rpmdeb-repository)
  - [linux rpm/deb standalone package](#linux-rpmdeb-standalone-package)
- [Usage](#usage)
  - [$ emd -help](#-emd--help)
  - [$ emd gen -help](#-emd-gen--help)
- [Cli examples](#cli-examples)
- [Templates helper](#templates-helper)
  - [Define data](#define-data)
  - [Data](#data)
  - [Function](#function)
  - [Templates](#templates)
- [API example](#api-example)
  - [> main_test.go](#-main_testgo)
- [Recipes](#recipes)
  - [Generate HTML content](#generate-html-content)
  - [Release the project](#release-the-project)
- [History](#history)

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

#### $ emd -help
```sh
emd - 0.0.9-beta3

Usage
  -h	Show help
  -help
    	Show help
  -version
    	Show version

Commands
	gen	Process an emd file.
```

#### $ emd gen -help
```sh
emd - 0.0.9-beta3

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

```yml
---
title: "Easygen - Easy to use universal code/text generator"
date: "2016-01-01T22:13:12-05:00"
categories: ["Tech"]
tags: ["go","programming","easygen","CLI"]
---
```

Then you can access thos data by their keys `{ {.categories} }`, `{ {.tags} }`

directly followed by yout content.

The keys are injected into the template `dot`, the value are `json` decoded.

#### Data

| Key | Description |
| --- | --- |
| __Name__ | Project name based on the cwd (example: emd). |
| __User__ | User name based on the cwd (example: mh-cbon). |
| __ProviderURL__ | The vcs provider url (example: github.com). |
| __ProviderName__ | The vcs provider name (example: github). |
| __URL__ | Project url as determined by the cwd (example: github.com/mh-cbon/emd). |
| __Branch__ | Current vcs branch name (defaults to master). |

#### Function

| Name | Description | Options |
| --- | --- | -- |
| __color__(color string, content string]) | Embed given content with triple backquote syntax colorizer support. | |
| __cat__(f string) | Displays a file header.<br/>Reads and returns the file body. | `emd_cat_pre: "### > "`: string to show right before the file path. |
| __exec__(bin string, args ...string) | Displays a command line header.<br/>Executes and returns its response. | `emd_exec_pre: "### > "`:  string to show right before the command line. |
| __shell__(s string) | Displays a command line header.<br/>Executes the command on a shell, and returns the its response. | `emd_shell_pre: "### > "`: string to show right before the command line. |
| __toc__(maxImportance int, title ...string) | Displays a `TOC` of the `README` file being processed.<br/>`maxImportance` defines the titles to select by their numbers of `#`.<br/>`titles` define the title to display, defaults to `TOC`.<br/>Titles displayed before the call to `{ {toc x}}` are automatically ignored.| |
| __pkgdoc__(files ...string) | Reads the first of the files, or `main.go`, lookup for its package comment and returns it as plain text. | |
| __gotest__(rpkg string, run string, args ...string) | Runs `go test <rpkg> -v -run <run> <args>`, returns its output. <br/>`rpkg` can be a path to a relative folder like `./emd`. It will resolve to <br/>`github.com/mh-cbon/emd/emd`| `emd_gotest_pre: "### $ "` defines a sring to show right before the `go test` command line. |
| __render__(name string, data interface{}, keyValues ...interface{}) | Renders given template name, using data as its data.<br/> Additionnal data values can be declared using `keyValues ...interface{}` signature, such as <br/>`render("x", data, "key1", "val1", "key2", "val2")`. | | |

Options are keys to define into the `prelude`:

```yaml
---
emd_cat_pre: "### > "
emd_gotest_pre: "### $ "
emd_exec_pre: "### $ "
emd_shell_pre: "### $ "
---
```

__deprecated helpers__

| Name | Description |
| --- | --- |
| __file__(f string[, colorizer string]) | read and display a file enclosed with triples backquotes. If `colorizer` is empty, it defaults to the file extension. |
| __cli__(bin string, args ...string) | execute and display a command line enclosed with triples backquotes. The highlight defaults to `sh`. |

#### Templates

##### std

| Name | Description | Params |
| --- | --- | --- |
| __gh/releases__ | Show a text to link the release page. | |
| __badge/travis__ | Show a travis badge. | |
| __badge/appveyor__ | Show an appveyor badge. | |
| __badge/codeship__ | Show a codeship badge. | __CsProjectID__: The codeship project ID (*123465*).<br/> __CsUUID__: the codeship project UUID (*654654-6465-54...*).<br/>Within your `e.md` file use the `render` function, `{render "badge/codeship" . "CsUUID" "xx" "CsProjectID" "yyy"}`. <br/>Via cli, add it with `--data '{"CsUUID": "xx", "CsProjectID":"yy"}'`. |
| __choco/install__ | Show an sh snippet to install the package with chocolatey. | |
| __linux/gh_src_repo__ | Show an sh snippet to install the package via linux repositories (deb/rpm). | |
| __linux/gh_pkg__ | Show an sh snippet to install the package via linux packages (deb/rpm). | |
| __license/shields__ | Show a license badge. | __License__: The license name like `MIT`, `BSD`.<br/>__LicenseFile__: The path to the license file. |

##### go

| Name | Description | Params |
| --- | --- | --- |
| __go/install__ | Show an sh snippet to install the package via `go get`. | |
| __badge/godoc__ | Show a godoc badge. | |
| __badge/goreport__ | Show a goreport badge. | |

##### go-nonstd

| Name | Description | Params |
| --- | --- | --- |
| __glide/install__ | Show an sh snippet to install the package via `glide`. | |


# API example

#### > main_test.go
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

# History

[CHANGELOG](CHANGELOG.md)

