pragma solidity ^0.4.23;

import "./IdentityRegistry.sol";

contract IdentityPromises is IdentityRegistry {
    string constant ISSUER_PREFIX = "Issuer prefix:";
    string constant RECEIVER_PREFIX = "Receiver prefix:";

    event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount);

    mapping(address => mapping(address => uint256)) public clearedPromises;

    constructor(address erc20address, uint256 registrationFee) public IdentityRegistry(erc20address, registrationFee) {

    }

    function clearPromise(address receiver, uint256 seq, uint256 amount,
        uint8 sender_V , bytes32 sender_R, bytes32 sender_S,
        uint8 receiver_V, bytes32 receiver_R, bytes32 receiver_S) public returns (bool) {

        bytes32 promiseHash = keccak256(ISSUER_PREFIX, receiver , seq , amount);

        address sender = ecrecover(promiseHash, sender_V , sender_R, sender_S);
        address recoveredReceiver = ecrecover(keccak256(RECEIVER_PREFIX, promiseHash , sender), receiver_V , receiver_R, receiver_S);

        require(sender > 0);
        require(recoveredReceiver > 0);
        require(recoveredReceiver == receiver);
        require(amount > 0);
        require(seq > clearedPromises[sender][receiver]);

        clearedPromises[sender][receiver] = seq;

        emit PromiseCleared(sender, receiver, seq, amount);
        return true;
    }
}
