#!/usr/bin/env bash

solc --bin contracts/IdentityPromises.sol --optimize --overwrite -o $1