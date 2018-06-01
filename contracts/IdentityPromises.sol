pragma solidity ^0.4.23;

import "./IdentityRegistry.sol";
import "./IdentityBalances.sol";
import "./Utils.sol";

contract IdentityPromises is IdentityRegistry , IdentityBalances {
    string constant ISSUER_PREFIX = "Issuer prefix:";
    string constant RECEIVER_PREFIX = "Receiver prefix:";

    event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount);

    mapping(address => mapping(address => uint256)) public clearedPromises;

    constructor(address erc20address, uint256 registrationFee) public IdentityRegistry(erc20address, registrationFee) IdentityBalances(erc20address) {

    }

    function clearPromise(bytes32 receiverAndSigns, bytes32 extraDataHash, uint256 seq, uint256 amount,
        bytes32 sender_R, bytes32 sender_S,
        bytes32 receiver_R, bytes32 receiver_S) public returns (bool) {
        (address receiver, uint8 sender_V, uint8 receiver_V) = Utils.unpackSignatureAndSigns(receiverAndSigns);
        bytes32 promiseHash = keccak256(abi.encodePacked(ISSUER_PREFIX, extraDataHash, receiver, seq , amount));

        address sender = ecrecover(promiseHash, sender_V , sender_R, sender_S);
        address recoveredReceiver = ecrecover(keccak256(abi.encodePacked(RECEIVER_PREFIX, promiseHash , sender)), receiver_V , receiver_R, receiver_S);

        require(sender > 0);
        require(recoveredReceiver > 0);
        require(recoveredReceiver == receiver);
        require(isRegistered(sender));
        require(isRegistered(receiver));
        require(amount > 0);
        return internalClearPromise(receiver , sender , seq, amount);
    }

    function internalClearPromise(address receiver, address sender, uint256 seq, uint256 amount) private returns (bool) {
        require(seq > clearedPromises[sender][receiver]);

        clearedPromises[sender][receiver] = seq;
        uint256 transfered = internalTransfer(sender, receiver, amount);
        require(transfered == amount);

        emit PromiseCleared(sender, receiver, seq, amount);
        return true;
    }
}
