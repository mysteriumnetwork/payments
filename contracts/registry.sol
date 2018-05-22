pragma solidity ^0.4.23;

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

    constructor(address tokenAddress , uint256 regFee) public {
        ERC20Token = ERC20(tokenAddress);
        registrationFee = regFee;
    }

    function RegisterIdentity(bytes32 pubKeyPart1, bytes32 pubKeyPart2, uint8 v, bytes32 r, bytes32 s) public {
        address identityFromPubKey = address(keccak256(pubKeyPart1 , pubKeyPart2));
        address identity = ecrecover(keccak256(REGISTER_PREFIX, pubKeyPart1 , pubKeyPart2), v, r, s);
        registeredIdentities[identity] = true;
        require(ERC20Token.transferFrom(msg.sender , this, registrationFee));
        emit Registered(identity);
    }

    function isRegistered(address identity) public constant returns (bool) {
        return registeredIdentities[identity];
    }

    function changeRegistrationFee(uint256 newFee) public onlyOwner {
        registrationFee = newFee;
    }

    function transferCollectedFeeTo(address receiver) public onlyOwner {
        uint256 balance = ERC20Token.balanceOf(this);
        require(ERC20Token.transfer(receiver , balance));
    }

}