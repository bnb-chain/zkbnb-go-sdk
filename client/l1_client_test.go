package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var l1Endpoint = "http://127.0.0.1:8545/"
var zkbnbContract = "0xC839FE71eA874Ce76fDf3C173fE18Ad724f77686"
var l1PrivateKey = "ac106f1fb2ca7cc6cc7b743b252b5681da468f7da45e310dc3be3e0ddcf8513d"
var l1Address = "0xCEbE78C663561624551Ac37C8d0333bB2F71a635"

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
	println("deposit bep 20 success, tx hash=%s", tx.Hash().String())
}
