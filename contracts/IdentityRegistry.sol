pragma solidity >=0.4.25 <0.6.0;

import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { ECRecovery } from "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import { SafeMath } from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import { ERC777Token } from "./erc777/ERC777Token.sol";
import { ERC820aClient } from "erc820a/contracts/ERC820aClient.sol";

contract Topup is ERC820aClient {
    IdentityRegistry registry;
    address identity;
    ERC777Token token;

    constructor(address _identity, address _token) {
        registry = IdentityRegistry(msg.sender);
        identity = _identity;
        token = ERC777Token(_token);

        setInterfaceImplementation("ERC777TokensRecipient", this);      
    }

    // ERC777 tokens receiver callback
    function tokensReceived(address _operator, address _from, address _to, uint256 _amount, bytes _data, bytes _operatorData) public {
        require(msg.sender == address(token)); // execute only if touched from MYST token
        require(registry.topUp(identity, _amount));
    }
}

contract IdentityRegistry is Ownable {
    using ECRecovery for bytes;
    using SafeMath for uint256;

    string constant REGISTER_PREFIX="Register prefix:";
    string constant WITHDRAW_PREFIX = "Withdraw request:";
    string constant ISSUER_PREFIX = "Issuer prefix:";

    ERC777Token public token;

    event Registered(address indexed identity);
    event ToppedUp(address indexed identity, uint256 amount, uint256 totalBalance);
    event Withdrawn(address indexed identity, uint256 amount, uint256 totalBalance);
    event PromiseClaimed(address indexed sender, address indexed receiver, uint256 amount, uint256 totalClaimed);

    modifier registered(address identity) {
        require(isRegistered(identity));
        _;
    }

    struct PublicKey {
        bytes32 part1;
        bytes32 part2;
    }

    struct Identity {
        PublicKey publicKey;
        uint256 balance;
        address topupAddress;
        uint256 widthdrawalNonce;
    }

    mapping(address => Identity) public registeredIdentities;
    mapping(address => mapping(address => uint256)) public claimedPromiseAmount;

    uint256 public registrationFee;
    uint256 public collectedFee;

    constructor(address tokenAddress, uint256 regFee) public {
        registrationFee = regFee;
        token = ERC777Token(tokenAddress);
    }

    function registerIdentity(bytes32 _pubKeyPart1, bytes32 _pubKeyPart2, bytes _signature) public returns (bool) {
        address _identity = recoverIdentity(_pubKeyPart1, _pubKeyPart2, _signature);
        address _identityFromPubKey = address(uint(keccak256(abi.encodePacked(_pubKeyPart1, _pubKeyPart2))));

        require(_identity != address(0));
        require(_identityFromPubKey == _identity);
        require(!isRegistered(_identity));

        collectedFee = collectedFee + registrationFee;

        registeredIdentities[_identity].publicKey = PublicKey({ part1: _pubKeyPart1, part2: _pubKeyPart2});
        token.operatorSend(msg.sender, address(this), registrationFee, "");

        // Deploy simple smart contract for user friendly topups.
        registeredIdentities[_identity].topupAddress = address(new Topup(_identity));

        emit Registered(_identity);
        return true;
    }

    // QUESTION: maybe claim and widthdrawal is same action? We could potentially not deposit tokens into this 
    // smart contract, but store them directry in identity. Then we would just do operatorSend()
    // instead of updating balances.
    function claimPromise(address _issuer, address _receiver, bytes32 _extraDataHash, uint256 _amount, bytes _senderSignature, bytes _receiverSigner) public returns (bool) {
        // Check if signature is valid, remember last running sum
        address _identity = keccak256(abi.encodePacked(ISSUER_PREFIX, _issuer, _receiver, _amount)).toEthSignedMessageHash.recover(_signature);
        require(_identity == _issuer && _identity != address(0));

        // Calculate amount of tokens to be claimed
        uint256 _unclaimedAmount = _amount.sub(claimedPromiseAmount[_receiver][_issuer]);
        require(_unclaimedAmount > 0);

        // If signer has less tokens than asked to transfer, we can transfer as much as he has already
        // and rest tokens can be transferred via same cheque but in another tx 
        // when signer will top up his balance.
        if (_unclaimedAmount > balances[_identity]) {
            _unclaimedAmount = balances[_identity];
        }

        // Increase already paid amount
        uint256 _totalClaimedAmount = claimedPromiseAmount[_receiver][_issuer].add(_unclaimedAmount);
        claimedPromiseAmount[_receiver][_issuer] = _totalClaimedAmount;

        // Internal token transfer
        balances[_issuer] = balances[_issuer].sub(_unclaimedAmount);
        balances[_receiver] = balances[_receiver].add(_unclaimedAmount);

        emit PromiseClaimed(_issuer, _receiver, _unclaimedAmount, _totalClaimedAmount);

        return true;
    }

    function topUp(address _identity, uint256 _amount) public returns (bool) {
        require(_amount > 0);

        balances[_identity] += _amount;
        require(token.operatorSend(msg.sender, address(this), amount));

        emit ToppedUp(identity, amount, balances[identity]);
        return true;
    }

    // TODO add timelock to delay withdrawals for 24h
    function withdraw(address _beneficiary, uint256 _amount, uint256 _nonce, bytes _signature) public returns (bool) {
        // Extract identity which signed this request
        address _identity = keccak256(abi.encodePacked(WITHDRAW_PREFIX, _beneficiary, _amount)).toEthSignedMessageHash.recover(_signature);
        require(_identity != address(0));
        
        // Send tokens to benefciary
        require(_beneficiary != address(0));
        require(_amount > 0);
        require(balances[_identity] >= _amount);

        balances[_identity] -= _amount;
        token.send(_beneficiary, _amount, "");

        // Mark this withdraw request as used
        require(_nonce > withdrawalNonces[_identity]);
        withdrawalNonces[_identity] = _nonce;

        emit Withdrawn(_identity, _amount, balances[_identity]);
        return true;
    }

    function recoverIdentity(bytes32 _pubKeyPart1, bytes32 _pubKeyPart2, bytes _signature) private pure returns (address) {
        return keccak256(abi.encodePacked(REGISTER_PREFIX, _pubKeyPart1, _pubKeyPart2)).toEthSignedMessageHash().recover(_signature);
    }

    function isRegistered(address _identity) public view returns (bool) {
        return registeredIdentities(_identity).topupAddress != address(0x0);
    }

    function getPublicKey(address _identity) public view returns (bytes32, bytes32) {
        PublicKey storage _pubKey = registeredIdentities[_identity];
        return (_pubKey.part1, _pubKey.part2);
    }

    function changeRegistrationFee(uint256 _newFee) public onlyOwner {
        registrationFee = _newFee;
    }

    function transferCollectedFeeTo(address _beneficiary) public onlyOwner returns (bool) {
        require(collectedFee > 0);
        uint256 _transferAmount = collectedFee;
        collectedFee = 0;
        require(token.send(_beneficiary, _transferAmount, ""));
        return true;
    }
}
