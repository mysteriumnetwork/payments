package utils

import (
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/MysteriumNetwork/payments/utils/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var abiList, _ = test_utils.ParseAbis(test_utils.AbiMap{
	"TestUtilsContract": {
		generated.TestUtilsContractABI,
		generated.TestUtilsContractBin,
	},
},
)

func TestSignle32BytesParamIsUnpackedCorrectly(t *testing.T) {
	deployer := test_utils.Deployer
	backend := test_utils.NewSimulatedBackend(deployer.Address, 1*1000*1000*1000)
	loggingBackend := test_utils.LoggingBackend(backend, abiList)

	_, _, testUtilsC, err := generated.DeployTestUtilsContract(deployer.Transactor, loggingBackend)
	assert.NoError(t, err)
	loggingBackend.Commit()

	addr := common.HexToAddress("0x0011223344556677889900112233445566778899")
	var signs [2]byte
	signs[0] = byte(0x28)
	signs[1] = byte(0x73)
	var data [32]byte
	dataSlice := append(signs[:], addr.Bytes()...)
	copy(data[10:32], dataSlice)
	receivedAddr, receivedV1, receivedV2, err := testUtilsC.GiveMeData(&bind.CallOpts{}, data)
	assert.NoError(t, err)

	assert.Equal(t, addr, receivedAddr)
	assert.Equal(t, signs[0], receivedV1)
	assert.Equal(t, signs[1], receivedV2)
}
