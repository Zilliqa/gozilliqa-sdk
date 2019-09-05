### LaksaGo

The golang version of zilliqa blockchain library

#### Requirements

golang environment: 

* [download golang](https://golang.org/dl/)
* [installation instructions](https://golang.org/doc/install)

#### Installation

The sdk is using `go mod` to manager it's dependent libraries, so if you do want to hack the source code of this repository, make sure you have the minimal `golang` version that support `go mod` and enable it.

Install the dependent libraries:

```go
go get ./...
```

sdk it self cannot been built into a binary cause it does't have any `main` function, you can directly add it to your own project as a library. Also, we recommend that you can run the `golang uint test` or go through 
the section `quick start` first to get a basic understanding before you start to use this sdk.

#### Supports

##### Account API

- [x] fromFile
- [x] toFile

##### Wallet API

- [x] createAccount
- [x] addByPrivateKey addByKeyStore
- [x] remove
- [x] setDefault
- [x] signTransaction (default account)
- [x] signTransactionWith (specific account)

##### TransactionFactory Transaction

- [x] sendTransaction
- [x] trackTx
- [x] confirm
- [x] isPending isInitialised isConfirmed isRejected

##### ContractFactory Contract

- [x] deploy
- [x] call
- [x] isInitialised isDeployed isRejected
- [x] getState
- [x] getAddressForContract


##### Crypto API

- [x] getDerivedKey (PBKDF2 and Scrypt)
- [x] generatePrivateKey
- [x] Schnorr.sign
- [x] Schnorr.verify
- [x] getPublicKeyFromPrivateKey
- [x] getAddressFromPublicKey
- [x] getAddressFromPrivateKey
- [x] encryptPrivateKey
- [x] decryptPrivateKey

##### JSON-RPC API

Blockchain-related methods

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

Transaction-related methods

- [x] createTransaction
- [x] getTransaction
- [x] getRecentTransactions
- [x] getTransactionsForTxBlock
- [x] getNumTxnsTxEpoch
- [x] getNumTxnsDSEpoch
- [x] getMinimumGasPrice

Contract-related methods

- [x] getSmartContractCode
- [x] getSmartContractInit
- [x] getSmartContractState
- [x] getSmartContracts
- [x] getContractAddressFromTransactionID

Account-related methods

- [x] getBalance

##### Validation

- [x] isAddress
- [x] isPublicjKey
- [x] isPrivateKey
- [x] isSignature

##### Util

- [x] byteArrayToHexString
- [x] hexStringToByteArray
- [x] generateMac
- [x] isByteString
- [x] encodeTransactionProto
- [x] toChecksumAddress
- [x] isValidChecksumAddress
- [x] bech32 encode decode
- [x] isBech32
- [x] fromBech32Address toBech32Address

#### Quick Start

##### Generate a large number of private keys:

```go
func TestGeneratePrivateKey(t *testing.T) {
	for i := 0; i < 100000; i++ {
		privateKey, err := GeneratePrivateKey()
		if err != nil {
			panic("cannot generate private key")
		}

		prikeys := LaksaGo.EncodeHex(privateKey[:])
		if len(prikeys) != 64 {
			panic("generate private key error")
		}
		println("private key = " + prikeys)
		publickKey := GetPublicKeyFromPrivateKey(LaksaGo.DecodeHex(prikeys), true)
		pubkeys := LaksaGo.EncodeHex(publickKey)
		if len(pubkeys) != 66 {
			panic("generate public key error")
		}
		println("public key = " + pubkeys)

	}
}
```

##### Encrypt private key to a keystore file:

```go
func TestKeystore_EncryptPrivateKey(t *testing.T) {
	ks := NewDefaultKeystore()
	kv, err := ks.EncryptPrivateKey(util.DecodeHex("24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"), []byte("xiaohuo"), 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		println(kv)
	}
}
```

##### Decrypt private key from a keystore file:

```go
func TestKeystore_DecryptPrivateKey(t *testing.T) {
	json := "{\"address\":\"b5c2cdd79c37209c3cb59e04b7c4062a8f5d5271\",\"id\":\"979daaf9-daf1-4002-8656-3cea134c9518\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"26be10cdae0f397bdeead38e7fcc179957dd5e7ef95a1f0f53f37b7ad1355159\",\"kdf\":\"pbkdf2\",\"mac\":\"81d8e60bc08237e4ba154c0b27ad08562821d8c602ee8a492434128de48b66bc\",\"cipherparams\":{\"iv\":\"fc714ad6267c35a2df4cb3f8b8b3cc0d\"},\"kdfparams\":{\"n\":8192,\"c\":262144,\"r\":8,\"p\":1,\"dklen\":32,\"salt\":\"e22ef8a67a59299cee1532b6c6967bdfb0e75ca3c5dff852f9d8daa04683b0c1\"}}}"

	ks := NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(json, "xiaohuo")
	if err != nil {
		t.Error(err.Error())
	} else {
		if strings.Compare(strings.ToLower(privateKey), "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9") != 0 {
			t.Error("decrypt private key failed")
		}
	}
}
```

##### Send a transfer transaction

```go
func TestSendTransaction(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	tx := &transaction.Transaction{
		Version:      strconv.FormatInt(int64(LaksaGo.Pack(333, 2)), 10),
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		ToAddr:       "4BAF5faDA8e5Db92C3d3242618c5B47133AE003C",
		Amount:       "10000000",
		GasPrice:     "1000000000",
		GasLimit:     "1",
		Code:         "",
		Data:         "",
		Priority:     false,
	}

	err := wallet.Sign(tx, *provider)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	rsp := provider.CreateTransaction(tx.ToTransactionPayload())

	if rsp.Error != nil {
		fmt.Println(rsp.Error)
		t.Error(err)
	} else {
		result := rsp.Result.(map[string]interface{})
		hash := result["TranID"].(string)
		fmt.Printf("hash is %s", hash)
		tx.Confirm(hash, 1000, 3, provider)
	}
}
```

##### Deploy a smart contract

```go
func TestContract_Deploy(t *testing.T) {
	wallet := account.NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	code, _ := ioutil.ReadFile("./fungible.scilla")
	init := []Value{
		{
			"_scilla_version",
			"Uint32",
			"0",
		},
		{
			"owner",
			"ByStr20",
			"0x9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a",
		},
		{
			"total_tokens",
			"Uint128",
			"1000000000",
		},
		{
			"decimals",
			"Uint32",
			"0",
		},
		{
			"name",
			"String",
			"BobCoin",
		},
		{
			"symbol",
			"String",
			"BOB",
		},
	}
	contract := Contract{
		Code:     string(code),
		Init:     init,
		Singer:   wallet,
		Provider: provider,
	}

	nonce, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a").Result.(map[string]interface{})["nonce"].(json.Number).Int64()

	deployParams := DeployParams{
		Version:      strconv.FormatInt(int64(LaksaGo.Pack(333, 1)), 10),
		Nonce:        strconv.FormatInt(nonce+1, 10),
		GasPrice:     "10000000000",
		GasLimit:     "10000",
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
	}

	tx, err := contract.Deploy(deployParams)

	if err != nil {
		panic(err.Error())
	}

	tx.Confirm(tx.ID, 1000, 10, provider)
}
```

##### Call a smart contract

```go
func TestContract_Call(t *testing.T) {
	wallet := account.NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	contract := Contract{
		Address:  "bd7198209529dC42320db4bC8508880BcD22a9f2",
		Singer:   wallet,
		Provider: provider,
	}

	args := []Value{
		{
			"to",
			"ByStr20",
			"0x381f4008505e940ad7681ec3468a719060caf796",
		},
		{
			"tokens",
			"Uint128",
			"10",
		},
	}

	nonce, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a").Result.(map[string]interface{})["nonce"].(json.Number).Int64()
	n := nonce + 1
	params := CallParams{
		Nonce:        strconv.FormatInt(n, 10),
		Version:      strconv.FormatInt(int64(LaksaGo.Pack(333, 1)), 10),
		GasPrice:     "1000000000",
		GasLimit:     "1000",
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		Amount:       "0",
	}

	err, tx := contract.Call("Transfer", args, params, true, 1000, 3)
	if err != nil {
		fmt.Printf(err.Error())
	}

	tx.Confirm(tx.ID, 1000, 3, provider)

}
```