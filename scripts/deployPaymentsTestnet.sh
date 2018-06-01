#!/usr/bin/env bash
set -e
scripts/release.sh
scripts/deployer.sh --cmd deploy \
--keystore.passphrase $1 \
--contract.name payments2 \
--payments.registrationFee=100 \
--payments.erc20address=0x453c11c058F13B36a35e1AEe504b20c1A09667De \
--payments.contractPath=build/bin/IdentityPromises.bin \
--payments.abiPath=build/abi/IdentityPromises.abi
