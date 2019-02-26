pragma solidity ^0.5.0;

import { ECDSA } from "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import { SafeMath } from "openzeppelin-solidity/contracts/math/SafeMath.sol";
import { IERC20 } from "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";

import { ERC820aClient } from "./erc820a/ERC820aClient.sol";

// NOTE: it can be implemented as pure ERC20 token as well. It will not have ToppedUp event however.
contract IdentityImplementation is ERC820aClient {
    using ECDSA for bytes32;
    using SafeMath for uint256;

    IERC20 public token;  // NODE: token can be actually constant or be received from config
    address public identityHash;

    string constant WITHDRAW_PREFIX = "Withdraw request:";
    string constant ISSUER_PREFIX = "Issuer prefix:";

    mapping(address => uint256) public paidAmounts;

    event ToppedUp(address indexed identityHash, uint256 amount, uint256 totalBalance);
    event Withdrawn(address indexed receiver, uint256 amount, uint256 totalBalance);
    event PromiseSettled(address indexed receiver, uint256 amount, uint256 totalSettled);

    // Because of proxy pattern this function is used insted of constructor.
    // Have to be called right after proxy deployment.
    function initialize(address _token, address _identityHash) public {
        require(!isInitialized(), "Is already initialized");
        require(_identityHash != address(0));
        require(_token != address(0));

        token = IERC20(_token);
        identityHash = _identityHash;

        // Register as ERC777 recipient in ERC820 registry
        setInterfaceImplementation("ERC777TokensRecipient", address(this));
    }

    function isInitialized() public view returns (bool) {
        return identityHash != address(0);
    }

    // Settle promise. Withdrawal can be also be done using same function. 
    function settlePromise(address _to, uint256 _amount, uint256 _fee, bytes32 _extraDataHash, bytes memory _signature) public {
        uint256 _paidAmount = sendByCheque(_to, _amount, _fee, _extraDataHash, _signature);
        emit PromiseSettled(_to, _paidAmount, paidAmounts[_to]);
    }

    function withdraw(address _to, uint256 _amount, bytes memory _signature) public {
        uint256 _paidAmount = sendByCheque(_to, _amount, 0, "", _signature);
        emit Withdrawn(_to, _paidAmount, token.balanceOf(address(this)));
    }

    function sendByCheque(address _to, uint256 _amount, uint256 _fee, bytes32 _extraDataHash, bytes memory _signature) internal returns (uint256) {
        require(_to != address(0));

        address _signer = keccak256(abi.encodePacked(ISSUER_PREFIX, _to, _amount, _fee, _extraDataHash)).recover(_signature);
        require(_signer == identityHash, "Have to be signed by proper identity");

        // Calculate amount of tokens to be claimed.
        uint256 _unpaidAmount = _amount.sub(paidAmounts[_to]).sub(_fee);
        require(_unpaidAmount > 0, "Amount to settle is less that zero");

        // If signer has less tokens than asked to transfer, we can transfer as much as he has already
        // and rest tokens can be transferred via same cheque but in another tx 
        // when signer will top up his balance.
        if (_unpaidAmount > token.balanceOf(address(this))) {
            _unpaidAmount = token.balanceOf(address(this)).sub(_fee);
        }

        // Increase already paid amount
        uint256 _totalPaidAmount = paidAmounts[_to].add(_unpaidAmount);
        paidAmounts[_to] = _totalPaidAmount;

        // Send tokens
        token.transfer(_to, _amount);

        if (_fee > 0) {
            token.transfer(msg.sender, _fee);
        }

        return _unpaidAmount;
    }

    // ERC777 tokens recipient callback
    function tokensReceived(address _operator, address _from, address _to, uint256 _amount, bytes memory _data, bytes memory _operatorData) public {
        emit ToppedUp(identityHash, _amount, token.balanceOf(address(this)));
    }
}
