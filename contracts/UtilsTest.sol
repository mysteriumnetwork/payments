pragma solidity ^0.5.0;

import "./Utils.sol";


contract TestUtilsContract {

    function GiveMeData(bytes32 param) public view returns( address, uint8 , uint8) {
        return Utils.unpackSignatureAndSigns(param);
    }
}
