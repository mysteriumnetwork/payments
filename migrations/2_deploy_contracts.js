const ERC820aRegistry = require("erc820a");

const IdentityRegistry = artifacts.require("IdentityRegistry");
const MystToken = artifacts.require("MystToken");


module.exports = async function(deployer, network, accounts) {
    // Don't run migrations for tests
	if (network === 'development') {
		const Web3Js = require("web3")
        const web3 = new Web3Js('http://localhost:7545')
        const registry = await ERC820aRegistry.deploy(web3, accounts[9])

	  	const token = await deployer.deploy(MystToken);
	  	console.log('============', token.address, '============')
		const identityRegistry = await deployer.deploy(IdentityRegistry, token.address, 0);
	}
};
