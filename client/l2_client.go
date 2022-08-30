package client

import (
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/txutils"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

const defaultExpireTime = time.Minute * 10

var (
	dialer = &net.Dialer{
		Timeout:   1 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	transport = &http.Transport{
		DialContext:         dialer.DialContext,
		MaxConnsPerHost:     1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	HttpClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
)

type l2Client struct {
	endpoint   string
	keyManager accounts.KeyManager
}

func (c *l2Client) SetKeyManager(keyManager accounts.KeyManager) {
	c.keyManager = keyManager
}

func (c *l2Client) KeyManager() accounts.KeyManager {
	return c.keyManager
}

func (c *l2Client) GetTxsByAccountPk(accountPk string, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/accountTxs?by=account_pk&value=%s&offset=%d&limit=%d",
			accountPk, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/accountTxs?by=account_name&value=%s&offset=%d&limit=%d",
			accountName, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/txs?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/accountTxs?by=account_index&value=%d&offset=%d&limit=%d", accountIndex, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) Search(keyword string) (*types.Search, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/search?keyword=%s", keyword))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Search{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccounts(offset, limit uint32) (*types.Accounts, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/accounts?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Accounts{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetGasFeeAssets() (*types.GasFeeAssets, error) {
	resp, err := HttpClient.Get(c.endpoint + "/api/v1/gasFeeAssets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.GasFeeAssets{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (*big.Int, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/withdrawGasFee?asset_id=%d&withdraw_asset_id=%d&withdraw_amount=%d",
			assetId, withdrawAssetId, withdrawAmount))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.GasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	var price big.Int
	price.SetString(result.GasFee, 10)
	return &price, nil
}

func (c *l2Client) GetGasFee(assetId int64) (*big.Int, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/gasFee?asset_id=%d", assetId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.GasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	var price big.Int
	price.SetString(result.GasFee, 10)
	return &price, nil
}

func (c *l2Client) GetAssets(offset, limit uint32) (*types.Assets, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/assets?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Assets{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetLayer2BasicInfo() (*types.Layer2BasicInfo, error) {
	resp, err := HttpClient.Get(c.endpoint + "/api/v1/layer2BasicInfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Layer2BasicInfo{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetBlockByCommitment(blockCommitment string) (*types.Block, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/block?by=commitment&value=%s", blockCommitment))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Block{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccountByIndex(accountIndex int64) (*types.Account, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account?by=index&value=%d", accountIndex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Account{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccountByPk(accountPk string) (*types.Account, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account?by=pk&value=%s", accountPk))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Account{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetCurrencyPrice(symbol string) (*types.CurrencyPrice, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/currencyPrice?symbol=%s", symbol))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.CurrencyPrice{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetCurrencyPrices(offset, limit uint32) (*types.CurrencyPrices, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/currencyPrices?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.CurrencyPrices{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetSwapAmount(pairIndex, assetId int64, assetAmount string, isFrom bool) (*types.SwapAmount, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/swapAmount?pair_index=%d&asset_id=%d&asset_amount=%s&is_from=%v",
			pairIndex, assetId, assetAmount, isFrom))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.SwapAmount{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetPairs(offset, limit uint32) (*types.Pairs, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pairs?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Pairs{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetLpValue(pairIndex uint32, lpAmount string) (*types.LpValue, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/lpValue?pair_index=%d&lp_amount=%s", pairIndex, lpAmount))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.LpValue{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetPair(index uint32) (*types.Pair, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pair?index=%d", index))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Pair{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetTx(hash string) (*types.EnrichedTx, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx?hash=%s", hash))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	txResp := &types.EnrichedTx{}
	if err := json.Unmarshal(body, txResp); err != nil {
		return nil, err
	}
	return txResp, nil
}

func (c *l2Client) GetMempoolTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/mempoolTxs?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &types.MempoolTxs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}

func (c *l2Client) GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint + "/api/v1/accountMempoolTxs?by=account_name&value=" + accountName)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &types.MempoolTxs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}

func (c *l2Client) GetAccountByName(accountName string) (*types.Account, error) {
	resp, err := HttpClient.Get(c.endpoint + "/api/v1/account?by=name&value=" + accountName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	account := &types.Account{}
	if err := json.Unmarshal(body, account); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *l2Client) GetNextNonce(accountIdx int64) (int64, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/nextNonce?account_index=%d", accountIdx))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &types.NextNonce{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return int64(result.Nonce), nil
}

func (c *l2Client) GetTxsByBlockHeight(blockHeight uint32) ([]*types.Tx, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/blockTxs?by=block_height&value=%d", blockHeight))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result.Txs, nil
}

func (c *l2Client) GetMaxOfferId(accountIndex int64) (uint64, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/maxOfferId?account_index=%d", accountIndex))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &types.MaxOfferId{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return result.OfferId, nil
}

func (c *l2Client) GetBlockByHeight(blockHeight int64) (*types.Block, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/block?by=height&value=%d", blockHeight))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	res := &types.Block{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) GetBlocks(offset, limit int64) (uint32, []*types.Block, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/blocks?limit=%d&offset=%d", limit, offset))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	res := &types.Blocks{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, nil, err
	}
	return res.Total, res.Blocks, nil
}

func (c *l2Client) GetGasAccount() (*types.GasAccount, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/gasAccount"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	res := &types.GasAccount{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) GetNftsByAccountIndex(accountIndex, offset, limit int64) (*types.Nfts, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/accountNfts?by=account_index&value=%d&limit=%d&offset=%d", accountIndex, limit, offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	res := &types.Nfts{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) SendRawTx(txType uint32, txInfo string) (string, error) {
	resp, err := HttpClient.PostForm(c.endpoint+"/api/v1/sendTx",
		url.Values{"tx_type": {strconv.Itoa(int(txType))}, "tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &types.TxHash{}
	if err := json.Unmarshal(body, res); err != nil {
		return "", err
	}
	return res.TxHash, nil
}

func (c *l2Client) MintNft(tx *types.MintNftTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	ops, err = c.fullFillToAddrOps(ops, tx.To)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructMintNftTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeMintNft, txInfo)
}

func (c *l2Client) CreateCollection(tx *types.CreateCollectionReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructCreateCollectionTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeCreateCollection, txInfo)
}

func (c *l2Client) CancelOffer(tx *types.CancelOfferReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}
	txInfo, err := txutils.ConstructCancelOfferTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeCancelOffer, txInfo)
}

func (c *l2Client) AtomicMatch(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}
	txInfo, err := txutils.ConstructAtomicMatchTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(types.TxTypeAtomicMatch, txInfo)
}

func (c *l2Client) WithdrawNft(tx *types.WithdrawNftTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructWithdrawNftTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeWithdrawNft, txInfo)
}

func (c *l2Client) TransferNft(tx *types.TransferNftTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}
	ops, err = c.fullFillToAddrOps(ops, tx.To)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructTransferNftTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeTransferNft, txInfo)
}

func (c *l2Client) Withdraw(tx *types.WithdrawReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructWithdrawTxInfo(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeWithdraw, txInfo)
}

func (c *l2Client) RemoveLiquidity(tx *types.RemoveLiquidityReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructRemoveLiquidityTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeRemoveLiquidity, txInfo)
}

func (c *l2Client) AddLiquidity(tx *types.AddLiquidityReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}

	txInfo, err := txutils.ConstructAddLiquidityTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeAddLiquidity, txInfo)
}

func (c *l2Client) Swap(tx *types.SwapTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}
	txInfo, err := txutils.ConstructSwapTx(c.keyManager, tx, ops)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeSwap, txInfo)
}

func (c *l2Client) Transfer(tx *types.TransferTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return "", err
	}
	ops, err = c.fullFillToAddrOps(ops, tx.ToAccountName)
	if err != nil {
		return "", err
	}
	txInfo, err := txutils.ConstructTransferTx(c.keyManager, ops, tx)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(types.TxTypeTransfer, txInfo)
}

func (c *l2Client) fullFillToAddrOps(ops *types.TransactOpts, to string) (*types.TransactOpts, error) {
	toAccount, err := c.GetAccountByName(to)
	if err != nil {
		return nil, err
	}
	toAccountNameHash, err := txutils.AccountNameHash(to)
	if err != nil {
		return nil, err
	}
	ops.ToAccountIndex = toAccount.Index
	ops.ToAccountNameHash = toAccountNameHash
	return ops, nil
}

func (c *l2Client) fullFillDefaultOps(ops *types.TransactOpts) (*types.TransactOpts, error) {
	if ops == nil {
		ops = new(types.TransactOpts)
	}
	if ops.GasAccountIndex == 0 {
		gasAccount, err := c.GetGasAccount()
		if err != nil {
			return nil, err
		}
		if gasAccount.Index == 0 {
			return nil, fmt.Errorf("get gas account error, gas account index is %d", gasAccount.Index)
		}
		ops.GasAccountIndex = gasAccount.Index
	}
	if ops.ExpiredAt == 0 {
		ops.ExpiredAt = time.Now().Add(defaultExpireTime).UnixMilli()
	}
	if ops.FromAccountIndex == 0 {
		l2Account, err := c.GetAccountByPk(hex.EncodeToString(c.keyManager.PubKey().Bytes()))
		if err != nil {
			return nil, err
		}
		ops.FromAccountIndex = l2Account.Index
	}
	if ops.Nonce == 0 {
		nonce, err := c.GetNextNonce(ops.FromAccountIndex)
		if err != nil {
			return nil, err
		}
		ops.Nonce = nonce
	}
	if len(ops.CallDataHash) == 0 {
		hFunc := mimc.NewMiMC()
		ops.CallDataHash = hFunc.Sum([]byte(ops.CallData))
	}
	if ops.GasFeeAssetAmount == nil {
		// TODO, need change when it is a withdraw tx
		gas, err := c.GetGasFee(ops.GasFeeAssetId)
		if err != nil {
			return nil, err
		}
		ops.GasFeeAssetAmount = gas
	}
	return ops, nil
}
