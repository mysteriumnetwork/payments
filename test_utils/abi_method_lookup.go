package test_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"errors"
	"reflect"
)

type AbiMap map[string]string

func ParseAbis(abiMap AbiMap) (*AbiList, error) {
	var abis = AbiList{
		make(map[string]abi.ABI),
	}
	for key , val := range abiMap {
		parsed , err := abi.JSON(strings.NewReader(val))
		if err != nil {
			return nil , err
		}
		abis.Register(key , parsed)
	}
	return &abis , nil
}

type AbiList struct {
	nameToAbi map[string]abi.ABI
}

func (abis *AbiList) Register(name string , abi abi.ABI) {
	abis.nameToAbi[name] = abi
}

var constructorId = []byte{96, 128, 96, 64}

func (abis * AbiList) Lookup(methodId []byte) ( string , *abi.Method , error ) {
	for key , val := range abis.nameToAbi {
		if reflect.DeepEqual(methodId, constructorId) {  //TODO a better way to convert both slices to uint4 ?
			return key +" (First taken)" , &val.Constructor , nil
		}

		method , err := val.MethodById(methodId)
		if err != nil {
			continue
		}
		return key, method, nil
	}
	return "" , nil , errors.New("no method found")
}