#!/usr/bin/env bash
set -e

version=$1
path=$2
path=${path:-"contracts/deps"}

releaseurl="https://github.com/OpenZeppelin/openzeppelin-solidity/archive/$version.tar.gz"


fullPath=$path/OpenZeppelin
rm -rf $fullPath && mkdir -p $fullPath

curl --location $releaseurl | tar -xz --strip-components 1 -C $fullPath 'openzeppelin-solidity-*/contracts/*'
echo "Fetched version: $version" > $fullPath/version.txt