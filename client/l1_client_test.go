package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var l1Endpoint = "https://data-seed-prebsc-1-s1.binance.org:8545"
var zkbnbContract = "0x308fC6afE1A0738C8BAD2cAf5255c47A051e000e"
var l1PrivateKey = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"
var l1Address = "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

var chainNetworkId uint64 = 97

func TestDepositBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.DepositBNB(l1Address, big.NewInt(1e18))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestFullExitBNB(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.RequestFullExit(2, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestFullExitNFT(t *testing.T) {
	client, _ := NewZkBNBL1Client(l1Endpoint, zkbnbContract)
	err := client.SetPrivateKey(l1PrivateKey)
	assert.NoError(t, err)
	hash, err := client.RequestFullExitNft(2, 2)
	assert.NoError(t, err)
	fmt.Println(hash)
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
	txHash, err := client.DepositBEP20(common.HexToAddress("0x92AC3dBcA5AA61e43bD74ef59F5f3acd1E724730"), l1Address, big.NewInt(1000000))
	if err != nil {
		println(err.Error())
		return
	}
	println("deposit bep 20 success, tx hash=", txHash.String())
}
