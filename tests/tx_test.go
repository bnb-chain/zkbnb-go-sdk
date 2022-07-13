package tests

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/bnb-chain/zkbas-go-sdk/sdk"
)

func getSdkClient() sdk.ZkbasSDK {
	return sdk.NewZkbasSDK("http://172.22.41.148:8888")
}

func keccakHash(value []byte) []byte {
	hashVal := crypto.Keccak256Hash(value)
	return hashVal[:]
}

func accountNameHash(accountName string) (res string, err error) {
	words := strings.Split(accountName, ".")
	if len(words) != 2 {
		return "", errors.New("[AccountNameHash] invalid account name")
	}
	buf := make([]byte, 32)
	label := keccakHash([]byte(words[0]))
	res = common.Bytes2Hex(
		keccakHash(append(
			keccakHash(append(buf,
				keccakHash([]byte(words[1]))...)), label...)))
	return res, nil
}

func TestCreateCollection(t *testing.T) {
	keyManager, err := sdk.NewSeedKeyManager("28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b")
	if err != nil {
		println("new key manager error")
		return
	}

	sdkClient := getSdkClient()
	sdkClient.SetKeyManager(keyManager)

	accountName := "sher.legend"
	account, err := sdkClient.GetAccountInfoByAccountName(accountName)
	if err != nil {
		panic(err)
	}

	nonce, err := sdkClient.GetNextNonce(int64(account.Index))
	if err != nil {
		println(err.Error())
		return
	}

	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	txInfo := &sdk.CreateCollectionTxInfo{
		AccountIndex:      int64(account.Index),
		Name:              fmt.Sprintf("Nft Collection - %d", nonce),
		Introduction:      "Great Nft!",
		GasAccountIndex:   1,
		GasFeeAssetId:     2,
		GasFeeAssetAmount: big.NewInt(5000),
		ExpiredAt:         expiredAt,
		Nonce:             nonce,
	}

	collectionId, err := sdkClient.SignAndSendCreateCollectionTx(txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("create collection success, collection_id=%d \n", collectionId)
}

func TestMintNft(t *testing.T) {
	keyManager, err := sdk.NewSeedKeyManager("28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b")
	if err != nil {
		println("new key manager error")
		return
	}

	sdkClient := getSdkClient()
	sdkClient.SetKeyManager(keyManager)

	accountName := "sher.legend"
	account, err := sdkClient.GetAccountInfoByAccountName(accountName)
	if err != nil {
		panic(err)
	}

	accountIndex := int64(account.Index)
	nonce, err := sdkClient.GetNextNonce(accountIndex)
	if err != nil {
		println(err.Error())
		return
	}

	nameHash, err := accountNameHash(accountName)
	if err != nil {
		panic(err)
	}
	fmt.Println("nameHash ", nameHash)

	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	txInfo := &sdk.MintNftTxInfo{
		CreatorAccountIndex: accountIndex,
		ToAccountIndex:      accountIndex,
		ToAccountNameHash:   nameHash,
		NftContentHash:      "content_hash",
		NftCollectionId:     1,
		CreatorTreasuryRate: 0,
		GasAccountIndex:     1,
		GasFeeAssetId:       2,
		GasFeeAssetAmount:   big.NewInt(5000),
		ExpiredAt:           expiredAt,
		Nonce:               nonce,
	}

	nftId, err := sdkClient.SignAndSendMintNftTx(txInfo)
	if err != nil {
		println(err.Error())
		return
	}

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

	txId, err := sdkClient.SendTx(sdk.TxTypeAtomicMatch, txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("send atomic match tx success, tx_id=%s \n", txId)
}

func PrepareAtomicMatchInfo(buyerSeed, sellerSeed string, nftIndex, buyerIndex, buyerOfferId, sellerIndex, sellerOfferId, sellerNonce int64) string {
	buyerKey, err := sdk.NewSeedKeyManager(buyerSeed)
	if err != nil {
		panic(err)
	}

	listedAt := time.Now().UnixMilli()
	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	buyOffer := &sdk.OfferTxInfo{
		Type:         sdk.BuyOfferType,
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

	buyTx, err := sdk.ConstructOfferTx(buyerKey, buyOffer)
	if err != nil {
		panic(err)
	}

	sellerKey, err := sdk.NewSeedKeyManager(sellerSeed)
	if err != nil {
		panic(err)
	}
	sellOffer := &sdk.OfferTxInfo{
		Type:         sdk.SellOfferType,
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

	sellTx, err := sdk.ConstructOfferTx(sellerKey, sellOffer)
	if err != nil {
		panic(err)
	}

	signedBuyOffer, _ := sdk.ParseOfferTxInfo(buyTx)
	signedSellOffer, _ := sdk.ParseOfferTxInfo(sellTx)

	txInfo := &sdk.AtomicMatchTxInfo{
		AccountIndex:      sellerIndex,
		BuyOffer:          signedBuyOffer,
		SellOffer:         signedSellOffer,
		GasAccountIndex:   1,
		GasFeeAssetId:     0,
		GasFeeAssetAmount: big.NewInt(5000),
		TreasuryAmount:    big.NewInt(5000),
		Nonce:             sellerNonce,
		ExpiredAt:         expiredAt,
		Sig:               nil,
	}

	tx, err := sdk.ConstructAtomicMatchTx(sellerKey, txInfo)
	if err != nil {
		panic(err)
	}
	return tx
}

func TestTransferNft(t *testing.T) {
	accountSeed := "28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"
	accountName := "sher.legend"

	toAccountIndex := 3
	toAccountName := "gavin.legend"

	sdkClient := getSdkClient()

	account, err := sdkClient.GetAccountInfoByAccountName(accountName)
	if err != nil {
		panic(err)
	}

	nonce, err := sdkClient.GetNextNonce(int64(account.Index))
	if err != nil {
		panic(err)
	}

	nftIndex := int64(3)
	txInfo := PrepareTransferNftTxInfo(accountSeed, int64(account.Index), nonce, nftIndex, toAccountName, int64(toAccountIndex))

	txId, err := sdkClient.SendTx(sdk.TxTypeTransferNft, txInfo)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Printf("send transfer nft tx success, tx_id=%s \n", txId)
}

func PrepareTransferNftTxInfo(seed string, accountIndex, accountNonce, nftIndex int64, toAccountName string, toAccountIndex int64) string {
	key, err := sdk.NewSeedKeyManager(seed)
	if err != nil {
		panic(err)
	}

	nameHash, err := accountNameHash(toAccountName)
	if err != nil {
		panic(err)
	}
	fmt.Println("nameHash ", nameHash)

	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	txInfo := &sdk.TransferNftTxInfo{
		FromAccountIndex:  accountIndex,
		ToAccountIndex:    toAccountIndex,
		ToAccountNameHash: nameHash,
		NftIndex:          nftIndex,
		GasAccountIndex:   1,
		GasFeeAssetId:     2,
		GasFeeAssetAmount: big.NewInt(5000),
		ExpiredAt:         expiredAt,
		Nonce:             accountNonce,
		CallData:          "",
		CallDataHash:      nil,
		Sig:               nil,
	}

	tx, err := sdk.ConstructTransferNftTx(key, txInfo)
	if err != nil {
		panic(err)
	}
	return tx
}

func TestCancelOfferTx(t *testing.T) {
	accountSeed := "28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b"
	accountName := "sher.legend"

	sdkClient := getSdkClient()

	account, err := sdkClient.GetAccountInfoByAccountName(accountName)
	if err != nil {
		println(err.Error())
		return
	}

	offerId, err := sdkClient.GetMaxOfferId(account.Index)
	if err != nil {
		println(err.Error())
		return
	}

	txInfo := PrepareCancelOfferTxInfo(accountSeed, int64(account.Index), account.Nonce, int64(offerId))

	txId, err := sdkClient.SendTx(sdk.TxTypeCancelOffer, txInfo)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("send cancel offer success, tx_id=%s \n", txId)
}

func PrepareCancelOfferTxInfo(seed string, accountIndex, accountNonce, offerId int64) string {
	key, err := sdk.NewSeedKeyManager(seed)
	if err != nil {
		panic(err)
	}

	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	txInfo := &sdk.CancelOfferTxInfo{
		AccountIndex:      accountIndex,
		OfferId:           offerId,
		GasAccountIndex:   1,
		GasFeeAssetId:     2,
		GasFeeAssetAmount: big.NewInt(5000),
		ExpiredAt:         expiredAt,
		Nonce:             accountNonce,
		Sig:               nil,
	}

	tx, err := sdk.ConstructCancelOfferTx(key, txInfo)
	if err != nil {
		panic(err)
	}
	return tx
}
