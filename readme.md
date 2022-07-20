# ZkBAS Go SDK

The ZkBAS Go SDK provides a thin wrapper around thin all the apis provided by ZkBAS, including a simple key manager for signing
txs and sending signed txs to ZkBAS.

## Usage

### Importing

```go
import (
    "github.com/bnb-chain/zkbas-go-sdk" latest
)
```

### Key manager

The key manager is a simple interface to help sign the txs with the provided private key.

```go
type KeyManager interface {
    Sign(message []byte, hFunc hash.Hash) ([]byte, error)
    PubKey() signature.PublicKey
    PubKeyPoint() [2][]byte
}
```

Currently, we only support key manager generated from seed.

Examples: 

```go
keyManager, _ := NewSeedKeyManager("you private key seed")
```

### ZkBAS Client

```go
type ZkBASClient interface {
    ZkBASQuerier
    ZkBASTxSender
}
```

The ZkBAS go sdk wraps the supported apis and also provides methods to sign txs with the key manager.

#### Init sdk 

The ZkBAS client can be initiated by an API endpoint.

```go
client := NewZkBASClient("The ZkBAS endpoint")
```

#### Queries

You can perform the query methods directly:

```go
// query the layer2 basic info

layer2Info, err := client.GetLayer2BasicInfo()
...
```

#### Send txs

To send txs, you need to init the key manager first and set the key manager to client.

```go
client.SetSetKeyManager(keyManager)
```

Then you can sign txs and send with the sdk:

```go
// prepare tx info
txInfo := &OfferTxInfo{
    Type:         int64(offerType),
    OfferId:      offerId,
    AccountIndex: accountIndex,
    NftIndex:     nftIndex,
    AssetId:      0, //payment asset id
    AssetAmount:  big.NewInt(10000),
    ListedAt:     listedAt,
    ExpiredAt:    expiredAt,
    TreasuryRate: 200,
    Sig:          nil,
}

txId, err := client.Offer(txInfo)
```

You can also sign the raw transaction by yourself and send with the `SendTx` api:

```go
txInfo := &OfferTxInfo{
    Type:         int64(offerType),
    OfferId:      offerId,
    AccountIndex: accountIndex,
    NftIndex:     nftIndex,
    AssetId:      0, //payment asset id
    AssetAmount:  big.NewInt(10000),
    ListedAt:     listedAt,
    ExpiredAt:    expiredAt,
    TreasuryRate: 200,
    Sig:          nil,
}

txInfo, err := ConstructOfferTx(keyManager, txInfo)
if err != nil {
    return "", err
}

client.SendTx(TxTypeOffer, txInfo)
```

### ZkBAS L1 Client

The ZkBASL1Client is used to interact with ZkBAS proxy contract in l1. 

#### Interface 

```go
type ZkBASL1Client interface {
	// DepositBNB will deposit specific amount bnb to l2
	DepositBNB(accountName string, amount *big.Int) (common.Hash, error)

	// DepositBEP20 will deposit specific amount of bep20 token to l2
	DepositBEP20(token common.Address, accountName string, amount *big.Int) (common.Hash, error)

	// DepositNft will deposit specific nft to l2
	DepositNft(nftL1Address common.Address, accountName string, nftL1TokenId *big.Int) (common.Hash, error)

	// RegisterZNS will register account in l2
	RegisterZNS(name string, owner common.Address, pubKeyX [32]byte, pubKeyY [32]byte) (common.Hash, error)

	// CreatePair will create swap pair in l2
	CreatePair(tokenA common.Address, tokenB common.Address) (common.Hash, error)

	// RequestFullExit will request full exit from l2
	RequestFullExit(accountName string, asset common.Address) (common.Hash, error)

	// RequestFullExitNft will request full nft exit from l2
	RequestFullExitNft(accountName string, nftIndex uint32) (common.Hash, error)

	// UpdatePairRate will update pair info in l2
	UpdatePairRate(pairInfo abi.ZkbasPairInfo) (common.Hash, error)
}
```

#### Init

```go
client := NewZkBASL1Client("l1 provider", "zkbas proxy contract address")
```

#### Send tx

Before you send tx, you need to set a private key to sign the tx:

```go
client.SetPrivateKey("private key")
```

Then you can send txs.
