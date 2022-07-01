package sdk

import (
	"encoding/json"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
)

func ConstructCreateCollectionTx(key KeyManager, tx *CreateCollectionTxInfo) (string, error) {
	convertedTx := ConvertCreateCollectionTxInfo(tx)
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

func ConstructTransferNftTx(key KeyManager, tx *TransferNftTxInfo) (string, error) {
	convertedTx := ConvertTransferNftTxInfo(tx)
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

func ConstructWithdrawNftTx(key KeyManager, tx *WithdrawNftTxInfo) (string, error) {
	convertedTx := ConvertWithdrawNftTxInfo(tx)
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

func ConstructOfferTx(key KeyManager, tx *OfferTxInfo) (string, error) {
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

func ConstructMintNftTx(key KeyManager, tx *MintNftTxInfo) (string, error) {
	convertedTx := ConvertMintNftTxInfo(tx)
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

func ConstructAtomicMatchTx(key KeyManager, tx *AtomicMatchTxInfo) (string, error) {
	convertedTx := ConvertAtomicMatchTxInfo(tx)
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

func ConstructCancelOfferTx(key KeyManager, tx *CancelOfferTxInfo) (string, error) {
	convertedTx := ConvertCancelOfferTxInfo(tx)
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
