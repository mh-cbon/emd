---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
Name: emd
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

See [emd README file](https://raw.githubusercontent.com/mh-cbon/emd/master/README.e.md)

# {{toc 5}}

# Install

{{template "gh/releases" .}}

#### glide
{{template "glide/install" .}}

#### Chocolatey
{{template "choco/install" .}}

#### linux rpm/deb repository
{{template "linux/gh_src_repo" .}}

#### linux rpm/deb standalone package
{{template "linux/gh_pkg" .}}

# Usage

#### $ {{exec "emd" "-help" | color "sh"}}

#### $ {{shell "emd gen -help" | color "sh"}}

#### $ {{shell "emd init -help" | color "sh"}}

# Cli examples

```sh

# Init a basic emd file to get started.
emd init

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

Then you can access thos data by their keys `{{echo "{{.categories}}" "{{.tags}}" }}`

directly followed by yout content.

The keys are injected into the template `dot`, the value are `json` decoded.

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

Functions can be invoked like this `{{echo "{{func \"arg1\" \"arg2\"}}" }}`

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
| __cat__(f string) | Displays a file header.<br/>Reads and returns the file body. | `emd_cat_pre: "### > "`: string to show right before the file path. |
| __read__(f string) | Reads and returns the file body. |  |
| __yaml__(f string, keypaths ...string) | Parses given file as yaml, locate given path, build a new map, yaml encode it, returns its string. |  |

#### Templates functions

| Name | Description | Options |
| --- | --- | -- |
| __render__(name string, data interface{}, keyValues ...interface{}) | Renders given template name, using data as its data.<br/> Additionnal data values can be declared using `keyValues ...interface{}` signature, such as <br/>`render("x", data, "key1", "val1", "key2", "val2")`. | | |
| __set__(name string, x interface{}) | Save given value `x` as `name` on dot `.`. |  |

#### Go utils functions

| Name | Description | Options |
| --- | --- | -- |
| __pkgdoc__(files ...string) | Reads the first of the files, or `main.go`, lookup for its package comment and returns it as plain text. | |
| __gotest__(rpkg string, run string, args ...string) | Runs `go test <rpkg> -v -run <run> <args>`, returns its output. <br/>`rpkg` can be a path to a relative folder like `./emd`. It will resolve to <br/>`github.com/mh-cbon/emd/emd`| `emd_gotest_pre: "### $ "` defines a sring to show right before the `go test` command line. |

#### Markdown functions

| Name | Description | Options |
| --- | --- | -- |
| __color__(color string, content string]) string | Embed given content with triple backquote syntax colorizer support. | |
| __toc__(maxImportance int, title ...string) string | Displays a `TOC` of the `README` file being processed.<br/>`maxImportance` defines the titles to select by their numbers of `#`.<br/>`titles` define the title to display, defaults to `TOC`.<br/>Titles displayed before the call to `{{echo "{{toc x}}" }}` are automatically ignored.| |
| __preline__(pre string, content string) string | Prepends every line of content with pre. |  |
| __echo__(f string) string | Prints given string, usefull to print strings which contaisn the template tokens. |  |
| __link__(url string, text ...string) string | Prints markdown link. |  |
| __img__(url string, alt ...string) string | Prints markdown image. |  |
| __concat__(x ...string) string | Concat given arguments. |  |
| __pathjoin__(x ...string) string | Join given arguments with `/`. |  |

#### Cli functions

| Name | Description | Options |
| --- | --- | -- |
| __exec__(bin string, args ...string) | Displays a command line header.<br/>Executes and returns its response. | `emd_exec_pre: "### > "`:  string to show right before the command line. |
| __shell__(s string) | Displays a command line header.<br/>Executes the command on a shell, and returns the its response. | `emd_shell_pre: "### > "`: string to show right before the command line. |

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
| __badge/codeship__ | Show a codeship badge. | __CsProjectID__: The codeship project ID (*123465*).<br/> __CsUUID__: the codeship project UUID (*654654-6465-54...*).<br/>Within your `e.md` file use the `render` function, `{{echo "{{render \"badge/codeship\" . \"CsUUID\" \"xx\" \"CsProjectID\" \"yyy\"}}"}}`.<br/>Via cli, add it with `--data '{"CsUUID": "xx", "CsProjectID":"yy"}'`. |
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

#### > {{cat "main_test.go" | color "go"}}

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
