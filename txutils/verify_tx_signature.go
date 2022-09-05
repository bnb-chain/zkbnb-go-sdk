package txutils

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"

	"github.com/bnb-chain/zkbnb-crypto/wasm/legend/legendTxTypes"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
)

type PublicKey = eddsa.PublicKey

func parsePk(pkStr string) (pk *PublicKey, err error) {
	pkBytes, err := hex.DecodeString(pkStr)
	if err != nil {
		return nil, err
	}
	pk = new(PublicKey)
	size, err := pk.SetBytes(pkBytes)
	if err != nil {
		return nil, err
	}
	if size != 32 {
		return nil, errors.New("invalid public key")
	}
	return pk, nil
}

func ConvertTransferNftTxInfo(tx *types.TransferNftTxReq, ops *types.TransactOpts) *legendTxTypes.TransferNftTxInfo {
	return &legendTxTypes.TransferNftTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		ToAccountIndex:    ops.ToAccountIndex,
		ToAccountNameHash: ops.ToAccountNameHash,
		NftIndex:          tx.NftIndex,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		CallData:          ops.CallData,
		CallDataHash:      ops.CallDataHash,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertWithdrawNftTxInfo(tx *types.WithdrawNftTxReq, ops *types.TransactOpts) *legendTxTypes.WithdrawNftTxInfo {
	return &legendTxTypes.WithdrawNftTxInfo{
		AccountIndex:      tx.AccountIndex,
		NftIndex:          tx.NftIndex,
		ToAddress:         tx.ToAddress,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertOfferTxInfo(tx *types.OfferTxInfo) *legendTxTypes.OfferTxInfo {
	return &legendTxTypes.OfferTxInfo{
		Type:         tx.Type,
		OfferId:      tx.OfferId,
		AccountIndex: tx.AccountIndex,
		NftIndex:     tx.NftIndex,
		AssetId:      tx.AssetId,
		AssetAmount:  tx.AssetAmount,
		ListedAt:     tx.ListedAt,
		ExpiredAt:    tx.ExpiredAt,
		TreasuryRate: tx.TreasuryRate,
		Sig:          tx.Sig,
	}
}

func ConvertMintNftTxInfo(tx *types.MintNftTxReq, ops *types.TransactOpts) *legendTxTypes.MintNftTxInfo {
	return &legendTxTypes.MintNftTxInfo{
		CreatorAccountIndex: ops.FromAccountIndex,
		ToAccountIndex:      ops.ToAccountIndex,
		ToAccountNameHash:   ops.ToAccountNameHash,
		NftContentHash:      tx.NftContentHash,
		NftCollectionId:     tx.NftCollectionId,
		CreatorTreasuryRate: tx.CreatorTreasuryRate,
		GasAccountIndex:     ops.GasAccountIndex,
		GasFeeAssetId:       ops.GasFeeAssetId,
		GasFeeAssetAmount:   ops.GasFeeAssetAmount,
		ExpiredAt:           ops.ExpiredAt,
		Nonce:               ops.Nonce,
	}
}

func ConvertTransferTx(tx *types.TransferTxReq, ops *types.TransactOpts) *legendTxTypes.TransferTxInfo {
	return &legendTxTypes.TransferTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		ToAccountIndex:    ops.ToAccountIndex,
		ToAccountNameHash: ops.ToAccountNameHash,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		Memo:              ops.Memo,
		CallData:          ops.CallData,
		CallDataHash:      ops.CallDataHash,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertSwapTx(tx *types.SwapTxReq, ops *types.TransactOpts) *legendTxTypes.SwapTxInfo {
	return &legendTxTypes.SwapTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAAmount:      tx.AssetAAmount,
		AssetBId:          tx.AssetBId,
		AssetBMinAmount:   tx.AssetBMinAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertAddLiquidityTx(tx *types.AddLiquidityReq, ops *types.TransactOpts) *legendTxTypes.AddLiquidityTxInfo {
	return &legendTxTypes.AddLiquidityTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAAmount:      tx.AssetAAmount,
		AssetBAmount:      tx.AssetBAmount,
		LpAmount:          tx.LpAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertRemoveLiquidityTx(tx *types.RemoveLiquidityReq, ops *types.TransactOpts) *legendTxTypes.RemoveLiquidityTxInfo {
	return &legendTxTypes.RemoveLiquidityTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAMinAmount:   tx.AssetAMinAmount,
		AssetBMinAmount:   tx.AssetBMinAmount,
		LpAmount:          tx.LpAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertWithdrawTx(tx *types.WithdrawReq, ops *types.TransactOpts) *legendTxTypes.WithdrawTxInfo {
	return &legendTxTypes.WithdrawTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ToAddress:         tx.ToAddress,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertCreateCollectionTxInfo(tx *types.CreateCollectionReq, ops *types.TransactOpts) *legendTxTypes.CreateCollectionTxInfo {
	return &legendTxTypes.CreateCollectionTxInfo{
		AccountIndex:      ops.FromAccountIndex,
		Name:              tx.Name,
		Introduction:      tx.Introduction,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertAtomicMatchTxInfo(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) *legendTxTypes.AtomicMatchTxInfo {
	return &legendTxTypes.AtomicMatchTxInfo{
		AccountIndex: ops.FromAccountIndex,
		BuyOffer: &legendTxTypes.OfferTxInfo{
			Type:         tx.BuyOffer.Type,
			OfferId:      tx.BuyOffer.OfferId,
			AccountIndex: tx.BuyOffer.AccountIndex,
			NftIndex:     tx.BuyOffer.NftIndex,
			AssetId:      tx.BuyOffer.AssetId,
			AssetAmount:  tx.BuyOffer.AssetAmount,
			ListedAt:     tx.BuyOffer.ListedAt,
			ExpiredAt:    tx.BuyOffer.ExpiredAt,
			TreasuryRate: tx.BuyOffer.TreasuryRate,
			Sig:          tx.BuyOffer.Sig,
		},
		SellOffer: &legendTxTypes.OfferTxInfo{
			Type:         tx.SellOffer.Type,
			OfferId:      tx.SellOffer.OfferId,
			AccountIndex: tx.SellOffer.AccountIndex,
			NftIndex:     tx.SellOffer.NftIndex,
			AssetId:      tx.SellOffer.AssetId,
			AssetAmount:  tx.SellOffer.AssetAmount,
			ListedAt:     tx.SellOffer.ListedAt,
			ExpiredAt:    tx.SellOffer.ExpiredAt,
			TreasuryRate: tx.SellOffer.TreasuryRate,
			Sig:          tx.SellOffer.Sig,
		},
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		Nonce:             ops.Nonce,
		ExpiredAt:         ops.ExpiredAt,
	}
}

func ConvertCancelOfferTxInfo(tx *types.CancelOfferReq, ops *types.TransactOpts) *legendTxTypes.CancelOfferTxInfo {
	return &legendTxTypes.CancelOfferTxInfo{
		AccountIndex:      ops.FromAccountIndex,
		OfferId:           tx.OfferId,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func VerifyCancelOfferTxSig(pubKey string, tx *types.CancelOfferTxInfo) error {
	message, err := legendTxTypes.ComputeCancelOfferMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyWithdrawNftTxSig(pubKey string, tx *types.WithdrawNftTxInfo) error {
	message, err := legendTxTypes.ComputeWithdrawNftMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyTransferNftTxSig(pubKey string, tx *types.TransferNftTxInfo) error {
	message, err := legendTxTypes.ComputeTransferNftMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyOfferTxSig(pubKey string, tx *types.OfferTxInfo) error {
	convertedTx := ConvertOfferTxInfo(tx)
	message, err := legendTxTypes.ComputeOfferMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyMintNftTxSig(pubKey string, tx *types.MintNftTxInfo) error {
	message, err := legendTxTypes.ComputeMintNftMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyCreateCollectionTxSig(pubKey string, tx *types.CreateCollectionTxInfo) error {
	message, err := legendTxTypes.ComputeCreateCollectionMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyAtomicMatchTxSig(pubKey string, tx *types.AtomicMatchTxInfo) error {
	message, err := legendTxTypes.ComputeAtomicMatchMsgHash(tx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}
