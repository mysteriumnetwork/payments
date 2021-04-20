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

// RewarderABI is the input ABI used to generate the binding from.
const RewarderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_custody\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDropped\",\"type\":\"uint256\"}],\"name\":\"Airdrop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUnclaimed\",\"type\":\"uint256\"}],\"name\":\"ClaimedChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalEarnings\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalEarned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custody\",\"outputs\":[{\"internalType\":\"contractCustody\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalEarned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"isValidProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRootBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalPayoutsFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalReward\",\"type\":\"uint256\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewarderBin is the compiled bytecode used for deploying new contracts.
var RewarderBin = "0x60806040523480156200001157600080fd5b50604051620024dc380380620024dc8339818101604052810190620000379190620001ed565b600062000049620001b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550436006819055507f156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f36000436000604051620001a79392919062000261565b60405180910390a150506200036d565b600033905090565b600081519050620001d08162000339565b92915050565b600081519050620001e78162000353565b92915050565b600080604083850312156200020157600080fd5b60006200021185828601620001bf565b92505060206200022485828601620001d6565b9150509250929050565b6200023981620002fa565b82525050565b6200024a8162000318565b82525050565b6200025b81620002f0565b82525050565b60006060820190506200027860008301866200022e565b62000287602083018562000250565b6200029660408301846200023f565b949350505050565b6000620002ab82620002d0565b9050919050565b6000620002bf82620002d0565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000620003116200030b83620002c6565b6200032c565b9050919050565b60006200032582620002f0565b9050919050565b60008160001b9050919050565b62000344816200029e565b81146200035057600080fd5b50565b6200035e81620002b2565b81146200036a57600080fd5b50565b61215f806200037d6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063d54ad2a111610066578063d54ad2a114610235578063dda79b7514610253578063f2fde38b14610271578063fc0c546a1461028d576100ea565b80638da5cb5b146101c95780638e59d422146101e7578063ca5a278014610205576100ea565b80634d676fa7116100c85780634d676fa7146101435780636724348214610173578063715018a61461018f5780637558b6ef14610199576100ea565b8063056097ac146100ef57806312b6af241461010b578063172bd6de14610127575b600080fd5b61010960048036038101906101049190611739565b6102ab565b005b610125600480360381019061012091906116ea565b6104d8565b005b610141600480360381019061013c91906115cc565b6107e1565b005b61015d600480360381019061015891906115a3565b6109db565b60405161016a9190611c65565b60405180910390f35b61018d6004803603810190610188919061164c565b6109f3565b005b610197610ee4565b005b6101b360048036038101906101ae9190611775565b61101e565b6040516101c09190611a9d565b60405180910390f35b6101d1611036565b6040516101de9190611a3e565b60405180910390f35b6101ef61105f565b6040516101fc9190611c65565b60405180910390f35b61021f600480360381019061021a91906115cc565b611065565b60405161022c9190611a82565b60405180910390f35b61023d611118565b60405161024a9190611c65565b60405180910390f35b61025b61111e565b6040516102689190611aef565b60405180910390f35b61028b600480360381019061028691906115a3565b611144565b005b6102956112ed565b6040516102a29190611b0a565b60405180910390f35b6102b3611313565b73ffffffffffffffffffffffffffffffffffffffff166102d1611036565b73ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e90611bc5565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156103b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103af90611be5565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103f39190611a3e565b60206040518083038186803b15801561040b57600080fd5b505afa15801561041f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610443919061179e565b90508273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610480929190611a59565b602060405180830381600087803b15801561049a57600080fd5b505af11580156104ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d291906116c1565b50505050565b6104e0611313565b73ffffffffffffffffffffffffffffffffffffffff166104fe611036565b73ffffffffffffffffffffffffffffffffffffffff1614610554576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054b90611bc5565b60405180910390fd5b438210610596576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058d90611ba5565b60405180910390fd5b60065482116105da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d190611c45565b60405180910390fd5b600554811161061e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061590611b85565b60405180910390fd5b60006005548261062e9190611ce7565b90506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161068d9190611a3e565b60206040518083038186803b1580156106a557600080fd5b505afa1580156106b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106dd919061179e565b90508082111561078057600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d82846107319190611ce7565b6040518263ffffffff1660e01b815260040161074d9190611c65565b600060405180830381600087803b15801561076757600080fd5b505af115801561077b573d6000803e3d6000fd5b505050505b836006819055508460036000868152602001908152602001600020819055507f156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f38585856040516107d293929190611ab8565b60405180910390a15050505050565b6107ee8585858585611065565b61082d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082490611c05565b60405180910390fd5b6000600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050808510156108b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ab90611b65565b60405180910390fd5b600081866108c29190611ce7565b905060008114156108d45750506109d4565b85600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506109218161131b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb88836040518363ffffffff1660e01b815260040161097e929190611a59565b602060405180830381600087803b15801561099857600080fd5b505af11580156109ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d091906116c1565b5050505b5050505050565b60046020528060005260406000206000915090505481565b6109fb611313565b73ffffffffffffffffffffffffffffffffffffffff16610a19611036565b73ffffffffffffffffffffffffffffffffffffffff1614610a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6690611bc5565b60405180910390fd5b818190508484905014610ab7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aae90611b25565b60405180910390fd5b60008282905067ffffffffffffffff811115610afc577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610b2a5781602001602082028036833780820191505090505b5090506000805b86869050811015610d2b576000878783818110610b77577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002016020810190610b8c91906115a3565b90506000868684818110610bc9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013590506000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905080821015610c59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5090611c25565b60405180910390fd5b60008183610c679190611ce7565b90506000811415610c7b5750505050610d18565b80878681518110610cb5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250508086610ccd9190611c91565b955082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505b8080610d2390611dc7565b915050610b31565b506000811415610d3c575050610ede565b610d458161131b565b60005b86869050811015610ea357600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb888884818110610dca577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002016020810190610ddf91906115a3565b858481518110610e18577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b8152600401610e3d929190611a59565b602060405180830381600087803b158015610e5757600080fd5b505af1158015610e6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8f91906116c1565b508080610e9b90611dc7565b915050610d48565b507fd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e381604051610ed39190611c65565b60405180910390a150505b50505050565b610eec611313565b73ffffffffffffffffffffffffffffffffffffffff16610f0a611036565b73ffffffffffffffffffffffffffffffffffffffff1614610f60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5790611bc5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60036020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60065481565b60008046905060008787833060405160200161108494939291906119c4565b60405160208183030381529060405280519060200120905060006003600088815260200190815260200160002054905061110a8183888880806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505061136b9092919063ffffffff16565b935050505095945050505050565b60055481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61114c611313565b73ffffffffffffffffffffffffffffffffffffffff1661116a611036565b73ffffffffffffffffffffffffffffffffffffffff16146111c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b790611bc5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611230576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122790611b45565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b806005546113299190611c91565b6005819055507fc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b6656005546040516113609190611c65565b60405180910390a150565b60008082905060005b85518110156114395760008682815181106113b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190508083116113f95782816040516020016113dc929190611a12565b604051602081830303815290604052805190602001209250611425565b808360405160200161140c929190611a12565b6040516020818303038152906040528051906020012092505b50808061143190611dc7565b915050611374565b508381149150509392505050565b600081359050611456816120b6565b92915050565b60008083601f84011261146e57600080fd5b8235905067ffffffffffffffff81111561148757600080fd5b60208301915083602082028301111561149f57600080fd5b9250929050565b60008083601f8401126114b857600080fd5b8235905067ffffffffffffffff8111156114d157600080fd5b6020830191508360208202830111156114e957600080fd5b9250929050565b60008083601f84011261150257600080fd5b8235905067ffffffffffffffff81111561151b57600080fd5b60208301915083602082028301111561153357600080fd5b9250929050565b600081519050611549816120cd565b92915050565b60008135905061155e816120e4565b92915050565b600081359050611573816120fb565b92915050565b60008135905061158881612112565b92915050565b60008151905061159d81612112565b92915050565b6000602082840312156115b557600080fd5b60006115c384828501611447565b91505092915050565b6000806000806000608086880312156115e457600080fd5b60006115f288828901611447565b955050602061160388828901611579565b945050604061161488828901611579565b935050606086013567ffffffffffffffff81111561163157600080fd5b61163d888289016114a6565b92509250509295509295909350565b6000806000806040858703121561166257600080fd5b600085013567ffffffffffffffff81111561167c57600080fd5b6116888782880161145c565b9450945050602085013567ffffffffffffffff8111156116a757600080fd5b6116b3878288016114f0565b925092505092959194509250565b6000602082840312156116d357600080fd5b60006116e18482850161153a565b91505092915050565b6000806000606084860312156116ff57600080fd5b600061170d8682870161154f565b935050602061171e86828701611579565b925050604061172f86828701611579565b9150509250925092565b6000806040838503121561174c57600080fd5b600061175a85828601611564565b925050602061176b85828601611447565b9150509250929050565b60006020828403121561178757600080fd5b600061179584828501611579565b91505092915050565b6000602082840312156117b057600080fd5b60006117be8482850161158e565b91505092915050565b6117d081611d1b565b82525050565b6117e76117e282611d1b565b611e10565b82525050565b6117f681611d2d565b82525050565b61180581611d39565b82525050565b61181c61181782611d39565b611e22565b82525050565b61182b81611d7f565b82525050565b61183a81611da3565b82525050565b600061184d601483611c80565b915061185882611e84565b602082019050919050565b6000611870602683611c80565b915061187b82611ead565b604082019050919050565b6000611893600c83611c80565b915061189e82611efc565b602082019050919050565b60006118b6602e83611c80565b91506118c182611f25565b604082019050919050565b60006118d9603983611c80565b91506118e482611f74565b604082019050919050565b60006118fc602083611c80565b915061190782611fc3565b602082019050919050565b600061191f601f83611c80565b915061192a82611fec565b602082019050919050565b6000611942600d83611c80565b915061194d82612015565b602082019050919050565b6000611965600d83611c80565b91506119708261203e565b602082019050919050565b6000611988603483611c80565b915061199382612067565b604082019050919050565b6119a781611d75565b82525050565b6119be6119b982611d75565b611e3e565b82525050565b60006119d082876117d6565b6014820191506119e082866119ad565b6020820191506119f082856119ad565b602082019150611a0082846117d6565b60148201915081905095945050505050565b6000611a1e828561180b565b602082019150611a2e828461180b565b6020820191508190509392505050565b6000602082019050611a5360008301846117c7565b92915050565b6000604082019050611a6e60008301856117c7565b611a7b602083018461199e565b9392505050565b6000602082019050611a9760008301846117ed565b92915050565b6000602082019050611ab260008301846117fc565b92915050565b6000606082019050611acd60008301866117fc565b611ada602083018561199e565b611ae7604083018461199e565b949350505050565b6000602082019050611b046000830184611822565b92915050565b6000602082019050611b1f6000830184611831565b92915050565b60006020820190508181036000830152611b3e81611840565b9050919050565b60006020820190508181036000830152611b5e81611863565b9050919050565b60006020820190508181036000830152611b7e81611886565b9050919050565b60006020820190508181036000830152611b9e816118a9565b9050919050565b60006020820190508181036000830152611bbe816118cc565b9050919050565b60006020820190508181036000830152611bde816118ef565b9050919050565b60006020820190508181036000830152611bfe81611912565b9050919050565b60006020820190508181036000830152611c1e81611935565b9050919050565b60006020820190508181036000830152611c3e81611958565b9050919050565b60006020820190508181036000830152611c5e8161197b565b9050919050565b6000602082019050611c7a600083018461199e565b92915050565b600082825260208201905092915050565b6000611c9c82611d75565b9150611ca783611d75565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611cdc57611cdb611e48565b5b828201905092915050565b6000611cf282611d75565b9150611cfd83611d75565b925082821015611d1057611d0f611e48565b5b828203905092915050565b6000611d2682611d55565b9050919050565b60008115159050919050565b6000819050919050565b6000611d4e82611d1b565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611d8a82611d91565b9050919050565b6000611d9c82611d55565b9050919050565b6000611dae82611db5565b9050919050565b6000611dc082611d55565b9050919050565b6000611dd282611d75565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e0557611e04611e48565b5b600182019050919050565b6000611e1b82611e2c565b9050919050565b6000819050919050565b6000611e3782611e77565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160601b9050919050565b7f496e76616c6964206172726179206c656e677468000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f416c726561647920706169640000000000000000000000000000000000000000600082015250565b7f546f74616c20726577617264206d75737420626520626967676572207468616e60008201527f20746f74616c20636c61696d6564000000000000000000000000000000000000602082015250565b7f476976656e20626c6f636b206e756d626572206d757374206265206c6573732060008201527f7468616e2063757272656e7420626c6f636b206e756d62657200000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f596f752063616e2774207265636f7665722064656661756c7420746f6b656e00600082015250565b7f496e76616c69642070726f6f6600000000000000000000000000000000000000600082015250565b7f496e76616c696420626174636800000000000000000000000000000000000000600082015250565b7f476976656e20626c6f636b206e756d626572206d757374206265206d6f72652060008201527f7468616e206c61737420726f6f7420626c6f636b000000000000000000000000602082015250565b6120bf81611d1b565b81146120ca57600080fd5b50565b6120d681611d2d565b81146120e157600080fd5b50565b6120ed81611d39565b81146120f857600080fd5b50565b61210481611d43565b811461210f57600080fd5b50565b61211b81611d75565b811461212657600080fd5b5056fea26469706673582212207841c9994766935460ced5cc08ae01d60095f86e8c7d5570358f497ed7dec28564736f6c63430008030033"

// DeployRewarder deploys a new Ethereum contract, binding an instance of Rewarder to it.
func DeployRewarder(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _custody common.Address) (common.Address, *types.Transaction, *Rewarder, error) {
	parsed, err := abi.JSON(strings.NewReader(RewarderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewarderBin), backend, _token, _custody)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rewarder{RewarderCaller: RewarderCaller{contract: contract}, RewarderTransactor: RewarderTransactor{contract: contract}, RewarderFilterer: RewarderFilterer{contract: contract}}, nil
}

// Rewarder is an auto generated Go binding around an Ethereum contract.
type Rewarder struct {
	RewarderCaller     // Read-only binding to the contract
	RewarderTransactor // Write-only binding to the contract
	RewarderFilterer   // Log filterer for contract events
}

// RewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewarderSession struct {
	Contract     *Rewarder         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewarderCallerSession struct {
	Contract *RewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewarderTransactorSession struct {
	Contract     *RewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewarderRaw struct {
	Contract *Rewarder // Generic contract binding to access the raw methods on
}

// RewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewarderCallerRaw struct {
	Contract *RewarderCaller // Generic read-only contract binding to access the raw methods on
}

// RewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewarderTransactorRaw struct {
	Contract *RewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewarder creates a new instance of Rewarder, bound to a specific deployed contract.
func NewRewarder(address common.Address, backend bind.ContractBackend) (*Rewarder, error) {
	contract, err := bindRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewarder{RewarderCaller: RewarderCaller{contract: contract}, RewarderTransactor: RewarderTransactor{contract: contract}, RewarderFilterer: RewarderFilterer{contract: contract}}, nil
}

// NewRewarderCaller creates a new read-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderCaller(address common.Address, caller bind.ContractCaller) (*RewarderCaller, error) {
	contract, err := bindRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderCaller{contract: contract}, nil
}

// NewRewarderTransactor creates a new write-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*RewarderTransactor, error) {
	contract, err := bindRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderTransactor{contract: contract}, nil
}

// NewRewarderFilterer creates a new log filterer instance of Rewarder, bound to a specific deployed contract.
func NewRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*RewarderFilterer, error) {
	contract, err := bindRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewarderFilterer{contract: contract}, nil
}

// bindRewarder binds a generic wrapper to an already deployed contract.
func bindRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.RewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transact(opts, method, params...)
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderCaller) ClaimRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "claimRoots", arg0)
	return *ret0, err
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderSession) ClaimRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rewarder.Contract.ClaimRoots(&_Rewarder.CallOpts, arg0)
}

// ClaimRoots is a free data retrieval call binding the contract method 0x7558b6ef.
//
// Solidity: function claimRoots(uint256 ) view returns(bytes32)
func (_Rewarder *RewarderCallerSession) ClaimRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rewarder.Contract.ClaimRoots(&_Rewarder.CallOpts, arg0)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "custody")
	return *ret0, err
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderSession) Custody() (common.Address, error) {
	return _Rewarder.Contract.Custody(&_Rewarder.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_Rewarder *RewarderCallerSession) Custody() (common.Address, error) {
	return _Rewarder.Contract.Custody(&_Rewarder.CallOpts)
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderCaller) IsValidProof(opts *bind.CallOpts, _recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "isValidProof", _recipient, _totalEarned, _blockNumber, _proof)
	return *ret0, err
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderSession) IsValidProof(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	return _Rewarder.Contract.IsValidProof(&_Rewarder.CallOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// IsValidProof is a free data retrieval call binding the contract method 0xca5a2780.
//
// Solidity: function isValidProof(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) view returns(bool)
func (_Rewarder *RewarderCallerSession) IsValidProof(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (bool, error) {
	return _Rewarder.Contract.IsValidProof(&_Rewarder.CallOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderCaller) LastRootBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "lastRootBlock")
	return *ret0, err
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderSession) LastRootBlock() (*big.Int, error) {
	return _Rewarder.Contract.LastRootBlock(&_Rewarder.CallOpts)
}

// LastRootBlock is a free data retrieval call binding the contract method 0x8e59d422.
//
// Solidity: function lastRootBlock() view returns(uint256)
func (_Rewarder *RewarderCallerSession) LastRootBlock() (*big.Int, error) {
	return _Rewarder.Contract.LastRootBlock(&_Rewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderSession) Owner() (common.Address, error) {
	return _Rewarder.Contract.Owner(&_Rewarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewarder *RewarderCallerSession) Owner() (common.Address, error) {
	return _Rewarder.Contract.Owner(&_Rewarder.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderSession) Token() (common.Address, error) {
	return _Rewarder.Contract.Token(&_Rewarder.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewarder *RewarderCallerSession) Token() (common.Address, error) {
	return _Rewarder.Contract.Token(&_Rewarder.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderCaller) TotalClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "totalClaimed")
	return *ret0, err
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderSession) TotalClaimed() (*big.Int, error) {
	return _Rewarder.Contract.TotalClaimed(&_Rewarder.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xd54ad2a1.
//
// Solidity: function totalClaimed() view returns(uint256)
func (_Rewarder *RewarderCallerSession) TotalClaimed() (*big.Int, error) {
	return _Rewarder.Contract.TotalClaimed(&_Rewarder.CallOpts)
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderCaller) TotalPayoutsFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rewarder.contract.Call(opts, out, "totalPayoutsFor", arg0)
	return *ret0, err
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderSession) TotalPayoutsFor(arg0 common.Address) (*big.Int, error) {
	return _Rewarder.Contract.TotalPayoutsFor(&_Rewarder.CallOpts, arg0)
}

// TotalPayoutsFor is a free data retrieval call binding the contract method 0x4d676fa7.
//
// Solidity: function totalPayoutsFor(address ) view returns(uint256)
func (_Rewarder *RewarderCallerSession) TotalPayoutsFor(arg0 common.Address) (*big.Int, error) {
	return _Rewarder.Contract.TotalPayoutsFor(&_Rewarder.CallOpts, arg0)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderTransactor) Airdrop(opts *bind.TransactOpts, _beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "airdrop", _beneficiaries, _totalEarnings)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderSession) Airdrop(_beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.Airdrop(&_Rewarder.TransactOpts, _beneficiaries, _totalEarnings)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] _beneficiaries, uint256[] _totalEarnings) returns()
func (_Rewarder *RewarderTransactorSession) Airdrop(_beneficiaries []common.Address, _totalEarnings []*big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.Airdrop(&_Rewarder.TransactOpts, _beneficiaries, _totalEarnings)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderTransactor) Claim(opts *bind.TransactOpts, _recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "claim", _recipient, _totalEarned, _blockNumber, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderSession) Claim(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.Contract.Claim(&_Rewarder.TransactOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x172bd6de.
//
// Solidity: function claim(address _recipient, uint256 _totalEarned, uint256 _blockNumber, bytes32[] _proof) returns()
func (_Rewarder *RewarderTransactorSession) Claim(_recipient common.Address, _totalEarned *big.Int, _blockNumber *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Rewarder.Contract.Claim(&_Rewarder.TransactOpts, _recipient, _totalEarned, _blockNumber, _proof)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderTransactor) RecoverTokens(opts *bind.TransactOpts, _erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "recoverTokens", _erc20, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderSession) RecoverTokens(_erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.RecoverTokens(&_Rewarder.TransactOpts, _erc20, _to)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x056097ac.
//
// Solidity: function recoverTokens(address _erc20, address _to) returns()
func (_Rewarder *RewarderTransactorSession) RecoverTokens(_erc20 common.Address, _to common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.RecoverTokens(&_Rewarder.TransactOpts, _erc20, _to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceOwnership(&_Rewarder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewarder *RewarderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceOwnership(&_Rewarder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.TransferOwnership(&_Rewarder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewarder *RewarderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.TransferOwnership(&_Rewarder.TransactOpts, newOwner)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderTransactor) UpdateRoot(opts *bind.TransactOpts, _claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "updateRoot", _claimRoot, _blockNumber, _totalReward)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderSession) UpdateRoot(_claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.UpdateRoot(&_Rewarder.TransactOpts, _claimRoot, _blockNumber, _totalReward)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0x12b6af24.
//
// Solidity: function updateRoot(bytes32 _claimRoot, uint256 _blockNumber, uint256 _totalReward) returns()
func (_Rewarder *RewarderTransactorSession) UpdateRoot(_claimRoot [32]byte, _blockNumber *big.Int, _totalReward *big.Int) (*types.Transaction, error) {
	return _Rewarder.Contract.UpdateRoot(&_Rewarder.TransactOpts, _claimRoot, _blockNumber, _totalReward)
}

// RewarderAirdropIterator is returned from FilterAirdrop and is used to iterate over the raw logs and unpacked data for Airdrop events raised by the Rewarder contract.
type RewarderAirdropIterator struct {
	Event *RewarderAirdrop // Event containing the contract specifics and raw log

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
func (it *RewarderAirdropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderAirdrop)
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
		it.Event = new(RewarderAirdrop)
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
func (it *RewarderAirdropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderAirdropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderAirdrop represents a Airdrop event raised by the Rewarder contract.
type RewarderAirdrop struct {
	TotalDropped *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAirdrop is a free log retrieval operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) FilterAirdrop(opts *bind.FilterOpts) (*RewarderAirdropIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "Airdrop")
	if err != nil {
		return nil, err
	}
	return &RewarderAirdropIterator{contract: _Rewarder.contract, event: "Airdrop", logs: logs, sub: sub}, nil
}

// WatchAirdrop is a free log subscription operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) WatchAirdrop(opts *bind.WatchOpts, sink chan<- *RewarderAirdrop) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "Airdrop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderAirdrop)
				if err := _Rewarder.contract.UnpackLog(event, "Airdrop", log); err != nil {
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

// ParseAirdrop is a log parse operation binding the contract event 0xd0ecdf4854f39daba34ba0e2c1ed0132a7023a5434bffc90b24f3335fb90e5e3.
//
// Solidity: event Airdrop(uint256 totalDropped)
func (_Rewarder *RewarderFilterer) ParseAirdrop(log types.Log) (*RewarderAirdrop, error) {
	event := new(RewarderAirdrop)
	if err := _Rewarder.contract.UnpackLog(event, "Airdrop", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderClaimedChangedIterator is returned from FilterClaimedChanged and is used to iterate over the raw logs and unpacked data for ClaimedChanged events raised by the Rewarder contract.
type RewarderClaimedChangedIterator struct {
	Event *RewarderClaimedChanged // Event containing the contract specifics and raw log

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
func (it *RewarderClaimedChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderClaimedChanged)
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
		it.Event = new(RewarderClaimedChanged)
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
func (it *RewarderClaimedChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderClaimedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderClaimedChanged represents a ClaimedChanged event raised by the Rewarder contract.
type RewarderClaimedChanged struct {
	TotalUnclaimed *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimedChanged is a free log retrieval operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) FilterClaimedChanged(opts *bind.FilterOpts) (*RewarderClaimedChangedIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "ClaimedChanged")
	if err != nil {
		return nil, err
	}
	return &RewarderClaimedChangedIterator{contract: _Rewarder.contract, event: "ClaimedChanged", logs: logs, sub: sub}, nil
}

// WatchClaimedChanged is a free log subscription operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) WatchClaimedChanged(opts *bind.WatchOpts, sink chan<- *RewarderClaimedChanged) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "ClaimedChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderClaimedChanged)
				if err := _Rewarder.contract.UnpackLog(event, "ClaimedChanged", log); err != nil {
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

// ParseClaimedChanged is a log parse operation binding the contract event 0xc58f861f50ca7e2541250d781773a4614f235d08e72ecc0d2f25028b70a1b665.
//
// Solidity: event ClaimedChanged(uint256 totalUnclaimed)
func (_Rewarder *RewarderFilterer) ParseClaimedChanged(log types.Log) (*RewarderClaimedChanged, error) {
	event := new(RewarderClaimedChanged)
	if err := _Rewarder.contract.UnpackLog(event, "ClaimedChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewarder contract.
type RewarderOwnershipTransferredIterator struct {
	Event *RewarderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewarderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderOwnershipTransferred)
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
		it.Event = new(RewarderOwnershipTransferred)
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
func (it *RewarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderOwnershipTransferred represents a OwnershipTransferred event raised by the Rewarder contract.
type RewarderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewarder *RewarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewarderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewarderOwnershipTransferredIterator{contract: _Rewarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewarder *RewarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewarderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderOwnershipTransferred)
				if err := _Rewarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rewarder *RewarderFilterer) ParseOwnershipTransferred(log types.Log) (*RewarderOwnershipTransferred, error) {
	event := new(RewarderOwnershipTransferred)
	if err := _Rewarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewarderRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the Rewarder contract.
type RewarderRootUpdatedIterator struct {
	Event *RewarderRootUpdated // Event containing the contract specifics and raw log

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
func (it *RewarderRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderRootUpdated)
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
		it.Event = new(RewarderRootUpdated)
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
func (it *RewarderRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderRootUpdated represents a RootUpdated event raised by the Rewarder contract.
type RewarderRootUpdated struct {
	Root        [32]byte
	BlockNumber *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*RewarderRootUpdatedIterator, error) {

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &RewarderRootUpdatedIterator{contract: _Rewarder.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *RewarderRootUpdated) (event.Subscription, error) {

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderRootUpdated)
				if err := _Rewarder.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// ParseRootUpdated is a log parse operation binding the contract event 0x156798a72d63b37f86ff1ecc41eec4f30e3c7b325c8a410c2671f2e7fc0c30f3.
//
// Solidity: event RootUpdated(bytes32 root, uint256 blockNumber, uint256 _totalAmount)
func (_Rewarder *RewarderFilterer) ParseRootUpdated(log types.Log) (*RewarderRootUpdated, error) {
	event := new(RewarderRootUpdated)
	if err := _Rewarder.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
