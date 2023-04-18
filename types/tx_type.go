package types

import (
	"encoding/json"
	"math/big"

	"github.com/bnb-chain/zkbnb-crypto/wasm/txtypes"
)

type (
	AtomicMatchTxInfo      = txtypes.AtomicMatchTxInfo
	CancelOfferTxInfo      = txtypes.CancelOfferTxInfo
	CreateCollectionTxInfo = txtypes.CreateCollectionTxInfo
	TransferNftTxInfo      = txtypes.TransferNftTxInfo
	MintNftTxInfo          = txtypes.MintNftTxInfo
	TransferTxInfo         = txtypes.TransferTxInfo
	WithdrawNftTxInfo      = txtypes.WithdrawNftTxInfo
	WithdrawTxInfo         = txtypes.WithdrawTxInfo
	OfferTxInfo            = txtypes.OfferTxInfo
)

const (
	TxTypeEmpty = iota
	TxTypeChangePubKey
	TxTypeDeposit
	TxTypeDepositNft
	TxTypeTransfer
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
	TxTypeUpdateNFT
)

const (
	BuyOfferType  = 0
	SellOfferType = 1
)

type TransactOpts struct {
	TxType            int
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
	ToAccountIndex   int64
	ToAccountAddress string
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

func ParseCreateCollectionTxInfo(txInfoStr string) (txInfo *CreateCollectionTxInfo, err error) {
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

func ParseOfferTxInfo(txInfoStr string) (txInfo *OfferTxInfo, err error) {
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

func ParseTransferTxInfo(txInfoStr string) (txInfo *TransferTxInfo, err error) {
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

func ParseWithdrawTxInfo(txInfoStr string) (txInfo *WithdrawTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type ChangePubKeyTxInfo struct {
	TxType       uint8
	AccountIndex int64
	L1Address    string
	PubKeyX      string
	PubKeyY      string
}

func ParseChangePubKeyTxInfo(txInfoStr string) (txInfo *ChangePubKeyTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type DepositTxInfo struct {
	TxType       uint8
	AccountIndex int64
	L1Address    string
	AssetId      int64
	AssetAmount  *big.Int
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
	NftContentType      int8
	NftL1TokenId        *big.Int
	L1Address           string
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
	TxType       uint8
	AccountIndex int64
	L1Address    string
	AssetId      int64
	AssetAmount  *big.Int
}

func ParseFullExitTxInfo(txInfoStr string) (txInfo *FullExitTxInfo, err error) {
	err = json.Unmarshal([]byte(txInfoStr), &txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

type FullExitNftTxInfo struct {
	TxType              uint8
	AccountIndex        int64
	CreatorAccountIndex int64
	CreatorTreasuryRate int64
	NftIndex            int64
	CollectionId        int64
	NftL1Address        string
	L1Address           string
	CreatorL1Address    string
	NftContentHash      []byte
	NftContentType      int8
	NftL1TokenId        *big.Int
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
