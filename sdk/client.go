package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type client struct {
	zecreyLegendURL string
}

func (c *client) GetAccountInfoByAccountName(accountName string) (*AccountInfo, error) {
	resp, err := http.Get(c.zecreyLegendURL + "/api/v1/account/getAccountInfoByAccountName?account_name=" + accountName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	account := &AccountInfo{}
	if err := json.Unmarshal([]byte(string(body)), &account); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *client) GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error) {
	resp, err := http.Get(c.zecreyLegendURL +
		fmt.Sprintf("/api/v1/tx/getTxsListByBlockHeight?block_height=%d&limit=%d&offset=%d", blockHeight, 0, 0))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &RespGetTxsListByBlockHeight{}
	if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
		return nil, err
	}
	return result.Txs, nil
}

func (c *client) GetMaxOfferId(accountIndex uint32) (uint64, error) {
	resp, err := http.Get(c.zecreyLegendURL +
		fmt.Sprintf("/api/v1/nft/getMaxOfferId?account_index=%d", accountIndex))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	result := &RespGetMaxOfferId{}
	if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
		return 0, err
	}
	return result.OfferId, nil
}

func (c *client) GetBlockByBlockHeight(blockHeight int64) (*Block, error) {
	num, blocks, err := c.GetBlocks(blockHeight, 1)
	if err != nil {
		return nil, err
	}
	if num != 1 {
		return nil, fmt.Errorf("block does not exist")
	}

	return blocks[0], nil
}

func (c *client) GetBlocks(offset, limit int64) (uint32, []*Block, error) {
	resp, err := http.Get(c.zecreyLegendURL +
		fmt.Sprintf("/api/v1/block/getBlocks?limit=%d&offset=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	res := &RespGetBlocks{}
	if err := json.Unmarshal([]byte(string(body)), &res); err != nil {
		return 0, nil, err
	}
	return res.Total, res.Blocks, nil
}

func (c *client) SendTx(txType uint32, txInfo string) (string, error) {
	resp, err := http.Get(c.zecreyLegendURL +
		fmt.Sprintf("/api/v1/tx/sendTx?tx_type=%d&tx_info=%s", txType, txInfo))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := &RespSendTx{}
	if err := json.Unmarshal([]byte(string(body)), &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}
