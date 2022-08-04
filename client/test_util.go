package client

import (
	"errors"
	"math/big"

	"github.com/bnb-chain/zkbas-crypto/ffmath"
)

var (
	// 2^35 - 1
	PackedAmountMaxMantissa = big.NewInt(34359738367)
	// 2^11 - 1
	PackedAmountMaxAmount = ffmath.Multiply(big.NewInt(34359738367), new(big.Int).Exp(big.NewInt(10), big.NewInt(31), nil))
	ZeroBigInt            = big.NewInt(0)
)

func ComputeEmptyLpAmount(
	assetAAmount *big.Int,
	assetBAmount *big.Int,
) (lpAmount *big.Int, err error) {
	lpSquare := ffmath.Multiply(assetAAmount, assetBAmount)
	lpFloat := ffmath.FloatSqrt(ffmath.IntToFloat(lpSquare))
	lpAmount, err = CleanPackedAmount(ffmath.FloatToInt(lpFloat))
	if err != nil {
		return nil, err
	}
	return lpAmount, nil
}

func CleanPackedAmount(amount *big.Int) (nAmount *big.Int, err error) {
	if amount.Cmp(ZeroBigInt) < 0 || amount.Cmp(PackedAmountMaxAmount) > 0 {
		return nil, errors.New("[ToPackedAmount] invalid amount")
	}
	oAmount := new(big.Int).Set(amount)
	exponent := int64(0)
	for oAmount.Cmp(PackedAmountMaxMantissa) > 0 {
		oAmount = ffmath.Div(oAmount, big.NewInt(10))
		exponent++
	}
	nAmount = ffmath.Multiply(oAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(exponent), nil))
	return nAmount, nil
}
