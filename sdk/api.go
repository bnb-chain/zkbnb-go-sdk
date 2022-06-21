package sdk

type ZecreySDK interface {
	// GetTxsListByBlockHeight return txs in block
	GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error)

	// GetAccountInfoByAccountName returns account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*AccountInfo, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex uint32) (uint64, error)

	// GetBlocks return total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*Block, error)

	// GetBlockByBlockHeight returns block by height
	GetBlockByBlockHeight(blockHeight int64) (*Block, error)

	// SendTx sends raw transaction
	SendTx(txType uint32, txInfo string) (string, error)
}

func NewZecreySDK(url string) ZecreySDK {
	return &client{
		zecreyLegendURL: url,
	}
}
