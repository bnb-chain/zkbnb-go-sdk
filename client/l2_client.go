package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bnb-chain/zkbnb-crypto/wasm/txtypes"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
	"github.com/bnb-chain/zkbnb-go-sdk/signer"
	"github.com/bnb-chain/zkbnb-go-sdk/txutils"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
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
	endpoint    string
	privateKey  string
	address     string
	chainId     uint64
	channelName string
	l1Signer    signer.L1Signer
	keyManager  accounts.KeyManager
}

func (c *l2Client) KeyManager() accounts.KeyManager {
	return c.keyManager
}

func (c *l2Client) GetCurrentHeight() (int64, error) {
	resp, err := HttpClient.Get(c.endpoint + fmt.Sprintf("/api/v1/currentHeight"))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf(string(body))
	}
	if err = c.parseResultStatus(body); err != nil {
		return -1, err
	}
	result := &types.CurrentHeight{}
	if err := json.Unmarshal(body, result); err != nil {
		return -1, err
	}
	return result.Height, nil
}

func (c *l2Client) GetTxsByL1Address(l1Address string, offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error) {
	opt := &getTxOption{}
	for _, f := range options {
		f(opt)
	}

	path := fmt.Sprintf("/api/v1/accountTxs?by=l1_address&value=%s&offset=%d&limit=%d", l1Address, offset, limit)
	if len(opt.Types) > 0 {
		txTypes, _ := json.Marshal(opt.Types)
		path += fmt.Sprintf("&types=%s", string(txTypes))
	}

	resp, err := HttpClient.Get(c.endpoint + path)
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
	}
	result := &types.Txs{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsByAccountIndex(accountIndex int64, offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error) {
	opt := &getTxOption{}
	for _, f := range options {
		f(opt)
	}

	path := fmt.Sprintf("/api/v1/accountTxs?by=account_index&value=%d&offset=%d&limit=%d", accountIndex, offset, limit)
	if len(opt.Types) > 0 {
		txTypes, _ := json.Marshal(opt.Types)
		path += fmt.Sprintf("&types=%s", string(txTypes))
	}

	resp, err := HttpClient.Get(c.endpoint + path)
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.GasFeeAssets{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetGasFee(assetId int64, txType int) (*big.Int, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/gasFee?asset_id=%d&tx_type=%d", assetId, txType))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.GasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	var price big.Int
	price.SetString(result.GasFee, 10)
	return &price, nil
}

func (c *l2Client) GetAssetById(id uint32) (*types.Asset, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/asset?by=id&value=%d", id))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.Asset{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAssetBySymbol(symbol string) (*types.Asset, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/asset?by=symbol&value=%s", symbol))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.Asset{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.Assets{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetProtocolRate() (int64, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/getProtocolRate"))
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, err
	}
	result := &types.ProtocolRate{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	platformFeeRate, err := strconv.ParseInt(result.ProtocolRate, 10, 64)
	if err != nil {
		return 0, err
	}
	return platformFeeRate, nil
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.Layer2BasicInfo{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetRollbacks(fromBlockHeight, offset, limit int64) (total uint32, rollbacks []*types.Rollback, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/rollbacks?from_block_height=%d&limit=%d&offset=%d", fromBlockHeight, limit, offset))
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
	}
	result := &types.Rollbacks{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Rollbacks, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.Account{}
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	txResp := &types.EnrichedTx{}
	if err := json.Unmarshal(body, txResp); err != nil {
		return nil, err
	}
	return txResp, nil
}

func (c *l2Client) GetPendingTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pendingTxs?offset=%d&limit=%d", offset, limit))
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
	}
	txsResp := &types.Txs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *l2Client) GetPendingTxsByL1Address(l1Address string, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error) {
	opt := &getTxOption{}
	for _, f := range options {
		f(opt)
	}

	path := "/api/v1/accountPendingTxs?by=l1_address&value=" + l1Address
	if len(opt.Types) > 0 {
		txTypes, _ := json.Marshal(opt.Types)
		path += fmt.Sprintf("&types=%s", string(txTypes))
	}

	resp, err := HttpClient.Get(c.endpoint + path)
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
	}
	txsResp := &types.Txs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *l2Client) GetExecutedTxs(offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error) {
	opt := &getTxOption{}
	for _, f := range options {
		f(opt)
	}

	path := fmt.Sprintf("/api/v1/executedTxs?offset=%d&limit=%d", offset, limit)
	if len(opt.FromHash) > 0 {
		path += fmt.Sprintf("&from_hash=%s", opt.FromHash)
	}

	resp, err := HttpClient.Get(c.endpoint + path)
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
	}
	txsResp := &types.Txs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *l2Client) GetAccountByL1Address(l1Address string) (*types.Account, error) {
	resp, err := HttpClient.Get(c.endpoint + "/api/v1/account?by=l1_address&value=" + l1Address)
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	res := &types.Nfts{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) GetNftByNftIndex(nftIndex int64) (*types.Nft, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/GetNftByNftIndex?nft_index=%d", nftIndex))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	res := &types.NftEntity{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res.Nft, nil
}

func (c *l2Client) getL2SignatureBody(txType uint32, txInfo string) (string, error) {
	resp, err := HttpClient.PostForm(c.endpoint+"/api/v1/l2Signature",
		url.Values{"tx_type": {strconv.Itoa(int(txType))}, "tx_info": {txInfo}, "tx_signature": {"-"}})
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
	if err = c.parseResultStatus(body); err != nil {
		return "", err
	}
	res := &types.SignBody{}
	if err := json.Unmarshal(body, res); err != nil {
		return "", err
	}
	return res.SignBody, nil
}

func (c *l2Client) GetMaxCollectionId(accountIndex int64) (*types.MaxCollectionId, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/maxCollectionId?account_index=%d", accountIndex))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.MaxCollectionId{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetNftByTxHash(txHash string) (*types.NftIndex, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/getNftByTxHash?tx_hash=%s", txHash))
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	result := &types.NftIndex{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) UpdateNftByIndex(nft *types.UpdateNftReq, signatureList ...string) (*types.Mutable, error) {
	txInfo, err := c.constructUpdateNFTTransaction(nft, nil)
	if err != nil {
		return nil, err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return nil, err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return nil, err
	}
	resp, err := HttpClient.PostForm(c.endpoint+"/api/v1/updateNftByIndex",
		url.Values{"tx_info": {string(txInfoBytes)}})
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
	if err = c.parseResultStatus(body); err != nil {
		return nil, err
	}
	res := &types.Mutable{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) GetNftNextNonce(nftIndex int64) (int64, error) {
	resp, err := HttpClient.Get(c.endpoint +
		fmt.Sprintf("/api/v1/nftNextNonce?nft_index=%d", nftIndex))
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
	if err = c.parseResultStatus(body); err != nil {
		return 0, err
	}
	result := &types.NextNonce{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return int64(result.Nonce), nil
}

func (c *l2Client) SendRawTx(txType uint32, txInfo string) (string, error) {
	data := url.Values{"tx_type": {strconv.Itoa(int(txType))}, "tx_info": {txInfo}}
	req, _ := http.NewRequest("POST", c.endpoint+"/api/v1/sendTx", strings.NewReader(data.Encode()))
	req.Header.Set("Channel-Name", c.channelName)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := HttpClient.Do(req)
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
	if err = c.parseResultStatus(body); err != nil {
		return "", err
	}
	res := &types.TxHash{}
	if err := json.Unmarshal(body, res); err != nil {
		return "", err
	}
	return res.TxHash, nil
}

func (c *l2Client) ChangePubKey(tx *types.ChangePubKeyReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}
	txInfo, err := c.constructChangePubKeyTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) MintNft(tx *types.MintNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructMintNftTransaction(tx, ops)
	if err != nil {
		return "", err
	}

	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) CreateCollection(tx *types.CreateCollectionTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructCreateCollectionTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) CancelOffer(tx *types.CancelOfferTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructCancelOfferTransaction(tx, ops)
	if err != nil {
		return "", err
	}

	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) AtomicMatch(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}
	if ops == nil {
		ops = new(types.TransactOpts)
	}
	txInfo, err := c.constructAtomicMatchTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) WithdrawNft(tx *types.WithdrawNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructWithdrawNftTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}

	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) TransferNft(tx *types.TransferNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructTransferNftTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) Withdraw(tx *types.WithdrawTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructWithdrawTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}

	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) Transfer(tx *types.TransferTxReq, ops *types.TransactOpts, signatureList ...string) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	if ops == nil {
		ops = new(types.TransactOpts)
	}

	txInfo, err := c.constructTransferTransaction(tx, ops)
	if err != nil {
		return "", err
	}
	signature, err := c.generateSignature(txInfo, signatureList)
	if err != nil {
		return "", err
	}
	txInfo.L1Sig = signature
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", err
	}
	return c.SendRawTx(uint32(txInfo.GetTxType()), string(txInfoBytes))
}

func (c *l2Client) fullFillToAddrOps(ops *types.TransactOpts, to string) (*types.TransactOpts, error) {
	toAccount, err := c.GetAccountByL1Address(to)
	if err != nil {
		return nil, err
	}
	ops.ToAccountIndex = toAccount.Index
	ops.ToAccountAddress = toAccount.L1Address
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
		l2Account, err := c.GetAccountByL1Address(c.address)
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
		var x fr.Element
		_ = x.SetBytes([]byte(ops.CallData))
		b := x.Bytes()
		hFunc.Write(b[:])
		ops.CallDataHash = hFunc.Sum(nil)
	}
	if ops.GasFeeAssetAmount == nil {
		gas, err := c.GetGasFee(ops.GasFeeAssetId, ops.TxType)
		if err != nil {
			return nil, err
		}
		ops.GasFeeAssetAmount = gas
	}
	return ops, nil
}

func (c *l2Client) generateSignature(txInfo txtypes.TxInfo, signatureList []string) (string, error) {
	if len(signatureList) == 0 {
		if c.l1Signer == nil {
			return "", errors.New("PrivateKey has not been initialized correctly, signature is expected to be passed instead")
		}

		signBody := txInfo.GetL1SignatureBody()
		signHex, err := c.l1Signer.Sign(signBody)
		if err != nil {
			return "", err
		}
		return signHex, nil
	} else if len(signatureList) == 1 {
		return signatureList[0], nil
	} else {
		return "", errors.New("the passed signatureList contains more than one signature value and it is illegal")
	}
}

func (c *l2Client) GenerateSignBody(txData interface{}, ops *types.TransactOpts) (string, error) {
	txInfo, err := c.constructTransaction(txData, ops)
	if err != nil {
		return "", err
	}
	signatureBody := txInfo.GetL1SignatureBody()
	return signatureBody, nil
}

func (c *l2Client) GenerateSignature(privateKey string, txData interface{}, ops *types.TransactOpts) (string, error) {
	l1Signer, err := signer.NewL1Singer(privateKey)
	if err != nil {
		return "", err
	}
	txInfo, err := c.constructTransaction(txData, ops)
	if err != nil {
		return "", err
	}
	signBody := txInfo.GetL1SignatureBody()
	if err != nil {
		return "", err
	}
	signHex, err := l1Signer.Sign(signBody)
	if err != nil {
		return "", err
	}
	return signHex, nil
}

func (c *l2Client) constructTransaction(tx interface{}, ops *types.TransactOpts) (txtypes.TxInfo, error) {
	if ops == nil {
		ops = new(types.TransactOpts)
	}
	if value, ok := tx.(*types.MintNftTxReq); ok {
		return c.constructMintNftTransaction(value, ops)
	} else if value, ok := tx.(*types.CreateCollectionTxReq); ok {
		return c.constructCreateCollectionTransaction(value, ops)
	} else if value, ok := tx.(*types.CancelOfferTxReq); ok {
		return c.constructCancelOfferTransaction(value, ops)
	} else if value, ok := tx.(*types.OfferTxInfo); ok {
		return c.constructOfferTxInfoTransaction(value, ops)
	} else if value, ok := tx.(*types.TransferTxReq); ok {
		return c.constructTransferTransaction(value, ops)
	} else if value, ok := tx.(*types.TransferNftTxReq); ok {
		return c.constructTransferNftTransaction(value, ops)
	} else if value, ok := tx.(*types.WithdrawTxReq); ok {
		return c.constructWithdrawTransaction(value, ops)
	} else if value, ok := tx.(*types.WithdrawNftTxReq); ok {
		return c.constructWithdrawNftTransaction(value, ops)
	} else if value, ok := tx.(*types.UpdateNftReq); ok {
		return c.constructUpdateNFTTransaction(value, ops)
	} else if value, ok := tx.(*types.ChangePubKeyReq); ok {
		return c.constructChangePubKeyTransaction(value, ops)
	}
	return nil, errors.New("invalid tx type is passed")
}

func (c *l2Client) constructChangePubKeyTransaction(tx *types.ChangePubKeyReq, ops *types.TransactOpts) (*txtypes.ChangePubKeyInfo, error) {
	ops.TxType = txtypes.TxTypeChangePubKey
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	txInfo, err := txutils.ConstructChangePubKeyTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructMintNftTransaction(tx *types.MintNftTxReq, ops *types.TransactOpts) (*txtypes.MintNftTxInfo, error) {
	ops.TxType = txtypes.TxTypeMintNft
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	ops.ToAccountAddress = tx.To

	txInfo, err := txutils.ConstructMintNftTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructCancelOfferTransaction(tx *types.CancelOfferTxReq, ops *types.TransactOpts) (*txtypes.CancelOfferTxInfo, error) {
	ops.TxType = txtypes.TxTypeCancelOffer
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	txInfo, err := txutils.ConstructCancelOfferTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructCreateCollectionTransaction(tx *types.CreateCollectionTxReq, ops *types.TransactOpts) (*txtypes.CreateCollectionTxInfo, error) {
	ops.TxType = txtypes.TxTypeCreateCollection
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	txInfo, err := txutils.ConstructCreateCollectionTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructAtomicMatchTransaction(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (*txtypes.AtomicMatchTxInfo, error) {
	ops.TxType = txtypes.TxTypeAtomicMatch
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	txInfo, err := txutils.ConstructAtomicMatchTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructOfferTxInfoTransaction(tx *types.OfferTxInfo, ops *types.TransactOpts) (*txtypes.OfferTxInfo, error) {
	return &txtypes.OfferTxInfo{
		Type:                tx.Type,
		OfferId:             tx.OfferId,
		AccountIndex:        tx.AccountIndex,
		NftIndex:            tx.NftIndex,
		AssetId:             tx.AssetId,
		AssetAmount:         tx.AssetAmount,
		ListedAt:            tx.ListedAt,
		ExpiredAt:           tx.ExpiredAt,
		RoyaltyRate:         tx.RoyaltyRate,
		ChannelAccountIndex: tx.ChannelAccountIndex,
		ChannelRate:         tx.ChannelRate,
		ProtocolRate:        tx.ProtocolRate,
		ProtocolAmount:      tx.ProtocolAmount,
		Sig:                 tx.Sig,
	}, nil
}

func (c *l2Client) constructTransferNftTransaction(tx *types.TransferNftTxReq, ops *types.TransactOpts) (*txtypes.TransferNftTxInfo, error) {
	ops.TxType = txtypes.TxTypeTransferNft
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	ops.ToAccountAddress = tx.To
	txInfo, err := txutils.ConstructTransferNftTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructTransferTransaction(tx *types.TransferTxReq, ops *types.TransactOpts) (*txtypes.TransferTxInfo, error) {
	ops.TxType = txtypes.TxTypeTransfer
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	ops.ToAccountAddress = tx.To
	txInfo, err := txutils.ConstructTransferTx(c.keyManager, ops, tx)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructWithdrawTransaction(tx *types.WithdrawTxReq, ops *types.TransactOpts) (*txtypes.WithdrawTxInfo, error) {
	ops.TxType = txtypes.TxTypeWithdraw
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}
	txInfo, err := txutils.ConstructWithdrawTxInfo(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}

	return txInfo, nil
}

func (c *l2Client) constructWithdrawNftTransaction(tx *types.WithdrawNftTxReq, ops *types.TransactOpts) (*txtypes.WithdrawNftTxInfo, error) {
	ops.TxType = txtypes.TxTypeWithdrawNft
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return nil, err
	}

	txInfo, err := txutils.ConstructWithdrawNftTx(c.keyManager, tx, ops)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func (c *l2Client) constructUpdateNFTTransaction(req *types.UpdateNftReq, ops *types.TransactOpts) (*txtypes.UpdateNFTTxInfo, error) {
	if req.AccountIndex == 0 {
		l2Account, err := c.GetAccountByL1Address(c.address)
		if err != nil {
			return nil, err
		}
		req.AccountIndex = l2Account.Index
	}
	if req.Nonce == 0 {
		nonce, err := c.GetNftNextNonce(req.NftIndex)
		if err != nil {
			return nil, err
		}
		req.Nonce = nonce
	}
	updateNFTTxInfo, err := txutils.ConstructUpdateNFTTx(req, ops)
	if err != nil {
		return nil, err
	}
	return updateNFTTxInfo, nil
}

func (c *l2Client) parseResultStatus(respBody []byte) error {
	resultStatus := &types.Result{}
	if err := json.Unmarshal(respBody, resultStatus); err != nil {
		return err
	}
	if resultStatus.Code != types.CodeOK {
		return errors.New(resultStatus.Message)
	}
	return nil
}
