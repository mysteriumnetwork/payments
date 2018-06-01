# Mysterium network smart contracts for payments
[![Build Status](https://travis-ci.com/MysteriumNetwork/payments.svg?token=t9FwiYsxwDxkJWnSMpfr&branch=master)](https://travis-ci.com/MysteriumNetwork/payments)

Requires solc 0.4.23 or later, go 1.9.2 or later

First we need some external tooling:
1. Solidity:
```bash
brew tap ethereum/ethereum
brew install solidity
```
2. Ethereum (geth, abigen):
```bash
brew install ethereum
```

### Building
```bash
scripts/deps.sh ensure && scripts/build.sh
```
### Testing
```bash
scripts/test.sh
```

### Current deployment (ethereum Ropsten testnet)
ERC20 Token (Mintable a la myst token): [https://ropsten.etherscan.io/address/0x453c11c058f13b36a35e1aee504b20c1a09667de](https://ropsten.etherscan.io/address/0x453c11c058f13b36a35e1aee504b20c1a09667de)

Payments: [https://ropsten.etherscan.io/address/0xbe5F9CCea12Df756bF4a5Baf4c29A10c3ee7C83B](https://ropsten.etherscan.io/address/0xbe5F9CCea12Df756bF4a5Baf4c29A10c3ee7C83B)