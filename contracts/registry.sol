pragma solidity ^0.4.23;

import "./deps/OpenZeppelin/contracts/ECRecovery.sol";

contract IdentityRegistry {
    string constant REGISTER_PREFIX="Register prefix:";

    event Registered(address indexed identity);

    modifier registered(address identity) {
        require(registeredIdentities[identity]);
        _;
    }

    mapping(address => bool)registeredIdentities;

    function RegisterIdentity(uint64 randomNumber, bytes signature) public {
        address identity = ECRecovery.recover(keccak256(REGISTER_PREFIX, randomNumber), signature);
        registeredIdentities[identity] = true;
        emit Registered(identity);
    }

    function isRegistered(address identity) public constant returns (bool) {
        return registeredIdentities[identity];
    }
}