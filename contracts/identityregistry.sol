pragma solidity ^0.4.23;

contract IdentityRegistry {
    event Registered(address indexed identity);

    mapping(address => bool)registeredIdentities;



    function RegisterIdentity(address identity) public {
        registeredIdentities[identity] = true;
        emit Registered(identity);
    }

    function isRegistered(address identity) public constant returns (bool) {
        return registeredIdentities[identity];
    }
}