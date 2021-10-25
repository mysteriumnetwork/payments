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

// ChannelImplementationMetaData contains all meta data concerning the ChannelImplementation contract.
var ChannelImplementationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"ExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSettled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lock\",\"type\":\"bytes32\"}],\"name\":\"PromiseSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hermes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"settled\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_identity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settlePromise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"requestExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_operatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_hermesSignature\",\"type\":\"bytes\"}],\"name\":\"fastExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"setFundsDestinationByCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060018055612402806100246000396000f3fe6080604052600436106101025760003560e01c80638da5cb5b11610095578063f2fde38b11610064578063f2fde38b1461044f578063f4b3a1971461046f578063f58c5b6e146104b0578063f7013ef6146104ce578063fc0c546a146104ee57600080fd5b80638da5cb5b146103a0578063d8092c92146103be578063df8de3e71461040f578063e9e8ad8b1461042f57600080fd5b8063570ca735116100d1578063570ca735146103135780636931b5501461034b5780636a2b76ad146103605780636f1746301461038057600080fd5b806307e8ec1f1461028d578063182f3488146102a4578063238e130a146102c4578063392e53cd146102e457600080fd5b36610288576040805160028082526060820183526000926020830190803683375050600b54604080516315ab88c960e31b815290519394506001600160a01b039091169263ad5c464892506004808301926020929190829003018186803b15801561016c57600080fd5b505afa158015610180573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a49190611e1b565b816000815181106101b7576101b7611e3f565b6001600160a01b0392831660209182029290920101526003548251911690829060019081106101e8576101e8611e3f565b6001600160a01b039283166020918202929092010152600b54604051637ff36ab560e01b8152911690637ff36ab590349061022e90600090869030904290600401611e55565b6000604051808303818588803b15801561024757600080fd5b505af115801561025b573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526102849190810190611ef0565b5050005b600080fd5b34801561029957600080fd5b506102a261050e565b005b3480156102b057600080fd5b506102a26102bf366004612006565b61071d565b3480156102d057600080fd5b506102a26102df36600461205f565b610a34565b3480156102f057600080fd5b50600a546040516001600160a01b03909116151581526020015b60405180910390f35b34801561031f57600080fd5b50600a54610333906001600160a01b031681565b6040516001600160a01b03909116815260200161030a565b34801561035757600080fd5b506102a2610b11565b34801561036c57600080fd5b506102a261037b36600461207c565b610bbe565b34801561038c57600080fd5b506102a261039b3660046120cc565b610d0d565b3480156103ac57600080fd5b506000546001600160a01b0316610333565b3480156103ca57600080fd5b506007546008546009546103e9926001600160a01b0390811692169083565b604080516001600160a01b0394851681529390921660208401529082015260600161030a565b34801561041b57600080fd5b506102a261042a36600461205f565b61109c565b34801561043b57600080fd5b506102a261044a366004612126565b611280565b34801561045b57600080fd5b506102a261046a36600461205f565b611632565b34801561047b57600080fd5b5060055460065461049391906001600160a01b031682565b604080519283526001600160a01b0390911660208301520161030a565b3480156104bc57600080fd5b506002546001600160a01b0316610333565b3480156104da57600080fd5b506102a26104e93660046121b8565b611760565b3480156104fa57600080fd5b50600354610333906001600160a01b031681565b6005541580159061052157506005544210155b6105a35760405162461bcd60e51b815260206004820152604260248201527f4368616e6e656c3a2065786974206861766520746f206265207265717565737460448201527f656420616e642074696d656c6f636b206861766520746f20626520696e2070616064820152611cdd60f21b608482015260a4015b60405180910390fd5b6003546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b1580156105e757600080fd5b505afa1580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f919061221c565b60035460065460405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b15801561067157600080fd5b505af1158015610685573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a99190612235565b50600654604080516001600160a01b039092168252602082018390527f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364910160405180910390a1506040805180820190915260008082526020909101819052600555600680546001600160a01b0319169055565b6000610727611a6e565b600554909150156107ab5760405162461bcd60e51b815260206004820152604260248201527f4368616e6e656c3a206e657720657869742063616e206265207265717565737460448201527f6564206f6e6c79207768656e206f6c64206f6e65207761732066696e616c6973606482015261195960f21b608482015260a40161059a565b4283101561083a5760405162461bcd60e51b815260206004820152605060248201527f4368616e6e656c3a2076616c696420756e74696c206861766520746f2062652060448201527f67726561746572207468616e206f7220657175616c20746f2063757272656e7460648201526f020626c6f636b2074696d657374616d760841b608482015260a40161059a565b8281116108af5760405162461bcd60e51b815260206004820152603c60248201527f4368616e6e656c3a2072657175657374206861766520746f2062652076616c6960448201527f642073686f72746572207468616e2044454c41595f5345434f4e445300000000606482015260840161059a565b6001600160a01b0384166109185760405162461bcd60e51b815260206004820152602a60248201527f4368616e6e656c3a2062656e65666963696172792063616e2774206265207a65604482015269726f206164647265737360b01b606482015260840161059a565b600a546001600160a01b031633146109c85760003090506000610995846040518060400160405280600d81526020016c22bc34ba103932b8bab2b9ba1d60991b8152508489896040516020016109719493929190612292565b60405160208183030381529060405280519060200120611a8290919063ffffffff16565b600a549091506001600160a01b038083169116146109c55760405162461bcd60e51b815260040161059a906122d0565b50505b6040805180820182528281526001600160a01b03861660209182018190526005849055600680546001600160a01b031916909117905590518281527fe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db910160405180910390a150505050565b6000546001600160a01b0316331480610a5657506000546001600160a01b0316155b610aa25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161059a565b6001600160a01b038116610ab557600080fd5b6002546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600280546001600160a01b0319166001600160a01b0392909216919091179055565b60026001541415610b645760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161059a565b60026001819055546001600160a01b0316610b7e57600080fd5b6002546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610bb7573d6000803e3d6000fd5b5060018055565b6001600160a01b038216610bd157600080fd5b60003090506000610c37836040518060400160405280601681526020017529b2ba10333ab73239903232b9ba34b730ba34b7b71d60511b815250848760046000815480929190610c209061232c565b919050556040516020016109719493929190612292565b600a549091506001600160a01b03808316911614610cad5760405162461bcd60e51b815260206004820152602d60248201527f4368616e6e656c3a206861766520746f206265207369676e656420627920707260448201526c6f706572206964656e7469747960981b606482015260840161059a565b6002546040516001600160a01b038087169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a35050600280546001600160a01b0319166001600160a01b03939093169290921790915550565b600082604051602001610d2291815260200190565b60408051808303601f1901815291905280516020909101209050306000610d7b84466040805160208101929092526001600160a01b03861690820152606081018a90526080810189905260a0810186905260c001610971565b600a549091506001600160a01b03808316911614610de95760405162461bcd60e51b815260206004820152602560248201527f6861766520746f206265207369676e6564206279206368616e6e656c206f70656044820152643930ba37b960d91b606482015260840161059a565b600954600090610df99089612347565b905060008111610e715760405162461bcd60e51b815260206004820152603760248201527f616d6f756e7420746f20736574746c652073686f756c6420626520677265617460448201527f6572207468617420616c726561647920736574746c6564000000000000000000606482015260840161059a565b6003546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a082319060240160206040518083038186803b158015610eb757600080fd5b505afa158015610ecb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eef919061221c565b905080821115610efd578091505b600954610f0b90839061235e565b6009556003546008546001600160a01b039182169163a9059cbb9116610f318b86612347565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401602060405180830381600087803b158015610f7757600080fd5b505af1158015610f8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faf9190612235565b50871561103c5760035460405163a9059cbb60e01b8152336004820152602481018a90526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b15801561100257600080fd5b505af1158015611016573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103a9190612235565b505b600854600954604080516001600160a01b03909316835260208301859052820152606081018890527f7f7697dfbfc8e203dd21764a16312ebed4de12b161eca0b41922de085b2eccd19060800160405180910390a1505050505050505050565b600260015414156110ef5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161059a565b60026001819055546001600160a01b031661110957600080fd5b6003546001600160a01b03828116911614156111755760405162461bcd60e51b815260206004820152602560248201527f6e617469766520746f6b656e2066756e64732063616e2774206265207265636f6044820152641d995c995960da1b606482015260840161059a565b6040516370a0823160e01b81523060048201526000906001600160a01b038316906370a082319060240160206040518083038186803b1580156111b757600080fd5b505afa1580156111cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ef919061221c565b60025460405163a9059cbb60e01b81526001600160a01b0391821660048201526024810183905291925083169063a9059cbb90604401602060405180830381600087803b15801561123f57600080fd5b505af1158015611253573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112779190612235565b50506001805550565b4283101561130f5760405162461bcd60e51b815260206004820152605060248201527f4368616e6e656c3a205f76616c6964556e74696c206861766520746f2062652060448201527f67726561746572207468616e206f7220657175616c20746f2063757272656e7460648201526f020626c6f636b2074696d657374616d760841b608482015260a40161059a565b60408051808201909152600d81526c22bc34ba103932b8bab2b9ba1d60991b6020820152309060009046836001600160a01b03168a8a8a6001600160a01b03168a600460008154809291906113639061232c565b9190505560405160200161137e989796959493929190612376565b60408051601f198184030181529190528051602090910120905060006113a48286611a82565b600a549091506001600160a01b038083169116146113d45760405162461bcd60e51b815260040161059a906122d0565b60006113e08386611a82565b6007549091506001600160a01b0380831691161461144c5760405162461bcd60e51b8152602060048201526024808201527f4368616e6e656c3a206861766520746f206265207369676e6564206279206865604482015263726d657360e01b606482015260840161059a565b881561154e57888a10156114c85760405162461bcd60e51b815260206004820152603e60248201527f4368616e6e656c3a207472616e736163746f72206665652063616e277420626560448201527f206269676765722074686174207769746864726177616c20616d6f756e740000606482015260840161059a565b60035460405163a9059cbb60e01b8152336004820152602481018b90526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b15801561151457600080fd5b505af1158015611528573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154c9190612235565b505b600061155a8a8c612347565b60035460405163a9059cbb60e01b81526001600160a01b038c811660048301526024820184905292935091169063a9059cbb90604401602060405180830381600087803b1580156115aa57600080fd5b505af11580156115be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e29190612235565b50604080516001600160a01b038b168152602081018390527f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364910160405180910390a15050505050505050505050565b6000546001600160a01b031633148061165457506000546001600160a01b0316155b6116a05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161059a565b6001600160a01b0381166117055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161059a565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600a546001600160a01b0316156117b25760405162461bcd60e51b8152602060048201526016602482015275125cc8185b1c9958591e481a5b9a5d1a585b1a5e995960521b604482015260640161059a565b6001600160a01b0383166118015760405162461bcd60e51b81526020600482015260166024820152754964656e746974792063616e2774206265207a65726f60501b604482015260640161059a565b6001600160a01b0382166118505760405162461bcd60e51b81526020600482015260166024820152754865726d657349442063616e2774206265207a65726f60501b604482015260640161059a565b6001600160a01b0385166118b75760405162461bcd60e51b815260206004820152602860248201527f546f6b656e2063616e2774206265206465706c6f796420696e746f207a65726f604482015267206164647265737360c01b606482015260840161059a565b600380546001600160a01b038088166001600160a01b031992831617909255600b80549287169290911691909117905580156119735760035460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b15801561193957600080fd5b505af115801561194d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119719190612235565b505b600a80546001600160a01b0319166001600160a01b03851690811790915561199a90611632565b6040518060600160405280836001600160a01b031663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b1580156119de57600080fd5b505afa1580156119f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a169190611e1b565b6001600160a01b03908116825293841660208083019190915260006040928301528251600780549187166001600160a01b03199283161790559083015160088054919096169116179093559091015160095550505050565b6000611a7d620546004261235e565b905090565b6000806000611a918585611aa6565b91509150611a9e81611b16565b509392505050565b600080825160411415611add5760208301516040840151606085015160001a611ad187828585611cd4565b94509450505050611b0f565b825160401415611b075760208301516040840151611afc868383611dc1565b935093505050611b0f565b506000905060025b9250929050565b6000816004811115611b2a57611b2a6123b6565b1415611b335750565b6001816004811115611b4757611b476123b6565b1415611b955760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161059a565b6002816004811115611ba957611ba96123b6565b1415611bf75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161059a565b6003816004811115611c0b57611c0b6123b6565b1415611c645760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161059a565b6004816004811115611c7857611c786123b6565b1415611cd15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161059a565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611d0b5750600090506003611db8565b8460ff16601b14158015611d2357508460ff16601c14155b15611d345750600090506004611db8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611d88573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611db157600060019250925050611db8565b9150600090505b94509492505050565b6000806001600160ff1b03831660ff84901c601b01611de287828885611cd4565b935093505050935093915050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611cd157600080fd5b600060208284031215611e2d57600080fd5b8151611e3881611e06565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060808201868352602060808185015281875180845260a086019150828901935060005b81811015611e9f5784516001600160a01b031683529383019391830191600101611e7a565b50506001600160a01b039690961660408501525050506060015292915050565b604051601f8201601f1916810167ffffffffffffffff81118282101715611ee857611ee8611df0565b604052919050565b60006020808385031215611f0357600080fd5b825167ffffffffffffffff80821115611f1b57600080fd5b818501915085601f830112611f2f57600080fd5b815181811115611f4157611f41611df0565b8060051b9150611f52848301611ebf565b8181529183018401918481019088841115611f6c57600080fd5b938501935b83851015611f8a57845182529385019390850190611f71565b98975050505050505050565b600082601f830112611fa757600080fd5b813567ffffffffffffffff811115611fc157611fc1611df0565b611fd4601f8201601f1916602001611ebf565b818152846020838601011115611fe957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561201b57600080fd5b833561202681611e06565b925060208401359150604084013567ffffffffffffffff81111561204957600080fd5b61205586828701611f96565b9150509250925092565b60006020828403121561207157600080fd5b8135611e3881611e06565b6000806040838503121561208f57600080fd5b823561209a81611e06565b9150602083013567ffffffffffffffff8111156120b657600080fd5b6120c285828601611f96565b9150509250929050565b600080600080608085870312156120e257600080fd5b843593506020850135925060408501359150606085013567ffffffffffffffff81111561210e57600080fd5b61211a87828801611f96565b91505092959194509250565b60008060008060008060c0878903121561213f57600080fd5b8635955060208701359450604087013561215881611e06565b935060608701359250608087013567ffffffffffffffff8082111561217c57600080fd5b6121888a838b01611f96565b935060a089013591508082111561219e57600080fd5b506121ab89828a01611f96565b9150509295509295509295565b600080600080600060a086880312156121d057600080fd5b85356121db81611e06565b945060208601356121eb81611e06565b935060408601356121fb81611e06565b9250606086013561220b81611e06565b949793965091946080013592915050565b60006020828403121561222e57600080fd5b5051919050565b60006020828403121561224757600080fd5b81518015158114611e3857600080fd5b6000815160005b81811015612278576020818501810151868301520161225e565b81811115612287576000828601525b509290920192915050565b600061229e8287612257565b6bffffffffffffffffffffffff19606096871b811682529490951b909316601485015250602883015250604801919050565b60208082526026908201527f4368616e6e656c3a206861766520746f206265207369676e6564206279206f7060408201526532b930ba37b960d11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060001982141561234057612340612316565b5060010190565b60008282101561235957612359612316565b500390565b6000821982111561237157612371612316565b500190565b6000612382828b612257565b9889525050602087019590955260408601939093526060850191909152608084015260a083015260c082015260e001919050565b634e487b7160e01b600052602160045260246000fdfea264697066735822122068303ffeec681607c42bf5915842e8be95777e6a78a27616f63b326cd62820f464736f6c63430008090033",
}

// ChannelImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use ChannelImplementationMetaData.ABI instead.
var ChannelImplementationABI = ChannelImplementationMetaData.ABI

// ChannelImplementationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChannelImplementationMetaData.Bin instead.
var ChannelImplementationBin = ChannelImplementationMetaData.Bin

// DeployChannelImplementation deploys a new Ethereum contract, binding an instance of ChannelImplementation to it.
func DeployChannelImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChannelImplementation, error) {
	parsed, err := ChannelImplementationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChannelImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelImplementation{ChannelImplementationCaller: ChannelImplementationCaller{contract: contract}, ChannelImplementationTransactor: ChannelImplementationTransactor{contract: contract}, ChannelImplementationFilterer: ChannelImplementationFilterer{contract: contract}}, nil
}

// ChannelImplementation is an auto generated Go binding around an Ethereum contract.
type ChannelImplementation struct {
	ChannelImplementationCaller     // Read-only binding to the contract
	ChannelImplementationTransactor // Write-only binding to the contract
	ChannelImplementationFilterer   // Log filterer for contract events
}

// ChannelImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelImplementationSession struct {
	Contract     *ChannelImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChannelImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelImplementationCallerSession struct {
	Contract *ChannelImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ChannelImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelImplementationTransactorSession struct {
	Contract     *ChannelImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ChannelImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelImplementationRaw struct {
	Contract *ChannelImplementation // Generic contract binding to access the raw methods on
}

// ChannelImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelImplementationCallerRaw struct {
	Contract *ChannelImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelImplementationTransactorRaw struct {
	Contract *ChannelImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannelImplementation creates a new instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementation(address common.Address, backend bind.ContractBackend) (*ChannelImplementation, error) {
	contract, err := bindChannelImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementation{ChannelImplementationCaller: ChannelImplementationCaller{contract: contract}, ChannelImplementationTransactor: ChannelImplementationTransactor{contract: contract}, ChannelImplementationFilterer: ChannelImplementationFilterer{contract: contract}}, nil
}

// NewChannelImplementationCaller creates a new read-only instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationCaller(address common.Address, caller bind.ContractCaller) (*ChannelImplementationCaller, error) {
	contract, err := bindChannelImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationCaller{contract: contract}, nil
}

// NewChannelImplementationTransactor creates a new write-only instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelImplementationTransactor, error) {
	contract, err := bindChannelImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationTransactor{contract: contract}, nil
}

// NewChannelImplementationFilterer creates a new log filterer instance of ChannelImplementation, bound to a specific deployed contract.
func NewChannelImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelImplementationFilterer, error) {
	contract, err := bindChannelImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationFilterer{contract: contract}, nil
}

// bindChannelImplementation binds a generic wrapper to an already deployed contract.
func bindChannelImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementation *ChannelImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChannelImplementation.Contract.ChannelImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementation *ChannelImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ChannelImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementation *ChannelImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ChannelImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChannelImplementation *ChannelImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChannelImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChannelImplementation *ChannelImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChannelImplementation *ChannelImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.contract.Transact(opts, method, params...)
}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationCaller) ExitRequest(opts *bind.CallOpts) (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "exitRequest")

	outstruct := new(struct {
		Timelock    *big.Int
		Beneficiary common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timelock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Beneficiary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationSession) ExitRequest() (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	return _ChannelImplementation.Contract.ExitRequest(&_ChannelImplementation.CallOpts)
}

// ExitRequest is a free data retrieval call binding the contract method 0xf4b3a197.
//
// Solidity: function exitRequest() view returns(uint256 timelock, address beneficiary)
func (_ChannelImplementation *ChannelImplementationCallerSession) ExitRequest() (struct {
	Timelock    *big.Int
	Beneficiary common.Address
}, error) {
	return _ChannelImplementation.Contract.ExitRequest(&_ChannelImplementation.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "getFundsDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) GetFundsDestination() (common.Address, error) {
	return _ChannelImplementation.Contract.GetFundsDestination(&_ChannelImplementation.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) GetFundsDestination() (common.Address, error) {
	return _ChannelImplementation.Contract.GetFundsDestination(&_ChannelImplementation.CallOpts)
}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationCaller) Hermes(opts *bind.CallOpts) (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "hermes")

	outstruct := new(struct {
		Operator        common.Address
		ContractAddress common.Address
		Settled         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Settled = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationSession) Hermes() (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	return _ChannelImplementation.Contract.Hermes(&_ChannelImplementation.CallOpts)
}

// Hermes is a free data retrieval call binding the contract method 0xd8092c92.
//
// Solidity: function hermes() view returns(address operator, address contractAddress, uint256 settled)
func (_ChannelImplementation *ChannelImplementationCallerSession) Hermes() (struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}, error) {
	return _ChannelImplementation.Contract.Hermes(&_ChannelImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationSession) IsInitialized() (bool, error) {
	return _ChannelImplementation.Contract.IsInitialized(&_ChannelImplementation.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChannelImplementation *ChannelImplementationCallerSession) IsInitialized() (bool, error) {
	return _ChannelImplementation.Contract.IsInitialized(&_ChannelImplementation.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Operator() (common.Address, error) {
	return _ChannelImplementation.Contract.Operator(&_ChannelImplementation.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Operator() (common.Address, error) {
	return _ChannelImplementation.Contract.Operator(&_ChannelImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Owner() (common.Address, error) {
	return _ChannelImplementation.Contract.Owner(&_ChannelImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Owner() (common.Address, error) {
	return _ChannelImplementation.Contract.Owner(&_ChannelImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChannelImplementation.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationSession) Token() (common.Address, error) {
	return _ChannelImplementation.Contract.Token(&_ChannelImplementation.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ChannelImplementation *ChannelImplementationCallerSession) Token() (common.Address, error) {
	return _ChannelImplementation.Contract.Token(&_ChannelImplementation.CallOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationSession) ClaimEthers() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimEthers(&_ChannelImplementation.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimEthers(&_ChannelImplementation.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimTokens(&_ChannelImplementation.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.ClaimTokens(&_ChannelImplementation.TransactOpts, _token)
}

// FastExit is a paid mutator transaction binding the contract method 0xe9e8ad8b.
//
// Solidity: function fastExit(uint256 _amount, uint256 _transactorFee, address _beneficiary, uint256 _validUntil, bytes _operatorSignature, bytes _hermesSignature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) FastExit(opts *bind.TransactOpts, _amount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _validUntil *big.Int, _operatorSignature []byte, _hermesSignature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "fastExit", _amount, _transactorFee, _beneficiary, _validUntil, _operatorSignature, _hermesSignature)
}

// FastExit is a paid mutator transaction binding the contract method 0xe9e8ad8b.
//
// Solidity: function fastExit(uint256 _amount, uint256 _transactorFee, address _beneficiary, uint256 _validUntil, bytes _operatorSignature, bytes _hermesSignature) returns()
func (_ChannelImplementation *ChannelImplementationSession) FastExit(_amount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _validUntil *big.Int, _operatorSignature []byte, _hermesSignature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FastExit(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _beneficiary, _validUntil, _operatorSignature, _hermesSignature)
}

// FastExit is a paid mutator transaction binding the contract method 0xe9e8ad8b.
//
// Solidity: function fastExit(uint256 _amount, uint256 _transactorFee, address _beneficiary, uint256 _validUntil, bytes _operatorSignature, bytes _hermesSignature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) FastExit(_amount *big.Int, _transactorFee *big.Int, _beneficiary common.Address, _validUntil *big.Int, _operatorSignature []byte, _hermesSignature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FastExit(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _beneficiary, _validUntil, _operatorSignature, _hermesSignature)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationTransactor) FinalizeExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "finalizeExit")
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationSession) FinalizeExit() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FinalizeExit(&_ChannelImplementation.TransactOpts)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x07e8ec1f.
//
// Solidity: function finalizeExit() returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) FinalizeExit() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.FinalizeExit(&_ChannelImplementation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dexAddress, address _identity, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _dexAddress common.Address, _identity common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "initialize", _token, _dexAddress, _identity, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dexAddress, address _identity, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationSession) Initialize(_token common.Address, _dexAddress common.Address, _identity common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dexAddress, _identity, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dexAddress, address _identity, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) Initialize(_token common.Address, _dexAddress common.Address, _identity common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dexAddress, _identity, _hermesId, _fee)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) RequestExit(opts *bind.TransactOpts, _beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "requestExit", _beneficiary, _validUntil, _signature)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) RequestExit(_beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RequestExit(&_ChannelImplementation.TransactOpts, _beneficiary, _validUntil, _signature)
}

// RequestExit is a paid mutator transaction binding the contract method 0x182f3488.
//
// Solidity: function requestExit(address _beneficiary, uint256 _validUntil, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) RequestExit(_beneficiary common.Address, _validUntil *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.RequestExit(&_ChannelImplementation.TransactOpts, _beneficiary, _validUntil, _signature)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestination(&_ChannelImplementation.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestination(&_ChannelImplementation.TransactOpts, _newDestination)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x6a2b76ad.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SetFundsDestinationByCheque(opts *bind.TransactOpts, _newDestination common.Address, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "setFundsDestinationByCheque", _newDestination, _signature)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x6a2b76ad.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) SetFundsDestinationByCheque(_newDestination common.Address, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestinationByCheque(&_ChannelImplementation.TransactOpts, _newDestination, _signature)
}

// SetFundsDestinationByCheque is a paid mutator transaction binding the contract method 0x6a2b76ad.
//
// Solidity: function setFundsDestinationByCheque(address _newDestination, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SetFundsDestinationByCheque(_newDestination common.Address, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SetFundsDestinationByCheque(&_ChannelImplementation.TransactOpts, _newDestination, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) SettlePromise(opts *bind.TransactOpts, _amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "settlePromise", _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationSession) SettlePromise(_amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SettlePromise(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _lock, _signature)
}

// SettlePromise is a paid mutator transaction binding the contract method 0x6f174630.
//
// Solidity: function settlePromise(uint256 _amount, uint256 _transactorFee, bytes32 _lock, bytes _signature) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) SettlePromise(_amount *big.Int, _transactorFee *big.Int, _lock [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.SettlePromise(&_ChannelImplementation.TransactOpts, _amount, _transactorFee, _lock, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.TransferOwnership(&_ChannelImplementation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.TransferOwnership(&_ChannelImplementation.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelImplementation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationSession) Receive() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Receive(&_ChannelImplementation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) Receive() (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Receive(&_ChannelImplementation.TransactOpts)
}

// ChannelImplementationDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the ChannelImplementation contract.
type ChannelImplementationDestinationChangedIterator struct {
	Event *ChannelImplementationDestinationChanged // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationDestinationChanged)
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
		it.Event = new(ChannelImplementationDestinationChanged)
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
func (it *ChannelImplementationDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationDestinationChanged represents a DestinationChanged event raised by the ChannelImplementation contract.
type ChannelImplementationDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*ChannelImplementationDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationDestinationChangedIterator{contract: _ChannelImplementation.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *ChannelImplementationDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationDestinationChanged)
				if err := _ChannelImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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
func (_ChannelImplementation *ChannelImplementationFilterer) ParseDestinationChanged(log types.Log) (*ChannelImplementationDestinationChanged, error) {
	event := new(ChannelImplementationDestinationChanged)
	if err := _ChannelImplementation.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelImplementationExitRequestedIterator is returned from FilterExitRequested and is used to iterate over the raw logs and unpacked data for ExitRequested events raised by the ChannelImplementation contract.
type ChannelImplementationExitRequestedIterator struct {
	Event *ChannelImplementationExitRequested // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationExitRequested)
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
		it.Event = new(ChannelImplementationExitRequested)
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
func (it *ChannelImplementationExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationExitRequested represents a ExitRequested event raised by the ChannelImplementation contract.
type ChannelImplementationExitRequested struct {
	Timelock *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitRequested is a free log retrieval operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterExitRequested(opts *bind.FilterOpts) (*ChannelImplementationExitRequestedIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "ExitRequested")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationExitRequestedIterator{contract: _ChannelImplementation.contract, event: "ExitRequested", logs: logs, sub: sub}, nil
}

// WatchExitRequested is a free log subscription operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchExitRequested(opts *bind.WatchOpts, sink chan<- *ChannelImplementationExitRequested) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "ExitRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationExitRequested)
				if err := _ChannelImplementation.contract.UnpackLog(event, "ExitRequested", log); err != nil {
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

// ParseExitRequested is a log parse operation binding the contract event 0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db.
//
// Solidity: event ExitRequested(uint256 timelock)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseExitRequested(log types.Log) (*ChannelImplementationExitRequested, error) {
	event := new(ChannelImplementationExitRequested)
	if err := _ChannelImplementation.contract.UnpackLog(event, "ExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelImplementationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChannelImplementation contract.
type ChannelImplementationOwnershipTransferredIterator struct {
	Event *ChannelImplementationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationOwnershipTransferred)
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
		it.Event = new(ChannelImplementationOwnershipTransferred)
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
func (it *ChannelImplementationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationOwnershipTransferred represents a OwnershipTransferred event raised by the ChannelImplementation contract.
type ChannelImplementationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChannelImplementationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationOwnershipTransferredIterator{contract: _ChannelImplementation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChannelImplementationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationOwnershipTransferred)
				if err := _ChannelImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChannelImplementation *ChannelImplementationFilterer) ParseOwnershipTransferred(log types.Log) (*ChannelImplementationOwnershipTransferred, error) {
	event := new(ChannelImplementationOwnershipTransferred)
	if err := _ChannelImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelImplementationPromiseSettledIterator is returned from FilterPromiseSettled and is used to iterate over the raw logs and unpacked data for PromiseSettled events raised by the ChannelImplementation contract.
type ChannelImplementationPromiseSettledIterator struct {
	Event *ChannelImplementationPromiseSettled // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationPromiseSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationPromiseSettled)
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
		it.Event = new(ChannelImplementationPromiseSettled)
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
func (it *ChannelImplementationPromiseSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationPromiseSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationPromiseSettled represents a PromiseSettled event raised by the ChannelImplementation contract.
type ChannelImplementationPromiseSettled struct {
	Beneficiary  common.Address
	Amount       *big.Int
	TotalSettled *big.Int
	Lock         [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPromiseSettled is a free log retrieval operation binding the contract event 0x7f7697dfbfc8e203dd21764a16312ebed4de12b161eca0b41922de085b2eccd1.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled, bytes32 lock)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterPromiseSettled(opts *bind.FilterOpts) (*ChannelImplementationPromiseSettledIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "PromiseSettled")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationPromiseSettledIterator{contract: _ChannelImplementation.contract, event: "PromiseSettled", logs: logs, sub: sub}, nil
}

// WatchPromiseSettled is a free log subscription operation binding the contract event 0x7f7697dfbfc8e203dd21764a16312ebed4de12b161eca0b41922de085b2eccd1.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled, bytes32 lock)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchPromiseSettled(opts *bind.WatchOpts, sink chan<- *ChannelImplementationPromiseSettled) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "PromiseSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationPromiseSettled)
				if err := _ChannelImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
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

// ParsePromiseSettled is a log parse operation binding the contract event 0x7f7697dfbfc8e203dd21764a16312ebed4de12b161eca0b41922de085b2eccd1.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled, bytes32 lock)
func (_ChannelImplementation *ChannelImplementationFilterer) ParsePromiseSettled(log types.Log) (*ChannelImplementationPromiseSettled, error) {
	event := new(ChannelImplementationPromiseSettled)
	if err := _ChannelImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelImplementationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ChannelImplementation contract.
type ChannelImplementationWithdrawIterator struct {
	Event *ChannelImplementationWithdraw // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationWithdraw)
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
		it.Event = new(ChannelImplementationWithdraw)
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
func (it *ChannelImplementationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationWithdraw represents a Withdraw event raised by the ChannelImplementation contract.
type ChannelImplementationWithdraw struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address beneficiary, uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterWithdraw(opts *bind.FilterOpts) (*ChannelImplementationWithdrawIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationWithdrawIterator{contract: _ChannelImplementation.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address beneficiary, uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ChannelImplementationWithdraw) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationWithdraw)
				if err := _ChannelImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address beneficiary, uint256 amount)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseWithdraw(log types.Log) (*ChannelImplementationWithdraw, error) {
	event := new(ChannelImplementationWithdraw)
	if err := _ChannelImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
