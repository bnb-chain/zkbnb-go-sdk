package types

import "math/big"

type RemoveLiquidityReq struct {
	PairIndex         int64
	AssetAId          int64
	AssetAMinAmount   *big.Int
	AssetBId          int64
	AssetBMinAmount   *big.Int
	LpAmount          *big.Int
	AssetAAmountDelta *big.Int
	AssetBAmountDelta *big.Int
	KLast             *big.Int
	TreasuryAmount    *big.Int
}

type SwapTxReq struct {
	PairIndex         int64
	AssetAId          int64
	AssetAAmount      *big.Int
	AssetBId          int64
	AssetBMinAmount   *big.Int
	AssetBAmountDelta *big.Int
}

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
	BuyOffer       *OfferTxInfo
	SellOffer      *OfferTxInfo
	CreatorAmount  *big.Int
	TreasuryAmount *big.Int
}

type CancelOfferReq struct {
	OfferId int64
}

type CreateCollectionReq struct {
	Name         string
	Introduction string
}

type MintNftTxReq struct {
	To                  string
	NftContentHash      string
	NftCollectionId     int64
	CreatorTreasuryRate int64
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
	AccountIndex           int64
	CreatorAccountIndex    int64
	CreatorAccountNameHash []byte
	CreatorTreasuryRate    int64
	NftIndex               int64
	NftContentHash         []byte
	NftL1Address           string
	NftL1TokenId           *big.Int
	CollectionId           int64
	ToAddress              string
}

type WithdrawReq struct {
	AssetId     int64
	AssetAmount *big.Int
	ToAddress   string
}

type AddLiquidityReq struct {
	PairIndex      int64
	AssetAId       int64
	AssetAAmount   *big.Int
	AssetBId       int64
	AssetBAmount   *big.Int
	LpAmount       *big.Int
	KLast          *big.Int
	TreasuryAmount *big.Int
}
