#!/usr/bin/env bash

set -e
#directory to which to put results (must exist)
dir=$1

#solc version 0.4.25 (upgrading to 0.5.x should be considered after OpenZeppelin project fixes up it's contracts
url='https://github.com/ethereum/solidity/releases/download/v0.4.25/solc-static-linux'
curl --location --create-dirs --output $dir/solc $url && chmod +x $dir/solc
