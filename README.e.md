# {{.Name}}

{{template "badge/travis" .}}{{template "badge/godoc" .}}

{{pkgdoc}}

# Install

{{template "go/install" .}}

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
