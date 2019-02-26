pragma solidity ^0.5.0;

import { Ownable } from "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import { ERC777ERC20BaseToken } from "./erc777/ERC777ERC20BaseToken.sol";

//The purpose of this contract is to have sample of potential updated "MYST" into ERC777
//for testing purposes of easier than ERC777 interactions. If this will work well in testnet,
//we can think of migration of mainnet token into ERC777.

contract MystToken is ERC777ERC20BaseToken, Ownable {
	address private mBurnOperator;
	
	// NOTE: IdentityRegistry should be set as default operator.
	constructor() public ERC777ERC20BaseToken() {
    	mName = "Myst Credit";
        mSymbol = "MYSTC";
        mGranularity = 1;
        mTotalSupply = 0;
    }

    function mint(address _tokenHolder, uint256 _amount, bytes memory _operatorData) public onlyOwner {
        requireMultiple(_amount);
        mTotalSupply = mTotalSupply.add(_amount);
        mBalances[_tokenHolder] = mBalances[_tokenHolder].add(_amount);

        callRecipient(msg.sender, address(0x0), _tokenHolder, _amount, "", _operatorData, true);

        emit Minted(msg.sender, _tokenHolder, _amount, _operatorData);
    }

    function burn(uint256 _amount, bytes memory _data) public onlyOwner {
        super.burn(_amount, _data);
    }

    function operatorBurn(address _tokenHolder, uint256 _amount, bytes memory _data, bytes memory _operatorData) public {
        require(msg.sender == mBurnOperator, "Not a burn operator");
        super.operatorBurn(_tokenHolder, _amount, _data, _operatorData);
    }
}
