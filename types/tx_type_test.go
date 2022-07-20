package types

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAddLiquidityTxInfo(t *testing.T) {
	txInfo := "{\"FromAccountIndex\":0,\"PairIndex\":0,\"AssetAId\":0,\"AssetAAmount\":10000,\"AssetBId\":0,\"AssetBAmount\":100,\"LpAmount\":995,\"KLast\":50000,\"TreasuryAmount\":3,\"GasAccountIndex\":0,\"GasFeeAssetId\":0,\"GasFeeAssetAmount\":200,\"ExpiredAt\":1654656781000,\"Nonce\":1,\"Sig\":\"UWdrVERiRXEzUHE3QWppZG9vUHlmSG1sU2ExVnVCQWdxdjU3WGpPVDd5UUM2T3pOQnY2WVFMU202VTFCbVBLQS9xekZoZnBuVkZSOGpMNjRrWC9XK2c9PQ==\"}"
	var addLiquidityTx *AddLiquidityTxInfo

	err := json.Unmarshal([]byte(txInfo), &addLiquidityTx)
	assert.Nil(t, err)
	log.Println(addLiquidityTx)
}
