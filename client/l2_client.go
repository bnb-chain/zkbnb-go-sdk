package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/txutils"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

type l2Client struct {
	endpoint   string
	keyManager accounts.KeyManager
}

func (c *l2Client) SetKeyManager(keyManager accounts.KeyManager) {
	c.keyManager = keyManager
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

func (c *l2Client) GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getWithdrawGasFee?asset_id=%d&withdraw_asset_id=%d&withdraw_amount=%d",
			assetId, withdrawAssetId, withdrawAmount))
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
	result := &types.RespGetGasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
}

func (c *l2Client) GetGasFee(assetId uint32) (int64, error) {
	resp, err := http.Get(c.endpoint +
		fmt.Sprintf("/api/v1/info/getGasFee?asset_id=%d", assetId))
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
	result := &types.RespGetGasFee{}
	if err := json.Unmarshal(body, result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
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

func (c *l2Client) GetMaxOfferId(accountIndex uint32) (uint64, error) {
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
		fmt.Sprintf("/api/v1/block/getBlocks?limit=%d&offset=%d", offset, limit))
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

func (c *l2Client) MintNft(tx *types.MintNftTxInfo) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructMintNftTx(c.keyManager, tx)
	if err != nil {
		return 0, err
	}

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

func (c *l2Client) CreateCollection(tx *types.CreateCollectionTxInfo) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructCreateCollectionTx(c.keyManager, tx)
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

func (c *l2Client) CancelOffer(tx *types.CancelOfferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructCancelOfferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeCancelOffer, txInfo)
}

func (c *l2Client) AtomicMatch(tx *types.AtomicMatchTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructAtomicMatchTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeAtomicMatch, txInfo)
}

func (c *l2Client) Offer(tx *types.OfferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructOfferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeOffer, txInfo)
}

func (c *l2Client) WithdrawNft(tx *types.WithdrawNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructWithdrawNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeWithdrawNft, txInfo)
}

func (c *l2Client) SendTransferNft(tx *types.TransferNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructTransferNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeTransferNft, txInfo)
}

func (c *l2Client) Withdraw(tx *types.WithdrawTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructWithdrawTxInfo(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeWithdraw, txInfo)
}

func (c *l2Client) RemoveLiquidity(tx *types.RemoveLiquidityTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructRemoveLiquidityTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeRemoveLiquidity, txInfo)
}

func (c *l2Client) AddLiquidity(tx *types.AddLiquidityTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructAddLiquidityTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeAddLiquidity, txInfo)
}

func (c *l2Client) Swap(tx *types.SwapTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructSwapTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeSwap, txInfo)
}

func (c *l2Client) Transfer(tx *types.TransferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := txutils.ConstructTransferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendRawTx(types.TxTypeTransfer, txInfo)
}
