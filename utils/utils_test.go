package utils

import (
	"testing"
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/MysteriumNetwork/payments/utils/generated"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
)

func TestSignle32BytesParamIsUnpackedCorrectly(t *testing.T) {
	deployer := test_utils.Deployer
	backend := test_utils.NewSimulatedBackend(deployer.Address, 1 * 1000 * 1000 * 1000)
	loggingBackend := test_utils.LoggingBackend(backend , &test_utils.AbiList{})

	_ , _ , testUtilsC , err := generated.DeployTestUtilsContract(deployer.Transactor , loggingBackend)
	assert.NoError(t , err)
	loggingBackend.Commit()

	addr := common.HexToAddress("0x0011223344556677889900112233445566778899")
	v1 := byte(0x28)
	v2 := byte(0x73)
	var data [32]byte
	dataSlice := append(addr.Bytes() , v1 ,v2)
	copy(data[0:22] , dataSlice)
	fmt.Printf("Data: %v\n" , common.ToHex(data[:]))
	receivedAddr , receivedV1, receivedV2 , err := testUtilsC.GiveMeData(&bind.CallOpts{} , data)
	assert.NoError(t , err)

	assert.Equal(t, addr, receivedAddr)
	assert.Equal(t, v1, receivedV1)
	assert.Equal(t, v2, receivedV2)
}
