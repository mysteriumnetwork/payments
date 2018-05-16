pragma solidity ^0.4.23;


contract IdentityRegistry {

    event Registered(address indexed identity);

    modifier registered(address identity) {
        require(registeredIdentities[identity]);
        _;
    }

    mapping(address => bool)registeredIdentities;

    function RegisterIdentity(address identity) public {
        registeredIdentities[identity] = true;
        emit Registered(identity);
    }

    function isRegistered(address identity) public constant returns (bool) {
        return registeredIdentities[identity];
    }
}