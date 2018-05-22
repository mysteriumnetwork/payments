pragma solidity ^0.4.23;

import "./IdentityRegistry.sol";
import "./deps/OpenZeppelin/contracts/ECRecovery.sol";

contract IdentityPromises is IdentityRegistry {
    string constant PROMISE_PREFIX = "Promise prefix:";

    event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount);

    constructor(address erc20address, uint256 registrationFee) public IdentityRegistry(erc20address, registrationFee) {

    }

    function clearPromise(uint seqNo, uint amount , bytes receiverSig, bytes payerSig) public returns (bool) {
        address receiver = ECRecovery.recover(keccak256(PROMISE_PREFIX,"abc"), receiverSig);
        address payer = ECRecovery.recover(keccak256(PROMISE_PREFIX,"abc"), payerSig);
        require(super.isRegistered(receiver));
        require(super.isRegistered(payer));
        emit PromiseCleared(payer, receiver, seqNo, amount);
        return true;
    }
}
