package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
)

var l1Endpoint = "http://127.0.0.1:8545/"
var zkbnbContract = "0x500697ae94eBd64F6288690c530595E490C1E3E0"
var l1PrivateKey = "d15ea705a4b60356e158c77d8f60e25bbbb9e5f8e6c6e077f0a5d58e24b8d1aa"
var l1Address = "0x5F6dc9EaD0EE1Ad6aB5aFD90eb3Aea55659f1012"

// Random seed
var l2KeyManager, _ = accounts.NewSeedKeyManager("30e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b")
var l2Name = "walt.legend"

func TestRegisterZNS(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
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
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.DepositBNB("walt", big.NewInt(1e18))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestFullExitBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.RequestFullExit("walt", common.HexToAddress("0x0000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestDepositBep20(t *testing.T) {
	assetPrivateKey := "dc3543c9c912db587693f9b27e4d221c367772cc905cbb4b76c9f30050d2534c"

	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	client.SetPrivateKey(assetPrivateKey)
	txHash, err := client.DepositBEP20(common.HexToAddress("0x92AC3dBcA5AA61e43bD74ef59F5f3acd1E724730"), "sher", big.NewInt(1000000))
	if err != nil {
		println(err.Error())
		return
	}
	println("deposit bep 20 success, tx hash=", txHash.String())
}
