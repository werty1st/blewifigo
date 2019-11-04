#!/bin/bash
# This is the entry point for configuring the system.
#####################################################

#install basic tools
sudo DEBIAN_FRONTEND=noninteractive apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confnew" update
sudo DEBIAN_FRONTEND=noninteractive apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confnew" install curl wget git make wireless-tools

#get golang 1.9.1
curl -O https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz

#unzip the archive 
tar -xvf go1.13.4.linux-amd64.tar.gz

#move the go lib to local folder
mv go /usr/local

#delete the source file
rm  go1.13.4.linux-amd64.tar.gz

#only full path will work
touch /home/vagrant/.bash_profile

echo "export PATH=$PATH:/usr/local/go/bin:/go/bin" >> /home/vagrant/.bash_profile
echo "export GOPATH=/go" >> /home/vagrant/.bash_profile

export GOPATH=/go
export PATH=$PATH:/usr/local/go/bin:/$GOPATH/bin

mkdir -p "$GOPATH/bin"
mkdir -p "$GOPATH/src/github.com/werty1st/blewifigo"

# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0

golangci-lint --version

chown -R vagrant: /home/vagrant
chown -R vagrant: /go

