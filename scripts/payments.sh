#!/usr/bin/env bash

go run $GOPATH/src/github.com/mysteriumnetwork/payments/cli/payments/*.go \
--geth.url "https://ropsten.infura.io" \
--payments.contract=0xbe5F9CCea12Df756bF4a5Baf4c29A10c3ee7C83B \
$@