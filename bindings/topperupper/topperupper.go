// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package topperupper

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TopperUpperMetaData contains all meta data concerning the TopperUpper contract.
var TopperUpperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limitsNative\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limitsToken\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_blocksWindow\",\"type\":\"uint256[]\"}],\"name\":\"approveAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"native\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"disapproveAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validTill\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"}],\"name\":\"removeManagers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"}],\"name\":\"setManagers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validTill\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"topupNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"topupNatives\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"topupToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"topupTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002edc38038062002edc8339818101604052810190620000379190620001d5565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000207565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b620001af8162000190565b8114620001bb57600080fd5b50565b600081519050620001cf81620001a4565b92915050565b600060208284031215620001ee57620001ed6200016b565b5b6000620001fe84828501620001be565b91505092915050565b612cc580620002176000396000f3fe6080604052600436106101235760003560e01c806391b571cf116100a0578063f2fde38b11610064578063f2fde38b146103c6578063f4856230146103ef578063f58c5b6e14610406578063fc0c546a14610431578063fdff9b4d1461045c5761012a565b806391b571cf146102ce578063d4f2b668146102f7578063df8de3e714610320578063e7fc94c414610349578063f136a874146103875761012a565b80634784e21b116100e75780634784e21b146102115780634f85d3db1461023a578063714280b314610263578063715018a61461028c5780638da5cb5b146102a35761012a565b80630d4e1a1e1461012f57806310f3ee2914610158578063238e130a14610181578063367caa09146101aa578063417a0698146101d35761012a565b3661012a57005b600080fd5b34801561013b57600080fd5b5061015660048036038101906101519190611f43565b610499565b005b34801561016457600080fd5b5061017f600480360381019061017a919061211a565b61053a565b005b34801561018d57600080fd5b506101a860048036038101906101a39190612163565b610686565b005b3480156101b657600080fd5b506101d160048036038101906101cc919061211a565b610832565b005b3480156101df57600080fd5b506101fa60048036038101906101f59190612190565b610987565b6040516102089291906121cc565b60405180910390f35b34801561021d57600080fd5b50610238600480360381019061023391906122b8565b6109ab565b005b34801561024657600080fd5b50610261600480360381019061025c919061211a565b610b94565b005b34801561026f57600080fd5b5061028a6004803603810190610285919061238f565b610dc3565b005b34801561029857600080fd5b506102a1610e64565b005b3480156102af57600080fd5b506102b8610eec565b6040516102c591906123de565b60405180910390f35b3480156102da57600080fd5b506102f560048036038101906102f091906123f9565b610f15565b005b34801561030357600080fd5b5061031e60048036038101906103199190612534565b61104e565b005b34801561032c57600080fd5b5061034760048036038101906103429190612190565b611187565b005b34801561035557600080fd5b50610370600480360381019061036b9190612190565b6113eb565b60405161037e9291906121cc565b60405180910390f35b34801561039357600080fd5b506103ae60048036038101906103a99190612190565b61140f565b6040516103bd939291906125ac565b60405180910390f35b3480156103d257600080fd5b506103ed60048036038101906103e89190612190565b611439565b005b3480156103fb57600080fd5b50610404611531565b005b34801561041257600080fd5b5061041b61162e565b60405161042891906123de565b60405180910390f35b34801561043d57600080fd5b506104466116d4565b6040516104539190612642565b60405180910390f35b34801561046857600080fd5b50610483600480360381019061047e9190612190565b6116fa565b6040516104909190612678565b60405180910390f35b600360006104a561171a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661052c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610523906126f0565b60405180910390fd5b6105368282611722565b5050565b61054261171a565b73ffffffffffffffffffffffffffffffffffffffff16610560610eec565b73ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad9061275c565b60405180910390fd5b60008151116105fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f1906127c8565b60405180910390fd5b60005b8151811015610682576003600083838151811061061d5761061c6127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055808061067a90612846565b9150506105fd565b5050565b61068e61171a565b73ffffffffffffffffffffffffffffffffffffffff166106ac610eec565b73ffffffffffffffffffffffffffffffffffffffff1614610702576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f99061275c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610772576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610769906128db565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61083a61171a565b73ffffffffffffffffffffffffffffffffffffffff16610858610eec565b73ffffffffffffffffffffffffffffffffffffffff16146108ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a59061275c565b60405180910390fd5b60008151116108f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e990612947565b60405180910390fd5b60005b815181101561098357600160036000848481518110610917576109166127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808061097b90612846565b9150506108f5565b5050565b60056020528060005260406000206000915090508060000154908060010154905082565b6109b361171a565b73ffffffffffffffffffffffffffffffffffffffff166109d1610eec565b73ffffffffffffffffffffffffffffffffffffffff1614610a27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1e9061275c565b60405180910390fd5b82518451148015610a39575081518351145b8015610a46575081518151145b610a85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7c906127c8565b60405180910390fd5b60005b8451811015610b8d5760006040518060600160405280868481518110610ab157610ab06127e8565b5b60200260200101518152602001858481518110610ad157610ad06127e8565b5b60200260200101518152602001848481518110610af157610af06127e8565b5b602002602001015181525090508060046000888581518110610b1657610b156127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050508080610b8590612846565b915050610a88565b5050505050565b610b9c61171a565b73ffffffffffffffffffffffffffffffffffffffff16610bba610eec565b73ffffffffffffffffffffffffffffffffffffffff1614610c10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c079061275c565b60405180910390fd5b6000815111610c54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4b906127c8565b60405180910390fd5b60005b8151811015610dbf5760046000838381518110610c7757610c766127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160009055505060066000838381518110610ceb57610cea6127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160009055505060056000838381518110610d5757610d566127e8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160009055600182016000905550508080610db790612846565b915050610c57565b5050565b60036000610dcf61171a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4d906126f0565b60405180910390fd5b610e608282611a48565b5050565b610e6c61171a565b73ffffffffffffffffffffffffffffffffffffffff16610e8a610eec565b73ffffffffffffffffffffffffffffffffffffffff1614610ee0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed79061275c565b60405180910390fd5b610eea6000611dd7565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60036000610f2161171a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610fa8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9f906126f0565b60405180910390fd5b8151815114610fec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe3906127c8565b60405180910390fd5b60005b81518110156110495761103683828151811061100e5761100d6127e8565b5b6020026020010151838381518110611029576110286127e8565b5b6020026020010151611a48565b808061104190612846565b915050610fef565b505050565b6003600061105a61171a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166110e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d8906126f0565b60405180910390fd5b8151815114611125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111c906127c8565b60405180910390fd5b60005b82518110156111825761116f838281518110611147576111466127e8565b5b6020026020010151838381518110611162576111616127e8565b5b6020026020010151610499565b808061117a90612846565b915050611128565b505050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611219576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611210906128db565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156112aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a1906129d9565b60405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016112e591906123de565b60206040518083038186803b1580156112fd57600080fd5b505afa158015611311573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113359190612a0e565b90508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401611394929190612a5c565b602060405180830381600087803b1580156113ae57600080fd5b505af11580156113c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e69190612ab1565b505050565b60066020528060005260406000206000915090508060000154908060010154905082565b60046020528060005260406000206000915090508060000154908060010154908060020154905083565b61144161171a565b73ffffffffffffffffffffffffffffffffffffffff1661145f610eec565b73ffffffffffffffffffffffffffffffffffffffff16146114b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ac9061275c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611525576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151c90612b50565b60405180910390fd5b61152e81611dd7565b50565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156115c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ba906128db565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561162b573d6000803e3d6000fd5b50565b600061163861171a565b73ffffffffffffffffffffffffffffffffffffffff16611656610eec565b73ffffffffffffffffffffffffffffffffffffffff16146116ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116a39061275c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b600033905090565b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015443111561191e5780600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156117f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e890612bbc565b60405180910390fd5b600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201544361183f9190612bdc565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546118d39190612c32565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506119fd565b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156119a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161199a90612bbc565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546119f59190612c32565b925050819055505b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611a43573d6000803e3d6000fd5b505050565b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154431115611c445780600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101541015611b17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b0e90612bbc565b60405180910390fd5b600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015443611b659190612bdc565b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154611bf99190612c32565b600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550611d23565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015611cc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cc090612bbc565b60405180910390fd5b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254611d1b9190612c32565b925050819055505b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401611d80929190612c66565b602060405180830381600087803b158015611d9a57600080fd5b505af1158015611dae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dd29190612ab1565b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611eda82611eaf565b9050919050565b611eea81611ecf565b8114611ef557600080fd5b50565b600081359050611f0781611ee1565b92915050565b6000819050919050565b611f2081611f0d565b8114611f2b57600080fd5b50565b600081359050611f3d81611f17565b92915050565b60008060408385031215611f5a57611f59611ea5565b5b6000611f6885828601611ef8565b9250506020611f7985828601611f2e565b9150509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611fd182611f88565b810181811067ffffffffffffffff82111715611ff057611fef611f99565b5b80604052505050565b6000612003611e9b565b905061200f8282611fc8565b919050565b600067ffffffffffffffff82111561202f5761202e611f99565b5b602082029050602081019050919050565b600080fd5b600061205082611eaf565b9050919050565b61206081612045565b811461206b57600080fd5b50565b60008135905061207d81612057565b92915050565b600061209661209184612014565b611ff9565b905080838252602082019050602084028301858111156120b9576120b8612040565b5b835b818110156120e257806120ce888261206e565b8452602084019350506020810190506120bb565b5050509392505050565b600082601f83011261210157612100611f83565b5b8135612111848260208601612083565b91505092915050565b6000602082840312156121305761212f611ea5565b5b600082013567ffffffffffffffff81111561214e5761214d611eaa565b5b61215a848285016120ec565b91505092915050565b60006020828403121561217957612178611ea5565b5b600061218784828501611ef8565b91505092915050565b6000602082840312156121a6576121a5611ea5565b5b60006121b48482850161206e565b91505092915050565b6121c681611f0d565b82525050565b60006040820190506121e160008301856121bd565b6121ee60208301846121bd565b9392505050565b600067ffffffffffffffff8211156122105761220f611f99565b5b602082029050602081019050919050565b600061223461222f846121f5565b611ff9565b9050808382526020820190506020840283018581111561225757612256612040565b5b835b81811015612280578061226c8882611f2e565b845260208401935050602081019050612259565b5050509392505050565b600082601f83011261229f5761229e611f83565b5b81356122af848260208601612221565b91505092915050565b600080600080608085870312156122d2576122d1611ea5565b5b600085013567ffffffffffffffff8111156122f0576122ef611eaa565b5b6122fc878288016120ec565b945050602085013567ffffffffffffffff81111561231d5761231c611eaa565b5b6123298782880161228a565b935050604085013567ffffffffffffffff81111561234a57612349611eaa565b5b6123568782880161228a565b925050606085013567ffffffffffffffff81111561237757612376611eaa565b5b6123838782880161228a565b91505092959194509250565b600080604083850312156123a6576123a5611ea5565b5b60006123b48582860161206e565b92505060206123c585828601611f2e565b9150509250929050565b6123d881612045565b82525050565b60006020820190506123f360008301846123cf565b92915050565b600080604083850312156124105761240f611ea5565b5b600083013567ffffffffffffffff81111561242e5761242d611eaa565b5b61243a858286016120ec565b925050602083013567ffffffffffffffff81111561245b5761245a611eaa565b5b6124678582860161228a565b9150509250929050565b600067ffffffffffffffff82111561248c5761248b611f99565b5b602082029050602081019050919050565b60006124b06124ab84612471565b611ff9565b905080838252602082019050602084028301858111156124d3576124d2612040565b5b835b818110156124fc57806124e88882611ef8565b8452602084019350506020810190506124d5565b5050509392505050565b600082601f83011261251b5761251a611f83565b5b813561252b84826020860161249d565b91505092915050565b6000806040838503121561254b5761254a611ea5565b5b600083013567ffffffffffffffff81111561256957612568611eaa565b5b61257585828601612506565b925050602083013567ffffffffffffffff81111561259657612595611eaa565b5b6125a28582860161228a565b9150509250929050565b60006060820190506125c160008301866121bd565b6125ce60208301856121bd565b6125db60408301846121bd565b949350505050565b6000819050919050565b60006126086126036125fe84611eaf565b6125e3565b611eaf565b9050919050565b600061261a826125ed565b9050919050565b600061262c8261260f565b9050919050565b61263c81612621565b82525050565b60006020820190506126576000830184612633565b92915050565b60008115159050919050565b6126728161265d565b82525050565b600060208201905061268d6000830184612669565b92915050565b600082825260208201905092915050565b7f43616c6c6572206973206e6f742061206d616e61676572000000000000000000600082015250565b60006126da601783612693565b91506126e5826126a4565b602082019050919050565b60006020820190508181036000830152612709816126cd565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000612746602083612693565b915061275182612710565b602082019050919050565b6000602082019050818103600083015261277581612739565b9050919050565b7f496e76616c6964206172726179206c656e677468000000000000000000000000600082015250565b60006127b2601483612693565b91506127bd8261277c565b602082019050919050565b600060208201905081810360008301526127e1816127a5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061285182611f0d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561288457612883612817565b5b600182019050919050565b7f6164647265737320697320307830000000000000000000000000000000000000600082015250565b60006128c5600e83612693565b91506128d08261288f565b602082019050919050565b600060208201905081810360008301526128f4816128b8565b9050919050565b7f506c656173652070617373206174206c65617374206f6e65206d616e61676572600082015250565b6000612931602083612693565b915061293c826128fb565b602082019050919050565b6000602082019050818103600083015261296081612924565b9050919050565b7f6e617469766520746f6b656e2066756e64732063616e2774206265207265636f60008201527f7665726564000000000000000000000000000000000000000000000000000000602082015250565b60006129c3602583612693565b91506129ce82612967565b604082019050919050565b600060208201905081810360008301526129f2816129b6565b9050919050565b600081519050612a0881611f17565b92915050565b600060208284031215612a2457612a23611ea5565b5b6000612a32848285016129f9565b91505092915050565b6000612a468261260f565b9050919050565b612a5681612a3b565b82525050565b6000604082019050612a716000830185612a4d565b612a7e60208301846121bd565b9392505050565b612a8e8161265d565b8114612a9957600080fd5b50565b600081519050612aab81612a85565b92915050565b600060208284031215612ac757612ac6611ea5565b5b6000612ad584828501612a9c565b91505092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000612b3a602683612693565b9150612b4582612ade565b604082019050919050565b60006020820190508181036000830152612b6981612b2d565b9050919050565b7f5061796f7574206c696d69747320657863656564656400000000000000000000600082015250565b6000612ba6601683612693565b9150612bb182612b70565b602082019050919050565b60006020820190508181036000830152612bd581612b99565b9050919050565b6000612be782611f0d565b9150612bf283611f0d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612c2757612c26612817565b5b828201905092915050565b6000612c3d82611f0d565b9150612c4883611f0d565b925082821015612c5b57612c5a612817565b5b828203905092915050565b6000604082019050612c7b60008301856123cf565b612c8860208301846121bd565b939250505056fea26469706673582212203b14fac4064189dffae036bec2db6544f9b9270e536874abcdf7c784b29b290364736f6c63430008090033",
}

// TopperUpperABI is the input ABI used to generate the binding from.
// Deprecated: Use TopperUpperMetaData.ABI instead.
var TopperUpperABI = TopperUpperMetaData.ABI

// TopperUpperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TopperUpperMetaData.Bin instead.
var TopperUpperBin = TopperUpperMetaData.Bin

// DeployTopperUpper deploys a new Ethereum contract, binding an instance of TopperUpper to it.
func DeployTopperUpper(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *TopperUpper, error) {
	parsed, err := TopperUpperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TopperUpperBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TopperUpper{TopperUpperCaller: TopperUpperCaller{contract: contract}, TopperUpperTransactor: TopperUpperTransactor{contract: contract}, TopperUpperFilterer: TopperUpperFilterer{contract: contract}}, nil
}

// TopperUpper is an auto generated Go binding around an Ethereum contract.
type TopperUpper struct {
	TopperUpperCaller     // Read-only binding to the contract
	TopperUpperTransactor // Write-only binding to the contract
	TopperUpperFilterer   // Log filterer for contract events
}

// TopperUpperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TopperUpperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopperUpperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TopperUpperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopperUpperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TopperUpperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TopperUpperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TopperUpperSession struct {
	Contract     *TopperUpper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TopperUpperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TopperUpperCallerSession struct {
	Contract *TopperUpperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TopperUpperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TopperUpperTransactorSession struct {
	Contract     *TopperUpperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TopperUpperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TopperUpperRaw struct {
	Contract *TopperUpper // Generic contract binding to access the raw methods on
}

// TopperUpperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TopperUpperCallerRaw struct {
	Contract *TopperUpperCaller // Generic read-only contract binding to access the raw methods on
}

// TopperUpperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TopperUpperTransactorRaw struct {
	Contract *TopperUpperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTopperUpper creates a new instance of TopperUpper, bound to a specific deployed contract.
func NewTopperUpper(address common.Address, backend bind.ContractBackend) (*TopperUpper, error) {
	contract, err := bindTopperUpper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TopperUpper{TopperUpperCaller: TopperUpperCaller{contract: contract}, TopperUpperTransactor: TopperUpperTransactor{contract: contract}, TopperUpperFilterer: TopperUpperFilterer{contract: contract}}, nil
}

// NewTopperUpperCaller creates a new read-only instance of TopperUpper, bound to a specific deployed contract.
func NewTopperUpperCaller(address common.Address, caller bind.ContractCaller) (*TopperUpperCaller, error) {
	contract, err := bindTopperUpper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TopperUpperCaller{contract: contract}, nil
}

// NewTopperUpperTransactor creates a new write-only instance of TopperUpper, bound to a specific deployed contract.
func NewTopperUpperTransactor(address common.Address, transactor bind.ContractTransactor) (*TopperUpperTransactor, error) {
	contract, err := bindTopperUpper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TopperUpperTransactor{contract: contract}, nil
}

// NewTopperUpperFilterer creates a new log filterer instance of TopperUpper, bound to a specific deployed contract.
func NewTopperUpperFilterer(address common.Address, filterer bind.ContractFilterer) (*TopperUpperFilterer, error) {
	contract, err := bindTopperUpper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TopperUpperFilterer{contract: contract}, nil
}

// bindTopperUpper binds a generic wrapper to an already deployed contract.
func bindTopperUpper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TopperUpperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TopperUpper *TopperUpperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TopperUpper.Contract.TopperUpperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TopperUpper *TopperUpperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopperUpperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TopperUpper *TopperUpperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopperUpperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TopperUpper *TopperUpperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TopperUpper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TopperUpper *TopperUpperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopperUpper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TopperUpper *TopperUpperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TopperUpper.Contract.contract.Transact(opts, method, params...)
}

// ApprovedAddresses is a free data retrieval call binding the contract method 0xf136a874.
//
// Solidity: function approvedAddresses(address ) view returns(uint256 native, uint256 token, uint256 blocksWindow)
func (_TopperUpper *TopperUpperCaller) ApprovedAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
	Native       *big.Int
	Token        *big.Int
	BlocksWindow *big.Int
}, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "approvedAddresses", arg0)

	outstruct := new(struct {
		Native       *big.Int
		Token        *big.Int
		BlocksWindow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Native = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlocksWindow = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApprovedAddresses is a free data retrieval call binding the contract method 0xf136a874.
//
// Solidity: function approvedAddresses(address ) view returns(uint256 native, uint256 token, uint256 blocksWindow)
func (_TopperUpper *TopperUpperSession) ApprovedAddresses(arg0 common.Address) (struct {
	Native       *big.Int
	Token        *big.Int
	BlocksWindow *big.Int
}, error) {
	return _TopperUpper.Contract.ApprovedAddresses(&_TopperUpper.CallOpts, arg0)
}

// ApprovedAddresses is a free data retrieval call binding the contract method 0xf136a874.
//
// Solidity: function approvedAddresses(address ) view returns(uint256 native, uint256 token, uint256 blocksWindow)
func (_TopperUpper *TopperUpperCallerSession) ApprovedAddresses(arg0 common.Address) (struct {
	Native       *big.Int
	Token        *big.Int
	BlocksWindow *big.Int
}, error) {
	return _TopperUpper.Contract.ApprovedAddresses(&_TopperUpper.CallOpts, arg0)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_TopperUpper *TopperUpperCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "getFundsDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_TopperUpper *TopperUpperSession) GetFundsDestination() (common.Address, error) {
	return _TopperUpper.Contract.GetFundsDestination(&_TopperUpper.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_TopperUpper *TopperUpperCallerSession) GetFundsDestination() (common.Address, error) {
	return _TopperUpper.Contract.GetFundsDestination(&_TopperUpper.CallOpts)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_TopperUpper *TopperUpperCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "managers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_TopperUpper *TopperUpperSession) Managers(arg0 common.Address) (bool, error) {
	return _TopperUpper.Contract.Managers(&_TopperUpper.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_TopperUpper *TopperUpperCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _TopperUpper.Contract.Managers(&_TopperUpper.CallOpts, arg0)
}

// NativeLimits is a free data retrieval call binding the contract method 0xe7fc94c4.
//
// Solidity: function nativeLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperCaller) NativeLimits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "nativeLimits", arg0)

	outstruct := new(struct {
		Amount    *big.Int
		ValidTill *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidTill = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NativeLimits is a free data retrieval call binding the contract method 0xe7fc94c4.
//
// Solidity: function nativeLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperSession) NativeLimits(arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	return _TopperUpper.Contract.NativeLimits(&_TopperUpper.CallOpts, arg0)
}

// NativeLimits is a free data retrieval call binding the contract method 0xe7fc94c4.
//
// Solidity: function nativeLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperCallerSession) NativeLimits(arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	return _TopperUpper.Contract.NativeLimits(&_TopperUpper.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TopperUpper *TopperUpperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TopperUpper *TopperUpperSession) Owner() (common.Address, error) {
	return _TopperUpper.Contract.Owner(&_TopperUpper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TopperUpper *TopperUpperCallerSession) Owner() (common.Address, error) {
	return _TopperUpper.Contract.Owner(&_TopperUpper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TopperUpper *TopperUpperCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TopperUpper *TopperUpperSession) Token() (common.Address, error) {
	return _TopperUpper.Contract.Token(&_TopperUpper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TopperUpper *TopperUpperCallerSession) Token() (common.Address, error) {
	return _TopperUpper.Contract.Token(&_TopperUpper.CallOpts)
}

// TokenLimits is a free data retrieval call binding the contract method 0x417a0698.
//
// Solidity: function tokenLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperCaller) TokenLimits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "tokenLimits", arg0)

	outstruct := new(struct {
		Amount    *big.Int
		ValidTill *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidTill = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenLimits is a free data retrieval call binding the contract method 0x417a0698.
//
// Solidity: function tokenLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperSession) TokenLimits(arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	return _TopperUpper.Contract.TokenLimits(&_TopperUpper.CallOpts, arg0)
}

// TokenLimits is a free data retrieval call binding the contract method 0x417a0698.
//
// Solidity: function tokenLimits(address ) view returns(uint256 amount, uint256 validTill)
func (_TopperUpper *TopperUpperCallerSession) TokenLimits(arg0 common.Address) (struct {
	Amount    *big.Int
	ValidTill *big.Int
}, error) {
	return _TopperUpper.Contract.TokenLimits(&_TopperUpper.CallOpts, arg0)
}

// ApproveAddresses is a paid mutator transaction binding the contract method 0x4784e21b.
//
// Solidity: function approveAddresses(address[] _addrs, uint256[] _limitsNative, uint256[] _limitsToken, uint256[] _blocksWindow) returns()
func (_TopperUpper *TopperUpperTransactor) ApproveAddresses(opts *bind.TransactOpts, _addrs []common.Address, _limitsNative []*big.Int, _limitsToken []*big.Int, _blocksWindow []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "approveAddresses", _addrs, _limitsNative, _limitsToken, _blocksWindow)
}

// ApproveAddresses is a paid mutator transaction binding the contract method 0x4784e21b.
//
// Solidity: function approveAddresses(address[] _addrs, uint256[] _limitsNative, uint256[] _limitsToken, uint256[] _blocksWindow) returns()
func (_TopperUpper *TopperUpperSession) ApproveAddresses(_addrs []common.Address, _limitsNative []*big.Int, _limitsToken []*big.Int, _blocksWindow []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.ApproveAddresses(&_TopperUpper.TransactOpts, _addrs, _limitsNative, _limitsToken, _blocksWindow)
}

// ApproveAddresses is a paid mutator transaction binding the contract method 0x4784e21b.
//
// Solidity: function approveAddresses(address[] _addrs, uint256[] _limitsNative, uint256[] _limitsToken, uint256[] _blocksWindow) returns()
func (_TopperUpper *TopperUpperTransactorSession) ApproveAddresses(_addrs []common.Address, _limitsNative []*big.Int, _limitsToken []*big.Int, _blocksWindow []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.ApproveAddresses(&_TopperUpper.TransactOpts, _addrs, _limitsNative, _limitsToken, _blocksWindow)
}

// ClaimNative is a paid mutator transaction binding the contract method 0xf4856230.
//
// Solidity: function claimNative() returns()
func (_TopperUpper *TopperUpperTransactor) ClaimNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "claimNative")
}

// ClaimNative is a paid mutator transaction binding the contract method 0xf4856230.
//
// Solidity: function claimNative() returns()
func (_TopperUpper *TopperUpperSession) ClaimNative() (*types.Transaction, error) {
	return _TopperUpper.Contract.ClaimNative(&_TopperUpper.TransactOpts)
}

// ClaimNative is a paid mutator transaction binding the contract method 0xf4856230.
//
// Solidity: function claimNative() returns()
func (_TopperUpper *TopperUpperTransactorSession) ClaimNative() (*types.Transaction, error) {
	return _TopperUpper.Contract.ClaimNative(&_TopperUpper.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_TopperUpper *TopperUpperTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_TopperUpper *TopperUpperSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.ClaimTokens(&_TopperUpper.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_TopperUpper *TopperUpperTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.ClaimTokens(&_TopperUpper.TransactOpts, _token)
}

// DisapproveAddresses is a paid mutator transaction binding the contract method 0x4f85d3db.
//
// Solidity: function disapproveAddresses(address[] _addrs) returns()
func (_TopperUpper *TopperUpperTransactor) DisapproveAddresses(opts *bind.TransactOpts, _addrs []common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "disapproveAddresses", _addrs)
}

// DisapproveAddresses is a paid mutator transaction binding the contract method 0x4f85d3db.
//
// Solidity: function disapproveAddresses(address[] _addrs) returns()
func (_TopperUpper *TopperUpperSession) DisapproveAddresses(_addrs []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.DisapproveAddresses(&_TopperUpper.TransactOpts, _addrs)
}

// DisapproveAddresses is a paid mutator transaction binding the contract method 0x4f85d3db.
//
// Solidity: function disapproveAddresses(address[] _addrs) returns()
func (_TopperUpper *TopperUpperTransactorSession) DisapproveAddresses(_addrs []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.DisapproveAddresses(&_TopperUpper.TransactOpts, _addrs)
}

// RemoveManagers is a paid mutator transaction binding the contract method 0x10f3ee29.
//
// Solidity: function removeManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperTransactor) RemoveManagers(opts *bind.TransactOpts, _managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "removeManagers", _managers)
}

// RemoveManagers is a paid mutator transaction binding the contract method 0x10f3ee29.
//
// Solidity: function removeManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperSession) RemoveManagers(_managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.RemoveManagers(&_TopperUpper.TransactOpts, _managers)
}

// RemoveManagers is a paid mutator transaction binding the contract method 0x10f3ee29.
//
// Solidity: function removeManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperTransactorSession) RemoveManagers(_managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.RemoveManagers(&_TopperUpper.TransactOpts, _managers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TopperUpper *TopperUpperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TopperUpper *TopperUpperSession) RenounceOwnership() (*types.Transaction, error) {
	return _TopperUpper.Contract.RenounceOwnership(&_TopperUpper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TopperUpper *TopperUpperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TopperUpper.Contract.RenounceOwnership(&_TopperUpper.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_TopperUpper *TopperUpperTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_TopperUpper *TopperUpperSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetFundsDestination(&_TopperUpper.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_TopperUpper *TopperUpperTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetFundsDestination(&_TopperUpper.TransactOpts, _newDestination)
}

// SetManagers is a paid mutator transaction binding the contract method 0x367caa09.
//
// Solidity: function setManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperTransactor) SetManagers(opts *bind.TransactOpts, _managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "setManagers", _managers)
}

// SetManagers is a paid mutator transaction binding the contract method 0x367caa09.
//
// Solidity: function setManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperSession) SetManagers(_managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetManagers(&_TopperUpper.TransactOpts, _managers)
}

// SetManagers is a paid mutator transaction binding the contract method 0x367caa09.
//
// Solidity: function setManagers(address[] _managers) returns()
func (_TopperUpper *TopperUpperTransactorSession) SetManagers(_managers []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetManagers(&_TopperUpper.TransactOpts, _managers)
}

// TopupNative is a paid mutator transaction binding the contract method 0x0d4e1a1e.
//
// Solidity: function topupNative(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupNative(opts *bind.TransactOpts, _to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupNative", _to, _amounts)
}

// TopupNative is a paid mutator transaction binding the contract method 0x0d4e1a1e.
//
// Solidity: function topupNative(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupNative(_to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupNative(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupNative is a paid mutator transaction binding the contract method 0x0d4e1a1e.
//
// Solidity: function topupNative(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupNative(_to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupNative(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupNatives is a paid mutator transaction binding the contract method 0xd4f2b668.
//
// Solidity: function topupNatives(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupNatives(opts *bind.TransactOpts, _to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupNatives", _to, _amounts)
}

// TopupNatives is a paid mutator transaction binding the contract method 0xd4f2b668.
//
// Solidity: function topupNatives(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupNatives(_to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupNatives(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupNatives is a paid mutator transaction binding the contract method 0xd4f2b668.
//
// Solidity: function topupNatives(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupNatives(_to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupNatives(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupToken is a paid mutator transaction binding the contract method 0x714280b3.
//
// Solidity: function topupToken(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupToken(opts *bind.TransactOpts, _to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupToken", _to, _amounts)
}

// TopupToken is a paid mutator transaction binding the contract method 0x714280b3.
//
// Solidity: function topupToken(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupToken(_to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupToken(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupToken is a paid mutator transaction binding the contract method 0x714280b3.
//
// Solidity: function topupToken(address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupToken(_to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupToken(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x91b571cf.
//
// Solidity: function topupTokens(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupTokens(opts *bind.TransactOpts, _to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupTokens", _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x91b571cf.
//
// Solidity: function topupTokens(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupTokens(_to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupTokens(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x91b571cf.
//
// Solidity: function topupTokens(address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupTokens(_to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupTokens(&_TopperUpper.TransactOpts, _to, _amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TopperUpper *TopperUpperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TopperUpper *TopperUpperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.TransferOwnership(&_TopperUpper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TopperUpper *TopperUpperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.TransferOwnership(&_TopperUpper.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TopperUpper *TopperUpperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TopperUpper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TopperUpper *TopperUpperSession) Receive() (*types.Transaction, error) {
	return _TopperUpper.Contract.Receive(&_TopperUpper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TopperUpper *TopperUpperTransactorSession) Receive() (*types.Transaction, error) {
	return _TopperUpper.Contract.Receive(&_TopperUpper.TransactOpts)
}

// TopperUpperDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the TopperUpper contract.
type TopperUpperDestinationChangedIterator struct {
	Event *TopperUpperDestinationChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TopperUpperDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TopperUpperDestinationChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TopperUpperDestinationChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TopperUpperDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TopperUpperDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TopperUpperDestinationChanged represents a DestinationChanged event raised by the TopperUpper contract.
type TopperUpperDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_TopperUpper *TopperUpperFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*TopperUpperDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _TopperUpper.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &TopperUpperDestinationChangedIterator{contract: _TopperUpper.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_TopperUpper *TopperUpperFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *TopperUpperDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _TopperUpper.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TopperUpperDestinationChanged)
				if err := _TopperUpper.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDestinationChanged is a log parse operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_TopperUpper *TopperUpperFilterer) ParseDestinationChanged(log types.Log) (*TopperUpperDestinationChanged, error) {
	event := new(TopperUpperDestinationChanged)
	if err := _TopperUpper.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TopperUpperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TopperUpper contract.
type TopperUpperOwnershipTransferredIterator struct {
	Event *TopperUpperOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TopperUpperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TopperUpperOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TopperUpperOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TopperUpperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TopperUpperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TopperUpperOwnershipTransferred represents a OwnershipTransferred event raised by the TopperUpper contract.
type TopperUpperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TopperUpper *TopperUpperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TopperUpperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TopperUpper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TopperUpperOwnershipTransferredIterator{contract: _TopperUpper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TopperUpper *TopperUpperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TopperUpperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TopperUpper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TopperUpperOwnershipTransferred)
				if err := _TopperUpper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TopperUpper *TopperUpperFilterer) ParseOwnershipTransferred(log types.Log) (*TopperUpperOwnershipTransferred, error) {
	event := new(TopperUpperOwnershipTransferred)
	if err := _TopperUpper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
