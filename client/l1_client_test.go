package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var l1Endpoint = "https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd"
var zkbnbContract = "0x2b7Bc4406Ae01dC2f1b13b878D52431672BbCdCd"
var l1PrivateKey = "18a13221591a3873c30d87af364347f5850891c0b16155848c34ac1a0162332c"
var l1Address = "0xF2aFfafFf1929f684Bf8aAC80F6347E93590Bb7B"

var chainNetworkId uint64 = 97

func TestDepositBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	tx, err := client.DepositBNB(l1Address, big.NewInt(1e16))
	assert.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestFullExitBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	tx, err := client.RequestFullExit(2, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestFullExitNFT(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	tx, err := client.RequestFullExitNft(2, 2)
	assert.NoError(t, err)
	fmt.Println(tx.Hash())
}

func TestDepositNft(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.DepositNft(common.HexToAddress("0xBeABc8291d54eC257184B7C42Fde848166e372BB"), l1Address, big.NewInt(2))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestDepositBep20(t *testing.T) {
	assetPrivateKey := "dc3543c9c912db587693f9b27e4d221c367772cc905cbb4b76c9f30050d2534c"

	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	client.SetPrivateKey(assetPrivateKey)
	tx, err := client.DepositBEP20(common.HexToAddress("0x92AC3dBcA5AA61e43bD74ef59F5f3acd1E724730"), l1Address, big.NewInt(1000000))
	if err != nil {
		println(err.Error())
		return
	}
	println("deposit bep 20 success, tx hash=", tx.Hash())
}
