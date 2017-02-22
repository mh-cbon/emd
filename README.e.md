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

to generate a README file,
```sh
emd gen -out README.md
```

# Templates helper

#### Data

__Name__: Project directory name (filepath.Base(cwd)).

__User__: Project directory name (filepath.Base(cwd)).

__ProviderURL__: The vcs provider url (example: github.com).

__ProviderName__: The vcs provider name (example: github).

__URL__: Project url as determined by the cwd (example: github.com/mh-cbon/emd).

__Branch__: Current vcs branch name (defaults to master).

#### Function

__file(f string)__: read and display a file enclosed with triples backquotes. The highlight defaults to `go`.

__cli(bin string, args ...string)__: execute and display a command line enclosed with triples backquotes. The highlight defaults to `sh`.

__pkgdoc(files ...string)__: reads the first of files, lookup for its package comment and shows it as plain text.

#### Templates

__gh/releases__: Show a text to link the release page.

__badge/travis__: Show a travis badge.

__badge/appveyor__: Show an appveyor badge.

__choco/install__: Show an sh snippet to install the package with chocolatey.

__linux/gh_src_repo__: Show an sh snippet to install the package via linux repositories (deb/rpm).

__linux/gh_pkg__: Show an sh snippet to install the package via linux packages (deb/rpm).

__glide/install__: Show an sh snippet to install the package via `glide`.

__go/install__: Show an sh snippet to install the package via `go get`.

# API example

{{file "main_example.go"}}
