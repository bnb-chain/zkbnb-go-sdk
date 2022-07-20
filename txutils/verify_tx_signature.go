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

func ConvertTransferNftTxInfo(tx *types.TransferNftTxInfo, ops *types.TransactOpts) *legendTxTypes.TransferNftTxInfo {
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

func ConvertWithdrawNftTxInfo(tx *types.WithdrawNftTxInfo, ops *types.TransactOpts) *legendTxTypes.WithdrawNftTxInfo {
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
		GasAccountIndex:        ops.GasAccountIndex,
		GasFeeAssetId:          ops.GasFeeAssetId,
		GasFeeAssetAmount:      ops.GasFeeAssetAmount,
		ExpiredAt:              ops.ExpiredAt,
		Nonce:                  ops.Nonce,
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

func ConvertMintNftTxInfo(tx *types.MintNftTxInfo, ops *types.TransactOpts) *legendTxTypes.MintNftTxInfo {
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

func ConvertTransferTx(tx *types.TransferTxInfo, ops *types.TransactOpts) *legendTxTypes.TransferTxInfo {
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

func ConvertSwapTx(tx *types.SwapTxInfo, ops *types.TransactOpts) *legendTxTypes.SwapTxInfo {
	return &legendTxTypes.SwapTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAAmount:      tx.AssetAAmount,
		AssetBId:          tx.AssetBId,
		AssetBMinAmount:   tx.AssetBMinAmount,
		AssetBAmountDelta: tx.AssetBAmountDelta,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertAddLiquidityTx(tx *types.AddLiquidityTxInfo, ops *types.TransactOpts) *legendTxTypes.AddLiquidityTxInfo {
	return &legendTxTypes.AddLiquidityTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		PairIndex:         tx.PairIndex,
		AssetAId:          tx.AssetAId,
		AssetAAmount:      tx.AssetAAmount,
		AssetBId:          tx.AssetBId,
		AssetBAmount:      tx.AssetBAmount,
		LpAmount:          tx.LpAmount,
		KLast:             tx.KLast,
		TreasuryAmount:    tx.TreasuryAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertRemoveLiquidityTx(tx *types.RemoveLiquidityTxInfo, ops *types.TransactOpts) *legendTxTypes.RemoveLiquidityTxInfo {
	return &legendTxTypes.RemoveLiquidityTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
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
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertWithdrawTx(tx *types.WithdrawTxInfo, ops *types.TransactOpts) *legendTxTypes.WithdrawTxInfo {
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

func ConvertCreateCollectionTxInfo(tx *types.CreateCollectionTxInfo, ops *types.TransactOpts) *legendTxTypes.CreateCollectionTxInfo {
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

func ConvertAtomicMatchTxInfo(tx *types.AtomicMatchTxInfo, ops *types.TransactOpts) *legendTxTypes.AtomicMatchTxInfo {
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
		CreatorAmount:     tx.CreatorAmount,
		TreasuryAmount:    tx.TreasuryAmount,
		Nonce:             ops.Nonce,
		ExpiredAt:         ops.ExpiredAt,
	}
}

func ConvertCancelOfferTxInfo(tx *types.CancelOfferTxInfo, ops *types.TransactOpts) *legendTxTypes.CancelOfferTxInfo {
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

func VerifyCancelOfferTxSig(pubKey string, tx *legendTxTypes.CancelOfferTxInfo) error {
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

func VerifyWithdrawNftTxSig(pubKey string, tx *legendTxTypes.WithdrawNftTxInfo) error {
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

func VerifyTransferNftTxSig(pubKey string, tx *legendTxTypes.TransferNftTxInfo) error {
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

func VerifyOfferTxSig(pubKey string, tx *legendTxTypes.OfferTxInfo) error {
	message, err := legendTxTypes.ComputeOfferMsgHash(tx, mimc.NewMiMC())
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

func VerifyMintNftTxSig(pubKey string, tx *legendTxTypes.MintNftTxInfo) error {
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

func VerifyCreateCollectionTxSig(pubKey string, tx *legendTxTypes.CreateCollectionTxInfo) error {
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

func VerifyAtomicMatchTxSig(pubKey string, tx *legendTxTypes.AtomicMatchTxInfo) error {
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
