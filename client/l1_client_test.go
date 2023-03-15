package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var l1Endpoint = "http://127.0.0.1:8545/"
var zkbnbContract = "0xF23d9b81Ce9e7b60e8E5748B4909467877a96E3c"
var l1PrivateKey = "355c102f0c8fb7efd0a2d66d70895e7cb0c4580eabc59073adb928d3e7315641"
var l1Address = "0xF792CC80193Ea942820C945F010051dE5CF6975A"

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
