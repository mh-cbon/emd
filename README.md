# emd

[![travis Status](https://travis-ci.org/mh-cbon/emd.svg?branch=master)](https://travis-ci.org/mh-cbon/emd) [![Appveyor Status](https://ci.appveyor.com/api/projects/status/github/mh-cbon/emd?branch=master&svg=true)](https://ci.appveyor.com/project/mh-cbon/emd) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/emd)](https://goreportcard.com/report/github.com/mh-cbon/emd) [![GoDoc](https://godoc.org/github.com/mh-cbon/emd?status.svg)](http://godoc.org/github.com/mh-cbon/emd) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Enhanced Markdown template processor.


See [emd README file](https://raw.githubusercontent.com/mh-cbon/emd/master/README.e.md)

# TOC
- [Install](#install)
  - [glide](#glide)
  - [Bintray](#bintray)
  - [Chocolatey](#chocolatey)
  - [linux rpm/deb repository](#linux-rpmdeb-repository)
  - [linux rpm/deb standalone package](#linux-rpmdeb-standalone-package)
- [Usage](#usage)
  - [$ emd -help](#-emd--help)
  - [$ emd gen -help](#-emd-gen--help)
  - [$ emd init -help](#-emd-init--help)
- [Cli examples](#cli-examples)
- [Templates helper](#templates-helper)
  - [Define data](#define-data)
  - [Data](#data)
  - [Functions](#functions)
  - [Files functions](#files-functions)
  - [Templates functions](#templates-functions)
  - [Go utils functions](#go-utils-functions)
  - [Markdown functions](#markdown-functions)
  - [Cli functions](#cli-functions)
  - [Deprecated function](#deprecated-function)
  - [Templates](#templates)
- [API example](#api-example)
  - [> main_test.go](#-main_testgo)
- [Recipes](#recipes)
  - [Generate HTML content](#generate-html-content)
  - [Release the project](#release-the-project)
- [History](#history)

# Install

Check the [release page](https://github.com/mh-cbon/emd/releases)!

#### glide
```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/emd
cd $GOPATH/src/github.com/mh-cbon/emd
git clone https://github.com/mh-cbon/emd.git .
glide install
go install
```

#### Bintray
```sh
choco source add -n=mh-cbon -s="https://api.bintray.com/nuget/mh-cbon/choco"
choco install emd
```

#### Chocolatey
```sh
choco install emd
```

#### linux rpm/deb repository
```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/bintray.sh \
| GH=mh-cbon/emd sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/bintray.sh \
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
emd - 0.0.0

Usage
  -h	Show help
  -help
    	Show help
  -version
    	Show version

Commands
	gen	Process an emd file.
	init	Init a basic emd file.
```

#### $ emd gen -help
```sh
emd - 0.0.0

Command "gen": Process an emd file.
  -data string
    	JSON map of data
  -h	Show help
  -help
    	Show help
  -in string
    	Input src file
  -out string
    	Output destination, defaults to stdout (default "-")
```

#### $ emd init -help
```sh
emd - 0.0.0

Command "init": Init a basic emd file.
  -force
    	Force write
  -h	Show help
  -help
    	Show help
  -out string
    	Out file (default "README.e.md")
```

# Cli examples

```sh

# Init a basic emd file to get started.
emd init

# Reads content of README.e.md, outputs to README.md
emd gen -out README.md

# same with data injections,
emd gen -out README.md --data='{"CsUUID":"xxxx"}'

# use verbose mode
VERBOSE=y emd gen
```

# Templates helper

#### Define data

Template data can be defined directly into the `README.e.md` file using a `prelude`,

```yml
---
title: "Easygen - Easy to use universal code/text generator"
date: "2016-01-01T22:13:12-05:00"
categories: ["Tech"]
tags: ["go","programming","easygen","CLI"]
---
```

This `prelude` must be inserted right before the regular `md` content.

The keys are injected into the template `dot`, the value are `json` decoded.

Template can access those data using name: `{{.categories}} {{.tags}}`


#### Data

| Key | Description |
| --- | --- |
| __ProviderURL__ | The vcs provider url (example: github.com). |
| __ProviderName__ | The vcs provider name (example: github). |
| __Name__ | Project name based on the cwd (example: emd). |
| __User__ | User name based on the cwd (example: mh-cbon). |
| __URL__ | Project url as determined by the cwd (example: github.com/mh-cbon/emd). |
| __ProjectURL__ | Project url as determined by the cwd + relative path (example: github.com/mh-cbon/emd/cmd). |
| __Branch__ | Current vcs branch name (defaults to master). |

#### Functions

Functions can be invoked like this `{{func "arg1" "arg2"}}`

Options are keys to define into the `prelude`:

```yaml
---
emd_cat_pre: "### > "
emd_gotest_pre: "### $ "
emd_exec_pre: "### $ "
emd_shell_pre: "### $ "
---
```

#### Files functions

| Name | Description | Options |
| --- | --- | -- |
| __cat__(f string) | Displays a file header.<br/>Read and return the file body. | `emd_cat_pre: "### > "`: string to show right before the file path. |
| __read__(f string) | Read and return the file body. |  |
| __yaml__(f string, keypaths ...string) | Parse given file as yaml, locate given path into the yaml content, yaml re encode it, return its string. |  |

#### Templates functions

| Name | Description | Options |
| --- | --- | -- |
| __render__(name string, data interface{}, keyValues ...interface{}) | Render given `template` name, using `data`.<br/> Additionnal data values can be declared using `keyValues ...interface{}` signature, such as <br/>`render("x", data, "key1", "val1", "key2", "val2")`. | | |
| __set__(name string, x interface{}) | Save given value `x` as `name` on dot `.`. |  |

#### Go utils functions

| Name | Description | Options |
| --- | --- | -- |
| __pkgdoc__(files ...string) | Read the first of `files`, or `main.go`, lookup for its package comment and return it as plain text. | |
| __gotest__(rpkg string, run string, args ...string) | Run `go test <rpkg> -v -run <run> <args>`, return its output. <br/>`rpkg` can be a path to a relative folder like `./emd`. It will resolve to <br/>`github.com/mh-cbon/emd/emd`| `emd_gotest_pre: "### $ "` defines a sring to show right before the `go test` command line. |

#### Markdown functions

| Name | Description | Options |
| --- | --- | -- |
| __color__(color string, content string]) string | Embed given content with triple backquote syntax colorizer support. | |
| __toc__(maxImportance int, title ...string) string | Displays a `TOC` of the `README` file being processed.<br/>`maxImportance` defines the titles to select by their numbers of `#`.<br/>`titles` define the title to display, defaults to `TOC`.<br/>Titles displayed before the call to `{{toc x}}` are automatically ignored.| |
| __preline__(pre string, content string) string | Prepend every line of `content` with `pre`. |  |
| __echo__(f string) string | Prints `f`, usefull to print strings which contains the template tokens. |  |
| __link__(url string, text ...string) string | Prints markdown link. |  |
| __img__(url string, alt ...string) string | Prints markdown image. |  |
| __concat__(x ...string) string | Concat all `x`. |  |
| __pathjoin__(x ...string) string | Join all `x` with `/`. |  |

#### Cli functions

| Name | Description | Options |
| --- | --- | -- |
| __exec__(bin string, args ...string) | Display a command line header.<br/>Execute and return its response. | `emd_exec_pre: "### > "`:  string to show right before the command line. |
| __shell__(s string) | Display a command line header.<br/>Execute the command on a shell, and return the its response. | `emd_shell_pre: "### > "`: string to show right before the command line. |

#### Deprecated function

| Name | Description |
| --- | --- |
| __file__(f string[, colorizer string]) | Read and display a file enclosed with triples backquotes. If `colorizer` is empty, it defaults to the file extension. |
| __cli__(bin string, args ...string) | Execute and display a command line enclosed with triples backquotes. The highlight defaults to `sh`. |

#### Templates

##### std

| Name | Description | Params |
| --- | --- | --- |
| __gh/releases__ | Show a text to link the release page. | |
| __badge/travis__ | Show a travis badge. | |
| __badge/appveyor__ | Show an appveyor badge. | |
| __badge/codeship__ | Show a codeship badge. | __CsProjectID__: The codeship project ID (*123465*).<br/> __CsUUID__: the codeship project UUID (*654654-6465-54...*).<br/>Within your `e.md` file use the `render` function, `{{render "badge/codeship" . "CsUUID" "xx" "CsProjectID" "yyy"}}`.<br/>Via cli, add it with `--data '{"CsUUID": "xx", "CsProjectID":"yy"}'`. |
| __choco_bintray/install__ | Show a snippet to install the package with chocolatey from bintray repos. | __BintrayRepo__: the name of the bintray repo (default: `choco`) |
| __choco/install__ | Show a snippet to install the package with chocolatey. | |
| __linux/gh_src_repo__ | Show an sh snippet to install the package via `rpm|deb|apt` repositories hosted on gh-pages. | |
| __linux/bintray_repo__ | Show an sh snippet to install the package via `rpm|deb|apt` repositories hosted on bintray. | |
| __linux/gh_pkg__ | Show an sh snippet to install the package via standalone packages (deb/rpm). | |
| __license/shields__ | Show a license badge. | __License__: The license name like `MIT`, `BSD`.<br/>__LicenseFile__: The path to the license file.<br/>__LicenseColor__: The color of the badge (defaults t o blue). |

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
