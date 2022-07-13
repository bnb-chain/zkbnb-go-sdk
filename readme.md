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