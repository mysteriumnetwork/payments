pragma solidity ^0.5.0;

import "./deps/OpenZeppelin/contracts/token/ERC20/ERC20.sol";


// With solidity 0.5.0, we're unable to specify base constructor arguments multiple times.
// Therefore, these are now split to two different contracts.
// From https://solidity.readthedocs.io/en/v0.5.0/050-breaking-changes.html
// `Specifying base constructor arguments multiple times in the same inheritance hierarchy is now disallowed.`

contract ERC20AwareRegistry {
    ERC20 public ERC20Token;

    constructor(address erc20tokenAddress) public {
        ERC20Token = ERC20(erc20tokenAddress);
    }
}

contract ERC20AwareBalance {
    ERC20 public ERC20Token;

    constructor(address erc20tokenAddress) public {
        ERC20Token = ERC20(erc20tokenAddress);
    }
}
