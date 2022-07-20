package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var l1Endpoint = "https://data-seed-prebsc-1-s1.binance.org:8545"
var zkBASContract = "0x52fd0267B7fAd7768c39d1abC4a1D9B930deF3D8"
var l1PrivateKey = "5265077abda355eaa59a3ba1189e0bf535be155e83232b21a47792a810f2d2db"
var l1Address = "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

// Random seed
var l2KeyManager, _ = accounts.NewSeedKeyManager("30e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b")
var l2Name = "walt.legend"

func TestRegisterZNS(t *testing.T) {
	client, _ := NewZkBASL1Client(l1Endpoint, zkBASContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	pk := l2KeyManager.PubKeyPoint()
	fmt.Println(pk[0])
	fmt.Println(pk[1])
	hash, err := client.RegisterZNS("walt", common.HexToAddress(l1Address), big.NewInt(1e17), pk[0], pk[1])
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestDepositBNB(t *testing.T) {
	client, _ := NewZkBASL1Client(l1Endpoint, zkBASContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.DepositBNB("walt", big.NewInt(1e18))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestFullExitBNB(t *testing.T) {
	client, _ := NewZkBASL1Client(l1Endpoint, zkBASContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.RequestFullExit("walt", common.HexToAddress("0x0000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestTransferInLayer2(t *testing.T) {
	l2Client := getSdkClient()
	l2Client.SetKeyManager(l2KeyManager)

	txInfo := types.TransferTxInfo{
		ToAccountName: "sher.legend",
		AssetId:       0,
		AssetAmount:   big.NewInt(1e17),
	}
	hash, err := l2Client.Transfer(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Println(hash)
}
