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

func (z *client) IfRollbacksOccurred() (blockHeight uint32, err error) {
	// todo implement
	return blockHeight, err
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
func (z *client) GetMaxOfferId(accountIndex uint32) (offerId uint64, err error) {
	// todo implement
	return offerId, err
}
func (c *client) GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error) {
	resp, err := http.Get(c.zecreyLegendURL +
		fmt.Sprintf("/api/v1/block/getTxsListByBlockHeight?block_height=%d&limit=%d&offset=%d", blockHeight, 0, 0))
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
