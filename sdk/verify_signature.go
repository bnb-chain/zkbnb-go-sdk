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
