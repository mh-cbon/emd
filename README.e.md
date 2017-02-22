# {{.Name}}

{{template "badge/travis" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

# Install

#### Go
{{template "go/install" .}}

#### Chocolatey

{{template "choco/install" .}}

#### linux rpm/deb repository

{{template "linux/gh_src_repo" .}}

#### linux rpm/deb standalone package

{{template "linux/gh_pkg" .}}

# Usage

{{cli "build/emd" "-help"}}

{{cli "build/emd" "gen" "-help"}}

# Cli examples

to generate a README file,
```sh
emd gen -out README.md
```

# API example

{{file "main_example.go"}}
