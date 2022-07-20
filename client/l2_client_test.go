package client

import (
	"encoding/hex"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/txutils"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

var testEndpoint = "http://172.22.41.67:8888"
var seed = "30e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"

func getSdkClient() *l2Client {
	c := &l2Client{
		endpoint: testEndpoint,
	}
	keyManager, _ := accounts.NewSeedKeyManager(seed)
	c.SetKeyManager(keyManager)
	return c
}

func TestCreateCollection(t *testing.T) {
	sdkClient := getSdkClient()
	txInfo := &types.CreateCollectionTxInfo{
		Name:         fmt.Sprintf("Nft Collection - my collection"),
		Introduction: "Great Nft!",
	}

	collectionId, err := sdkClient.CreateCollection(txInfo, nil)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("create collection success, collection_id=%d \n", collectionId)
}

// TODO: failed currently
func TestMintNft(t *testing.T) {
	sdkClient := getSdkClient()
	bz := mimc.NewMiMC().Sum([]byte("contend_hash"))
	txInfo := &types.MintNftTxInfo{
		To:                  "walt.legend",
		NftContentHash:      hex.EncodeToString(bz),
		NftCollectionId:     1,
		CreatorTreasuryRate: 0,
	}

	nftId, err := sdkClient.MintNft(txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("mint nft success, assetId=%d \n", nftId)
}

func TestAtomicMatchTx(t *testing.T) {
	sellerSeed := "28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"
	sellerName := "sher.legend"

	buyerSeed := "17673b9a9fdec6dc90c7cc1eb1c939134dfb659d2f08edbe071e5c45f343d008"
	buyerName := "gavin.legend"

	sdkClient := getSdkClient()

	buyer, err := sdkClient.GetAccountInfoByAccountName(buyerName)
	if err != nil {
		println(err.Error())
		return
	}

	seller, err := sdkClient.GetAccountInfoByAccountName(sellerName)
	if err != nil {
		println(err.Error())
		return
	}

	buyerOfferId, err := sdkClient.GetMaxOfferId(buyer.Index)
	if err != nil {
		println(err.Error())
		return
	}

	sellerOfferId, err := sdkClient.GetMaxOfferId(seller.Index)
	if err != nil {
		println(err.Error())
		return
	}

	nftIndex := int64(0)

	txInfo := PrepareAtomicMatchInfo(buyerSeed, sellerSeed, nftIndex, int64(buyer.Index), int64(buyerOfferId), int64(seller.Index), int64(sellerOfferId), seller.Nonce)

	txId, err := sdkClient.SendRawTx(types.TxTypeAtomicMatch, txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("send atomic match tx success, tx_id=%s \n", txId)
}

// TODO, test all transaction type.

func PrepareAtomicMatchInfo(buyerSeed, sellerSeed string, nftIndex, buyerIndex, buyerOfferId, sellerIndex, sellerOfferId, sellerNonce int64) string {
	buyerKey, err := accounts.NewSeedKeyManager(buyerSeed)
	if err != nil {
		panic(err)
	}

	listedAt := time.Now().UnixMilli()
	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	buyOffer := &types.OfferTxInfo{
		Type:         types.BuyOfferType,
		OfferId:      buyerOfferId,
		AccountIndex: buyerIndex,
		NftIndex:     nftIndex,
		AssetId:      0,
		AssetAmount:  big.NewInt(10000),
		ListedAt:     listedAt,
		ExpiredAt:    expiredAt,
		TreasuryRate: 200,
		Sig:          nil,
	}

	buyTx, err := txutils.ConstructOfferTx(buyerKey, buyOffer)
	if err != nil {
		panic(err)
	}

	sellerKey, err := accounts.NewSeedKeyManager(sellerSeed)
	if err != nil {
		panic(err)
	}
	sellOffer := &types.OfferTxInfo{
		Type:         types.SellOfferType,
		OfferId:      sellerOfferId,
		AccountIndex: sellerIndex,
		NftIndex:     nftIndex,
		AssetId:      0,
		AssetAmount:  big.NewInt(10000),
		ListedAt:     listedAt,
		ExpiredAt:    expiredAt,
		TreasuryRate: 200,
		Sig:          nil,
	}

	sellTx, err := txutils.ConstructOfferTx(sellerKey, sellOffer)
	if err != nil {
		panic(err)
	}

	signedBuyOffer, _ := types.ParseOfferTxInfo(buyTx)
	signedSellOffer, _ := types.ParseOfferTxInfo(sellTx)

	txInfo := &types.AtomicMatchTxInfo{
		BuyOffer:       signedBuyOffer,
		SellOffer:      signedSellOffer,
		TreasuryAmount: big.NewInt(5000),
	}

	tx, err := txutils.ConstructAtomicMatchTx(sellerKey, txInfo, nil)
	if err != nil {
		panic(err)
	}
	return tx
}

func TestTransferNft(t *testing.T) {
	toAccountName := "gavin.legend"

	sdkClient := getSdkClient()

	nftIndex := int64(3)
	txInfo := PrepareTransferNftTxInfo(sdkClient, nftIndex, toAccountName)

	txId, err := sdkClient.SendRawTx(types.TxTypeTransferNft, txInfo)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Printf("send transfer nft tx success, tx_id=%s \n", txId)
}

func PrepareTransferNftTxInfo(c *l2Client, nftIndex int64, toAccountName string) string {

	txInfo := &types.TransferNftTxInfo{
		NftIndex: nftIndex,
		To:       toAccountName,
	}
	ops := new(types.TransactOpts)
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		panic(err)
	}

	ops, err = c.fullFillToAddrOps(ops, toAccountName)
	if err != nil {
		panic(err)
	}
	tx, err := txutils.ConstructTransferNftTx(c.KeyManager(), txInfo, ops)
	if err != nil {
		panic(err)
	}
	return tx
}

func TestCancelOfferTx(t *testing.T) {
	sdkClient := getSdkClient()

	account, err := sdkClient.GetAccountInfoByPubKey(hex.EncodeToString(sdkClient.KeyManager().PubKey().Bytes()))
	if err != nil {
		println(err.Error())
		return
	}

	offerId, err := sdkClient.GetMaxOfferId(account.AccountIndex)
	if err != nil {
		println(err.Error())
		return
	}

	txInfo := PrepareCancelOfferTxInfo(sdkClient, int64(offerId))

	txId, err := sdkClient.SendRawTx(types.TxTypeCancelOffer, txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("send cancel offer success, tx_id=%s \n", txId)
}

func PrepareCancelOfferTxInfo(c *l2Client, offerId int64) string {

	txInfo := &types.CancelOfferTxInfo{
		OfferId: offerId,
	}

	ops := new(types.TransactOpts)
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		panic(err)
	}

	tx, err := txutils.ConstructCancelOfferTx(c.keyManager, txInfo, ops)
	if err != nil {
		panic(err)
	}
	return tx
}
