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
