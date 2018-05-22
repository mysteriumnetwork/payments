pragma solidity ^0.4.23;

import "./deps/OpenZeppelin/contracts/ownership/Ownable.sol";
import "./ERC20Aware.sol";

contract IdentityRegistry is Ownable, ERC20Aware {
    string constant REGISTER_PREFIX="Register prefix:";

    event Registered(address indexed identity);

    modifier registered(address identity) {
        require(isRegistered(identity));
        _;
    }

    struct PublicKey {
        bytes32 part1;
        bytes32 part2;
    }

    mapping(address => PublicKey)registeredIdentities;

    uint256 public registrationFee;

    constructor(address tokenAddress , uint256 regFee) public ERC20Aware(tokenAddress) {
        registrationFee = regFee;
    }

    function RegisterIdentity(bytes32 pubKeyPart1, bytes32 pubKeyPart2, uint8 v, bytes32 r, bytes32 s) public returns (bool) {
        address identityFromPubKey = address(keccak256(pubKeyPart1 , pubKeyPart2));
        address identity = ecrecover(keccak256(REGISTER_PREFIX, pubKeyPart1 , pubKeyPart2), v, r, s);

        require(identity > 0);
        require(identityFromPubKey == identity);
        require(!isRegistered(identity));

        registeredIdentities[identity] = PublicKey({ part1: pubKeyPart1 , part2: pubKeyPart2});
        require(ERC20Token.transferFrom(msg.sender , this, registrationFee));

        emit Registered(identity);
        return true;
    }

    function isRegistered(address identity) public constant returns (bool) {
        PublicKey storage pubKey = registeredIdentities[identity];
        return uint256(pubKey.part1) > 0 && uint256(pubKey.part2) > 0;
    }

    function getPublicKey(address identity) public constant returns (bytes32,bytes32) {
        PublicKey storage pubKey = registeredIdentities[identity];
        return (pubKey.part1 , pubKey.part2);
    }

    function changeRegistrationFee(uint256 newFee) public onlyOwner {
        registrationFee = newFee;
    }

    function transferCollectedFeeTo(address receiver) public onlyOwner returns (bool) {
        uint256 balance = ERC20Token.balanceOf(this);
        require(ERC20Token.transfer(receiver , balance));
        return true;
    }

}