# LaksaGo

Zilliqa Blockchain  Library

## Quick Start

This repo is still under active development, please do not use in production environment!

More docs can be found in https://apidocs.zilliqa.com/


## Supports

### Account API

- [x] fromFile
- [x] toFile

### Wallet API

- [x] createAccount
- [x] addByPrivateKey addByKeyStore
- [x] remove
- [x] setDefault
- [x] signTransaction (default account)
- [x] signTransactionWith (specific account)

### TransactionFactory Transaction

- [x] sendTransaction
- [x] trackTx
- [x] confirm
- [x] isPending isInitialised isConfirmed isRejected

### ContractFactory Contract

- [x] deploy
- [x] call
- [x] isInitialised isDeployed isRejected
- [x] getState
- [x] getAddressForContract


### Crypto API

- [x] getDerivedKey (PBKDF2 and Scrypt)
- [x] generatePrivateKey
- [x] Schnorr.sign
- [x] Schnorr.verify
- [x] getPublicKeyFromPrivateKey
- [x] getAddressFromPublicKey
- [x] getAddressFromPrivateKey
- [x] encryptPrivateKey
- [x] decryptPrivateKey

### JSON-RPC API

#### Blockchain-related methods

- [x] getNetworkId
- [x] getBlockchainInfo
- [x] getShardingStructure
- [x] getDsBlock
- [x] getLatestDsBlock
- [x] getNumDSBlocks
- [x] getDSBlockRate
- [x] getDSBlockListing
- [x] getTxBlock
- [x] getLatestTxBlock
- [x] getNumTxBlocks
- [x] getTxBlockRate
- [x] getTxBlockListing
- [x] getNumTransactions
- [x] getTransactionRate
- [x] getCurrentMiniEpoch
- [x] getCurrentDSEpoch
- [x] getPrevDifficulty
- [x] getPrevDSDifficulty

#### Transaction-related methods

- [ ] createTransaction
- [x] getTransaction
- [x] getRecentTransactions
- [x] getTransactionsForTxBlock
- [x] getNumTxnsTxEpoch
- [x] getNumTxnsDSEpoch
- [x] getMinimumGasPrice

#### Contract-related methods

- [x] getSmartContractCode
- [x] getSmartContractInit
- [x] getSmartContractState
- [x] getSmartContracts
- [x] getContractAddressFromTransactionID

#### Account-related methods

- [x] getBalance

### Validation

- [x] isAddress
- [x] isPublicjKey
- [x] isPrivateKey
- [x] isSignature

### Util

- [x] byteArrayToHexString
- [x] hexStringToByteArray
- [x] generateMac
- [x] isByteString
- [x] encodeTransactionProto
- [x] toChecksumAddress
- [x] isValidChecksumAddress