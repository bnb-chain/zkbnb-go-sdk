package types

import (
	"github.com/bnb-chain/zkbnb-crypto/wasm/txtypes"
	"math/big"
)

type OfferTxInfo struct {
	Type         int64
	OfferId      int64
	AccountIndex int64
	NftIndex     int64
	AssetId      int64
	AssetAmount  *big.Int
	ListedAt     int64
	ExpiredAt    int64
	TreasuryRate int64
	Sig          []byte
}

type AtomicMatchTxReq struct {
	BuyOffer  *OfferTxInfo
	SellOffer *OfferTxInfo
}

type CancelOfferReq struct {
	OfferId int64
}

type CreateCollectionReq struct {
	Name               string
	Introduction       string
	CollectionMetaData *txtypes.CollectionMetaData
}

type MintNftTxReq struct {
	To                  string
	NftCollectionId     int64
	CreatorTreasuryRate int64
	MetaData            *txtypes.NftMetaData
}

type UpdateNftReq struct {
	NftIndex int64
	Mutable  string
}

type TransferNftTxReq struct {
	To       string
	NftIndex int64
}

type TransferTxReq struct {
	ToAccountName string
	AssetId       int64
	AssetAmount   *big.Int
}

type WithdrawNftTxReq struct {
	AccountIndex int64
	NftIndex     int64
	ToAddress    string
}

type WithdrawReq struct {
	AssetId     int64
	AssetAmount *big.Int
	ToAddress   string
}

type AttributeStr struct {
	DisplayType string `json:"display_type"`
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
}

type AttributeInt struct {
	DisplayType string `json:"display_type"`
	TraitType   string `json:"trait_type"`
	Value       int64  `json:"value"`
	MaxValue    int64  `json:"max_value"`
}
