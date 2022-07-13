package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type client struct {
	zkbasURL   string
	keyManager KeyManager
}

func (c *client) SetKeyManager(keyManager KeyManager) {
	c.keyManager = keyManager
}

func (c *client) GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxsByPubKey?account_pk=%s&offset=%d&limit=%d",
			accountPk, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByPubKey{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountName?account_name=%s&offset=%d&limit=%d",
			accountName, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountIndexAndTxType?account_index=%d&tx_type=%d&offset=%d&limit=%d",
			accountIndex, txType, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByAccountIndexAndTxType{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxsListByAccountIndex?account_index=%d&offset=%d&limit=%d", accountIndex, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsListByAccountIndex{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) Search(info string) (*RespSearch, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/info/search?info=%s", info))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespSearch{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccounts(offset, limit uint32) (*RespGetAccounts, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/info/getAccounts?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccounts{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetGasFeeAssetList() (*RespGetGasFeeAssetList, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/info/getGasFeeAssetList")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetGasFeeAssetList{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/info/getWithdrawGasFee?asset_id=%d&withdraw_asset_id=%d&withdraw_amount=%d",
			assetId, withdrawAssetId, withdrawAmount))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetGasFee{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
}

func (c *client) GetGasFee(assetId uint32) (int64, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/info/getGasFee?asset_id=%d", assetId))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetGasFee{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
}

func (c *client) GetAssetsList() (*RespGetAssetsList, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/info/getAssetsList")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAssetsList{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetLayer2BasicInfo() (*RespGetLayer2BasicInfo, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/info/getLayer2BasicInfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetLayer2BasicInfo{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetBlockByCommitment(blockCommitment string) (*Block, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/block/getBlockByCommitment?block_commitment=%s", blockCommitment))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetBlockByCommitment{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result.Block, nil
}

func (c *client) GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/account/getBalanceByAssetIdAndAccountName?asset_id=%d&account_name=%s", assetId, accountName))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	result := &RespGetBalanceInfoByAssetIdAndAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Balance, nil
}

func (c *client) GetAccountStatusByAccountName(accountName string) (*RespGetAccountStatusByAccountName, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountName?account_name=%s", accountName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountStatusByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountInfoByAccountIndex(accountIndex int64) (*RespGetAccountInfoByAccountIndex, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/account/getAccountInfoByAccountIndex?account_index=%d", accountIndex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountInfoByAccountIndex{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountInfoByPubKey(accountPk string) (*RespGetAccountInfoByPubKey, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/account/getAccountInfoByPubKey?account_pk=%s", accountPk))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountInfoByPubKey{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountStatusByAccountPk(accountPk string) (*RespGetAccountStatusByAccountPk, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountPk?account_pk=%s", accountPk))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountStatusByAccountPk{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetCurrencyPriceBySymbol(symbol string) (*RespGetCurrencyPriceBySymbol, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/info/getCurrencyPriceBySymbol?symbol=%s", symbol))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetCurrencyPriceBySymbol{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetCurrencyPrices() (*RespGetCurrencyPrices, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/info/getCurrencyPrices")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetCurrencyPrices{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetSwapAmount(req ReqGetSwapAmount) (*RespGetSwapAmount, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/pair/getSwapAmount?pair_index=%d&asset_id=%d&asset_amount=%s&is_from=%v",
			req.PairIndex, req.AssetId, req.AssetAmount, req.IsFrom))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetSwapAmount{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAvailablePairs() (*RespGetAvailablePairs, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/pair/getAvailablePairs")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAvailablePairs{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetLPValue(pairIndex uint32, lpAmount string) (*RespGetLPValue, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/pair/getLPValue?pair_index=%d&lp_amount=%s", pairIndex, lpAmount))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetLPValue{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetPairInfo(pairIndex uint32) (*RespGetPairInfo, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/pair/getPairInfo?pair_index=%d", pairIndex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetPairInfo{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetTxByHash(txHash string) (*RespGetTxByHash, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxByHash?tx_hash=%s", txHash))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	txResp := &RespGetTxByHash{}
	if err := json.Unmarshal(body, &txResp); err != nil {
		return nil, err
	}
	return txResp, nil
}

func (c *client) GetMempoolTxs(offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getMempoolTxs?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &RespGetMempoolTxs{}
	if err := json.Unmarshal(body, &txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}

func (c *client) GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/tx/getmempoolTxsByAccountName?account_name=" + accountName)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &RespGetmempoolTxsByAccountName{}
	if err := json.Unmarshal(body, &txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *client) GetAccountInfoByAccountName(accountName string) (*AccountInfo, error) {
	resp, err := http.Get(c.zkbasURL + "/api/v1/account/getAccountInfoByAccountName?account_name=" + accountName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	account := &AccountInfo{}
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *client) GetNextNonce(accountIdx int64) (int64, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getNextNonce?account_index=%d", accountIdx))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetNextNonce{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Nonce, nil
}

func (c *client) GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/tx/getTxsListByBlockHeight?block_height=%d&limit=%d&offset=%d", blockHeight, 0, 0))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsListByBlockHeight{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Txs, nil
}

func (c *client) GetMaxOfferId(accountIndex uint32) (uint64, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/nft/getMaxOfferId?account_index=%d", accountIndex))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetMaxOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.OfferId, nil
}

func (c *client) GetBlockByBlockHeight(blockHeight int64) (*Block, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/block/getBlockByBlockHeight?block_height=%d", blockHeight))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	res := &RespGetBlockByBlockHeight{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res.Block, nil
}

func (c *client) GetBlocks(offset, limit int64) (uint32, []*Block, error) {
	resp, err := http.Get(c.zkbasURL +
		fmt.Sprintf("/api/v1/block/getBlocks?limit=%d&offset=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	res := &RespGetBlocks{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, nil, err
	}
	return res.Total, res.Blocks, nil
}

func (c *client) SendTx(txType uint32, txInfo string) (string, error) {
	if txType == TxTypeCreateCollection || txType == TxTypeMintNft {
		return "", fmt.Errorf("tx type not supported")
	}

	resp, err := http.PostForm(c.zkbasURL+"/api/v1/tx/sendTx",
		url.Values{"tx_type": {strconv.Itoa(int(txType))}, "tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SendCreateCollectionTx(txInfo string) (int64, error) {
	resp, err := http.PostForm(c.zkbasURL+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
}

func (c *client) SendMintNftTx(txInfo string) (int64, error) {
	resp, err := http.PostForm(c.zkbasURL+"/api/v1/tx/sendMintNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendMintNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.NftIndex, nil
}

func (c *client) SignAndSendMintNftTx(tx *MintNftTxInfo) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructMintNftTx(c.keyManager, tx)
	if err != nil {
		return 0, err
	}

	resp, err := http.PostForm(c.zkbasURL+"/api/v1/tx/sendMintNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendMintNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.NftIndex, nil
}

func (c *client) SignAndSendCreateCollectionTx(tx *CreateCollectionTxInfo) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructCreateCollectionTx(c.keyManager, tx)
	if err != nil {
		return 0, err
	}

	resp, err := http.PostForm(c.zkbasURL+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
}

func (c *client) SignAndSendCancelOfferTx(tx *CancelOfferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructCancelOfferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeCancelOffer, txInfo)
}

func (c *client) SignAndSendAtomicMatchTx(tx *AtomicMatchTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructAtomicMatchTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeAtomicMatch, txInfo)
}

func (c *client) SignAndSendOfferTx(tx *OfferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructOfferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeOffer, txInfo)
}

func (c *client) SignAndSendWithdrawNftTx(tx *WithdrawNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructWithdrawNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeWithdrawNft, txInfo)
}

func (c *client) SignAndSendTransferNftTx(tx *TransferNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructTransferNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeTransferNft, txInfo)
}

func (c *client) SignAndSendWithdrawTx(tx *WithdrawTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructWithdrawTxInfo(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeWithdraw, txInfo)
}

func (c *client) SignAndSendRemoveLiquidityTx(tx *RemoveLiquidityTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructRemoveLiquidityTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeRemoveLiquidity, txInfo)
}

func (c *client) SignAndSendAddLiquidityTx(tx *AddLiquidityTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructAddLiquidityTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeAddLiquidity, txInfo)
}

func (c *client) SignAndSendSwapTx(tx *SwapTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructSwapTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeSwap, txInfo)
}

func (c *client) SignAndTransfer(tx *TransferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructTransferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTx(TxTypeTransfer, txInfo)
}
