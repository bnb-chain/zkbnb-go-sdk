package types

import (
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

type CancelOfferTxReq struct {
	OfferId int64
}

type CreateCollectionTxReq struct {
	Name         string
	Introduction string
}

type MintNftTxReq struct {
	To                  string
	NftCollectionId     int64
	CreatorTreasuryRate int64
	MetaData            string
	MutableAttributes   string
}

type UpdateNftReq struct {
	NftIndex          int64
	MutableAttributes string
	AccountIndex      int64
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

type WithdrawTxReq struct {
	AssetId     int64
	AssetAmount *big.Int
	ToAddress   string
}
