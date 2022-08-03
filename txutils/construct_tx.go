package txutils

import (
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	curve "github.com/bnb-chain/zkbas-crypto/ecc/ztwistededwards/tebn254"
	"github.com/bnb-chain/zkbas-crypto/ffmath"
	"github.com/bnb-chain/zkbas-crypto/wasm/legend/legendTxTypes"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

func ConstructWithdrawTxInfo(key accounts.Signer, tx *types.WithdrawReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertWithdrawTx(tx, ops)
	err := legendTxTypes.ValidateWithdrawTxInfo(convertedTx)
	if err != nil {
		return "", err
	}

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

func ConstructRemoveLiquidityTx(key accounts.Signer, tx *types.RemoveLiquidityReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertRemoveLiquidityTx(tx, ops)
	err := legendTxTypes.ValidateRemoveLiquidityTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructAddLiquidityTx(key accounts.Signer, tx *types.AddLiquidityReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertAddLiquidityTx(tx, ops)
	err := legendTxTypes.ValidateAddLiquidityTxInfo(convertedTx)
	if err != nil {
		return "", err
	}

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

func ConstructSwapTx(key accounts.Signer, tx *types.SwapTxReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertSwapTx(tx, ops)
	err := legendTxTypes.ValidateSwapTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructTransferTx(key accounts.Signer, ops *types.TransactOpts, tx *types.TransferTxReq) (string, error) {
	convertedTx := ConvertTransferTx(tx, ops)
	err := legendTxTypes.ValidateTransferTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructCreateCollectionTx(key accounts.Signer, tx *types.CreateCollectionReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertCreateCollectionTxInfo(tx, ops)
	err := legendTxTypes.ValidateCreateCollectionTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructTransferNftTx(key accounts.Signer, tx *types.TransferNftTxReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertTransferNftTxInfo(tx, ops)
	err := legendTxTypes.ValidateTransferNftTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructWithdrawNftTx(key accounts.Signer, tx *types.WithdrawNftTxReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertWithdrawNftTxInfo(tx, ops)
	err := legendTxTypes.ValidateWithdrawNftTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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
	err := legendTxTypes.ValidateOfferTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructMintNftTx(key accounts.Signer, tx *types.MintNftTxReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertMintNftTxInfo(tx, ops)
	err := legendTxTypes.ValidateMintNftTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructAtomicMatchTx(key accounts.Signer, tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertAtomicMatchTxInfo(tx, ops)
	err := legendTxTypes.ValidateAtomicMatchTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

func ConstructCancelOfferTx(key accounts.Signer, tx *types.CancelOfferReq, ops *types.TransactOpts) (string, error) {
	convertedTx := ConvertCancelOfferTxInfo(tx, ops)
	err := legendTxTypes.ValidateCancelOfferTxInfo(convertedTx)
	if err != nil {
		return "", err
	}
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

	q, _ := big.NewInt(0).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

	rootNode := make([]byte, 32)
	hashOfBaseNode := keccakHash(append(rootNode, keccakHash([]byte(words[1]))...))

	baseNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(hashOfBaseNode), q)
	baseNodeBytes := make([]byte, 32)
	baseNode.FillBytes(baseNodeBytes)

	nameHash := keccakHash([]byte(words[0]))
	subNameHash := keccakHash(append(baseNodeBytes, nameHash...))

	subNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(subNameHash), q)
	subNodeBytes := make([]byte, 32)
	subNode.FillBytes(subNodeBytes)

	res = common.Bytes2Hex(subNodeBytes)
	return res, nil
}

func NftContentHash(nftContent string) string {
	return common.Bytes2Hex(ffmath.Mod(new(big.Int).SetBytes(common.FromHex(nftContent)), curve.Modulus).FillBytes(make([]byte, 32)))
}
