package sdk

import (
	"encoding/hex"
	"errors"

	"github.com/bnb-chain/zkbas-crypto/wasm/legend/legendTxTypes"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
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

func VerifyTransferNftTxSig(pubKey string, tx *TransferNftTxInfo) error {
	convertedTx := legendTxTypes.TransferNftTxInfo{
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
	message, err := legendTxTypes.ComputeTransferNftMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyWithdrawNftTxSig(pubKey string, tx *WithdrawNftTxInfo) error {
	convertedTx := legendTxTypes.WithdrawNftTxInfo{
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
	message, err := legendTxTypes.ComputeWithdrawNftMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyOfferTxSig(pubKey string, offerTx *OfferTxInfo) error {
	convertedTx := legendTxTypes.OfferTxInfo{
		Type:         offerTx.Type,
		OfferId:      offerTx.OfferId,
		AccountIndex: offerTx.AccountIndex,
		NftIndex:     offerTx.NftIndex,
		AssetId:      offerTx.AssetId,
		AssetAmount:  offerTx.AssetAmount,
		ListedAt:     offerTx.ListedAt,
		ExpiredAt:    offerTx.ExpiredAt,
		TreasuryRate: offerTx.TreasuryRate,
		Sig:          offerTx.Sig,
	}
	message, err := legendTxTypes.ComputeOfferMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyMintNftTxSig(pubKey string, tx *MintNftTxInfo) error {
	convertedTx := legendTxTypes.MintNftTxInfo{
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
	message, err := legendTxTypes.ComputeMintNftMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyCreateCollectionTxSig(pubKey string, tx *CreateCollectionTxInfo) error {
	convertedTx := legendTxTypes.CreateCollectionTxInfo{
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
	message, err := legendTxTypes.ComputeCreateCollectionMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyAtomicMatchTxSig(pubKey string, tx *AtomicMatchTxInfo) error {
	convertedTx := legendTxTypes.AtomicMatchTxInfo{
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
	message, err := legendTxTypes.ComputeAtomicMatchMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}

func VerifyCancelOfferTxSig(pubKey string, tx *CancelOfferTxInfo) error {
	convertedTx := legendTxTypes.CancelOfferTxInfo{
		AccountIndex:      tx.AccountIndex,
		OfferId:           tx.OfferId,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
	message, err := legendTxTypes.ComputeCancelOfferMsgHash(&convertedTx, mimc.NewMiMC())
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
		return err
	}
	return nil
}
