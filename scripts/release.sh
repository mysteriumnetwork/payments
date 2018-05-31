#!/usr/bin/env bash
set -e
mkdir -p build/src
./scripts/package_sol.sh ./contracts build/src
./scripts/generate_abi.sh build/src/IdentityPromises.sol build/abi
./scripts/generate_bin.sh build/src/IdentityPromises.sol build/bin