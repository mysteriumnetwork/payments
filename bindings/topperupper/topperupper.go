// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package topperupper

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TopperUpperABI is the input ABI used to generate the binding from.
const TopperUpperABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limitsNative\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limitsToken\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_blocksWindow\",\"type\":\"uint256[]\"}],\"name\":\"approveAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"native\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"native\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"disapproveAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"moderators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_moderators\",\"type\":\"address[]\"}],\"name\":\"removeModerators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_moderators\",\"type\":\"address[]\"}],\"name\":\"setModerators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"topupNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"topupNatives\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amounts\",\"type\":\"uint256\"}],\"name\":\"topupToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"topupTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TopperUpperBin is the compiled bytecode used for deploying new contracts.
var TopperUpperBin = "0x60806040523480156200001157600080fd5b5062000032620000266200003860201b60201c565b6200004060201b60201c565b62000104565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612d3380620001146000396000f3fe6080604052600436106101185760003560e01c80638da5cb5b116100a0578063f136a87411610064578063f136a8741461037d578063f2fde38b146103bc578063f4856230146103e5578063f58c5b6e146103fc578063fc0c546a146104275761011f565b80638da5cb5b146102ae5780639f2e07b3146102d9578063d4f2b66814610302578063df8de3e71461032b578063f0fdf5b3146103545761011f565b80634784e21b116100e75780634784e21b146101dc5780634f85d3db14610205578063681d57981461022e578063715018a61461025757806375a81c9e1461026e5761011f565b80630d4e1a1e1461012457806314d0f1ba1461014d5780631c990ecf1461018a578063238e130a146101b35761011f565b3661011f57005b600080fd5b34801561013057600080fd5b5061014b600480360381019061014691906121c6565b610452565b005b34801561015957600080fd5b50610174600480360381019061016f9190612174565b6104f3565b604051610181919061267a565b60405180910390f35b34801561019657600080fd5b506101b160048036038101906101ac9190612251565b610513565b005b3480156101bf57600080fd5b506101da60048036038101906101d5919061219d565b610685565b005b3480156101e857600080fd5b5061020360048036038101906101fe9190612329565b610831565b005b34801561021157600080fd5b5061022c60048036038101906102279190612251565b610ab2565b005b34801561023a57600080fd5b5061025560048036038101906102509190612251565b610cd1565b005b34801561026357600080fd5b5061026c610e4c565b005b34801561027a57600080fd5b5061029560048036038101906102909190612174565b610ed4565b6040516102a594939291906127e7565b60405180910390f35b3480156102ba57600080fd5b506102c3610f04565b6040516102d0919061260d565b60405180910390f35b3480156102e557600080fd5b5061030060048036038101906102fb9190612292565b610f2d565b005b34801561030e57600080fd5b50610329600480360381019061032491906123ec565b611101565b005b34801561033757600080fd5b50610352600480360381019061034d9190612174565b611286565b005b34801561036057600080fd5b5061037b60048036038101906103769190612202565b6114ea565b005b34801561038957600080fd5b506103a4600480360381019061039f9190612174565b61158d565b6040516103b3939291906127b0565b60405180910390f35b3480156103c857600080fd5b506103e360048036038101906103de9190612174565b6115b7565b005b3480156103f157600080fd5b506103fa6116af565b005b34801561040857600080fd5b506104116117ac565b60405161041e919061260d565b60405180910390f35b34801561043357600080fd5b5061043c611852565b6040516104499190612695565b60405180910390f35b6003600061045e611878565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc90612710565b60405180910390fd5b6104ef8282611880565b5050565b60036020528060005260406000206000915054906101000a900460ff1681565b61051b611878565b73ffffffffffffffffffffffffffffffffffffffff16610539610f04565b73ffffffffffffffffffffffffffffffffffffffff161461058f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058690612750565b60405180910390fd5b60008151116105d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ca906126b0565b60405180910390fd5b60005b8151811015610681576003600083838151811061061c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055808061067990612a55565b9150506105d6565b5050565b61068d611878565b73ffffffffffffffffffffffffffffffffffffffff166106ab610f04565b73ffffffffffffffffffffffffffffffffffffffff1614610701576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f890612750565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610771576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076890612730565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610839611878565b73ffffffffffffffffffffffffffffffffffffffff16610857610f04565b73ffffffffffffffffffffffffffffffffffffffff16146108ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a490612750565b60405180910390fd5b825184511480156108bf575081518351145b80156108cc575081518151145b61090b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610902906126b0565b60405180910390fd5b60005b8451811015610aab576000604051806060016040528086848151811061095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015181526020018584815181106109a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015181526020018484815181106109e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015181525090508060046000888581518110610a34577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050508080610aa390612a55565b91505061090e565b5050505050565b610aba611878565b73ffffffffffffffffffffffffffffffffffffffff16610ad8610f04565b73ffffffffffffffffffffffffffffffffffffffff1614610b2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2590612750565b60405180910390fd5b6000815111610b72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b69906126b0565b60405180910390fd5b60005b8151811015610ccd5760046000838381518110610bbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160009055505060056000838381518110610c55577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160009055600382016000905550508080610cc590612a55565b915050610b75565b5050565b610cd9611878565b73ffffffffffffffffffffffffffffffffffffffff16610cf7610f04565b73ffffffffffffffffffffffffffffffffffffffff1614610d4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4490612750565b60405180910390fd5b6000815111610d91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d88906126b0565b60405180910390fd5b60005b8151811015610e4857600160036000848481518110610ddc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508080610e4090612a55565b915050610d94565b5050565b610e54611878565b73ffffffffffffffffffffffffffffffffffffffff16610e72610f04565b73ffffffffffffffffffffffffffffffffffffffff1614610ec8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebf90612750565b60405180910390fd5b610ed260006119b1565b565b60056020528060005260406000206000915090508060000154908060010154908060020154908060030154905084565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60036000610f39611878565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610fc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb790612710565b60405180910390fd5b80518351148015610fd2575081518351145b611011576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611008906126b0565b60405180910390fd5b60005b83518110156110fb576110e8848281518110611059577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015184838151811061109a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518484815181106110db577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151611a75565b80806110f390612a55565b915050611014565b50505050565b6003600061110d611878565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611194576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118b90612710565b60405180910390fd5b81518151146111d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111cf906126b0565b60405180910390fd5b60005b82518110156112815761126e838281518110611220577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151838381518110611261577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151611880565b808061127990612a55565b9150506111db565b505050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611318576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130f90612730565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156113a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a0906126f0565b60405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016113e4919061260d565b60206040518083038186803b1580156113fc57600080fd5b505afa158015611410573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114349190612481565b90508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401611493929190612628565b602060405180830381600087803b1580156114ad57600080fd5b505af11580156114c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e59190612458565b505050565b600360006114f6611878565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661157d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157490612710565b60405180910390fd5b611588838383611a75565b505050565b60046020528060005260406000206000915090508060000154908060010154908060020154905083565b6115bf611878565b73ffffffffffffffffffffffffffffffffffffffff166115dd610f04565b73ffffffffffffffffffffffffffffffffffffffff1614611633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162a90612750565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156116a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169a906126d0565b60405180910390fd5b6116ac816119b1565b50565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611741576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173890612730565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156117a9573d6000803e3d6000fd5b50565b60006117b6611878565b73ffffffffffffffffffffffffffffffffffffffff166117d4610f04565b73ffffffffffffffffffffffffffffffffffffffff161461182a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161182190612750565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b61188982611bee565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541161190d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190490612790565b60405180910390fd5b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082825461195f919061293c565b925050819055508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156119ac573d6000803e3d6000fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611a7e82611bee565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015411611b02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611af990612790565b60405180910390fd5b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254611b54919061293c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401611b96929190612651565b602060405180830381600087803b158015611bb057600080fd5b505af1158015611bc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611be89190612458565b50505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015411611c73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c6a90612770565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201541480611d05575043600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154105b8015611d5357506000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154115b15611f465743600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015443611ded91906128e6565b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505b50565b6000611f5c611f5784612851565b61282c565b90508083825260208201905082856020860282011115611f7b57600080fd5b60005b85811015611fab5781611f91888261208d565b845260208401935060208301925050600181019050611f7e565b5050509392505050565b6000611fc8611fc38461287d565b61282c565b90508083825260208201905082856020860282011115611fe757600080fd5b60005b858110156120175781611ffd88826120a2565b845260208401935060208301925050600181019050611fea565b5050509392505050565b600061203461202f846128a9565b61282c565b9050808382526020820190508285602086028201111561205357600080fd5b60005b858110156120835781612069888261214a565b845260208401935060208301925050600181019050612056565b5050509392505050565b60008135905061209c81612ca1565b92915050565b6000813590506120b181612cb8565b92915050565b600082601f8301126120c857600080fd5b81356120d8848260208601611f49565b91505092915050565b600082601f8301126120f257600080fd5b8135612102848260208601611fb5565b91505092915050565b600082601f83011261211c57600080fd5b813561212c848260208601612021565b91505092915050565b60008151905061214481612ccf565b92915050565b60008135905061215981612ce6565b92915050565b60008151905061216e81612ce6565b92915050565b60006020828403121561218657600080fd5b60006121948482850161208d565b91505092915050565b6000602082840312156121af57600080fd5b60006121bd848285016120a2565b91505092915050565b600080604083850312156121d957600080fd5b60006121e7858286016120a2565b92505060206121f88582860161214a565b9150509250929050565b60008060006060848603121561221757600080fd5b60006122258682870161208d565b93505060206122368682870161208d565b92505060406122478682870161214a565b9150509250925092565b60006020828403121561226357600080fd5b600082013567ffffffffffffffff81111561227d57600080fd5b612289848285016120b7565b91505092915050565b6000806000606084860312156122a757600080fd5b600084013567ffffffffffffffff8111156122c157600080fd5b6122cd868287016120b7565b935050602084013567ffffffffffffffff8111156122ea57600080fd5b6122f6868287016120b7565b925050604084013567ffffffffffffffff81111561231357600080fd5b61231f8682870161210b565b9150509250925092565b6000806000806080858703121561233f57600080fd5b600085013567ffffffffffffffff81111561235957600080fd5b612365878288016120b7565b945050602085013567ffffffffffffffff81111561238257600080fd5b61238e8782880161210b565b935050604085013567ffffffffffffffff8111156123ab57600080fd5b6123b78782880161210b565b925050606085013567ffffffffffffffff8111156123d457600080fd5b6123e08782880161210b565b91505092959194509250565b600080604083850312156123ff57600080fd5b600083013567ffffffffffffffff81111561241957600080fd5b612425858286016120e1565b925050602083013567ffffffffffffffff81111561244257600080fd5b61244e8582860161210b565b9150509250929050565b60006020828403121561246a57600080fd5b600061247884828501612135565b91505092915050565b60006020828403121561249357600080fd5b60006124a18482850161215f565b91505092915050565b6124b3816129ca565b82525050565b6124c281612970565b82525050565b6124d181612994565b82525050565b6124e0816129dc565b82525050565b60006124f36014836128d5565b91506124fe82612b0d565b602082019050919050565b60006125166026836128d5565b915061252182612b36565b604082019050919050565b60006125396025836128d5565b915061254482612b85565b604082019050919050565b600061255c601b836128d5565b915061256782612bd4565b602082019050919050565b600061257f600e836128d5565b915061258a82612bfd565b602082019050919050565b60006125a26020836128d5565b91506125ad82612c26565b602082019050919050565b60006125c56014836128d5565b91506125d082612c4f565b602082019050919050565b60006125e86015836128d5565b91506125f382612c78565b602082019050919050565b612607816129c0565b82525050565b600060208201905061262260008301846124b9565b92915050565b600060408201905061263d60008301856124aa565b61264a60208301846125fe565b9392505050565b600060408201905061266660008301856124b9565b61267360208301846125fe565b9392505050565b600060208201905061268f60008301846124c8565b92915050565b60006020820190506126aa60008301846124d7565b92915050565b600060208201905081810360008301526126c9816124e6565b9050919050565b600060208201905081810360008301526126e981612509565b9050919050565b600060208201905081810360008301526127098161252c565b9050919050565b600060208201905081810360008301526127298161254f565b9050919050565b6000602082019050818103600083015261274981612572565b9050919050565b6000602082019050818103600083015261276981612595565b9050919050565b60006020820190508181036000830152612789816125b8565b9050919050565b600060208201905081810360008301526127a9816125db565b9050919050565b60006060820190506127c560008301866125fe565b6127d260208301856125fe565b6127df60408301846125fe565b949350505050565b60006080820190506127fc60008301876125fe565b61280960208301866125fe565b61281660408301856125fe565b61282360608301846125fe565b95945050505050565b6000612836612847565b90506128428282612a24565b919050565b6000604051905090565b600067ffffffffffffffff82111561286c5761286b612acd565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561289857612897612acd565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156128c4576128c3612acd565b5b602082029050602081019050919050565b600082825260208201905092915050565b60006128f1826129c0565b91506128fc836129c0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561293157612930612a9e565b5b828201905092915050565b6000612947826129c0565b9150612952836129c0565b92508282101561296557612964612a9e565b5b828203905092915050565b600061297b826129a0565b9050919050565b600061298d826129a0565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006129d582612a00565b9050919050565b60006129e7826129ee565b9050919050565b60006129f9826129a0565b9050919050565b6000612a0b82612a12565b9050919050565b6000612a1d826129a0565b9050919050565b612a2d82612afc565b810181811067ffffffffffffffff82111715612a4c57612a4b612acd565b5b80604052505050565b6000612a60826129c0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415612a9357612a92612a9e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e76616c6964206172726179206c656e677468000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f6e617469766520746f6b656e2066756e64732063616e2774206265207265636f60008201527f7665726564000000000000000000000000000000000000000000000000000000602082015250565b7f43616c6c6572206973206e6f7420746865206d6f64657261746f720000000000600082015250565b7f6164647265737320697320307830000000000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f41646472657373206e6f7420617070726f766564000000000000000000000000600082015250565b7f41646472657373206f7574206f66206c696d6974730000000000000000000000600082015250565b612caa81612970565b8114612cb557600080fd5b50565b612cc181612982565b8114612ccc57600080fd5b50565b612cd881612994565b8114612ce357600080fd5b50565b612cef816129c0565b8114612cfa57600080fd5b5056fea2646970667358221220e558bccf9da96465d53d29f83e72dbbc5e00908f62ed796d4b8ff1f3b6291b1a64736f6c63430008040033"

// DeployTopperUpper deploys a new Ethereum contract, binding an instance of TopperUpper to it.
func DeployTopperUpper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TopperUpper, error) {
	parsed, err := abi.JSON(strings.NewReader(TopperUpperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TopperUpperBin), backend)
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

// CurrentLimits is a free data retrieval call binding the contract method 0x75a81c9e.
//
// Solidity: function currentLimits(address ) view returns(uint256 native, uint256 token, uint256 startBlock, uint256 endBlock)
func (_TopperUpper *TopperUpperCaller) CurrentLimits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Native     *big.Int
	Token      *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "currentLimits", arg0)

	outstruct := new(struct {
		Native     *big.Int
		Token      *big.Int
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Native = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentLimits is a free data retrieval call binding the contract method 0x75a81c9e.
//
// Solidity: function currentLimits(address ) view returns(uint256 native, uint256 token, uint256 startBlock, uint256 endBlock)
func (_TopperUpper *TopperUpperSession) CurrentLimits(arg0 common.Address) (struct {
	Native     *big.Int
	Token      *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _TopperUpper.Contract.CurrentLimits(&_TopperUpper.CallOpts, arg0)
}

// CurrentLimits is a free data retrieval call binding the contract method 0x75a81c9e.
//
// Solidity: function currentLimits(address ) view returns(uint256 native, uint256 token, uint256 startBlock, uint256 endBlock)
func (_TopperUpper *TopperUpperCallerSession) CurrentLimits(arg0 common.Address) (struct {
	Native     *big.Int
	Token      *big.Int
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _TopperUpper.Contract.CurrentLimits(&_TopperUpper.CallOpts, arg0)
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

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_TopperUpper *TopperUpperCaller) Moderators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TopperUpper.contract.Call(opts, &out, "moderators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_TopperUpper *TopperUpperSession) Moderators(arg0 common.Address) (bool, error) {
	return _TopperUpper.Contract.Moderators(&_TopperUpper.CallOpts, arg0)
}

// Moderators is a free data retrieval call binding the contract method 0x14d0f1ba.
//
// Solidity: function moderators(address ) view returns(bool)
func (_TopperUpper *TopperUpperCallerSession) Moderators(arg0 common.Address) (bool, error) {
	return _TopperUpper.Contract.Moderators(&_TopperUpper.CallOpts, arg0)
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

// RemoveModerators is a paid mutator transaction binding the contract method 0x1c990ecf.
//
// Solidity: function removeModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperTransactor) RemoveModerators(opts *bind.TransactOpts, _moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "removeModerators", _moderators)
}

// RemoveModerators is a paid mutator transaction binding the contract method 0x1c990ecf.
//
// Solidity: function removeModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperSession) RemoveModerators(_moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.RemoveModerators(&_TopperUpper.TransactOpts, _moderators)
}

// RemoveModerators is a paid mutator transaction binding the contract method 0x1c990ecf.
//
// Solidity: function removeModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperTransactorSession) RemoveModerators(_moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.RemoveModerators(&_TopperUpper.TransactOpts, _moderators)
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

// SetModerators is a paid mutator transaction binding the contract method 0x681d5798.
//
// Solidity: function setModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperTransactor) SetModerators(opts *bind.TransactOpts, _moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "setModerators", _moderators)
}

// SetModerators is a paid mutator transaction binding the contract method 0x681d5798.
//
// Solidity: function setModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperSession) SetModerators(_moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetModerators(&_TopperUpper.TransactOpts, _moderators)
}

// SetModerators is a paid mutator transaction binding the contract method 0x681d5798.
//
// Solidity: function setModerators(address[] _moderators) returns()
func (_TopperUpper *TopperUpperTransactorSession) SetModerators(_moderators []common.Address) (*types.Transaction, error) {
	return _TopperUpper.Contract.SetModerators(&_TopperUpper.TransactOpts, _moderators)
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

// TopupToken is a paid mutator transaction binding the contract method 0xf0fdf5b3.
//
// Solidity: function topupToken(address _token, address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupToken", _token, _to, _amounts)
}

// TopupToken is a paid mutator transaction binding the contract method 0xf0fdf5b3.
//
// Solidity: function topupToken(address _token, address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupToken(_token common.Address, _to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupToken(&_TopperUpper.TransactOpts, _token, _to, _amounts)
}

// TopupToken is a paid mutator transaction binding the contract method 0xf0fdf5b3.
//
// Solidity: function topupToken(address _token, address _to, uint256 _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupToken(_token common.Address, _to common.Address, _amounts *big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupToken(&_TopperUpper.TransactOpts, _token, _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x9f2e07b3.
//
// Solidity: function topupTokens(address[] _tokens, address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactor) TopupTokens(opts *bind.TransactOpts, _tokens []common.Address, _to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.contract.Transact(opts, "topupTokens", _tokens, _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x9f2e07b3.
//
// Solidity: function topupTokens(address[] _tokens, address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperSession) TopupTokens(_tokens []common.Address, _to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupTokens(&_TopperUpper.TransactOpts, _tokens, _to, _amounts)
}

// TopupTokens is a paid mutator transaction binding the contract method 0x9f2e07b3.
//
// Solidity: function topupTokens(address[] _tokens, address[] _to, uint256[] _amounts) returns()
func (_TopperUpper *TopperUpperTransactorSession) TopupTokens(_tokens []common.Address, _to []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _TopperUpper.Contract.TopupTokens(&_TopperUpper.TransactOpts, _tokens, _to, _amounts)
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
