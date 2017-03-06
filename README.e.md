# {{.Name}}

{{template "badge/travis" .}}{{template "badge/appveyor" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

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

{{cli "emd" "-help"}}

{{cli "emd" "gen" "-help"}}

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
| Name | Project directory name (filepath.Base(cwd)). |
| User | Project directory name (filepath.Base(cwd)). |
| ProviderURL | The vcs provider url (example: github.com). |
| ProviderName | The vcs provider name (example: github). |
| URL | Project url as determined by the cwd (example: github.com/mh-cbon/emd). |
| Branch | Current vcs branch name (defaults to master). |

#### Function

| Name | Description |
| --- | --- |
| file(f string) | read and display a file enclosed with triples backquotes. The highlight defaults to `go`. |
| cli(bin string, args ...string) | execute and display a command line enclosed with triples backquotes. The highlight defaults to `sh`. |
| pkgdoc(files ...string) | reads the first of files, lookup for its package comment and shows it as plain text. |

#### Templates

| Name | Description | Params |
| --- | --- | --- |
| gh/releases | Show a text to link the release page. | |
| badge/travis | Show a travis badge. | |
| badge/appveyor | Show an appveyor badge. | |
| badge/codeship | Show acodeship badge. | __CsUUID__: the codeship project UUID. Add it with `--data '{"CsUUID": "xxxxxx"}'` |
| choco/install | Show an sh snippet to install the package with chocolatey. | |
| linux/gh_src_repo | Show an sh snippet to install the package via linux repositories (deb/rpm). | |
| linux/gh_pkg | Show an sh snippet to install the package via linux packages (deb/rpm). | |
| glide/install | Show an sh snippet to install the package via `glide`. | |
| go/install | Show an sh snippet to install the package via `go get`. | |

# API example

{{file "main_example.go"}}

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
