#!/usr/bin/env bash

set -e
#directory to which to put results (must exist)
dir=$1

url='https://github.com/ethereum/solidity/releases/download/v0.5.0/solc-static-linux'
curl --location --create-dirs --output $dir/solc $url && chmod +x $dir/solc
