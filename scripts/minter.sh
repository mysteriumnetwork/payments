#!/usr/bin/env bash

go run $GOPATH/src/github.com/MysteriumNetwork/payments/cli/minter/*.go \
--geth.url cli/testnet/geth.ipc \
--keystore.directory "cli/testnet" \
--ether.address=0xc08d190725659E8C6E5ec3eE845666CEE9487DD3 \
--erc20.address=0x453c11c058f13b36a35e1aee504b20c1a09667de \
$@