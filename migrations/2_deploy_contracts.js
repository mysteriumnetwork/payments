const MystToken = artifacts.require("MystToken");

module.exports = function(deployer) {
  deployer.deploy(MystToken);
};
