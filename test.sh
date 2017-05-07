#!/bin/sh

set -ex

rm $GOPATH/bin/emd
go install

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test
rm -fr ~/fake

# test 1: a project contained in GOPATH, aliased out of it
mkdir -p ~/fake/src/github.com/mh-cbon/test-emd
ln -s ~/fake/src/github.com/mh-cbon/test-emd $GOPATH/src/github.com/mh-cbon/emd-test

cd ~/fake/src/github.com/mh-cbon/test-emd

export VERBOSE=y

cat <<EOT | emd gen | grep "Name=test-emd" || exit 1;
Name={{.Name}}
EOT

export VERBOSE=""

cat <<EOT | emd gen | grep "ProviderName=github" || exit 1;
ProviderName={{.ProviderName}}
EOT

cat <<EOT | emd gen | grep "URL=github.com/mh-cbon/test-emd" || exit 1;
URL={{.URL}}
EOT

cat <<EOT | emd gen | grep "ProjectURL=github.com/mh-cbon/test-emd" || exit 1;
ProjectURL={{.ProjectURL}}
EOT

cat <<EOT | emd gen | grep "Branch=master" || exit 1;
Branch={{.Branch}}
EOT

cat <<EOT | emd gen | grep "User=mh-cbon" || exit 1;
User={{.User}}
EOT

cat <<EOT | emd gen | grep "ProviderURL=github.com" || exit 1;
ProviderURL={{.ProviderURL}}
EOT

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test
rm -fr ~/fake

# test 2: a project not contained in GOPATH, aliased in gopath
mkdir -p ~/fake/src/github.com/mh-cbon/
mkdir -p $GOPATH/src/github.com/mh-cbon/emd-test
ln -s $GOPATH/src/github.com/mh-cbon/emd-test ~/fake/src/github.com/mh-cbon/test-emd

cd  ~/fake/src/github.com/mh-cbon/test-emd

export VERBOSE=y

cat <<EOT | emd gen | grep "Name=test-emd" || exit 1;
Name={{.Name}}
EOT

export VERBOSE=""

cat <<EOT | emd gen | grep "ProviderName=github" || exit 1;
ProviderName={{.ProviderName}}
EOT

cat <<EOT | emd gen | grep "URL=github.com/mh-cbon/test-emd" || exit 1;
URL={{.URL}}
EOT

cat <<EOT | emd gen | grep "ProjectURL=github.com/mh-cbon/test-emd" || exit 1;
ProjectURL={{.ProjectURL}}
EOT

cat <<EOT | emd gen | grep "Branch=master" || exit 1;
Branch={{.Branch}}
EOT

cat <<EOT | emd gen | grep "User=mh-cbon" || exit 1;
User={{.User}}
EOT

cat <<EOT | emd gen | grep "ProviderURL=github.com" || exit 1;
ProviderURL={{.ProviderURL}}
EOT

cat <<EOT | emd gen | grep "https://travis-ci.org/mh-cbon/test-emd.svg?branch=master" || exit 1;
{{template "badge/travis" .}}
EOT
cat <<EOT | emd gen | grep "https://travis-ci.org/mh-cbon/test-emd" || exit 1;
{{template "badge/travis" .}}
EOT

cat <<EOT | emd gen | grep "https://ci.appveyor.com/api/projects/status/github/mh-cbon/test-emd?branch=master&svg=true" || exit 1;
{{template "badge/appveyor" .}}
EOT
cat <<EOT | emd gen | grep "https://ci.appveyor.com/projects/mh-cbon/test-emd" || exit 1;
{{template "badge/appveyor" .}}
EOT

cat <<EOT | emd gen | grep "https://goreportcard.com/badge/github.com/mh-cbon/test-emd" || exit 1;
{{template "badge/goreport" .}}
EOT


cat <<EOT | emd gen | grep "https://godoc.org/github.com/mh-cbon/test-emd?status.svg" || exit 1;
{{template "badge/godoc" .}}
EOT
cat <<EOT | emd gen | grep "http://godoc.org/github.com/mh-cbon/test-emd" || exit 1;
{{template "badge/godoc" .}}
EOT

cat <<EOT | emd gen | grep "MIT License" || exit 1;
{{render "license/shields" . "License" "MIT" "LicenseFile" "LICENSE" "LicenseColor" "yellow"}}
EOT
cat <<EOT | emd gen | grep "http://img.shields.io/badge/License-MIT-yellow.svg" || exit 1;
{{render "license/shields" . "License" "MIT" "LicenseFile" "LICENSE" "LicenseColor" "yellow"}}
EOT
cat <<EOT | emd gen | grep "(LICENSE)" || exit 1;
{{render "license/shields" . "License" "MIT" "LicenseFile" "LICENSE" "LicenseColor" "yellow"}}
EOT

cat <<EOT | emd gen | grep "[title]" || exit 1;
{{render "badge/codeship" . "CsUUID" "uuid" "CsProjectID" "projectID" "CsTitle" "title"}}
EOT
cat <<EOT | emd gen | grep "https://codeship.com/projects/uuid/status?branch=master" || exit 1;
{{render "badge/codeship" . "CsUUID" "uuid" "CsProjectID" "projectID" "CsTitle" "title"}}
EOT
cat <<EOT | emd gen | grep "https://codeship.com/projects/projectID" || exit 1;
{{render "badge/codeship" . "CsUUID" "uuid" "CsProjectID" "projectID" "CsTitle" "title"}}
EOT

cat <<EOT | emd gen | grep "GOPATH/src/github.com/mh-cbon/test-emd" || exit 1;
{{template "glide/install" . }}
EOT
cat <<EOT | emd gen | grep "git clone https://github.com/mh-cbon/test-emd.git ." || exit 1;
{{template "glide/install" . }}
EOT
cat <<EOT | emd gen | grep "glide install" || exit 1;
{{template "glide/install" . }}
EOT

cat <<EOT | emd gen | grep "choco install test-emd" || exit 1;
{{template "choco/install" . }}
EOT

cat <<EOT | emd gen | grep "https://github.com/mh-cbon/test-emd/releases" || exit 1;
{{template "gh/releases" . }}
EOT

cat <<EOT | emd gen | grep "go get github.com/mh-cbon/test-emd" || exit 1;
{{template "go/install" . }}
EOT

cat <<EOT | emd gen | grep "https://godoc.org/github.com/mh-cbon/test-emd?status.svg" || exit1
{{template "badge/godoc" . }}
EOT
cat <<EOT | emd gen | grep "http://godoc.org/github.com/mh-cbon/test-emd" || exit1
{{template "badge/godoc" . }}
EOT

cat <<EOT | emd gen | grep "https://goreportcard.com/badge/github.com/mh-cbon/test-emd" || exit 1;
{{template "badge/goreport" . }}
EOT

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test
rm -fr ~/fake

echo ""
echo "ALL GOOD!"
