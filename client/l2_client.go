package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
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

func (c *l2Client) GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxsByPubKey?account_pk=%s&offset=%d&limit=%d",
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
	result := &types.RespGetTxsByPubKey{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountName?account_name=%s&offset=%d&limit=%d",
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
	result := &types.RespGetTxsByAccountName{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountIndexAndTxType?account_index=%d&tx_type=%d&offset=%d&limit=%d",
			accountIndex, txType, offset, limit))
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
	result := &types.RespGetTxsByAccountIndexAndTxType{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxsListByAccountIndex?account_index=%d&offset=%d&limit=%d", accountIndex, offset, limit))
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
	result := &types.RespGetTxsListByAccountIndex{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *l2Client) Search(info string) (*types.RespSearch, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/search?info=%s", info))
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
	result := &types.RespSearch{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccounts(offset, limit uint32) (*types.RespGetAccounts, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getAccounts?offset=%d&limit=%d", offset, limit))
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
	result := &types.RespGetAccounts{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetGasFeeAssetList() (*types.RespGetGasFeeAssetList, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/info/getGasFeeAssetList")
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
	result := &types.RespGetGasFeeAssetList{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (*big.Int, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getWithdrawGasFee?asset_id=%d&withdraw_asset_id=%d&withdraw_amount=%d",
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
	result := &types.RespGetGasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	var price big.Int
	price.SetString(result.GasFee, 10)
	return &price, nil
}

func (c *l2Client) GetGasFee(assetId int64) (*big.Int, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getGasFee?asset_id=%d", assetId))
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
	result := &types.RespGetGasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	var price big.Int
	price.SetString(result.GasFee, 10)
	return &price, nil
}

func (c *l2Client) GetAssetsList() (*types.RespGetAssetsList, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/info/getAssetsList")
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
	result := &types.RespGetAssetsList{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetLayer2BasicInfo() (*types.RespGetLayer2BasicInfo, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/info/getLayer2BasicInfo")
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
	result := &types.RespGetLayer2BasicInfo{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetBlockByCommitment(blockCommitment string) (*types.Block, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/block/getBlockByCommitment?block_commitment=%s", blockCommitment))
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
	result := &types.RespGetBlockByCommitment{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return &result.Block, nil
}

func (c *l2Client) GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account/getBalanceByAssetIdAndAccountName?asset_id=%d&account_name=%s", assetId, accountName))
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
	result := &types.RespGetBalanceInfoByAssetIdAndAccountName{}
	if err := json.Unmarshal(body, result); err != nil {
		return "", err
	}
	return result.Balance, nil
}

func (c *l2Client) GetAccountStatusByAccountName(accountName string) (*types.RespGetAccountStatusByAccountName, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountName?account_name=%s", accountName))
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
	result := &types.RespGetAccountStatusByAccountName{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccountInfoByAccountIndex(accountIndex int64) (*types.RespGetAccountInfoByAccountIndex, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account/getAccountInfoByAccountIndex?account_index=%d", accountIndex))
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
	result := &types.RespGetAccountInfoByAccountIndex{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccountInfoByPubKey(accountPk string) (*types.RespGetAccountInfoByPubKey, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account/getAccountInfoByPubKey?account_pk=%s", accountPk))
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
	result := &types.RespGetAccountInfoByPubKey{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAccountStatusByAccountPk(accountPk string) (*types.RespGetAccountStatusByAccountPk, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountPk?account_pk=%s", accountPk))
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
	result := &types.RespGetAccountStatusByAccountPk{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetCurrencyPriceBySymbol(symbol string) (*types.RespGetCurrencyPriceBySymbol, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getCurrencyPriceBySymbol?symbol=%s", symbol))
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
	result := &types.RespGetCurrencyPriceBySymbol{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetCurrencyPrices() (*types.RespGetCurrencyPrices, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/info/getCurrencyPrices")
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
	result := &types.RespGetCurrencyPrices{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetSwapAmount(req *types.ReqGetSwapAmount) (*types.RespGetSwapAmount, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pair/getSwapAmount?pair_index=%d&asset_id=%d&asset_amount=%s&is_from=%v",
			req.PairIndex, req.AssetId, req.AssetAmount, req.IsFrom))
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
	result := &types.RespGetSwapAmount{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetAvailablePairs() (*types.RespGetAvailablePairs, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/pair/getAvailablePairs")
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
	result := &types.RespGetAvailablePairs{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetLPValue(pairIndex uint32, lpAmount string) (*types.RespGetLPValue, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pair/getLPValue?pair_index=%d&lp_amount=%s", pairIndex, lpAmount))
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
	result := &types.RespGetLPValue{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetPairInfo(pairIndex uint32) (*types.RespGetPairInfo, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/pair/getPairInfo?pair_index=%d", pairIndex))
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
	result := &types.RespGetPairInfo{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *l2Client) GetTxByHash(txHash string) (*types.RespGetTxByHash, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxByHash?tx_hash=%s", txHash))
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
	txResp := &types.RespGetTxByHash{}
	if err := json.Unmarshal(body, txResp); err != nil {
		return nil, err
	}
	return txResp, nil
}

func (c *l2Client) GetMempoolTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getMempoolTxs?offset=%d&limit=%d", offset, limit))
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
	txsResp := &types.RespGetMempoolTxs{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}

func (c *l2Client) GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*types.Tx, err error) {
	resp, err := http.Get(c.endpoint + "/api/v1/tx/getmempoolTxsByAccountName?account_name=" + accountName)
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
	txsResp := &types.RespGetmempoolTxsByAccountName{}
	if err := json.Unmarshal(body, txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *l2Client) GetAccountInfoByAccountName(accountName string) (*types.AccountInfo, error) {
	resp, err := http.Get(c.endpoint + "/api/v1/account/getAccountInfoByAccountName?account_name=" + accountName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		println("error code", resp.StatusCode)
		return nil, fmt.Errorf(string(body))
	}
	account := &types.AccountInfo{}
	if err := json.Unmarshal(body, account); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *l2Client) GetNextNonce(accountIdx int64) (int64, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getNextNonce?account_index=%d", accountIdx))
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
	result := &types.RespGetNextNonce{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return result.Nonce, nil
}

func (c *l2Client) GetTxsListByBlockHeight(blockHeight uint32) ([]*types.Tx, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/tx/getTxsListByBlockHeight?block_height=%d&limit=%d&offset=%d", blockHeight, 0, 0))
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
	result := &types.RespGetTxsListByBlockHeight{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result.Txs, nil
}

func (c *l2Client) GetMaxOfferId(accountIndex int64) (uint64, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/nft/getMaxOfferId?account_index=%d", accountIndex))
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
	result := &types.RespGetMaxOfferId{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return result.OfferId, nil
}

func (c *l2Client) GetBlockByHeight(blockHeight int64) (*types.Block, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/block/getBlockByBlockHeight?block_height=%d", blockHeight))
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
	res := &types.RespGetBlockByBlockHeight{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res.Block, nil
}

func (c *l2Client) GetBlocks(offset, limit int64) (uint32, []*types.Block, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/block/getBlocks?limit=%d&offset=%d", limit, offset))
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
	res := &types.RespGetBlocks{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, nil, err
	}
	return res.Total, res.Blocks, nil
}

func (c *l2Client) GetGasAccount() (*types.RespGetGasAccount, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getGasAccount"))
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
	res := &types.RespGetGasAccount{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) GetAccountNftList(accountIndex, offset, limit int64) (*types.RespGetAccountNftList, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/nft/getAccountNftList?account_index=%d&limit=%d&offset=%d", accountIndex, limit, offset))
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
	res := &types.RespGetAccountNftList{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *l2Client) SendRawTx(txType uint32, txInfo string) (string, error) {
	if txType == types.TxTypeCreateCollection || txType == types.TxTypeMintNft {
		return "", fmt.Errorf("tx type not supported")
	}

	resp, err := http.PostForm(c.endpoint+"/api/v1/tx/sendTx",
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
	res := &types.RespSendTx{}
	if err := json.Unmarshal(body, res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *l2Client) SendRawCreateCollectionTx(txInfo string) (int64, error) {
	resp, err := http.PostForm(c.endpoint+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
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
	res := &types.RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
}

func (c *l2Client) SendRawMintNftTx(txInfo string) (int64, error) {
	resp, err := http.PostForm(c.endpoint+"/api/v1/tx/sendMintNftTx",
		url.Values{"tx_info": {txInfo}})
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
	res := &types.RespSendMintNftTx{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, err
	}
	return res.NftIndex, nil
}

func (c *l2Client) MintNft(tx *types.MintNftTxReq, ops *types.TransactOpts) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}
	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return 0, err
	}

	ops, err = c.fullFillToAddrOps(ops, tx.To)
	if err != nil {
		return 0, err
	}

	txInfo, err := txutils.ConstructMintNftTx(c.keyManager, tx, ops)
	if err != nil {
		return 0, err
	}

	resp, err := http.PostForm(c.endpoint+"/api/v1/tx/sendMintNftTx", url.Values{"tx_info": {txInfo}})
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
	res := &types.RespSendMintNftTx{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, err
	}
	return res.NftIndex, nil
}

func (c *l2Client) CreateCollection(tx *types.CreateCollectionReq, ops *types.TransactOpts) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	ops, err := c.fullFillDefaultOps(ops)
	if err != nil {
		return 0, err
	}

	txInfo, err := txutils.ConstructCreateCollectionTx(c.keyManager, tx, ops)
	if err != nil {
		return 0, err
	}

	resp, err := http.PostForm(c.endpoint+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
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
	res := &types.RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
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

func (c *l2Client) SendTransferNft(tx *types.TransferNftTxReq, ops *types.TransactOpts) (string, error) {
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
	toAccount, err := c.GetAccountInfoByAccountName(to)
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
		if gasAccount.AccountIndex == 0 {
			return nil, fmt.Errorf("get gas account error, gas account index is %d", gasAccount.AccountIndex)
		}
		ops.GasAccountIndex = gasAccount.AccountIndex
	}
	if ops.ExpiredAt == 0 {
		ops.ExpiredAt = time.Now().Add(defaultExpireTime).UnixMilli()
	}
	if ops.FromAccountIndex == 0 {
		l2Account, err := c.GetAccountInfoByPubKey(hex.EncodeToString(c.keyManager.PubKey().Bytes()))
		if err != nil {
			return nil, err
		}
		ops.FromAccountIndex = l2Account.AccountIndex
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
