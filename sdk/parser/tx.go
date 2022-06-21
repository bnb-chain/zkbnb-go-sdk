package parser

import (
	"encoding/json"
	"math/big"

	"github.com/bnb-chain/zkbas-go-sdk/sdk"
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

type RegisterZnsTxInfo struct {
	TxType          uint8
	AccountIndex    int64
	AccountName     string
	AccountNameHash []byte
	PubKey          string
}

func Parse(tx *sdk.Tx) {

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

func ParseCreateCollectionTxInfo(txInfoStr string) (txInfo *CreateCollectionTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

// layer-2 transactions
func ParseTransferTxInfo(txInfoStr string) (txInfo *TransferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseSwapTxInfo(txInfoStr string) (txInfo *SwapTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseAddLiquidityTxInfo(txInfoStr string) (txInfo *AddLiquidityTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseRemoveLiquidityTxInfo(txInfoStr string) (txInfo *RemoveLiquidityTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseMintNftTxInfo(txInfoStr string) (txInfo *MintNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseTransferNftTxInfo(txInfoStr string) (txInfo *TransferNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseOfferTxInfo(txInfoStr string) (txInfo *OfferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseAtomicMatchTxInfo(txInfoStr string) (txInfo *AtomicMatchTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseCancelOfferTxInfo(txInfoStr string) (txInfo *CancelOfferTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseWithdrawTxInfo(txInfoStr string) (txInfo *WithdrawTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func ParseWithdrawNftTxInfo(txInfoStr string) (txInfo *WithdrawNftTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}
