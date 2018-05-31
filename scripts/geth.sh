#!/bin/bash
geth --ipcpath "geth.ipc" --datadir "cli/testnet" --testnet --syncmode light $@
