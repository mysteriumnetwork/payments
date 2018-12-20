pragma solidity ^0.4.23;

//The purpose of this contract is to have sample ERC20 a la "MYST" compatible mintable token
//for testing purposes of ERC20 interactions, actually we can use MintableToken directly,
//but for the sake of demonstrating that any compatible token will do
import "./deps/OpenZeppelin/contracts/token/ERC20/ERC20Mintable.sol";


contract MystToken is ERC20Mintable {

}
