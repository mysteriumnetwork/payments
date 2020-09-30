/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
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

// ChannelImplementationABI is the input ABI used to generate the binding from.
const ChannelImplementationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hermes\",\"type\":\"address\"}],\"name\":\"ChannelInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"}],\"name\":\"ExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSettled\",\"type\":\"uint256\"}],\"name\":\"PromiseSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hermes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"settled\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_identityHash\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hermesId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settlePromise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"requestExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transactorFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_operatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_hermesSignature\",\"type\":\"bytes\"}],\"name\":\"fastExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"setFundsDestinationByCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChannelImplementationBin is the compiled bytecode used for deploying new contracts.
var ChannelImplementationBin = "0x608060405234801561001057600080fd5b50612348806100206000396000f3fe6080604052600436106101025760003560e01c80638da5cb5b11610095578063f2fde38b11610064578063f2fde38b1461083b578063f4b3a1971461086e578063f58c5b6e146108a4578063f7013ef6146108b9578063fc0c546a1461090c57610354565b80638da5cb5b1461065d578063d8092c9214610672578063df8de3e7146106b1578063e9e8ad8b146106e457610354565b8063570ca735116100d1578063570ca735146104925780636931b550146104c35780636a2b76ad146104d85780636f1746301461059957610354565b806307e8ec1f14610359578063182f348814610370578063238e130a14610436578063392e53cd1461046957610354565b366103545760408051600280825260608083018452926020830190803683375050600a54604080516315ab88c960e31b815290519394506001600160a01b039091169263ad5c464892506004808301926020929190829003018186803b15801561016b57600080fd5b505afa15801561017f573d6000803e3d6000fd5b505050506040513d602081101561019557600080fd5b5051815182906000906101a457fe5b6001600160a01b0392831660209182029290920101526002548251911690829060019081106101cf57fe5b6001600160a01b03928316602091820292909201810191909152600a54604051637ff36ab560e01b8152600060048201818152306044840181905242606485018190526080602486019081528951608487015289519690981697637ff36ab597349795968b969495939460a49091019187810191028083838b5b83811015610261578181015183820152602001610249565b50505050905001955050505050506000604051808303818588803b15801561028857600080fd5b505af115801561029c573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405260208110156102c657600080fd5b8101908080516040519392919084600160201b8211156102e557600080fd5b9083019060208201858111156102fa57600080fd5b82518660208202830111600160201b8211171561031657600080fd5b82525081516020918201928201910280838360005b8381101561034357818101518382015260200161032b565b505050509050016040525050505050005b600080fd5b34801561036557600080fd5b5061036e610921565b005b34801561037c57600080fd5b5061036e6004803603606081101561039357600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156103c257600080fd5b8201836020820111156103d457600080fd5b803590602001918460018302840111600160201b831117156103f557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ae1945050505050565b34801561044257600080fd5b5061036e6004803603602081101561045957600080fd5b50356001600160a01b0316610d9b565b34801561047557600080fd5b5061047e610e7d565b604080519115158252519081900360200190f35b34801561049e57600080fd5b506104a7610e8e565b604080516001600160a01b039092168252519081900360200190f35b3480156104cf57600080fd5b5061036e610e9d565b3480156104e457600080fd5b5061036e600480360360408110156104fb57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561052557600080fd5b82018360208201111561053757600080fd5b803590602001918460018302840111600160201b8311171561055857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610eee945050505050565b3480156105a557600080fd5b5061036e600480360360808110156105bc57600080fd5b81359160208101359160408201359190810190608081016060820135600160201b8111156105e957600080fd5b8201836020820111156105fb57600080fd5b803590602001918460018302840111600160201b8311171561061c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061100d945050505050565b34801561066957600080fd5b506104a761132a565b34801561067e57600080fd5b50610687611339565b604080516001600160a01b0394851681529290931660208301528183015290519081900360600190f35b3480156106bd57600080fd5b5061036e600480360360208110156106d457600080fd5b50356001600160a01b0316611355565b3480156106f057600080fd5b5061036e600480360360c081101561070757600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b81111561074257600080fd5b82018360208201111561075457600080fd5b803590602001918460018302840111600160201b8311171561077557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156107c757600080fd5b8201836020820111156107d957600080fd5b803590602001918460018302840111600160201b831117156107fa57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114bb945050505050565b34801561084757600080fd5b5061036e6004803603602081101561085e57600080fd5b50356001600160a01b031661183f565b34801561087a57600080fd5b50610883611952565b604080519283526001600160a01b0390911660208301528051918290030190f35b3480156108b057600080fd5b506104a7611964565b3480156108c557600080fd5b5061036e600480360360a08110156108dc57600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135611973565b34801561091857600080fd5b506104a7611c97565b6004541580159061093457506004544310155b61096f5760405162461bcd60e51b8152600401808060200182810382526042815260200180611fd46042913960600191505060405180910390fd5b600254604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b1580156109ba57600080fd5b505afa1580156109ce573d6000803e3d6000fd5b505050506040513d60208110156109e457600080fd5b50516002546005546040805163a9059cbb60e01b81526001600160a01b03928316600482015260248101859052905193945091169163a9059cbb916044808201926020929091908290030181600087803b158015610a4157600080fd5b505af1158015610a55573d6000803e3d6000fd5b505050506040513d6020811015610a6b57600080fd5b5050600554604080516001600160a01b0390921682526020820183905280517f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243649281900390910190a1506040805180820190915260008082526020909101819052600455600580546001600160a01b0319169055565b6000610aeb611ca6565b60045490915015610b2d5760405162461bcd60e51b81526004018080602001828103825260428152602001806120636042913960600191505060405180910390fd5b438311610b6b5760405162461bcd60e51b81526004018080602001828103825260418152602001806120a56041913960600191505060405180910390fd5b828111610ba95760405162461bcd60e51b815260040180806020018281038252603b815260200180612153603b913960400191505060405180910390fd5b6001600160a01b038416610bee5760405162461bcd60e51b815260040180806020018281038252602a81526020018061223c602a913960400191505060405180910390fd5b6009546001600160a01b03163314610d2e5760003090506000610cdc846040518060400160405280600d81526020016c22bc34ba103932b8bab2b9ba1d60991b8152508489896040516020018085805190602001908083835b60208310610c665780518252601f199092019160209182019101610c47565b6001836020036101000a038019825116818451168082178552505050505050905001846001600160a01b031660601b8152601401836001600160a01b031660601b815260140182815260200194505050505060405160208183030381529060405280519060200120611cae90919063ffffffff16565b6009549091506001600160a01b03808316911614610d2b5760405162461bcd60e51b81526004018080602001828103825260268152602001806122166026913960400191505060405180910390fd5b50505b6040805180820182528281526001600160a01b03861660209182018190526004849055600580546001600160a01b0319169091179055815183815291517fe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db9281900390910190a150505050565b6000546001600160a01b0316331480610dbd57506000546001600160a01b0316155b610e0e576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b038116610e2157600080fd5b6001546040516001600160a01b038084169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b6009546001600160a01b0316151590565b6009546001600160a01b031681565b6001546001600160a01b0316610eb257600080fd5b6001546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610eeb573d6000803e3d6000fd5b50565b6001600160a01b038216610f0157600080fd5b60408051808201825260168082527529b2ba10333ab73239903232b9ba34b730ba34b7b71d60511b6020808401918252600380546001810190915594513095600095610f5e958995919489948c9493910191829190808383610c66565b6009549091506001600160a01b03808316911614610fad5760405162461bcd60e51b815260040180806020018281038252602d81526020018061229d602d913960400191505060405180910390fd5b6001546040516001600160a01b038087169216907fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad90600090a35050600180546001600160a01b0319166001600160a01b03939093169290921790915550565b604080516020808201859052825180830382018152828401845280519082012030606084018190526080840189905260a0840188905260c08085018390528551808603909101815260e0909401909452825192909101919091209091906000906110779085611cae565b6009549091506001600160a01b038083169116146110c65760405162461bcd60e51b81526004018080602001828103825260258152602001806122ee6025913960400191505060405180910390fd5b6008546000906110d7908990611e99565b9050600081116111185760405162461bcd60e51b81526004018080602001828103825260378152602001806122666037913960400191505060405180910390fd5b600254604080516370a0823160e01b81526001600160a01b038681166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b15801561116957600080fd5b505afa15801561117d573d6000803e3d6000fd5b505050506040513d602081101561119357600080fd5b50519050808211156111a3578091505b6008546111b09083611ee2565b6008556002546007546001600160a01b039182169163a9059cbb91166111d6858c611e99565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561121c57600080fd5b505af1158015611230573d6000803e3d6000fd5b505050506040513d602081101561124657600080fd5b505087156112cf576002546040805163a9059cbb60e01b8152336004820152602481018b905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b1580156112a257600080fd5b505af11580156112b6573d6000803e3d6000fd5b505050506040513d60208110156112cc57600080fd5b50505b600754600854604080516001600160a01b0390931683526020830185905282810191909152517f50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b9181900360600190a1505050505050505050565b6000546001600160a01b031690565b6006546007546008546001600160a01b03928316929091169083565b6001546001600160a01b031661136a57600080fd5b6002546001600160a01b03828116911614156113b75760405162461bcd60e51b815260040180806020018281038252602581526020018061210c6025913960400191505060405180910390fd5b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561140657600080fd5b505afa15801561141a573d6000803e3d6000fd5b505050506040513d602081101561143057600080fd5b50516001546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905290519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561148b57600080fd5b505af115801561149f573d6000803e3d6000fd5b505050506040513d60208110156114b557600080fd5b50505050565b438310156114fa5760405162461bcd60e51b815260040180806020018281038252604d815260200180612016604d913960600191505060405180910390fd5b600030905060006040518060400160405280600d81526020016c22bc34ba103932b8bab2b9ba1d60991b815250826001600160a01b03168989896001600160a01b03168960036000815480929190600101919050556040516020018088805190602001908083835b602083106115815780518252601f199092019160209182019101611562565b51815160209384036101000a600019018019909216911617905292019889525087810196909652506040808701949094526060860192909252608085015260a0808501919091528151808503909101815260c0909301905281519101209150600090506115ee8286611cae565b6009549091506001600160a01b0380831691161461163d5760405162461bcd60e51b81526004018080602001828103825260268152602001806122166026913960400191505060405180910390fd5b60006116498386611cae565b6006549091506001600160a01b038083169116146116985760405162461bcd60e51b81526004018080602001828103825260248152602001806122ca6024913960400191505060405180910390fd5b881561175e57888a10156116dd5760405162461bcd60e51b815260040180806020018281038252603e8152602001806121d8603e913960400191505060405180910390fd5b6002546040805163a9059cbb60e01b8152336004820152602481018c905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b15801561173157600080fd5b505af1158015611745573d6000803e3d6000fd5b505050506040513d602081101561175b57600080fd5b50505b600061176a8b8b611e99565b6002546040805163a9059cbb60e01b81526001600160a01b038d8116600483015260248201859052915193945091169163a9059cbb916044808201926020929091908290030181600087803b1580156117c257600080fd5b505af11580156117d6573d6000803e3d6000fd5b505050506040513d60208110156117ec57600080fd5b5050604080516001600160a01b038b1681526020810183905281517f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364929181900390910190a15050505050505050505050565b6000546001600160a01b031633148061186157506000546001600160a01b0316155b6118b2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166118f75760405162461bcd60e51b81526004018080602001828103825260268152602001806120e66026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6004546005546001600160a01b031682565b6001546001600160a01b031690565b61197b610e7d565b156119c6576040805162461bcd60e51b8152602060048201526016602482015275125cc8185b1c9958591e481a5b9a5d1a585b1a5e995960521b604482015290519081900360640190fd5b6001600160a01b038316611a1a576040805162461bcd60e51b81526020600482015260166024820152754964656e746974792063616e2774206265207a65726f60501b604482015290519081900360640190fd5b6001600160a01b038216611a6e576040805162461bcd60e51b81526020600482015260166024820152754865726d657349442063616e2774206265207a65726f60501b604482015290519081900360640190fd5b6001600160a01b038516611ab35760405162461bcd60e51b815260040180806020018281038252602881526020018061218e6028913960400191505060405180910390fd5b600280546001600160a01b038088166001600160a01b031992831617909255600a8054928716929091169190911790558015611b6a576002546040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015611b3d57600080fd5b505af1158015611b51573d6000803e3d6000fd5b505050506040513d6020811015611b6757600080fd5b50505b600980546001600160a01b0319166001600160a01b038581169190911791829055611b95911661183f565b6040518060600160405280836001600160a01b031663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015611bd957600080fd5b505afa158015611bed573d6000803e3d6000fd5b505050506040513d6020811015611c0357600080fd5b50516001600160a01b039081168252848116602083810182905260006040948501528451600680549185166001600160a01b031992831617905585820151600780549186169190921617905593830151600855825191871682529281019290925280517f9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a519281900390910190a15050505050565b6002546001600160a01b031681565b436146500190565b60008151604114611d06576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611d775760405162461bcd60e51b81526004018080602001828103825260228152602001806121316022913960400191505060405180910390fd5b8060ff16601b14158015611d8f57508060ff16601c14155b15611dcb5760405162461bcd60e51b81526004018080602001828103825260228152602001806121b66022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611e27573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611e8f576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b9695505050505050565b6000611edb83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611f3c565b9392505050565b600082820183811015611edb576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008184841115611fcb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611f90578181015183820152602001611f78565b50505050905090810190601f168015611fbd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4368616e6e656c3a2065786974206861766520746f2062652072657175657374656420616e642074696d656c6f636b206861766520746f20626520696e20706173744368616e6e656c3a205f76616c6964556e74696c206861766520746f2062652067726561746572207468616e206f7220657175616c20746f2063757272656e7420626c6f636b206e756d6265724368616e6e656c3a206e657720657869742063616e20626520726571756573746564206f6e6c79207768656e206f6c64206f6e65207761732066696e616c697365644368616e6e656c3a2076616c696420756e74696c206861766520746f2062652067726561746572207468616e2063757272656e7420626c6f636b206e756d6265724f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736e617469766520746f6b656e2066756e64732063616e2774206265207265636f766572656445434453413a20696e76616c6964207369676e6174757265202773272076616c75654368616e6e656c3a2072657175657374206861766520746f2062652076616c69642073686f72746572207468616e2044454c41595f424c4f434b53546f6b656e2063616e2774206265206465706c6f796420696e746f207a65726f206164647265737345434453413a20696e76616c6964207369676e6174757265202776272076616c75654368616e6e656c3a207472616e736163746f72206665652063616e2774206265206269676765722074686174207769746864726177616c20616d6f756e744368616e6e656c3a206861766520746f206265207369676e6564206279206f70657261746f724368616e6e656c3a2062656e65666963696172792063616e2774206265207a65726f2061646472657373616d6f756e7420746f20736574746c652073686f756c642062652067726561746572207468617420616c726561647920736574746c65644368616e6e656c3a206861766520746f206265207369676e65642062792070726f706572206964656e746974794368616e6e656c3a206861766520746f206265207369676e6564206279206865726d65736861766520746f206265207369676e6564206279206368616e6e656c206f70657261746f72a2646970667358221220fa8fdff7e17189846a35e37e8ddb420b7db77a39b7c25b366c15d7efee7bd8c864736f6c63430007010033"

// DeployChannelImplementation deploys a new Ethereum contract, binding an instance of ChannelImplementation to it.
func DeployChannelImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChannelImplementation, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelImplementationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelImplementationBin), backend)
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
func (_ChannelImplementation *ChannelImplementationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ChannelImplementation *ChannelImplementationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	ret := new(struct {
		Timelock    *big.Int
		Beneficiary common.Address
	})
	out := ret
	err := _ChannelImplementation.contract.Call(opts, out, "exitRequest")
	return *ret, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
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
	ret := new(struct {
		Operator        common.Address
		ContractAddress common.Address
		Settled         *big.Int
	})
	out := ret
	err := _ChannelImplementation.contract.Call(opts, out, "hermes")
	return *ret, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "isInitialized")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "operator")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "owner")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChannelImplementation.contract.Call(opts, out, "token")
	return *ret0, err
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
// Solidity: function initialize(address _token, address _dexAddress, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _dexAddress common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.contract.Transact(opts, "initialize", _token, _dexAddress, _identityHash, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dexAddress, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationSession) Initialize(_token common.Address, _dexAddress common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dexAddress, _identityHash, _hermesId, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _token, address _dexAddress, address _identityHash, address _hermesId, uint256 _fee) returns()
func (_ChannelImplementation *ChannelImplementationTransactorSession) Initialize(_token common.Address, _dexAddress common.Address, _identityHash common.Address, _hermesId common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _ChannelImplementation.Contract.Initialize(&_ChannelImplementation.TransactOpts, _token, _dexAddress, _identityHash, _hermesId, _fee)
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

// ChannelImplementationChannelInitialisedIterator is returned from FilterChannelInitialised and is used to iterate over the raw logs and unpacked data for ChannelInitialised events raised by the ChannelImplementation contract.
type ChannelImplementationChannelInitialisedIterator struct {
	Event *ChannelImplementationChannelInitialised // Event containing the contract specifics and raw log

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
func (it *ChannelImplementationChannelInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelImplementationChannelInitialised)
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
		it.Event = new(ChannelImplementationChannelInitialised)
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
func (it *ChannelImplementationChannelInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelImplementationChannelInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelImplementationChannelInitialised represents a ChannelInitialised event raised by the ChannelImplementation contract.
type ChannelImplementationChannelInitialised struct {
	Operator common.Address
	Hermes   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelInitialised is a free log retrieval operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterChannelInitialised(opts *bind.FilterOpts) (*ChannelImplementationChannelInitialisedIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "ChannelInitialised")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationChannelInitialisedIterator{contract: _ChannelImplementation.contract, event: "ChannelInitialised", logs: logs, sub: sub}, nil
}

// WatchChannelInitialised is a free log subscription operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) WatchChannelInitialised(opts *bind.WatchOpts, sink chan<- *ChannelImplementationChannelInitialised) (event.Subscription, error) {

	logs, sub, err := _ChannelImplementation.contract.WatchLogs(opts, "ChannelInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelImplementationChannelInitialised)
				if err := _ChannelImplementation.contract.UnpackLog(event, "ChannelInitialised", log); err != nil {
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

// ParseChannelInitialised is a log parse operation binding the contract event 0x9a7def6556351196c74c99e1cc8dcd284e9da181ea854c3e6367cc9fad882a51.
//
// Solidity: event ChannelInitialised(address operator, address hermes)
func (_ChannelImplementation *ChannelImplementationFilterer) ParseChannelInitialised(log types.Log) (*ChannelImplementationChannelInitialised, error) {
	event := new(ChannelImplementationChannelInitialised)
	if err := _ChannelImplementation.contract.UnpackLog(event, "ChannelInitialised", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPromiseSettled is a free log retrieval operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
func (_ChannelImplementation *ChannelImplementationFilterer) FilterPromiseSettled(opts *bind.FilterOpts) (*ChannelImplementationPromiseSettledIterator, error) {

	logs, sub, err := _ChannelImplementation.contract.FilterLogs(opts, "PromiseSettled")
	if err != nil {
		return nil, err
	}
	return &ChannelImplementationPromiseSettledIterator{contract: _ChannelImplementation.contract, event: "PromiseSettled", logs: logs, sub: sub}, nil
}

// WatchPromiseSettled is a free log subscription operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
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

// ParsePromiseSettled is a log parse operation binding the contract event 0x50c3491624aa1825a7653df63d067fecd5c8634ba63c99c4a7cf04ff1436070b.
//
// Solidity: event PromiseSettled(address beneficiary, uint256 amount, uint256 totalSettled)
func (_ChannelImplementation *ChannelImplementationFilterer) ParsePromiseSettled(log types.Log) (*ChannelImplementationPromiseSettled, error) {
	event := new(ChannelImplementationPromiseSettled)
	if err := _ChannelImplementation.contract.UnpackLog(event, "PromiseSettled", log); err != nil {
		return nil, err
	}
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
	return event, nil
}
