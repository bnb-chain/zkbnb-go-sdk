package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
)

var l1Endpoint = "https://data-seed-prebsc-1-s1.binance.org:8545"
var zkbnbContract = "0x144C8adf3443Df4Ad9738086fb7b9770D625c9aD"
var l1PrivateKey = "5a5b26d4ab5d2041b100785a3c7484e197c2346f78403faf844a43dd4be8cd34"
var l1Address = "0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896"

// Random seed
var l2KeyManager, _ = accounts.NewSeedKeyManager("30e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf243")
var l2Name = "gary.legend"

func TestRegisterZNS(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	pk := l2KeyManager.PubKeyPoint()
	fmt.Println(pk[0])
	fmt.Println(pk[1])
	hash, err := client.RegisterZNS("gary", common.HexToAddress(l1Address), big.NewInt(1e17), pk[0], pk[1])
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestDepositBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.DepositBNB("gary", big.NewInt(1e18))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestQueryMintNFT(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	//client.
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
