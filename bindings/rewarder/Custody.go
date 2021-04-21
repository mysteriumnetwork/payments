// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarder

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

// CustodyABI is the input ABI used to generate the binding from.
const CustodyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"authorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"forbid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"payout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CustodyBin is the compiled bytecode used for deploying new contracts.
var CustodyBin = "0x60806040523480156200001157600080fd5b5060405162001711380380620017118339818101604052810190620000379190620001dd565b6000620000496200019560201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060006200013d6200019d60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505062000257565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081519050620001d7816200023d565b92915050565b600060208284031215620001f057600080fd5b60006200020084828501620001c6565b91505092915050565b600062000216826200021d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620002488162000209565b81146200025457600080fd5b50565b6114aa80620002676000396000f3fe6080604052600436106100955760003560e01c8063b6a5d7de11610059578063b6a5d7de14610197578063b9181611146101c0578063c176e639146101fd578063f2fde38b14610226578063fc0c546a1461024f576100d5565b80632e1a7d4d146100da5780635705ae4314610103578063715018a61461012c5780637e95cd27146101435780638da5cb5b1461016c576100d5565b366100d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cc90611198565b60405180910390fd5b600080fd5b3480156100e657600080fd5b5061010160048036038101906100fc9190610f87565b61027a565b005b34801561010f57600080fd5b5061012a60048036038101906101259190610ead565b6103b9565b005b34801561013857600080fd5b506101416104c7565b005b34801561014f57600080fd5b5061016a60048036038101906101659190610e84565b610601565b005b34801561017857600080fd5b5061018161074e565b60405161018e91906110be565b60405180910390f35b3480156101a357600080fd5b506101be60048036038101906101b99190610e84565b610777565b005b3480156101cc57600080fd5b506101e760048036038101906101e29190610e84565b61084d565b6040516101f49190611102565b60405180910390f35b34801561020957600080fd5b50610224600480360381019061021f9190610ee9565b61086d565b005b34801561023257600080fd5b5061024d60048036038101906102489190610e84565b610a95565b005b34801561025b57600080fd5b50610264610bda565b604051610271919061111d565b60405180910390f35b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610306576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fd906111d8565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016103639291906110d9565b602060405180830381600087803b15801561037d57600080fd5b505af1158015610391573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b59190610f5e565b5050565b6103c1610c00565b73ffffffffffffffffffffffffffffffffffffffff166103df61074e565b73ffffffffffffffffffffffffffffffffffffffff1614610435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042c906111b8565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016104709291906110d9565b602060405180830381600087803b15801561048a57600080fd5b505af115801561049e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c29190610f5e565b505050565b6104cf610c00565b73ffffffffffffffffffffffffffffffffffffffff166104ed61074e565b73ffffffffffffffffffffffffffffffffffffffff1614610543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053a906111b8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b610609610c00565b73ffffffffffffffffffffffffffffffffffffffff1661062761074e565b73ffffffffffffffffffffffffffffffffffffffff161461067d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610674906111b8565b60405180910390fd5b61068561074e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ea90611178565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61077f610c00565b73ffffffffffffffffffffffffffffffffffffffff1661079d61074e565b73ffffffffffffffffffffffffffffffffffffffff16146107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ea906111b8565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915054906101000a900460ff1681565b610875610c00565b73ffffffffffffffffffffffffffffffffffffffff1661089361074e565b73ffffffffffffffffffffffffffffffffffffffff16146108e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e0906111b8565b60405180910390fd5b818190508484905014610931576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092890611138565b60405180910390fd5b60005b84849050811015610a8e57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8686848181106109b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020160208101906109cb9190610e84565b858585818110610a04577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201356040518363ffffffff1660e01b8152600401610a289291906110d9565b602060405180830381600087803b158015610a4257600080fd5b505af1158015610a56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7a9190610f5e565b508080610a8690611275565b915050610934565b5050505050565b610a9d610c00565b73ffffffffffffffffffffffffffffffffffffffff16610abb61074e565b73ffffffffffffffffffffffffffffffffffffffff1614610b11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b08906111b8565b60405180910390fd5b600060016000610b1f61074e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610b7981610c08565b6001806000610b8661074e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b610c10610c00565b73ffffffffffffffffffffffffffffffffffffffff16610c2e61074e565b73ffffffffffffffffffffffffffffffffffffffff1614610c84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7b906111b8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cf4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ceb90611158565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600081359050610dc08161142f565b92915050565b60008083601f840112610dd857600080fd5b8235905067ffffffffffffffff811115610df157600080fd5b602083019150836020820283011115610e0957600080fd5b9250929050565b60008083601f840112610e2257600080fd5b8235905067ffffffffffffffff811115610e3b57600080fd5b602083019150836020820283011115610e5357600080fd5b9250929050565b600081519050610e6981611446565b92915050565b600081359050610e7e8161145d565b92915050565b600060208284031215610e9657600080fd5b6000610ea484828501610db1565b91505092915050565b60008060408385031215610ec057600080fd5b6000610ece85828601610db1565b9250506020610edf85828601610e6f565b9150509250929050565b60008060008060408587031215610eff57600080fd5b600085013567ffffffffffffffff811115610f1957600080fd5b610f2587828801610dc6565b9450945050602085013567ffffffffffffffff811115610f4457600080fd5b610f5087828801610e10565b925092505092959194509250565b600060208284031215610f7057600080fd5b6000610f7e84828501610e5a565b91505092915050565b600060208284031215610f9957600080fd5b6000610fa784828501610e6f565b91505092915050565b610fb981611209565b82525050565b610fc88161121b565b82525050565b610fd781611251565b82525050565b6000610fea6014836111f8565b9150610ff5826112ed565b602082019050919050565b600061100d6026836111f8565b915061101882611316565b604082019050919050565b60006110306021836111f8565b915061103b82611365565b604082019050919050565b6000611053601d836111f8565b915061105e826113b4565b602082019050919050565b60006110766020836111f8565b9150611081826113dd565b602082019050919050565b6000611099600e836111f8565b91506110a482611406565b602082019050919050565b6110b881611247565b82525050565b60006020820190506110d36000830184610fb0565b92915050565b60006040820190506110ee6000830185610fb0565b6110fb60208301846110af565b9392505050565b60006020820190506111176000830184610fbf565b92915050565b60006020820190506111326000830184610fce565b92915050565b6000602082019050818103600083015261115181610fdd565b9050919050565b6000602082019050818103600083015261117181611000565b9050919050565b6000602082019050818103600083015261119181611023565b9050919050565b600060208201905081810360008301526111b181611046565b9050919050565b600060208201905081810360008301526111d181611069565b9050919050565b600060208201905081810360008301526111f18161108c565b9050919050565b600082825260208201905092915050565b600061121482611227565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061125c82611263565b9050919050565b600061126e82611227565b9050919050565b600061128082611247565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156112b3576112b26112be565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f496e76616c6964206172726179206c656e677468000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e6572206163636573732063616e6e6f7420626520666f7262696464656e60008201527f2100000000000000000000000000000000000000000000000000000000000000602082015250565b7f52656a656374696e672074782077697468206574686572732073656e74000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f4e6f7420617574686f72697a6564000000000000000000000000000000000000600082015250565b61143881611209565b811461144357600080fd5b50565b61144f8161121b565b811461145a57600080fd5b50565b61146681611247565b811461147157600080fd5b5056fea2646970667358221220093aa147483a4c6fd11d95ff9881aef72a0eb78001ff46e68de1e5481a36570c64736f6c63430008030033"

// DeployCustody deploys a new Ethereum contract, binding an instance of Custody to it.
func DeployCustody(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address) (common.Address, *types.Transaction, *Custody, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CustodyBin), backend, _tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Custody{CustodyCaller: CustodyCaller{contract: contract}, CustodyTransactor: CustodyTransactor{contract: contract}, CustodyFilterer: CustodyFilterer{contract: contract}}, nil
}

// Custody is an auto generated Go binding around an Ethereum contract.
type Custody struct {
	CustodyCaller     // Read-only binding to the contract
	CustodyTransactor // Write-only binding to the contract
	CustodyFilterer   // Log filterer for contract events
}

// CustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodySession struct {
	Contract     *Custody          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodyCallerSession struct {
	Contract *CustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodyTransactorSession struct {
	Contract     *CustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodyRaw struct {
	Contract *Custody // Generic contract binding to access the raw methods on
}

// CustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodyCallerRaw struct {
	Contract *CustodyCaller // Generic read-only contract binding to access the raw methods on
}

// CustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodyTransactorRaw struct {
	Contract *CustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustody creates a new instance of Custody, bound to a specific deployed contract.
func NewCustody(address common.Address, backend bind.ContractBackend) (*Custody, error) {
	contract, err := bindCustody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Custody{CustodyCaller: CustodyCaller{contract: contract}, CustodyTransactor: CustodyTransactor{contract: contract}, CustodyFilterer: CustodyFilterer{contract: contract}}, nil
}

// NewCustodyCaller creates a new read-only instance of Custody, bound to a specific deployed contract.
func NewCustodyCaller(address common.Address, caller bind.ContractCaller) (*CustodyCaller, error) {
	contract, err := bindCustody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodyCaller{contract: contract}, nil
}

// NewCustodyTransactor creates a new write-only instance of Custody, bound to a specific deployed contract.
func NewCustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodyTransactor, error) {
	contract, err := bindCustody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodyTransactor{contract: contract}, nil
}

// NewCustodyFilterer creates a new log filterer instance of Custody, bound to a specific deployed contract.
func NewCustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodyFilterer, error) {
	contract, err := bindCustody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodyFilterer{contract: contract}, nil
}

// bindCustody binds a generic wrapper to an already deployed contract.
func bindCustody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Custody *CustodyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Custody.Contract.CustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Custody *CustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custody.Contract.CustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Custody *CustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Custody.Contract.CustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Custody *CustodyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Custody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Custody *CustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Custody *CustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Custody.Contract.contract.Transact(opts, method, params...)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Custody *CustodyCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Custody.contract.Call(opts, out, "authorized", arg0)
	return *ret0, err
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Custody *CustodySession) Authorized(arg0 common.Address) (bool, error) {
	return _Custody.Contract.Authorized(&_Custody.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Custody *CustodyCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _Custody.Contract.Authorized(&_Custody.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Custody *CustodyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Custody.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Custody *CustodySession) Owner() (common.Address, error) {
	return _Custody.Contract.Owner(&_Custody.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Custody *CustodyCallerSession) Owner() (common.Address, error) {
	return _Custody.Contract.Owner(&_Custody.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Custody *CustodyCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Custody.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Custody *CustodySession) Token() (common.Address, error) {
	return _Custody.Contract.Token(&_Custody.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Custody *CustodyCallerSession) Token() (common.Address, error) {
	return _Custody.Contract.Token(&_Custody.CallOpts)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address _account) returns()
func (_Custody *CustodyTransactor) Authorize(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "authorize", _account)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address _account) returns()
func (_Custody *CustodySession) Authorize(_account common.Address) (*types.Transaction, error) {
	return _Custody.Contract.Authorize(&_Custody.TransactOpts, _account)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address _account) returns()
func (_Custody *CustodyTransactorSession) Authorize(_account common.Address) (*types.Transaction, error) {
	return _Custody.Contract.Authorize(&_Custody.TransactOpts, _account)
}

// Forbid is a paid mutator transaction binding the contract method 0x7e95cd27.
//
// Solidity: function forbid(address _account) returns()
func (_Custody *CustodyTransactor) Forbid(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "forbid", _account)
}

// Forbid is a paid mutator transaction binding the contract method 0x7e95cd27.
//
// Solidity: function forbid(address _account) returns()
func (_Custody *CustodySession) Forbid(_account common.Address) (*types.Transaction, error) {
	return _Custody.Contract.Forbid(&_Custody.TransactOpts, _account)
}

// Forbid is a paid mutator transaction binding the contract method 0x7e95cd27.
//
// Solidity: function forbid(address _account) returns()
func (_Custody *CustodyTransactorSession) Forbid(_account common.Address) (*types.Transaction, error) {
	return _Custody.Contract.Forbid(&_Custody.TransactOpts, _account)
}

// Payout is a paid mutator transaction binding the contract method 0xc176e639.
//
// Solidity: function payout(address[] _recipients, uint256[] _amounts) returns()
func (_Custody *CustodyTransactor) Payout(opts *bind.TransactOpts, _recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "payout", _recipients, _amounts)
}

// Payout is a paid mutator transaction binding the contract method 0xc176e639.
//
// Solidity: function payout(address[] _recipients, uint256[] _amounts) returns()
func (_Custody *CustodySession) Payout(_recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Payout(&_Custody.TransactOpts, _recipients, _amounts)
}

// Payout is a paid mutator transaction binding the contract method 0xc176e639.
//
// Solidity: function payout(address[] _recipients, uint256[] _amounts) returns()
func (_Custody *CustodyTransactorSession) Payout(_recipients []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Payout(&_Custody.TransactOpts, _recipients, _amounts)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address _tokenAddress, uint256 amount) returns()
func (_Custody *CustodyTransactor) Recover(opts *bind.TransactOpts, _tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "recover", _tokenAddress, amount)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address _tokenAddress, uint256 amount) returns()
func (_Custody *CustodySession) Recover(_tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Recover(&_Custody.TransactOpts, _tokenAddress, amount)
}

// Recover is a paid mutator transaction binding the contract method 0x5705ae43.
//
// Solidity: function recover(address _tokenAddress, uint256 amount) returns()
func (_Custody *CustodyTransactorSession) Recover(_tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Recover(&_Custody.TransactOpts, _tokenAddress, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Custody *CustodyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Custody *CustodySession) RenounceOwnership() (*types.Transaction, error) {
	return _Custody.Contract.RenounceOwnership(&_Custody.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Custody *CustodyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Custody.Contract.RenounceOwnership(&_Custody.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Custody *CustodyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Custody *CustodySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Custody.Contract.TransferOwnership(&_Custody.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Custody *CustodyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Custody.Contract.TransferOwnership(&_Custody.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Custody *CustodyTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Custody.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Custody *CustodySession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Withdraw(&_Custody.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Custody *CustodyTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Custody.Contract.Withdraw(&_Custody.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Custody *CustodyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custody.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Custody *CustodySession) Receive() (*types.Transaction, error) {
	return _Custody.Contract.Receive(&_Custody.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Custody *CustodyTransactorSession) Receive() (*types.Transaction, error) {
	return _Custody.Contract.Receive(&_Custody.TransactOpts)
}

// CustodyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Custody contract.
type CustodyOwnershipTransferredIterator struct {
	Event *CustodyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CustodyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodyOwnershipTransferred)
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
		it.Event = new(CustodyOwnershipTransferred)
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
func (it *CustodyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodyOwnershipTransferred represents a OwnershipTransferred event raised by the Custody contract.
type CustodyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Custody *CustodyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustodyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Custody.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustodyOwnershipTransferredIterator{contract: _Custody.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Custody *CustodyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustodyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Custody.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodyOwnershipTransferred)
				if err := _Custody.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Custody *CustodyFilterer) ParseOwnershipTransferred(log types.Log) (*CustodyOwnershipTransferred, error) {
	event := new(CustodyOwnershipTransferred)
	if err := _Custody.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
