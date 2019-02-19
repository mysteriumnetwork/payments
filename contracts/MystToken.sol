pragma solidity >=0.4.25 <0.6.0;

import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { ERC777ERC20BaseToken } from "./erc777/ERC777ERC20BaseToken.sol";

//The purpose of this contract is to have sample of potential updated "MYST" into ERC777
//for testing purposes of easier than ERC777 interactions. If this will work well in testnet,
//we can think of migration of mainnet token into ERC777.

contract MystToken is ERC777ERC20BaseToken, Ownable {
	event ERC20Enabled();
    event ERC20Disabled();

	address private mBurnOperator;

	// NOTE: IdentityRegistry should be set as default operator.
	constructor(
        string _name,
        string _symbol,
        uint256 _granularity,
        address[] _defaultOperators,
        address _burnOperator
    ) public ERC777ERC20BaseToken(_name, _symbol, _granularity, _defaultOperators) {
        mBurnOperator = _burnOperator;
    }

    function disableERC20() public onlyOwner {
        mErc20compatible = false;
        setInterfaceImplementation("ERC20Token", 0x0);
        emit ERC20Disabled();
    }

    function enableERC20() public onlyOwner {
        mErc20compatible = true;
        setInterfaceImplementation("ERC20Token", this);
        emit ERC20Enabled();
    }

    function mint(address _tokenHolder, uint256 _amount, bytes _operatorData) public onlyOwner {
        requireMultiple(_amount);
        mTotalSupply = mTotalSupply.add(_amount);
        mBalances[_tokenHolder] = mBalances[_tokenHolder].add(_amount);

        callRecipient(msg.sender, 0x0, _tokenHolder, _amount, "", _operatorData, true);

        emit Minted(msg.sender, _tokenHolder, _amount, _operatorData);
        if (mErc20compatible) { emit Transfer(0x0, _tokenHolder, _amount); }
    }

    function burn(uint256 _amount, bytes _data) public onlyOwner {
        super.burn(_amount, _data);
    }

    function operatorBurn(address _tokenHolder, uint256 _amount, bytes _data, bytes _operatorData) public {
        require(msg.sender == mBurnOperator, "Not a burn operator");
        super.operatorBurn(_tokenHolder, _amount, _data, _operatorData);
    }
}
