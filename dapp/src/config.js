var txDefaultOrig =
{
  websites: {
    "wallet": "https://wallet.gnosis.pm",
    "gnosis": "https://gnosis.pm",
    "ethGasStation": "https://safe-relay.gnosis.pm/api/v1/gas-station/"
  },
  resources : {
    "termsOfUse": "https://wallet.gnosis.pm/TermsofUseMultisig.pdf",
    "privacyPolicy": "https://gnosis.io/privacy-policy",
    "imprint": "https://wallet.gnosis.pm/imprint.html"
  },
  gasLimit: 3141592,
  gasPrice: 180000000000,
  ethereumNode: "https://mainnet.infura.io:443",
  connectionChecker: {
    method : "OPTIONS",
    url : "https://www.google.com",
    checkInterval: 5000
  },
  accountsChecker: {
    checkInterval: 5000
  },
  transactionChecker: {
    checkInterval: 15000
  },
  wallet: "injected",
  defaultChainID: null,
  // Mainnet
  //modified factory address 25.10.2022
  walletFactoryAddress: "0x86879e6E3D0cE01a61B7416f984B1d537CaF76ad",
  tokens: [
    {
      'address': '0xdAC17F958D2ee523a2206206994597C13D831ec7',
      'name': 'Tether USD',
      'symbol': 'USDT',
      'decimals': 18
    },
    {
      'address': '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      'name': 'USD Coin',
      'symbol': 'USDC',
      'decimals': 18
    },
    {
      'address': '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      'name': 'DAI Stablecoin',
      'symbol': 'DAI',
      'decimals': 18
    },
    {
      'address': '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599',
      'name': 'Wrapped Bitcoin',
      'symbol': '',
      'decimals': 18
    },
  ]
};

if (isElectron) {
  txDefaultOrig.wallet = "remotenode";
}

var txDefault = {
  ethereumNodes : [
    {
      url : "https://mainnet.infura.io:443",
      name: "Remote Mainnet"
    },
    {
      url : "https://ropsten.infura.io:443",
      name: "Remote Ropsten"
    },
    {
      url : "https://kovan.infura.io:443",
      name: "Remote Kovan"
    },
    {
      url : "https://rinkeby.infura.io:443",
      name: "Remote Rinkeby"
    },
    {
      url : "http://localhost:8545",
      name: "Local node"
    }
  ],
  walletFactoryAddresses: {
    'mainnet': {
      name: 'Mainnet',
      address: txDefaultOrig.walletFactoryAddress
    },
    'ropsten': {
      name: 'Ropsten',
      address: '0x5cb85db3e237cac78cbb3fd63e84488cac5bd3dd'
    },
    'kovan': {
      name: 'Kovan',
      address: '0x2c992817e0152a65937527b774c7a99a84603045'
    },
    'rinkeby': {
      name: 'Rinkeby',
      address: '0x19ba60816abca236baa096105df09260a4791418'
    },
    'goerli': {
      name: 'Goerli',
      address: '0x86879e6E3D0cE01a61B7416f984B1d537CaF76ad'
    }
  }
};

var oldWalletFactoryAddresses = [
  ("0x86879e6E3D0cE01a61B7416f984B1d537CaF76ad").toLowerCase(),
  ("0x86879e6E3D0cE01a61B7416f984B1d537CaF76ad").toLowerCase(),
  ("0x86879e6E3D0cE01a61B7416f984B1d537CaF76ad").toLowerCase()
];

/**
* Update the default wallet factory address in local storage
*/
function checkWalletFactoryAddress() {
  var userConfig = JSON.parse(localStorage.getItem("userConfig"));

  if (userConfig && oldWalletFactoryAddresses.indexOf(userConfig.walletFactoryAddress.toLowerCase()) >= 0) {
    userConfig.walletFactoryAddress = txDefaultOrig.walletFactoryAddress;
    localStorage.setItem("userConfig", JSON.stringify(userConfig));
  }
}

/**
* Reload configuration
*/
function loadConfiguration () {
  var userConfig = JSON.parse(localStorage.getItem("userConfig"));
  Object.assign(txDefault, txDefaultOrig, userConfig);
}

checkWalletFactoryAddress();
loadConfiguration();
