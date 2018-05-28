pragma solidity ^0.4.23;


library Utils {

    function unpackSignatureAndSigns(bytes32 inputParam) internal pure returns(address addr, uint8 v1, uint8 v2) {
        addr = address(uint256(inputParam)/256**12);
        v1 = uint8(inputParam[20]);
        v2 = uint8(inputParam[21]);
        return;
    }
}