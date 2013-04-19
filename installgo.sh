#!/bin/bash
set -xe
cd $HOME
wget eniac.cs.sun.ac.za/go1.0.3.linux-386.tar.gz
tar -xvf go1.0.3.linux-386.tar.gz 
export PATH=$PATH:$HOME/go/bin/
export GOROOT=$HOME/go
export GOPATH=$HOME/go
echo $PATH
go version
(cd go/src && ./all.bash )

