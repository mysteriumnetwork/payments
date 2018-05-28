pragma solidity ^0.4.23;

import "./Utils.sol";


contract TestUtilsContract {

    function GiveMeData(bytes32 param) public constant returns( address, uint8 , uint8) {
        return Utils.unpackSignatureAndSigns(param);
    }
}
