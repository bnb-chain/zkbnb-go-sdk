package sdk

import (
	"encoding/json"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAddLiquidity(t *testing.T) {
	var a = AddLiquidityTxInfo{
		FromAccountIndex:  0,
		PairIndex:         0,
		AssetAId:          0,
		AssetAAmount:      big.NewInt(10000),
		AssetBId:          0,
		AssetBAmount:      big.NewInt(100),
		LpAmount:          big.NewInt(995),
		KLast:             big.NewInt(50000),
		TreasuryAmount:    big.NewInt(3),
		GasAccountIndex:   0,
		GasFeeAssetId:     0,
		GasFeeAssetAmount: big.NewInt(200),
		ExpiredAt:         1654656781000,
		Nonce:             1,
		Sig:               []byte("QgkTDbEq3Pq7AjidooPyfHmlSa1VuBAgqv57XjOT7yQC6OzNBv6YQLSm6U1BmPKA/qzFhfpnVFR8jL64kX/W+g=="),
	}

	aBytes, err := json.Marshal(a)
	assert.Nil(t, err)
	log.Println(string(aBytes))
}

func TestParseAddLiquidityTxInfo(t *testing.T) {
	txInfo := "{\"FromAccountIndex\":0,\"PairIndex\":0,\"AssetAId\":0,\"AssetAAmount\":10000,\"AssetBId\":0,\"AssetBAmount\":100,\"LpAmount\":995,\"KLast\":50000,\"TreasuryAmount\":3,\"GasAccountIndex\":0,\"GasFeeAssetId\":0,\"GasFeeAssetAmount\":200,\"ExpiredAt\":1654656781000,\"Nonce\":1,\"Sig\":\"UWdrVERiRXEzUHE3QWppZG9vUHlmSG1sU2ExVnVCQWdxdjU3WGpPVDd5UUM2T3pOQnY2WVFMU202VTFCbVBLQS9xekZoZnBuVkZSOGpMNjRrWC9XK2c9PQ==\"}"
	var addLiquidityTx *AddLiquidityTxInfo

	err := json.Unmarshal([]byte(txInfo), &addLiquidityTx)
	assert.Nil(t, err)
	log.Println(addLiquidityTx)
}
