pragma solidity ^0.4.23;

import "./ERC20Aware.sol";

contract IdentityBalances is ERC20Aware {
    string constant WITHDRAW_PREFIX = "Withdraw request:";

    event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance);
    event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance);

    mapping(address => uint256) balances;

    constructor(address tokenAddress) public ERC20Aware(tokenAddress) {

    }

    function topUp(address identity , uint256 amount) public returns (bool) {
        require(amount > 0);

        balances[identity]+=amount;
        require(ERC20Token.transferFrom(msg.sender, this, amount));

        emit ToppedUp(identity , amount , balances[identity]);
        return true;
    }

    function withdraw(uint256 amount, uint8 v, bytes32 r, bytes32 s) public returns (bool) {
        address identity = ecrecover(keccak256(WITHDRAW_PREFIX, amount) , v , r ,s );

        require(identity > 0);
        require(amount > 0);
        require(balances[identity] >= amount);

        balances[identity]-=amount;
        require(ERC20Token.transfer(msg.sender, amount));

        emit Withdrawn(identity, amount, balances[identity]);
        return true;
    }

    function balanceOf(address identity) public constant returns (uint256) {
        return balances[identity];
    }
}