package sdk

import "math/big"

type AddLiquidityTxInfo struct {
	FromAccountIndex  int64
	PairIndex         int64
	AssetAId          int64
	AssetAAmount      *big.Int
	AssetBId          int64
	AssetBAmount      *big.Int
	LpAmount          *big.Int
	KLast             *big.Int
	TreasuryAmount    *big.Int
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type RemoveLiquidityTxInfo struct {
	FromAccountIndex  int64
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
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type SwapTxInfo struct {
	FromAccountIndex  int64
	PairIndex         int64
	AssetAId          int64
	AssetAAmount      *big.Int
	AssetBId          int64
	AssetBMinAmount   *big.Int
	AssetBAmountDelta *big.Int
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type AtomicMatchTxInfo struct {
	AccountIndex      int64
	BuyOffer          *OfferTxInfo
	SellOffer         *OfferTxInfo
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	CreatorAmount     *big.Int
	TreasuryAmount    *big.Int
	Nonce             int64
	ExpiredAt         int64
	Sig               []byte
}

type CancelOfferTxInfo struct {
	AccountIndex      int64
	OfferId           int64
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type CreateCollectionTxInfo struct {
	AccountIndex      int64
	CollectionId      int64
	Name              string
	Introduction      string
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type MintNftTxInfo struct {
	CreatorAccountIndex int64
	ToAccountIndex      int64
	ToAccountNameHash   string
	NftIndex            int64
	NftContentHash      string
	NftCollectionId     int64
	CreatorTreasuryRate int64
	GasAccountIndex     int64
	GasFeeAssetId       int64
	GasFeeAssetAmount   *big.Int
	ExpiredAt           int64
	Nonce               int64
	Sig                 []byte
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

type TransferNftTxInfo struct {
	FromAccountIndex  int64
	ToAccountIndex    int64
	ToAccountNameHash string
	NftIndex          int64
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	CallData          string
	CallDataHash      []byte
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type TransferTxInfo struct {
	FromAccountIndex  int64
	ToAccountIndex    int64
	ToAccountNameHash string
	AssetId           int64
	AssetAmount       *big.Int
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	Memo              string
	CallData          string
	CallDataHash      []byte
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type WithdrawNftTxInfo struct {
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
	GasAccountIndex        int64
	GasFeeAssetId          int64
	GasFeeAssetAmount      *big.Int
	ExpiredAt              int64
	Nonce                  int64
	Sig                    []byte
}

type WithdrawTxInfo struct {
	FromAccountIndex  int64
	AssetId           int64
	AssetAmount       *big.Int
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	ToAddress         string
	ExpiredAt         int64
	Nonce             int64
	Sig               []byte
}

type RegisterZnsTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountName     string
	AccountNameHash []byte
	PubKey          string
}

type CreatePairTxInfo struct {
	TxType               uint8
	PairIndex            int64
	AssetAId             int64
	AssetBId             int64
	FeeRate              int64
	TreasuryAccountIndex int64
	TreasuryRate         int64
}

type UpdatePairRateTxInfo struct {
	TxType               uint8
	PairIndex            int64
	FeeRate              int64
	TreasuryAccountIndex int64
	TreasuryRate         int64
}

type DepositTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountNameHash []byte
	AssetId         int64
	AssetAmount     *big.Int
}

type DepositNftTxInfo struct {
	TxType              uint8
	AccountIndex        int64
	NftIndex            int64
	NftL1Address        string
	CreatorAccountIndex int64
	CreatorTreasuryRate int64
	NftContentHash      []byte
	NftL1TokenId        *big.Int
	AccountNameHash     []byte
	CollectionId        int64
}

type FullExitTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountNameHash []byte
	AssetId         int64
	AssetAmount     *big.Int
}

type FullExitNftTxInfo struct {
	TxType                 uint8
	AccountIndex           int64
	CreatorAccountIndex    int64
	CreatorTreasuryRate    int64
	NftIndex               int64
	CollectionId           int64
	NftL1Address           string
	AccountNameHash        []byte
	CreatorAccountNameHash []byte
	NftContentHash         []byte
	NftL1TokenId           *big.Int
}
