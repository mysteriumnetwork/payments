#!/bin/bash

go run $GOPATH/src/github.com/MysteriumNetwork/payments/cli/deployer/*.go \
--geth.url cli/testnet/geth.ipc \
--keystore.directory "cli/testnet" \
--ether.address=0xc08d190725659E8C6E5ec3eE845666CEE9487DD3 \
$@