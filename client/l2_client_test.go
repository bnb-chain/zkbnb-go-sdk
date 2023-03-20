package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
	"github.com/bnb-chain/zkbnb-go-sdk/txutils"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/stretchr/testify/assert"
)

var testEndpoint = "http://127.0.0.1:8888"
var privateKey = l1PrivateKey

func prepareSdkClientWithPrivateKey() *l2Client {
	sdkClient, err := NewZkBNBClientWithPrivateKey(testEndpoint, privateKey, chainNetworkId)
	if err != nil {
		fmt.Errorf("error Occurred when Creating ZKBNB client! error:%s", err.Error())
		return nil
	}
	return sdkClient.(*l2Client)
}

func TestChangePubKey(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	pk := sdkClient.keyManager.PubKeyPoint()
	txInfo := &types.ChangePubKeyReq{
		L1Address: l1Address,
		PubKeyX:   pk[0],
		PubKeyY:   pk[1],
	}
	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("create ChangePubKey signature body:%s \n", signBody)

	// Generate the signature with private key and outside the ChangePubKey function
	signature, err := sdkClient.GenerateSignature(privateKey, txInfo)
	assert.NoError(t, err)

	txHash, err := sdkClient.ChangePubKey(txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("ChangePubKey success, tx_hash=%s \n", txHash)

}

func TestGetCurrentHeight(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	height, err := sdkClient.GetCurrentHeight()
	if err != nil {
		println(err.Error())
		return
	}

	println("current height: ", height)
}

func TestGetAsset(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	asset, err := sdkClient.GetAssetBySymbol("BNB")
	if err != nil {
		println(err.Error())
		return
	}

	println("bnb price: ", asset.Price)
}

func TestGetAccountNfts(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	nfts, err := sdkClient.GetNftsByAccountIndex(5, 0, 100)
	if err != nil {
		println(err.Error())
		return
	}

	println("nft total: ", nfts.Total)
	if len(nfts.Nfts) > 0 {
		println("creator: ", nfts.Nfts[0].CreatorL1Address)
		println("owner: ", nfts.Nfts[0].OwnerL1Address)
	}
}

func TestGetGasAccount(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	account, err := sdkClient.GetGasAccount()
	if err != nil {
		println(err.Error())
		return
	}

	println("gas account index: ", account.Index)
}

func TestGetNftsByAccountIndex(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
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
	sdkClient := prepareSdkClientWithPrivateKey()
	assetList, err := sdkClient.GetAssets(0, 50)
	if err != nil {
		println(err.Error())
		return
	}

	bz, _ := json.MarshalIndent(assetList, "", "  ")
	println(string(bz))
}

func TestGetTxs(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
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
	sdkClient := prepareSdkClientWithPrivateKey()
	txInfo := types.CreateCollectionTxReq{
		Name:         fmt.Sprintf("Nft Collection - my collection"),
		Introduction: "Great Nft!",
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("create collection signature body:%s \n", signBody)

	// Generate the signature with private key and outside the Create Collection function
	signature, err := sdkClient.GenerateSignature(privateKey, &txInfo)
	assert.NoError(t, err)

	txId, err := sdkClient.CreateCollection(&txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("mint nft success, tx_hash: %s \n", txId)
}

func TestGetAccountByL1Address(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	Account, err := sdkClient.GetAccountByL1Address(l1Address)
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(Account, "", "  ")
	println(string(bz))
}

func TestMintNft(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	txInfo := types.MintNftTxReq{
		To:                  l1Address,
		NftCollectionId:     0,
		NftContentType:      0,
		CreatorTreasuryRate: 0,
		MetaData:            "any information",
		MutableAttributes:   "any mutable attributes",
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("mint nft signature body:%s \n", signBody)

	// Generate the signature with private key and outside the MintNft function
	signature, err := sdkClient.GenerateSignature(privateKey, &txInfo)
	assert.NoError(t, err)

	txId, err := sdkClient.MintNft(&txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("mint nft success, tx_hash: %s \n", txId)

}

func TestGetMaxCollectionId(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	nft, err := sdkClient.GetMaxCollectionId(4)
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(nft, "", "  ")
	println(string(bz))
}

func TestGetNftByTxHash(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	nft, err := sdkClient.GetNftByTxHash("22b408110c9f376fafea6b0c5028121ed3cd389b4877e6cd7875c91288e46fa6")
	if err != nil {
		println(err.Error())
		return
	}
	bz, _ := json.MarshalIndent(nft, "", "  ")
	println(string(bz))
}

func TestUpdateNftByIndex(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()
	txInfo := types.UpdateNftReq{
		NftIndex:          1,
		MutableAttributes: "update information",
	}
	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("mint nft signature body:%s \n", signBody)

	// Generate the signature with private key and outside the updateNft function
	signature, err := sdkClient.GenerateSignature(privateKey, &txInfo)
	assert.NoError(t, err)

	assetList, err := sdkClient.UpdateNftByIndex(&txInfo, signature)
	if err != nil {
		println(err.Error())
		return
	}

	bz, _ := json.MarshalIndent(assetList, "", "  ")
	println(string(bz))
}

func TestAtomicMatchTx(t *testing.T) {

	sdkClient := prepareSdkClientWithPrivateKey()

	txInfo, err := PrepareAtomicMatchTxReq(sdkClient)
	assert.NoError(t, err)

	txId, err := sdkClient.AtomicMatch(txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("send atomic match tx success, tx_id=%s \n", txId)
}

func PrepareAtomicMatchTxReq(sdkClient *l2Client) (*types.AtomicMatchTxReq, error) {

	sellPrivateKey := "1699b28ce32c61721f02dec1fc04084fb627fa692bda8149351fe8adb95846d4"
	sellerSeed := "30ef6587bdd8211e1666613a767ec646545e47b1e66bfc90632e7c6caecb99ba3691b6ba295232cf7d7506f9996afad796895185d01a6d17740b7695e26fa4af1b"
	sellerAddress := "0x6086D0C71Bfa516FB976d9845eaB5109e10e580B"

	buyPrivateKey := "4112bf927b05689f877062212600de88d8e658d9a281f1517e09505729b1bc0d"
	buyerSeed := "ca5ca10fb553ed6d0ddc9db62aa6ae331feb7043aa6f2e52b0bbe90a3bba64292bedd8079bd9701e628dacc46aa6e7e56c67b7a4497047151de1536c038786211b"
	buyerAddress := "0xA64E522dCA2D1dA78569Ad7268096cb07A729867"

	buyer, err := sdkClient.GetAccountByL1Address(buyerAddress)
	if err != nil {
		return nil, err
	}

	seller, err := sdkClient.GetAccountByL1Address(sellerAddress)
	if err != nil {
		return nil, err
	}

	buyerOfferId, err := sdkClient.GetMaxOfferId(buyer.Index)
	if err != nil {
		return nil, err
	}

	sellerOfferId, err := sdkClient.GetMaxOfferId(seller.Index)
	if err != nil {
		return nil, err
	}

	nftIndex := int64(16)
	listedAt := time.Now().UnixMilli()
	expiredAt := time.Now().Add(time.Hour * 2).UnixMilli()
	buyOffer := &types.OfferTxInfo{
		Type:         types.BuyOfferType,
		OfferId:      int64(buyerOfferId),
		AccountIndex: buyer.Index,
		NftIndex:     nftIndex,
		AssetId:      0,
		AssetAmount:  big.NewInt(10000),
		ListedAt:     listedAt,
		ExpiredAt:    expiredAt,
		TreasuryRate: 200,
		Sig:          nil,
	}

	buyerKey, err := accounts.NewSeedKeyManager(buyerSeed)
	if err != nil {
		return nil, err
	}

	buyOfferSign, err := CalculateSignature(buyerKey, buyOffer)
	if err != nil {
		return nil, err
	}
	buyOffer.Sig = buyOfferSign

	// Generate the signature body for caller to calculate the signature
	buySignBody, err := sdkClient.GenerateSignBody(buyOffer, nil)

	fmt.Printf("create atomic match signature body:%s \n", buySignBody)

	// Generate the signature with private key and outside the Atomic Match function
	buySignature, err := sdkClient.GenerateSignature(buyPrivateKey, buyOffer)
	buyOffer.L1Sig = buySignature

	sellOffer := &types.OfferTxInfo{
		Type:         types.SellOfferType,
		OfferId:      int64(sellerOfferId),
		AccountIndex: seller.Index,
		NftIndex:     nftIndex,
		AssetId:      0,
		AssetAmount:  big.NewInt(10000),
		ListedAt:     listedAt,
		ExpiredAt:    expiredAt,
		TreasuryRate: 200,
		Sig:          nil,
	}

	sellerKey, err := accounts.NewSeedKeyManager(sellerSeed)
	if err != nil {
		return nil, err
	}

	sellOfferSign, err := CalculateSignature(sellerKey, sellOffer)
	if err != nil {
		return nil, err
	}
	sellOffer.Sig = sellOfferSign

	// Generate the signature body for caller to calculate the signature
	sellSignBody, err := sdkClient.GenerateSignBody(sellOffer, nil)
	fmt.Printf("create atomic match signature body:%s \n", sellSignBody)
	// Generate the signature with private key and outside the Atomic Match function
	sellSignature, err := sdkClient.GenerateSignature(sellPrivateKey, sellOffer)
	sellOffer.L1Sig = sellSignature

	txInfo := &types.AtomicMatchTxReq{
		BuyOffer:  buyOffer,
		SellOffer: sellOffer,
	}
	return txInfo, nil
}

func CalculateSignature(signer accounts.Signer, tx *types.OfferTxInfo) ([]byte, error) {
	convertedTx := txutils.ConvertOfferTxInfo(tx)
	err := convertedTx.Validate()
	if err != nil {
		return nil, err
	}
	hFunc := mimc.NewMiMC()
	msgHash, err := convertedTx.Hash(hFunc)
	if err != nil {
		return nil, err
	}
	hFunc.Reset()
	signature, err := signer.Sign(msgHash, hFunc)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func TestTransferNft(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	nftIndex := int64(8)
	txInfo := &types.TransferNftTxReq{
		NftIndex: nftIndex,
		To:       l1Address,
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("create transfer NFT signature body:%s \n", signBody)

	// Generate the signature with private key and outside the transferNft function
	signature, err := sdkClient.GenerateSignature(privateKey, txInfo)
	assert.NoError(t, err)

	txId, err := sdkClient.TransferNft(txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("send transfer nft tx success, tx_id=%s \n", txId)
}

func TestCancelOfferTx(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	account, err := sdkClient.GetAccountByPk(hex.EncodeToString(sdkClient.KeyManager().PubKey().Bytes()))
	assert.NoError(t, err)

	offerId, err := sdkClient.GetMaxOfferId(account.Index)
	assert.NoError(t, err)

	txInfo := types.CancelOfferTxReq{
		OfferId: int64(offerId),
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("create cancel offer signature body:%s \n", signBody)

	// Generate the signature with private key and outside the Cancel Offer function
	signature, err := sdkClient.GenerateSignature(privateKey, &txInfo)
	assert.NoError(t, err)

	txId, err := sdkClient.CancelOffer(&txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("cancel offer success, tx id: %s \n", txId)
}

func TestTransferInLayer2(t *testing.T) {
	l2Client := prepareSdkClientWithPrivateKey()

	txInfo := types.TransferTxReq{
		To:          l1Address,
		AssetId:     0,
		AssetAmount: big.NewInt(1),
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := l2Client.GenerateSignBody(&txInfo, nil)
	assert.NoError(t, err)
	fmt.Printf("create transfer signature body:%s \n", signBody)

	// Generate the signature with private key and outside the transfer function
	signature, err := l2Client.GenerateSignature(privateKey, &txInfo)
	hash, err := l2Client.Transfer(&txInfo, nil, signature)
	assert.NoError(t, err)
	fmt.Println("transfer success, tx id=", hash)
}

func TestWithdrawBNB(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawTxReq{
		AssetId:     0,
		AssetAmount: big.NewInt(100),
		ToAddress:   randomAddress,
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txReq, nil)
	assert.NoError(t, err)
	fmt.Printf("create withdraw BNB signature body:%s \n", signBody)

	// Generate the signature with private key and outside the Withdraw function
	signature, err := sdkClient.GenerateSignature(privateKey, &txReq)
	assert.NoError(t, err)

	txId, err := sdkClient.Withdraw(&txReq, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("withdraw success, tx id: %s \n", txId)
}

func TestWithdrawBEP20(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawTxReq{
		AssetId:     1,
		AssetAmount: big.NewInt(100),
		ToAddress:   randomAddress,
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txReq, nil)
	assert.NoError(t, err)
	fmt.Printf("create withdraw BEP signature body:%s \n", signBody)

	// Generate the signature with private key and outside the Withdraw function
	signature, err := sdkClient.GenerateSignature(privateKey, &txReq)
	assert.NoError(t, err)

	txId, err := sdkClient.Withdraw(&txReq, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("withdraw success, tx id: %s \n", txId)
}

func TestWithdrawNft(t *testing.T) {
	sdkClient := prepareSdkClientWithPrivateKey()

	randomAddress := "0x8b2C5A5744F42AA9269BaabDd05933a96D8EF911"

	txReq := types.WithdrawNftTxReq{
		AccountIndex: 88,
		NftIndex:     8,
		ToAddress:    randomAddress,
	}

	// Generate the signature body for caller to calculate the signature
	signBody, err := sdkClient.GenerateSignBody(&txReq, nil)
	assert.NoError(t, err)
	fmt.Printf("create withdraw NFT signature body:%s \n", signBody)

	// Generate the signature with private key and outside the WithdrawNft function
	signature, err := sdkClient.GenerateSignature(privateKey, &txReq)
	assert.NoError(t, err)

	txId, err := sdkClient.WithdrawNft(&txReq, nil, signature)
	assert.NoError(t, err)
	fmt.Printf("withdraw nft success, tx id: %s \n", txId)
}
