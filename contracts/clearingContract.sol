pragma solidity ^0.4.23;

import "./registry.sol";
import "./deps/OpenZeppelin/contracts/ECRecovery.sol";

contract ClearingContract is IdentityRegistry {
    string constant promisePrefix = "Promise prefix:";

    event PromiseCleared(address indexed from, address indexed to, uint256 seqNo, uint256 amount);

    function clearPromise(uint seqNo, uint amount , bytes receiverSig, bytes payerSig) public returns(bool) {
        address receiver = ECRecovery.recover(keccak256(promisePrefix,"abc"), receiverSig);
        address payer = ECRecovery.recover(keccak256(promisePrefix,"abc"), payerSig);
        //require(super.isRegistered(receiver));
        //require(super.isRegistered(payer));
        emit PromiseCleared(payer, receiver, seqNo, amount);
        return true;
    }
}
