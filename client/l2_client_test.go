package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
	"github.com/bnb-chain/zkbnb-go-sdk/txutils"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
)

var testEndpoint = "http://127.0.0.1:8888"
var seed = "30e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"

func getSdkClient() *l2Client {
	c := &l2Client{
		endpoint: testEndpoint,
	}
	keyManager, _ := accounts.NewSeedKeyManager(seed)
	c.SetKeyManager(keyManager)
	return c
}

func TestGetCurrentHeight(t *testing.T) {
	sdkClient := getSdkClient()
	height, err := sdkClient.GetCurrentHeight()
	if err != nil {
		println(err.Error())
		return
	}

	println("current height: ", height)
}

func TestGetAsset(t *testing.T) {
	sdkClient := getSdkClient()
	asset, err := sdkClient.GetAssetBySymbol("BNB")
	if err != nil {
		println(err.Error())
		return
	}

	println("bnb price: ", asset.Price)
}

func TestGetAccountNfts(t *testing.T) {
	sdkClient := getSdkClient()
	nfts, err := sdkClient.GetNftsByAccountIndex(5, 0, 100)
	if err != nil {
		println(err.Error())
		return
	}

	println("nft total: ", nfts.Total)
	if len(nfts.Nfts) > 0 {
		println("creator: ", nfts.Nfts[0].CreatorAccountName)
		println("owner: ", nfts.Nfts[0].OwnerAccountName)
	}
}

func TestGetGasAccount(t *testing.T) {
	sdkClient := getSdkClient()
	account, err := sdkClient.GetGasAccount()
	if err != nil {
		println(err.Error())
		return
	}

	println("gas account index: ", account.Index)
}

func TestGetNftsByAccountIndex(t *testing.T) {
	sdkClient := getSdkClient()
	account, err := sdkClient.GetNftsByAccountIndex(2, 0, 10)
	if err != nil {
		println(err.Error())
		return
	}

	println("account total nft count: ", account.Total)
	bz, _ := json.MarshalIndent(account.Nfts, "", "  ")
	println(string(bz))
}

func TestGetAssets(t *testing.T) {
	sdkClient := getSdkClient()
	assetList, err := sdkClient.GetAssets(0, 50)
	if err != nil {
		println(err.Error())
		return
	}

	bz, _ := json.MarshalIndent(assetList, "", "  ")
	println(string(bz))
}

func TestGetTxs(t *testing.T) {
	sdkClient := getSdkClient()
	total, txList, err := sdkClient.GetTxs(0, 10)
	if err != nil {
		println(err.Error())
		return
	}

	bz, _ := json.MarshalIndent(txList, "", "  ")
	println(total)
	println(string(bz))
}

func TestCreateCollection(t *testing.T) {
	sdkClient := getSdkClient()
	txInfo := &types.CreateCollectionReq{
		Name:         fmt.Sprintf("Nft Collection - my collection"),
		Introduction: "Great Nft!",
	}
	txHash, err := sdkClient.CreateCollection(txInfo, nil)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("create collection success, tx_hash=%s \n", txHash)
}

func TestGetAccountByName(t *testing.T) {
	sdkClient := getSdkClient()
	Account, err := sdkClient.GetAccountByName("walt.zkbnb")
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(Account, "", "  ")
	println(string(bz))
}

func TestMintNft(t *testing.T) {
	sdkClient := getSdkClient()
	txInfo := &types.MintNftTxReq{
		To:                  "walt.zkbnb",
		NftCollectionId:     0,
		CreatorTreasuryRate: 0,
		MetaData:            "any information",
		MutableAttributes:   "any mutable attributes",
	}
	txHash, err := sdkClient.MintNft(txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("mint nft success, tx_hash=%s \n", txHash)
}

func TestGetMaxCollectionId(t *testing.T) {
	sdkClient := getSdkClient()
	nft, err := sdkClient.GetMaxCollectionId(4)
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(nft, "", "  ")
	println(string(bz))
}

func TestGetNftByTxHash(t *testing.T) {
	sdkClient := getSdkClient()
	nft, err := sdkClient.GetNftByTxHash("22b408110c9f376fafea6b0c5028121ed3cd389b4877e6cd7875c91288e46fa6")
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(nft, "", "  ")
	println(string(bz))
}

func TestUpdateNftByIndex(t *testing.T) {
	sdkClient := getSdkClient()
	assetList, err := sdkClient.UpdateNftByIndex(&types.UpdateNftReq{
		NftIndex:          1,
		MutableAttributes: "update information",
	})
	if err != nil {
		println(err.Error())
		return
	}

	bz, _ := json.MarshalIndent(assetList, "", "  ")
	println(string(bz))
}

func TestAtomicMatchTx(t *testing.T) {
	sellerSeed := "28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"
	sellerName := "sher.zkbnb"

	buyerSeed := "17673b9a9fdec6dc90c7cc1eb1c939134dfb659d2f08edbe071e5c45f343d008"
	buyerName := "gavin.zkbnb"

	sdkClient := getSdkClient()

	buyer, err := sdkClient.GetAccountByName(buyerName)
	if err != nil {
		println(err.Error())
		return
	}

	seller, err := sdkClient.GetAccountByName(sellerName)
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

	txInfo := PrepareAtomicMatchInfo(sdkClient, buyerSeed, sellerSeed, nftIndex, buyer.Index, int64(buyerOfferId), seller.Index, int64(sellerOfferId))

	txId, err := sdkClient.SendRawTx(types.TxTypeAtomicMatch, txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("send atomic match tx success, tx_id=%s \n", txId)
}

func PrepareAtomicMatchInfo(c *l2Client, buyerSeed, sellerSeed string, nftIndex, buyerIndex, buyerOfferId, sellerIndex, sellerOfferId int64) string {
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

	txInfo := &types.AtomicMatchTxReq{
		BuyOffer:  signedBuyOffer,
		SellOffer: signedSellOffer,
	}

	ops := new(types.TransactOpts)
	ops, err = c.fullFillDefaultOps(ops)
	if err != nil {
		panic(err)
	}

	tx, err := txutils.ConstructAtomicMatchTx(sellerKey, txInfo, ops)
	if err != nil {
		panic(err)
	}
	return tx
}

func TestTransferNft(t *testing.T) {
	toAccountName := "gavin.zkbnb"

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
	txInfo := &types.TransferNftTxReq{
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

	account, err := sdkClient.GetAccountByPk(hex.EncodeToString(sdkClient.KeyManager().PubKey().Bytes()))
	if err != nil {
		println(err.Error())
		return
	}

	offerId, err := sdkClient.GetMaxOfferId(account.Index)
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
	txInfo := &types.CancelOfferReq{
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

func TestTransferInLayer2(t *testing.T) {
	l2Client := getSdkClient()

	txInfo := types.TransferTxReq{
		ToAccountName: "sher.zkbnb",
		AssetId:       0,
		AssetAmount:   big.NewInt(1),
	}
	hash, err := l2Client.Transfer(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Println("transfer success, tx id=", hash)
}

func TestWithdrawBNB(t *testing.T) {
	sdkClient := getSdkClient()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawReq{
		AssetId:     0,
		AssetAmount: big.NewInt(100),
		ToAddress:   randomAddress,
	}

	txId, err := sdkClient.Withdraw(&txReq, nil)
	if err != nil {
		println(err.Error())
		return
	}
	println("withdraw success, tx id: ", txId)
}

func TestWithdrawBEP20(t *testing.T) {
	sdkClient := getSdkClient()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawReq{
		AssetId:     1,
		AssetAmount: big.NewInt(100),
		ToAddress:   randomAddress,
	}

	txId, err := sdkClient.Withdraw(&txReq, nil)
	if err != nil {
		println(err.Error())
		return
	}
	println("withdraw success, tx id: ", txId)
}

func TestWithdrawNft(t *testing.T) {
	sdkClient := getSdkClient()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawNftTxReq{
		AccountIndex: 4,
		NftIndex:     17,
		ToAddress:    randomAddress,
	}

	txId, err := sdkClient.WithdrawNft(&txReq, nil)
	if err != nil {
		println(err.Error())
		return
	}
	println("withdraw nft success, tx id: ", txId)
}
