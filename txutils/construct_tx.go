package txutils

import (
	"encoding/json"
	"errors"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"strings"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

func ConstructWithdrawTxInfo(key accounts.Signer, tx *types.WithdrawTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertWithdrawTx(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeWithdrawMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructRemoveLiquidityTx(key accounts.Signer, tx *types.RemoveLiquidityTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertRemoveLiquidityTx(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeRemoveLiquidityMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructAddLiquidityTx(key accounts.Signer, tx *types.AddLiquidityTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertAddLiquidityTx(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeAddLiquidityMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructSwapTx(key accounts.Signer, tx *types.SwapTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertSwapTx(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeSwapMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructTransferTx(key accounts.Signer, ops *types.TransactOpts, tx *types.TransferTxInfo) (string, error) {
	convertedTx := ConvertTransferTx(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeTransferMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructCreateCollectionTx(key accounts.Signer, tx *types.CreateCollectionTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertCreateCollectionTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeCreateCollectionMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructTransferNftTx(key accounts.Signer, tx *types.TransferNftTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertTransferNftTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeTransferNftMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructWithdrawNftTx(key accounts.Signer, tx *types.WithdrawNftTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertWithdrawNftTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeWithdrawNftMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructOfferTx(key accounts.Signer, tx *types.OfferTxInfo) (string, error) {
	convertedTx := ConvertOfferTxInfo(tx)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeOfferMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructMintNftTx(key accounts.Signer, tx *types.MintNftTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertMintNftTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeMintNftMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructAtomicMatchTx(key accounts.Signer, tx *types.AtomicMatchTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertAtomicMatchTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeAtomicMatchMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func ConstructCancelOfferTx(key accounts.Signer, tx *types.CancelOfferTxInfo, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertCancelOfferTxInfo(tx, ops)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeCancelOfferMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}

func keccakHash(value []byte) []byte {
	hashVal := crypto.Keccak256Hash(value)
	return hashVal[:]
}

func AccountNameHash(accountName string) (res string, err error) {
	words := strings.Split(accountName, ".")
	if len(words) != 2 {
		return "", errors.New("[AccountNameHash] invalid account name")
	}
	buf := make([]byte, 32)
	label := keccakHash([]byte(words[0]))
	res = common.Bytes2Hex(
		keccakHash(append(
			keccakHash(append(buf,
				keccakHash([]byte(words[1]))...)), label...)))
	return res, nil
}
