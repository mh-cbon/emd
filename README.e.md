# {{.Name}}

{{template "badge/travis" .}}{{template "badge/appveyor" .}}{{template "badge/goreport" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

See [emd README file](https://raw.githubusercontent.com/mh-cbon/emd/master/README.e.md)

# Install

{{template "gh/releases" .}}

#### Go
{{template "go/install" .}}

#### Chocolatey
{{template "choco/install" .}}

#### linux rpm/deb repository
{{template "linux/gh_src_repo" .}}

#### linux rpm/deb standalone package
{{template "linux/gh_pkg" .}}

# Usage

{{exec "emd" "-help" | color "sh"}}

{{exec "emd" "gen" "-help" | color "sh"}}

# Cli examples

```sh

# Reads content of README.e.md, outputs to README.md
emd gen -out README.md

# same with data injections,
emd gen -out README.md --data='{"CsUUID":"xxxx"}'

```

# Templates helper

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

{{cat "main_test.go" | color "go"}}

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
