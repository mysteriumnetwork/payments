#!/usr/bin/env bash

set -e
#directory to which to put results (must exist)
dir=$1

#geth verison 1.8.20
url="https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-1.8.20-24d727b6.tar.gz"

curl $url | tar -xz --strip-components 1 -C $dir
