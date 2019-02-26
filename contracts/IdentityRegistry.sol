pragma solidity ^0.5.0;

import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { ECDSA } from "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import { SafeMath } from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import { ERC820aClient } from "./erc820a/ERC820aClient.sol";
import { ERC777Token } from "./erc777/ERC777Token.sol";
import { IdentityImplementation } from "./IdentityImplementation.sol";


contract IdentityRegistry is Ownable {
    using ECDSA for bytes32;
    using SafeMath for uint256;

    string constant REGISTER_PREFIX="Register prefix:";
    uint256 public registrationFee;
    ERC777Token public token;
    address identityContractImplementation;

    struct PublicKey {
        bytes32 part1;
        bytes32 part2;
    }

    mapping(address => PublicKey) registeredIdentities;
    
    event Registered(address indexed identity);

    constructor(address _tokenAddress, uint256 _regFee, address _implementation) public {
        registrationFee = _regFee;

        require(_tokenAddress != address(0));
        token = ERC777Token(_tokenAddress);
        
        require(_implementation != address(0));
        identityContractImplementation = _implementation;
    }

    function registerIdentity(bytes32 _pubKeyPart1, bytes32 _pubKeyPart2, bytes memory _signature) public {
        address _identity = keccak256(abi.encodePacked(REGISTER_PREFIX, _pubKeyPart1, _pubKeyPart2)).recover(_signature);
        address _identityFromPubKey = address(uint(keccak256(abi.encodePacked(_pubKeyPart1, _pubKeyPart2))));

        require(_identity != address(0));
        require(_identityFromPubKey == _identity);
        require(!isRegistered(_identity));

        registeredIdentities[_identity] = PublicKey({ part1: _pubKeyPart1, part2: _pubKeyPart2});

        if (registrationFee > 0) {
            token.operatorSend(msg.sender, address(this), registrationFee, "", "");
        }

        // Deploy identity contract (mini proxy which is pointing to implementation)
        // require(deployMiniProxy(uint256(_identity)) == getIdentityContractAddress(_identity), "Wrong identity contract address");
        IdentityImplementation _identityContract = IdentityImplementation(deployMiniProxy(uint256(_identity)));
        _identityContract.initialize(address(token), _identity);

        emit Registered(_identity);
    }

    // Alternative implementation if we don't need to store pubKeys
    // function registerIdentity(address _identityHash) public {
    //     require(_identityHash != address(0));
    //     require(!isRegistered(_identityHash));
    //     if (registrationFee > 0) {
    //         token.operatorSend(msg.sender, address(this), registrationFee, "", "");
    //     }
    //     deployMiniProxy(uint256(_identityHash));
    //     emit Registered(_identityHash);
    // }

    function deployMiniProxy(uint256 _salt) internal returns (address) {
        address _addr; 
        bytes memory _code = getProxyCode();

        assembly {
            _addr := create2(0, add(_code, 0x20), mload(_code), _salt)
            if iszero(extcodesize(_addr)) {
                revert(0, 0)
            }
        }
        
        return _addr;
    }

    function getProxyCode() public returns (bytes memory) {
        bytes memory _code = "0x363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf3";
        
        // NOTE: in final implementation this can be avoided by hardcoding `identityContractImplementation`.
        // We're using it now for easier testing.
        bytes20 _targetBytes = bytes20(identityContractImplementation);
        for (uint i = 0; i < 20; i++) {
            _code[20 + i] = _targetBytes[i];
        }

        return _code;
    }

    function getIdentityContractAddress(address _identity) public returns (address) {
        return address(uint256(keccak256(abi.encodePacked(uint256(0xff), address(this), _identity, keccak256(getProxyCode())))));
    }

    function isRegistered(address _identity) public view returns (bool) {
        return registeredIdentities[_identity].part1 != 0x0;
    }

    function getPublicKey(address _identity) public view returns (bytes32, bytes32) {
        PublicKey storage _pubKey = registeredIdentities[_identity];
        return (_pubKey.part1, _pubKey.part2);
    }

    function changeRegistrationFee(uint256 _newFee) public onlyOwner {
        registrationFee = _newFee;
    }

    function transferCollectedFeeTo(address _beneficiary) public onlyOwner{
        uint256 _collectedFee = token.balanceOf(address(this));
        require(_collectedFee > 0);
        token.send(_beneficiary, _collectedFee, "");
    }
}
