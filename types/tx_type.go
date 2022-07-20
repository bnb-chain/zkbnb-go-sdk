package types

import (
	"encoding/json"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"math/big"
)

const (
	TxTypeEmpty = iota
	TxTypeRegisterZns
	TxTypeCreatePair
	TxTypeUpdatePairRate
	TxTypeDeposit
	TxTypeDepositNft
	TxTypeTransfer
	TxTypeSwap
	TxTypeAddLiquidity
	TxTypeRemoveLiquidity
	TxTypeWithdraw
	TxTypeCreateCollection
	TxTypeMintNft
	TxTypeTransferNft
	TxTypeAtomicMatch
	TxTypeCancelOffer
	TxTypeWithdrawNft
	TxTypeFullExit
	TxTypeFullExitNft
	TxTypeOffer
)

const (
	BuyOfferType  = 0
	SellOfferType = 1
)

type TransactOpts struct {
	FromAccountIndex  int64
	GasAccountIndex   int64
	GasFeeAssetId     int64
	GasFeeAssetAmount *big.Int
	CallData          string
	CallDataHash      []byte
	ExpiredAt         int64
	Nonce             int64
	Memo              string

	// Optional
	ToAccountIndex    int64
	ToAccountNameHash string
}

type AddLiquidityTxInfo struct {
	PairIndex      int64
	AssetAId       int64
	AssetAAmount   *big.Int
	AssetBId       int64
	AssetBAmount   *big.Int
	LpAmount       *big.Int
	KLast          *big.Int
	TreasuryAmount *big.Int
}

func ParseAddLiquidityTxInfo(txInfoStr string) (txInfo *AddLiquidityTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type RemoveLiquidityTxInfo struct {
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

func ParseRemoveLiquidityTxInfo(txInfoStr string) (txInfo *RemoveLiquidityTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type SwapTxInfo struct {
	PairIndex         int64
	AssetAId          int64
	AssetAAmount      *big.Int
	AssetBId          int64
	AssetBMinAmount   *big.Int
	AssetBAmountDelta *big.Int
}

func ParseSwapTxInfo(txInfoStr string) (txInfo *SwapTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type AtomicMatchTxInfo struct {
	BuyOffer       *OfferTxInfo
	SellOffer      *OfferTxInfo
	CreatorAmount  *big.Int
	TreasuryAmount *big.Int
}

func ParseAtomicMatchTxInfo(txInfoStr string) (txInfo *AtomicMatchTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type CancelOfferTxInfo struct {
	OfferId int64
}

func ParseCancelOfferTxInfo(txInfoStr string) (txInfo *CancelOfferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type CreateCollectionTxInfo struct {
	Name         string
	Introduction string
}

func ParseCreateCollectionTxInfo(txInfoStr string) (txInfo *CreateCollectionTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type MintNftTxInfo struct {
	To                  string
	NftContentHash      string
	NftCollectionId     int64
	CreatorTreasuryRate int64
}

func ParseMintNftTxInfo(txInfoStr string) (txInfo *MintNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
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

func ParseOfferTxInfo(txInfoStr string) (txInfo *OfferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type TransferNftTxInfo struct {
	To       string
	NftIndex int64
}

func ParseTransferNftTxInfo(txInfoStr string) (txInfo *TransferNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type TransferTxInfo struct {
	ToAccountName string
	AssetId       int64
	AssetAmount   *big.Int
}

func ParseTransferTxInfo(txInfoStr string) (txInfo *legendTxTypes.TransferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
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
}

func ParseWithdrawNftTxInfo(txInfoStr string) (txInfo *legendTxTypes.WithdrawNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type WithdrawTxInfo struct {
	AssetId     int64
	AssetAmount *big.Int
	ToAddress   string
}

func ParseWithdrawTxInfo(txInfoStr string) (txInfo *legendTxTypes.WithdrawTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type RegisterZnsTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountName     string
	AccountNameHash []byte
	PubKey          string
}

func ParseRegisterZnsTxInfo(txInfoStr string) (txInfo *RegisterZnsTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
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

func ParseCreatePairTxInfo(txInfoStr string) (txInfo *CreatePairTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type UpdatePairRateTxInfo struct {
	TxType               uint8
	PairIndex            int64
	FeeRate              int64
	TreasuryAccountIndex int64
	TreasuryRate         int64
}

func ParseUpdatePairRateTxInfo(txInfoStr string) (txInfo *UpdatePairRateTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type DepositTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountNameHash []byte
	AssetId         int64
	AssetAmount     *big.Int
}

func ParseDepositTxInfo(txInfoStr string) (txInfo *DepositTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
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

func ParseDepositNftTxInfo(txInfoStr string) (txInfo *DepositNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type FullExitTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountNameHash []byte
	AssetId         int64
	AssetAmount     *big.Int
}

func ParseFullExitTxInfo(txInfoStr string) (txInfo *FullExitTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
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

func ParseFullExitNftTxInfo(txInfoStr string) (txInfo *FullExitNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func MarshalTxInfo(txInfo interface{}) (string, error) {
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}
