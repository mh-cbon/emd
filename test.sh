#!/bin/sh

set -ex

rm $GOPATH/bin/emd
go install

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test ~/test-emd
rm -fr ~/test-emd

mkdir -p ~/test-emd
ln -s ~/test-emd $GOPATH/src/github.com/mh-cbon/emd-test

cd $GOPATH/src/github.com/mh-cbon/emd-test

export VERBOSE=y

cat <<EOT | emd gen | grep "Name=emd-test" || exit 1;
Name={{.Name}}
EOT

export VERBOSE=""

cat <<EOT | emd gen | grep "ProviderName=github" || exit 1;
ProviderName={{.ProviderName}}
EOT

cat <<EOT | emd gen | grep "URL=github.com/mh-cbon/emd-test" || exit 1;
URL={{.URL}}
EOT

cat <<EOT | emd gen | grep "ProjectURL=github.com/mh-cbon/emd-test" || exit 1;
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

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test ~/test-emd
rm -fr ~/test-emd

# test 2
mkdir -p $GOPATH/src/github.com/mh-cbon/emd-test
ln -s $GOPATH/src/github.com/mh-cbon/emd-test ~/test-emd

cd  ~/test-emd

export VERBOSE=y

cat <<EOT | emd gen | grep "Name=emd-test" || exit 1;
Name={{.Name}}
EOT

export VERBOSE=""

cat <<EOT | emd gen | grep "ProviderName=github" || exit 1;
ProviderName={{.ProviderName}}
EOT

cat <<EOT | emd gen | grep "URL=github.com/mh-cbon/emd-test" || exit 1;
URL={{.URL}}
EOT

cat <<EOT | emd gen | grep "ProjectURL=github.com/mh-cbon/emd-test" || exit 1;
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

rm -fr $GOPATH/src/github.com/mh-cbon/emd-test ~/test-emd
rm -fr ~/test-emd
