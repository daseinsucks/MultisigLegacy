const MultisigFactory = artifacts.require("MultiSigWalletFactory.sol");
const Factory = artifacts.require("Factory.sol");
const Multisig = artifacts.require("MultiSigWallet.sol");
//const Multisig = artifacts.require("MultiSigWallet.sol");

module.exports = function(deployer) {
  deployer.deploy(Multisig, ["0xc905803BbC804fECDc36850281fEd6520A346AC5"], 1);
  //deployer.deploy(Multisig);
};