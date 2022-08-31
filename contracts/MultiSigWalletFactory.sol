//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./Factory.sol";
import "./MultiSigWallet.sol";


/// @title Multisignature wallet factory - Allows creation of multisig wallet.
/// @author Stefan George - <stefan.george@consensys.net>
contract MultiSigWalletFactory is Factory {

    /*
     * Public functions
     */
    /// @dev Allows verified creation of multisignature wallet.
    /// @param _owners List of initial owners.
    /// @param _required Number of required confirmations.
    function create(address[] memory _owners, uint _required)
        public
        returns (address)
    {
        MultiSigWallet wallet = new MultiSigWallet(_owners, _required);
        address walletAddress = address(wallet);
        register(walletAddress);
        return address(wallet);
    }
}
