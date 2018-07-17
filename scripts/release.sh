#!/usr/bin/env bash
set -e
mkdir -p build/src
./scripts/package_sol.sh ./contracts build/src
#MystToken for testnet purposes only
./scripts/generate_abi.sh build/src/MystToken.sol build/abi
./scripts/generate_bin.sh build/src/MystToken.sol build/bin
#Payments
./scripts/generate_abi.sh build/src/IdentityPromises.sol build/abi
./scripts/generate_bin.sh build/src/IdentityPromises.sol build/bin