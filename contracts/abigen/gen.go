package abigen

//go:generate abigen --sol ../IdentityPromises.sol --pkg abigen --out payments.go -exc ../Utils.sol:Utils,../deps/OpenZeppelin/contracts/ownership/Ownable.sol:Ownable,../deps/OpenZeppelin/contracts/math/Math.sol:Math,../deps/OpenZeppelin/contracts/token/ERC20/ERC20Basic.sol:ERC20Basic,../ERC20Aware.sol:ERC20Aware,../deps/OpenZeppelin/contracts/token/ERC20/ERC20.sol:ERC20,../IdentityRegistry.sol:IdentityRegistry,../IdentityBalances.sol:IdentityBalances
