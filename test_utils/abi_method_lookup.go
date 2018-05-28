package test_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
)

type AbiData struct {
	AbiData string
	BinData string
}

type AbiMap map[string]AbiData

func ParseAbis(abiMap AbiMap) (*AbiList, error) {
	var abis = AbiList{
		make([]NamedABI, 10),
		make(map[string]NamedABI),
	}
	for key , val := range abiMap {
		parsed , err := abi.JSON(strings.NewReader(val.AbiData))
		if err != nil {
			return nil , err
		}
		abis.Register(NamedABI{parsed , key} , codeToHash(common.FromHex(val.BinData)))
	}
	return &abis , nil
}

type NamedABI struct {
	Abi abi.ABI
	Name string
}

type AbiList struct {
	abiList []NamedABI
	codeHashToAbi map[string]NamedABI
}

func (abis *AbiList) Register(abi NamedABI, codeHash string) {
	abis.abiList = append(abis.abiList , abi)
	abis.codeHashToAbi[codeHash] = abi
}

func (abis * AbiList) Lookup(tx * types.Transaction) ( string , *abi.Method , error ) {
	if tx.To() == nil {   //contract creation is sending transaction to "0x00" account address (nil in this case)
		return abis.lookupByCodeHash(codeToHash(tx.Data()))
	}
	methodSignature := tx.Data()[0:4]
	for _ , val := range abis.abiList {
		method , err := val.Abi.MethodById(methodSignature)
		if err != nil {
			continue
		}
		return val.Name, method, nil
	}
	return "" , nil , errors.New("no method found")
}

func codeToHash(data []byte) string {
	return common.ToHex(crypto.Keccak256(data[0:128])) //let's take first 128 bytes and keep our fingers crossed that contracts are different
}

func (abis * AbiList) lookupByCodeHash(codeHash string) ( string , *abi.Method , error ) {
	abi , found := abis.codeHashToAbi[codeHash]
	if !found {
		return "", nil , errors.New("constructor not found by code hash " + codeHash)
	}
	return abi.Name+"(constructor)" , &abi.Abi.Constructor , nil
}