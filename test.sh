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

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test
rm -fr ~/fake

echo ""
echo "ALL GOOD!"
