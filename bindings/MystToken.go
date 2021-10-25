/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MystTokenMetaData contains all meta data concerning the MystToken contract.
var MystTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"FundsRecoveryDestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"UpgradeMasterSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUpgradeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"upgradeFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newUpgradeMaster\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"internalType\":\"enumMystToken.UpgradeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001ca938038062001ca983398101604081905262000034916200019c565b6001600160a01b0381166080819052604080516318160ddd60e01b815290516318160ddd91600480820192602092909190829003018186803b1580156200007a57600080fd5b505afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ce565b60a0908152600080546001600160a01b0319163317905560408051808201825260098152684d797374657269756d60b81b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f5a6b6ed89c9a1a7ef403de5fac83d634f25f39cd4318f8dabc0fdddabbba7bd1818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6606082015246608082015230818501528251808203909401845260c001909152815191012060055550620001e8565b600060208284031215620001af57600080fd5b81516001600160a01b0381168114620001c757600080fd5b9392505050565b600060208284031215620001e157600080fd5b5051919050565b60805160a051611a8d6200021c600039600081816103570152610a8801526000818161024d01526109e70152611a8d6000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c8063600440cb1161010f578063a9059cbb116100a2578063dd62ed3e11610071578063dd62ed3e1461048d578063df8de3e7146104c6578063f58c5b6e146104d9578063ffeb7d75146104ea57600080fd5b8063a9059cbb1461044c578063c752ff621461045f578063d505accf14610467578063d7e7088a1461047a57600080fd5b80637ecebe00116100de5780637ecebe00146103e15780638444b3911461040157806395d89b4114610416578063a457c2d71461043957600080fd5b8063600440cb1461038c57806361d3d7a61461039d57806370a08231146103a5578063753e88e5146103ce57600080fd5b8063313ce5671161018757806342966c681161015657806342966c681461032f57806345977d03146103425780634b2ba0dd146103555780635de4ccb01461037b57600080fd5b8063313ce567146102e65780633644e51514610300578063395093511461030957806340c10f191461031c57600080fd5b806318160ddd116101c357806318160ddd14610285578063238e130a1461029757806323b872dd146102ac57806330adf81f146102bf57600080fd5b806306fdde03146101ea578063095ea7b3146102285780630e7c1cb51461024b575b600080fd5b610212604051806040016040528060098152602001684d797374657269756d60b81b81525081565b60405161021f91906117a7565b60405180910390f35b61023b610236366004611818565b6104fd565b604051901515815260200161021f565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200161021f565b6004545b60405190815260200161021f565b6102aa6102a5366004611842565b610513565b005b61023b6102ba366004611864565b61064a565b6102897f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b6102ee601281565b60405160ff909116815260200161021f565b61028960055481565b61023b610317366004611818565b610748565b6102aa61032a366004611818565b61077f565b6102aa61033d3660046118a0565b6107f0565b6102aa6103503660046118a0565b6107fd565b7f0000000000000000000000000000000000000000000000000000000000000000610289565b6001546001600160a01b031661026d565b6000546001600160a01b031661026d565b61023b600181565b6102896103b3366004611842565b6001600160a01b031660009081526003602052604090205490565b6102aa6103dc366004611818565b6109dc565b6102896103ef366004611842565b60066020526000908152604090205481565b610409610b18565b60405161021f91906118cf565b61021260405180604001604052806004815260200163135654d560e21b81525081565b61023b610447366004611818565b610b51565b61023b61045a366004611818565b610b88565b600254610289565b6102aa6104753660046118f7565b610b95565b6102aa610488366004611842565b610dab565b61028961049b36600461196a565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b6102aa6104d4366004611842565b611131565b6008546001600160a01b031661026d565b6102aa6104f8366004611842565b61124d565b600061050a338484611370565b50600192915050565b6000546001600160a01b0316336001600160a01b0316146105915760405162461bcd60e51b815260206004820152602d60248201527f4d5953543a206f6e6c792061206d61737465722063616e207365742066756e6460448201526c39903232b9ba34b730ba34b7b760991b60648201526084015b60405180910390fd5b6001600160a01b0381166105fe5760405162461bcd60e51b815260206004820152602e60248201527f4d5953543a2066756e64732064657374696e6174696f6e2063616e277420626560448201526d207a65726f20616464726565737360901b6064820152608401610588565b600880546001600160a01b0319166001600160a01b03831690811790915560405181907f2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c80490600090a350565b60006001600160a01b0384166106ae5760405162461bcd60e51b8152602060048201526024808201527f4d5953543a207472616e736665722066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610588565b336001600160a01b03851681148015906106ef57506001600160a01b0380861660009081526007602090815260408083209385168352929052205460001914155b15610732576001600160a01b03808616600090815260076020908152604080832093851683529290522054610732908690839061072d9087906119b3565b611370565b61073d858585611494565b506001949350505050565b3360008181526007602090815260408083206001600160a01b0387168452909152812054909161050a91859061072d9086906119ca565b6000546001600160a01b0316336001600160a01b0316146107e25760405162461bcd60e51b815260206004820152601c60248201527f4d5953543a206f6e6c792061206d61737465722063616e206d696e74000000006044820152606401610588565b6107ec8282611554565b5050565b6107fa338261167b565b50565b6000610807610b18565b9050600381600581111561081d5761081d6118b9565b148061083a57506004816005811115610838576108386118b9565b145b6108945760405162461bcd60e51b815260206004820152602560248201527f4d5953543a20746f6b656e206973206e6f7420696e20757067726164696e6720604482015264737461746560d81b6064820152608401610588565b816108f75760405162461bcd60e51b815260206004820152602d60248201527f4d5953543a2075706772616461626c6520616d6f756e742073686f756c64206260448201526c065206d6f7265207468616e203609c1b6064820152608401610588565b33610902818461167b565b8260025461091091906119ca565b60025560015460405163753e88e560e01b81526001600160a01b038381166004830152602482018690529091169063753e88e590604401600060405180830381600087803b15801561096157600080fd5b505af1158015610975573d6000803e3d6000fd5b50505050806001600160a01b03167f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac6109b66001546001600160a01b031690565b604080516001600160a01b039092168252602082018790520160405180910390a2505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a655760405162461bcd60e51b815260206004820152602860248201527f6f6e6c79206f726967696e616c20746f6b656e2063616e2063616c6c207570676044820152677261646546726f6d60c01b6064820152608401610588565b610a7d82610a786402540be400846119e2565b611554565b610aac6402540be4007f00000000000000000000000000000000000000000000000000000000000000006119e2565b60045411156107ec5760405162461bcd60e51b815260206004820152603260248201527f63616e206e6f74206d696e74206d6f726520746f6b656e73207468616e20696e604482015271081bdc9a59da5b985b0818dbdb9d1c9858dd60721b6064820152608401610588565b6001546000906001600160a01b0316610b315750600290565b600254610b3e5750600390565b600454610b4b5750600590565b50600490565b3360008181526007602090815260408083206001600160a01b0387168452909152812054909161050a91859061072d9086906119b3565b600061050a338484611494565b42841015610bdc5760405162461bcd60e51b8152602060048201526014602482015273135654d50e8814195c9b5a5d08195e1c1a5c995960621b6044820152606401610588565b6005546001600160a01b038816600090815260066020526040812080549192917f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918b918b918b919087610c2f83611a01565b909155506040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810187905260e00160405160208183030381529060405280519060200120604051602001610ca892919061190160f01b81526002810192909252602282015260420190565b60408051601f198184030181528282528051602091820120600080855291840180845281905260ff88169284019290925260608301869052608083018590529092509060019060a0016020604051602081039080840390855afa158015610d13573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811615801590610d495750886001600160a01b0316816001600160a01b0316145b610d955760405162461bcd60e51b815260206004820152601760248201527f4d5953543a20696e76616c6964207369676e61747572650000000000000000006044820152606401610588565b610da0898989611370565b505050505050505050565b6000546001600160a01b0316336001600160a01b031614610e275760405162461bcd60e51b815260206004820152603060248201527f4d5953543a206f6e6c792061206d61737465722063616e2064657369676e617460448201526f19481d1a19481b995e1d081859d95b9d60821b6064820152608401610588565b6001600160a01b038116610e8f5760405162461bcd60e51b815260206004820152602960248201527f4d5953543a2075706772616465206167656e742063616e2774206265207a65726044820152686f206164647265737360b81b6064820152608401610588565b6004610e99610b18565b6005811115610eaa57610eaa6118b9565b1415610ef85760405162461bcd60e51b815260206004820152601f60248201527f4d5953543a20757067726164652068617320616c726561647920626567756e006044820152606401610588565b600180546001600160a01b0319166001600160a01b038316908117909155604080516330e9ebd360e11b815290516361d3d7a691600480820192602092909190829003018186803b158015610f4c57600080fd5b505afa158015610f60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f849190611a1c565b610fed5760405162461bcd60e51b815260206004820152603460248201527f4d5953543a206167656e742073686f756c6420696d706c656d656e742049557060448201527367726164654167656e7420696e7465726661636560601b6064820152608401610588565b600454600160009054906101000a90046001600160a01b03166001600160a01b0316634b2ba0dd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561103e57600080fd5b505afa158015611052573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110769190611a3e565b146110e05760405162461bcd60e51b815260206004820152603460248201527f4d5953543a2075706772616465206167656e742073686f756c64206b6e6f7720604482015273746f6b656e277320746f74616c20737570706c7960601b6064820152608401610588565b7f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc6111136001546001600160a01b031690565b6040516001600160a01b03909116815260200160405180910390a150565b6008546001600160a01b031661114657600080fd5b6040516370a0823160e01b81523060048201526000906001600160a01b038316906370a082319060240160206040518083038186803b15801561118857600080fd5b505afa15801561119c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c09190611a3e565b60085460405163a9059cbb60e01b81526001600160a01b0391821660048201526024810183905291925083169063a9059cbb90604401602060405180830381600087803b15801561121057600080fd5b505af1158015611224573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112489190611a1c565b505050565b6001600160a01b0381166112b65760405162461bcd60e51b815260206004820152602a60248201527f4d5953543a2075706772616465206d61737465722063616e2774206265207a65604482015269726f206164647265737360b01b6064820152608401610588565b6000546001600160a01b0316336001600160a01b03161461132b5760405162461bcd60e51b815260206004820152602960248201527f4d5953543a206f6e6c792075706772616465206d61737465722063616e20736560448201526874206e6577206f6e6560b81b6064820152608401610588565b600080546001600160a01b0319166001600160a01b0383169081179091557f0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d5090611113565b6001600160a01b0383166113d25760405162461bcd60e51b815260206004820152602360248201527f4d5953543a20617070726f76652066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610588565b6001600160a01b0382166114325760405162461bcd60e51b815260206004820152602160248201527f4d5953543a20617070726f766520746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610588565b6001600160a01b0383811660008181526007602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0382166114ac57611248838261167b565b6001600160a01b0383166000908152600360205260409020546114d09082906119b3565b6001600160a01b0380851660009081526003602052604080822093909355908416815220546115009082906119ca565b6001600160a01b0380841660008181526003602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906114879085815260200190565b6001600160a01b0382166115aa5760405162461bcd60e51b815260206004820152601e60248201527f4d5953543a206d696e7420746f20746865207a65726f206164647265737300006044820152606401610588565b806004546115b891906119ca565b6004556001600160a01b0382166000908152600360205260409020546115df9082906119ca565b6001600160a01b038316600081815260036020526040908190209290925590517f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe9061162e9084815260200190565b60405180910390a26040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6001600160a01b0382166116d15760405162461bcd60e51b815260206004820181905260248201527f4d5953543a206275726e2066726f6d20746865207a65726f20616464726573736044820152606401610588565b6001600160a01b0382166000908152600360205260409020546116f59082906119b3565b6001600160a01b03831660009081526003602052604090205560045461171c9082906119b3565b6004556040518181526000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3816001600160a01b03167f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df78260405161179b91815260200190565b60405180910390a25050565b600060208083528351808285015260005b818110156117d4578581018301518582016040015282016117b8565b818111156117e6576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461181357600080fd5b919050565b6000806040838503121561182b57600080fd5b611834836117fc565b946020939093013593505050565b60006020828403121561185457600080fd5b61185d826117fc565b9392505050565b60008060006060848603121561187957600080fd5b611882846117fc565b9250611890602085016117fc565b9150604084013590509250925092565b6000602082840312156118b257600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60208101600683106118f157634e487b7160e01b600052602160045260246000fd5b91905290565b600080600080600080600060e0888a03121561191257600080fd5b61191b886117fc565b9650611929602089016117fc565b95506040880135945060608801359350608088013560ff8116811461194d57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561197d57600080fd5b611986836117fc565b9150611994602084016117fc565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b6000828210156119c5576119c561199d565b500390565b600082198211156119dd576119dd61199d565b500190565b60008160001904831182151516156119fc576119fc61199d565b500290565b6000600019821415611a1557611a1561199d565b5060010190565b600060208284031215611a2e57600080fd5b8151801515811461185d57600080fd5b600060208284031215611a5057600080fd5b505191905056fea2646970667358221220648ff44524a84483a155252e91d43ce845b165c82019b2235db9aa2aa7efd5db64736f6c63430008090033",
}

// MystTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MystTokenMetaData.ABI instead.
var MystTokenABI = MystTokenMetaData.ABI

// MystTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MystTokenMetaData.Bin instead.
var MystTokenBin = MystTokenMetaData.Bin

// DeployMystToken deploys a new Ethereum contract, binding an instance of MystToken to it.
func DeployMystToken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *MystToken, error) {
	parsed, err := MystTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MystTokenBin), backend, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MystToken{MystTokenCaller: MystTokenCaller{contract: contract}, MystTokenTransactor: MystTokenTransactor{contract: contract}, MystTokenFilterer: MystTokenFilterer{contract: contract}}, nil
}

// MystToken is an auto generated Go binding around an Ethereum contract.
type MystToken struct {
	MystTokenCaller     // Read-only binding to the contract
	MystTokenTransactor // Write-only binding to the contract
	MystTokenFilterer   // Log filterer for contract events
}

// MystTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MystTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MystTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MystTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MystTokenSession struct {
	Contract     *MystToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MystTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MystTokenCallerSession struct {
	Contract *MystTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MystTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MystTokenTransactorSession struct {
	Contract     *MystTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MystTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MystTokenRaw struct {
	Contract *MystToken // Generic contract binding to access the raw methods on
}

// MystTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MystTokenCallerRaw struct {
	Contract *MystTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MystTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MystTokenTransactorRaw struct {
	Contract *MystTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMystToken creates a new instance of MystToken, bound to a specific deployed contract.
func NewMystToken(address common.Address, backend bind.ContractBackend) (*MystToken, error) {
	contract, err := bindMystToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MystToken{MystTokenCaller: MystTokenCaller{contract: contract}, MystTokenTransactor: MystTokenTransactor{contract: contract}, MystTokenFilterer: MystTokenFilterer{contract: contract}}, nil
}

// NewMystTokenCaller creates a new read-only instance of MystToken, bound to a specific deployed contract.
func NewMystTokenCaller(address common.Address, caller bind.ContractCaller) (*MystTokenCaller, error) {
	contract, err := bindMystToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MystTokenCaller{contract: contract}, nil
}

// NewMystTokenTransactor creates a new write-only instance of MystToken, bound to a specific deployed contract.
func NewMystTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MystTokenTransactor, error) {
	contract, err := bindMystToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MystTokenTransactor{contract: contract}, nil
}

// NewMystTokenFilterer creates a new log filterer instance of MystToken, bound to a specific deployed contract.
func NewMystTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MystTokenFilterer, error) {
	contract, err := bindMystToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MystTokenFilterer{contract: contract}, nil
}

// bindMystToken binds a generic wrapper to an already deployed contract.
func bindMystToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MystTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystToken *MystTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MystToken.Contract.MystTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystToken *MystTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystToken.Contract.MystTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystToken *MystTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystToken.Contract.MystTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MystToken *MystTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MystToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MystToken *MystTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MystToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MystToken *MystTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MystToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MystToken.Contract.DOMAINSEPARATOR(&_MystToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MystToken *MystTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MystToken.Contract.DOMAINSEPARATOR(&_MystToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MystToken.Contract.PERMITTYPEHASH(&_MystToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MystToken *MystTokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MystToken.Contract.PERMITTYPEHASH(&_MystToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_MystToken *MystTokenCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MystToken.Contract.Allowance(&_MystToken.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenCaller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_MystToken *MystTokenCallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _MystToken.Contract.BalanceOf(&_MystToken.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MystToken *MystTokenCallerSession) Decimals() (uint8, error) {
	return _MystToken.Contract.Decimals(&_MystToken.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "getFundsDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenSession) GetFundsDestination() (common.Address, error) {
	return _MystToken.Contract.GetFundsDestination(&_MystToken.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_MystToken *MystTokenCallerSession) GetFundsDestination() (common.Address, error) {
	return _MystToken.Contract.GetFundsDestination(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "getUpgradeState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_MystToken *MystTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _MystToken.Contract.GetUpgradeState(&_MystToken.CallOpts)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenCaller) IsUpgradeAgent(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "isUpgradeAgent")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenSession) IsUpgradeAgent() (bool, error) {
	return _MystToken.Contract.IsUpgradeAgent(&_MystToken.CallOpts)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_MystToken *MystTokenCallerSession) IsUpgradeAgent() (bool, error) {
	return _MystToken.Contract.IsUpgradeAgent(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MystToken *MystTokenCallerSession) Name() (string, error) {
	return _MystToken.Contract.Name(&_MystToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MystToken.Contract.Nonces(&_MystToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MystToken *MystTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MystToken.Contract.Nonces(&_MystToken.CallOpts, arg0)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenCaller) OriginalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "originalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenSession) OriginalSupply() (*big.Int, error) {
	return _MystToken.Contract.OriginalSupply(&_MystToken.CallOpts)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_MystToken *MystTokenCallerSession) OriginalSupply() (*big.Int, error) {
	return _MystToken.Contract.OriginalSupply(&_MystToken.CallOpts)
}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenCaller) OriginalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "originalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenSession) OriginalToken() (common.Address, error) {
	return _MystToken.Contract.OriginalToken(&_MystToken.CallOpts)
}

// OriginalToken is a free data retrieval call binding the contract method 0x0e7c1cb5.
//
// Solidity: function originalToken() view returns(address)
func (_MystToken *MystTokenCallerSession) OriginalToken() (common.Address, error) {
	return _MystToken.Contract.OriginalToken(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MystToken *MystTokenCallerSession) Symbol() (string, error) {
	return _MystToken.Contract.Symbol(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MystToken.Contract.TotalSupply(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "totalUpgraded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_MystToken *MystTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _MystToken.Contract.TotalUpgraded(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "upgradeAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _MystToken.Contract.UpgradeAgent(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MystToken.contract.Call(opts, &out, "upgradeMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_MystToken *MystTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _MystToken.Contract.UpgradeMaster(&_MystToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MystToken *MystTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Approve(&_MystToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Burn(&_MystToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MystToken *MystTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Burn(&_MystToken.TransactOpts, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenTransactor) ClaimTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "claimTokens", token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenSession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.ClaimTokens(&_MystToken.TransactOpts, token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_MystToken *MystTokenTransactorSession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.ClaimTokens(&_MystToken.TransactOpts, token)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.DecreaseAllowance(&_MystToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MystToken *MystTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.DecreaseAllowance(&_MystToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.IncreaseAllowance(&_MystToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MystToken *MystTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.IncreaseAllowance(&_MystToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MystToken *MystTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Mint(&_MystToken.TransactOpts, _account, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "permit", holder, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenSession) Permit(holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.Contract.Permit(&_MystToken.TransactOpts, holder, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address holder, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MystToken *MystTokenTransactorSession) Permit(holder common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MystToken.Contract.Permit(&_MystToken.TransactOpts, holder, spender, value, deadline, v, r, s)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenTransactor) SetFundsDestination(opts *bind.TransactOpts, newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setFundsDestination", newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenSession) SetFundsDestination(newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetFundsDestination(&_MystToken.TransactOpts, newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address newDestination) returns()
func (_MystToken *MystTokenTransactorSession) SetFundsDestination(newDestination common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetFundsDestination(&_MystToken.TransactOpts, newDestination)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenTransactor) SetUpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setUpgradeAgent", agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeAgent(&_MystToken.TransactOpts, agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_MystToken *MystTokenTransactorSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeAgent(&_MystToken.TransactOpts, agent)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "setUpgradeMaster", newUpgradeMaster)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenSession) SetUpgradeMaster(newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, newUpgradeMaster)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address newUpgradeMaster) returns()
func (_MystToken *MystTokenTransactorSession) SetUpgradeMaster(newUpgradeMaster common.Address) (*types.Transaction, error) {
	return _MystToken.Contract.SetUpgradeMaster(&_MystToken.TransactOpts, newUpgradeMaster)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Transfer(&_MystToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_MystToken *MystTokenTransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.TransferFrom(&_MystToken.TransactOpts, holder, recipient, amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenTransactor) Upgrade(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "upgrade", amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenSession) Upgrade(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 amount) returns()
func (_MystToken *MystTokenTransactorSession) Upgrade(amount *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.Upgrade(&_MystToken.TransactOpts, amount)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenTransactor) UpgradeFrom(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.contract.Transact(opts, "upgradeFrom", _account, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenSession) UpgradeFrom(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.UpgradeFrom(&_MystToken.TransactOpts, _account, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _account, uint256 _value) returns()
func (_MystToken *MystTokenTransactorSession) UpgradeFrom(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MystToken.Contract.UpgradeFrom(&_MystToken.TransactOpts, _account, _value)
}

// MystTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MystToken contract.
type MystTokenApprovalIterator struct {
	Event *MystTokenApproval // Event containing the contract specifics and raw log

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
func (it *MystTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenApproval)
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
		it.Event = new(MystTokenApproval)
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
func (it *MystTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenApproval represents a Approval event raised by the MystToken contract.
type MystTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MystToken *MystTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MystTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenApprovalIterator{contract: _MystToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MystToken *MystTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MystTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenApproval)
				if err := _MystToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MystToken *MystTokenFilterer) ParseApproval(log types.Log) (*MystTokenApproval, error) {
	event := new(MystTokenApproval)
	if err := _MystToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the MystToken contract.
type MystTokenBurnedIterator struct {
	Event *MystTokenBurned // Event containing the contract specifics and raw log

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
func (it *MystTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenBurned)
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
		it.Event = new(MystTokenBurned)
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
func (it *MystTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenBurned represents a Burned event raised by the MystToken contract.
type MystTokenBurned struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*MystTokenBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenBurnedIterator{contract: _MystToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *MystTokenBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenBurned)
				if err := _MystToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_MystToken *MystTokenFilterer) ParseBurned(log types.Log) (*MystTokenBurned, error) {
	event := new(MystTokenBurned)
	if err := _MystToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenFundsRecoveryDestinationChangedIterator is returned from FilterFundsRecoveryDestinationChanged and is used to iterate over the raw logs and unpacked data for FundsRecoveryDestinationChanged events raised by the MystToken contract.
type MystTokenFundsRecoveryDestinationChangedIterator struct {
	Event *MystTokenFundsRecoveryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenFundsRecoveryDestinationChanged)
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
		it.Event = new(MystTokenFundsRecoveryDestinationChanged)
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
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenFundsRecoveryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenFundsRecoveryDestinationChanged represents a FundsRecoveryDestinationChanged event raised by the MystToken contract.
type MystTokenFundsRecoveryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFundsRecoveryDestinationChanged is a free log retrieval operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) FilterFundsRecoveryDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*MystTokenFundsRecoveryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "FundsRecoveryDestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenFundsRecoveryDestinationChangedIterator{contract: _MystToken.contract, event: "FundsRecoveryDestinationChanged", logs: logs, sub: sub}, nil
}

// WatchFundsRecoveryDestinationChanged is a free log subscription operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) WatchFundsRecoveryDestinationChanged(opts *bind.WatchOpts, sink chan<- *MystTokenFundsRecoveryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "FundsRecoveryDestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenFundsRecoveryDestinationChanged)
				if err := _MystToken.contract.UnpackLog(event, "FundsRecoveryDestinationChanged", log); err != nil {
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

// ParseFundsRecoveryDestinationChanged is a log parse operation binding the contract event 0x2e1db88922daae16be4e3c1a1f4bfab0cf6741938844967bd985ac8b2a12c804.
//
// Solidity: event FundsRecoveryDestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_MystToken *MystTokenFilterer) ParseFundsRecoveryDestinationChanged(log types.Log) (*MystTokenFundsRecoveryDestinationChanged, error) {
	event := new(MystTokenFundsRecoveryDestinationChanged)
	if err := _MystToken.contract.UnpackLog(event, "FundsRecoveryDestinationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MystToken contract.
type MystTokenMintedIterator struct {
	Event *MystTokenMinted // Event containing the contract specifics and raw log

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
func (it *MystTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenMinted)
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
		it.Event = new(MystTokenMinted)
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
func (it *MystTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenMinted represents a Minted event raised by the MystToken contract.
type MystTokenMinted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*MystTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenMintedIterator{contract: _MystToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MystTokenMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenMinted)
				if err := _MystToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_MystToken *MystTokenFilterer) ParseMinted(log types.Log) (*MystTokenMinted, error) {
	event := new(MystTokenMinted)
	if err := _MystToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MystToken contract.
type MystTokenTransferIterator struct {
	Event *MystTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MystTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenTransfer)
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
		it.Event = new(MystTokenTransfer)
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
func (it *MystTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenTransfer represents a Transfer event raised by the MystToken contract.
type MystTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MystToken *MystTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MystTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenTransferIterator{contract: _MystToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MystToken *MystTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MystTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenTransfer)
				if err := _MystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MystToken *MystTokenFilterer) ParseTransfer(log types.Log) (*MystTokenTransfer, error) {
	event := new(MystTokenTransfer)
	if err := _MystToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the MystToken contract.
type MystTokenUpgradeIterator struct {
	Event *MystTokenUpgrade // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgrade)
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
		it.Event = new(MystTokenUpgrade)
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
func (it *MystTokenUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgrade represents a Upgrade event raised by the MystToken contract.
type MystTokenUpgrade struct {
	From  common.Address
	Agent common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, from []common.Address) (*MystTokenUpgradeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "Upgrade", fromRule)
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeIterator{contract: _MystToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *MystTokenUpgrade, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "Upgrade", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgrade)
				if err := _MystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
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

// ParseUpgrade is a log parse operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed from, address agent, uint256 _value)
func (_MystToken *MystTokenFilterer) ParseUpgrade(log types.Log) (*MystTokenUpgrade, error) {
	event := new(MystTokenUpgrade)
	if err := _MystToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenUpgradeAgentSetIterator is returned from FilterUpgradeAgentSet and is used to iterate over the raw logs and unpacked data for UpgradeAgentSet events raised by the MystToken contract.
type MystTokenUpgradeAgentSetIterator struct {
	Event *MystTokenUpgradeAgentSet // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeAgentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgradeAgentSet)
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
		it.Event = new(MystTokenUpgradeAgentSet)
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
func (it *MystTokenUpgradeAgentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeAgentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgradeAgentSet represents a UpgradeAgentSet event raised by the MystToken contract.
type MystTokenUpgradeAgentSet struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAgentSet is a free log retrieval operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) FilterUpgradeAgentSet(opts *bind.FilterOpts) (*MystTokenUpgradeAgentSetIterator, error) {

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeAgentSetIterator{contract: _MystToken.contract, event: "UpgradeAgentSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeAgentSet is a free log subscription operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) WatchUpgradeAgentSet(opts *bind.WatchOpts, sink chan<- *MystTokenUpgradeAgentSet) (event.Subscription, error) {

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgradeAgentSet)
				if err := _MystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
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

// ParseUpgradeAgentSet is a log parse operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_MystToken *MystTokenFilterer) ParseUpgradeAgentSet(log types.Log) (*MystTokenUpgradeAgentSet, error) {
	event := new(MystTokenUpgradeAgentSet)
	if err := _MystToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MystTokenUpgradeMasterSetIterator is returned from FilterUpgradeMasterSet and is used to iterate over the raw logs and unpacked data for UpgradeMasterSet events raised by the MystToken contract.
type MystTokenUpgradeMasterSetIterator struct {
	Event *MystTokenUpgradeMasterSet // Event containing the contract specifics and raw log

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
func (it *MystTokenUpgradeMasterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MystTokenUpgradeMasterSet)
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
		it.Event = new(MystTokenUpgradeMasterSet)
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
func (it *MystTokenUpgradeMasterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MystTokenUpgradeMasterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MystTokenUpgradeMasterSet represents a UpgradeMasterSet event raised by the MystToken contract.
type MystTokenUpgradeMasterSet struct {
	Master common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpgradeMasterSet is a free log retrieval operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) FilterUpgradeMasterSet(opts *bind.FilterOpts) (*MystTokenUpgradeMasterSetIterator, error) {

	logs, sub, err := _MystToken.contract.FilterLogs(opts, "UpgradeMasterSet")
	if err != nil {
		return nil, err
	}
	return &MystTokenUpgradeMasterSetIterator{contract: _MystToken.contract, event: "UpgradeMasterSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeMasterSet is a free log subscription operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) WatchUpgradeMasterSet(opts *bind.WatchOpts, sink chan<- *MystTokenUpgradeMasterSet) (event.Subscription, error) {

	logs, sub, err := _MystToken.contract.WatchLogs(opts, "UpgradeMasterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MystTokenUpgradeMasterSet)
				if err := _MystToken.contract.UnpackLog(event, "UpgradeMasterSet", log); err != nil {
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

// ParseUpgradeMasterSet is a log parse operation binding the contract event 0x0bae748e6d38d2b1532af619519837d91d74845ad693f6f229677b4ac20b2d50.
//
// Solidity: event UpgradeMasterSet(address master)
func (_MystToken *MystTokenFilterer) ParseUpgradeMasterSet(log types.Log) (*MystTokenUpgradeMasterSet, error) {
	event := new(MystTokenUpgradeMasterSet)
	if err := _MystToken.contract.UnpackLog(event, "UpgradeMasterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
