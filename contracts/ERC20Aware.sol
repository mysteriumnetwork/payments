pragma solidity ^0.4.23;

import "./deps/OpenZeppelin/contracts/token/ERC20/ERC20.sol";


contract ERC20Aware {
    ERC20 public ERC20Token;

    constructor(address erc20tokenAddress) public {
        ERC20Token = ERC20(erc20tokenAddress);
    }
}