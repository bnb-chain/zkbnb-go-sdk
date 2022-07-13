package txutils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bnb-chain/zkbas-go-sdk/types"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
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

func ConvertTransferNftTxInfo(tx *types.TransferNftTxInfo) *legendTxTypes.TransferNftTxInfo {
	return &legendTxTypes.TransferNftTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		ToAccountIndex:    tx.ToAccountIndex,
		ToAccountNameHash: tx.ToAccountNameHash,
		NftIndex:          tx.NftIndex,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		CallData:          tx.CallData,
		CallDataHash:      tx.CallDataHash,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func VerifyTransferNftTxSig(pubKey string, tx *types.TransferNftTxInfo) error {
	convertedTx := ConvertTransferNftTxInfo(tx)
	message, err := legendTxTypes.ComputeTransferNftMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertWithdrawNftTxInfo(tx *types.WithdrawNftTxInfo) *legendTxTypes.WithdrawNftTxInfo {
	return &legendTxTypes.WithdrawNftTxInfo{
		AccountIndex:           tx.AccountIndex,
		CreatorAccountIndex:    tx.CreatorAccountIndex,
		CreatorAccountNameHash: tx.CreatorAccountNameHash,
		CreatorTreasuryRate:    tx.CreatorTreasuryRate,
		NftIndex:               tx.NftIndex,
		NftContentHash:         tx.NftContentHash,
		NftL1Address:           tx.NftL1Address,
		NftL1TokenId:           tx.NftL1TokenId,
		CollectionId:           tx.CollectionId,
		ToAddress:              tx.ToAddress,
		GasAccountIndex:        tx.GasAccountIndex,
		GasFeeAssetId:          tx.GasFeeAssetId,
		GasFeeAssetAmount:      tx.GasFeeAssetAmount,
		ExpiredAt:              tx.ExpiredAt,
		Nonce:                  tx.Nonce,
		Sig:                    tx.Sig,
	}
}

func VerifyWithdrawNftTxSig(pubKey string, tx *types.WithdrawNftTxInfo) error {
	convertedTx := ConvertWithdrawNftTxInfo(tx)
	message, err := legendTxTypes.ComputeWithdrawNftMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
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
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertMintNftTxInfo(tx *types.MintNftTxInfo) *legendTxTypes.MintNftTxInfo {
	return &legendTxTypes.MintNftTxInfo{
		CreatorAccountIndex: tx.CreatorAccountIndex,
		ToAccountIndex:      tx.ToAccountIndex,
		ToAccountNameHash:   tx.ToAccountNameHash,
		NftIndex:            tx.NftIndex,
		NftContentHash:      tx.NftContentHash,
		NftCollectionId:     tx.NftCollectionId,
		CreatorTreasuryRate: tx.CreatorTreasuryRate,
		GasAccountIndex:     tx.GasAccountIndex,
		GasFeeAssetId:       tx.GasFeeAssetId,
		GasFeeAssetAmount:   tx.GasFeeAssetAmount,
		ExpiredAt:           tx.ExpiredAt,
		Nonce:               tx.Nonce,
		Sig:                 tx.Sig,
	}
}

func VerifyMintNftTxSig(pubKey string, tx *types.MintNftTxInfo) error {
	convertedTx := ConvertMintNftTxInfo(tx)
	message, err := legendTxTypes.ComputeMintNftMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertTransferTx(tx *types.TransferTxInfo) *legendTxTypes.TransferTxInfo {
	return &legendTxTypes.TransferTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		ToAccountIndex:    tx.ToAccountIndex,
		ToAccountNameHash: tx.ToAccountNameHash,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		Memo:              tx.Memo,
		CallData:          tx.CallData,
		CallDataHash:      tx.CallDataHash,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertSwapTx(tx *types.SwapTxInfo) *legendTxTypes.SwapTxInfo {
	return &legendTxTypes.SwapTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAAmount:      tx.AssetAAmount,
		AssetBId:          tx.AssetBId,
		AssetBMinAmount:   tx.AssetBMinAmount,
		AssetBAmountDelta: tx.AssetBAmountDelta,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertAddLiquidityTx(tx *types.AddLiquidityTxInfo) *legendTxTypes.AddLiquidityTxInfo {
	return &legendTxTypes.AddLiquidityTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAAmount:      tx.AssetAAmount,
		AssetBId:          tx.AssetBId,
		AssetBAmount:      tx.AssetBAmount,
		LpAmount:          tx.LpAmount,
		KLast:             tx.KLast,
		TreasuryAmount:    tx.TreasuryAmount,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertRemoveLiquidityTx(tx *types.RemoveLiquidityTxInfo) *legendTxTypes.RemoveLiquidityTxInfo {
	return &legendTxTypes.RemoveLiquidityTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAMinAmount:   tx.AssetAMinAmount,
		AssetBId:          tx.AssetBId,
		AssetBMinAmount:   tx.AssetBMinAmount,
		LpAmount:          tx.LpAmount,
		AssetAAmountDelta: tx.AssetAAmountDelta,
		AssetBAmountDelta: tx.AssetBAmountDelta,
		KLast:             tx.KLast,
		TreasuryAmount:    tx.TreasuryAmount,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertWithdrawTx(tx *types.WithdrawTxInfo) *legendTxTypes.WithdrawTxInfo {
	return &legendTxTypes.WithdrawTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ToAddress:         tx.ToAddress,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertCreateCollectionTxInfo(tx *types.CreateCollectionTxInfo) *legendTxTypes.CreateCollectionTxInfo {
	return &legendTxTypes.CreateCollectionTxInfo{
		AccountIndex:      tx.AccountIndex,
		CollectionId:      tx.CollectionId,
		Name:              tx.Name,
		Introduction:      tx.Introduction,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func VerifyCreateCollectionTxSig(pubKey string, tx *types.CreateCollectionTxInfo) error {
	convertedTx := ConvertCreateCollectionTxInfo(tx)
	message, err := legendTxTypes.ComputeCreateCollectionMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertAtomicMatchTxInfo(tx *types.AtomicMatchTxInfo) *legendTxTypes.AtomicMatchTxInfo {
	return &legendTxTypes.AtomicMatchTxInfo{
		AccountIndex: tx.AccountIndex,
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
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		CreatorAmount:     tx.CreatorAmount,
		TreasuryAmount:    tx.TreasuryAmount,
		Nonce:             tx.Nonce,
		ExpiredAt:         tx.ExpiredAt,
		Sig:               tx.Sig,
	}
}
func VerifyAtomicMatchTxSig(pubKey string, tx *types.AtomicMatchTxInfo) error {
	convertedTx := ConvertAtomicMatchTxInfo(tx)
	message, err := legendTxTypes.ComputeAtomicMatchMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertCancelOfferTxInfo(tx *types.CancelOfferTxInfo) *legendTxTypes.CancelOfferTxInfo {
	return &legendTxTypes.CancelOfferTxInfo{
		AccountIndex:      tx.AccountIndex,
		OfferId:           tx.OfferId,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func VerifyCancelOfferTxSig(pubKey string, tx *types.CancelOfferTxInfo) error {
	convertedTx := ConvertCancelOfferTxInfo(tx)
	message, err := legendTxTypes.ComputeCancelOfferMsgHash(convertedTx, mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(convertedTx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}
