pragma solidity ^0.4.23;

import "./deps/OpenZeppelin/contracts/ECRecovery.sol";
import "./deps/OpenZeppelin/contracts/token/ERC20/ERC20.sol";
import "./deps/OpenZeppelin/contracts/ownership/Ownable.sol";

contract IdentityRegistry is Ownable {
    string constant REGISTER_PREFIX="Register prefix:";

    event Registered(address indexed identity);

    modifier registered(address identity) {
        require(registeredIdentities[identity]);
        _;
    }

    mapping(address => bool)registeredIdentities;

    ERC20 public ERC20Token;

    uint256 public registrationFee;

    function IdentityRegistry(address tokenAddress , uint256 regFee) {
        ERC20Token =ERC20(tokenAddress);
        registrationFee = regFee;
    }

    function RegisterIdentity(uint64 randomNumber, bytes signature) public {
        address identity = ECRecovery.recover(keccak256(REGISTER_PREFIX, randomNumber), signature);
        registeredIdentities[identity] = true;
        emit Registered(identity);
    }

    function isRegistered(address identity) public constant returns (bool) {
        return registeredIdentities[identity];
    }

    function changeRegistrationFee(uint256 newFee) onlyOwner public {
        registrationFee = newFee;
    }
}