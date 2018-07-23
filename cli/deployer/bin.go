package main

import (
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DeploySmartContractFile(opts *bind.TransactOpts, path string, abiPath string, backend bind.ContractBackend, params ...interface{}) (common.Address, *types.Transaction, error) {
	contractContent, err := ioutil.ReadFile(path)
	if err != nil {
		return common.Address{}, nil, err
	}
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return common.Address{}, nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiContent)))
	if err != nil {
		return common.Address{}, nil, err
	}

	binBytes := common.FromHex(string(contractContent))
	addr, tx, _, err := bind.DeployContract(opts, contractAbi, binBytes, backend, params...)
	return addr, tx, err
}
