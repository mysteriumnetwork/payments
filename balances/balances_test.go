package balances

import (
	"testing"
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/MysteriumNetwork/payments/mysttoken"
	"github.com/stretchr/testify/assert"
	"math/big"
	"github.com/MysteriumNetwork/payments/registry"
	"github.com/MysteriumNetwork/payments/balances/generated"
	"time"
	generated2 "github.com/MysteriumNetwork/payments/mysttoken/generated"
)

var abiMap = test_utils.AbiMap{
	"MystToken" : generated2.MystTokenABI,
	"IdentityBalances" : generated.IdentityBalancesABI,
}

var abiList , _ = test_utils.ParseAbis(abiMap)

func TestTopUpActionAddsMystToBalanceAndEmitsToppedUpEvent(t *testing.T) {
	deployer := test_utils.Deployer
	simulator := test_utils.LoggingBackend(test_utils.NewSimulatedBackend(deployer.Address , 1000000000) , abiList)  //ether amount in gweis - if you get panics like 'out of gas' increase this

	erc20 , err := mysttoken.DeployMystERC20( deployer.Transactor , 2000 , simulator)
	assert.NoError(t , err)
	simulator.Commit()

	identityBalances , err := DeployIdentityBalances(deployer.Transactor , erc20.Address, simulator)
	assert.NoError(t , err)
	simulator.Commit()

	_ , err = erc20.Approve(identityBalances.Address , big.NewInt(1000)) //on behalf of myst token owner we allow balances contract to take some tokens from us
	assert.NoError( t , err)
	simulator.Commit()

	mystIdentity , err := registry.NewMystIdentity()
	assert.NoError(t , err)

	initialBalance , err := identityBalances.Balances(mystIdentity.Address)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), initialBalance.Uint64())

	topupChan := make(chan * generated.IdentityBalancesToppedUp, 10)
	subscription , err := identityBalances.BindForTopUpEvents(topupChan)
	assert.NoError(t, err)

	_ , err = identityBalances.TopUp(mystIdentity.Address , big.NewInt(500))
	assert.NoError(t, err)
	simulator.Commit()

	newBalance , err := identityBalances.Balances(mystIdentity.Address )
	assert.NoError(t ,err)
	assert.Equal(t, big.NewInt(500) , newBalance)

	erc20Balance, err := erc20.BalanceOf(deployer.Address)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(1500) , erc20Balance)

	select {
		case toppedUp := <- topupChan:
			assert.Equal(t, mystIdentity.Address , toppedUp.Identity)
		assert.Equal(t, uint64(500), toppedUp.Amount.Uint64())
		assert.Equal(t, uint64(500), toppedUp.TotalBalance.Uint64())
		case err = <- subscription.Err():
			assert.NoError(t, err)
		case <- time.After(100* time.Millisecond):
			assert.Fail(t, "ToppedUp event expected in 100 milliseconds")
	}
	subscription.Unsubscribe()
}

func TestWithdrawActionRemovesMystFromBalanceAndEmitsWithdrawnEvent(t *testing.T) {
	deployer := test_utils.Deployer
	simulator := test_utils.LoggingBackend(test_utils.NewSimulatedBackend(deployer.Address , 1000000000) , abiList)  //ether amount in gweis - if you get panics like 'out of gas' increase this

	erc20 , err := mysttoken.DeployMystERC20( deployer.Transactor , 2000 , simulator)
	assert.NoError(t , err)
	simulator.Commit()

	identityBalances , err := DeployIdentityBalances(deployer.Transactor , erc20.Address, simulator)
	assert.NoError(t , err)
	simulator.Commit()

	_ , err = erc20.Approve(identityBalances.Address , big.NewInt(1000)) //on behalf of myst token owner we allow balances contract to take some tokens from us
	assert.NoError( t , err)
	simulator.Commit()

	mystIdentity , err := registry.NewMystIdentity()
	assert.NoError(t , err)

	withdrawChan := make(chan * generated.IdentityBalancesWithdrawn, 10)
	subscription , err := identityBalances.BindForWithdrawEvents(withdrawChan)
	assert.NoError(t, err)

	_ , err = identityBalances.TopUp(mystIdentity.Address , big.NewInt(500))
	assert.NoError(t, err)
	simulator.Commit()

	request , err := NewWithdrawRequest(mystIdentity , 300)
	assert.NoError(t, err)

	err = identityBalances.Withdraw(request)
	assert.NoError(t, err)
	simulator.Commit()

	newBalance , err := identityBalances.Balances(mystIdentity.Address )
	assert.NoError(t ,err)
	assert.Equal(t, big.NewInt(200) , newBalance)

	erc20Balance, err := erc20.BalanceOf(deployer.Address)
	assert.NoError(t ,err)
	assert.Equal(t, big.NewInt(1800) , erc20Balance)

	select {
	case withdrawn := <- withdrawChan:
		assert.Equal(t, mystIdentity.Address , withdrawn.Identity)
		assert.Equal(t, big.NewInt(300), withdrawn.Amount)
		assert.Equal(t, big.NewInt(200), withdrawn.TotalBalance)
	case err = <- subscription.Err():
		assert.NoError(t, err)
	case <- time.After(100* time.Millisecond):
		assert.Fail(t, "Withdrawn event expected in 100 milliseconds")
	}
	subscription.Unsubscribe()
}