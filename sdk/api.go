package sdk

type ZecreySDK interface {
	// Send raw transaction
	GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error)

	// Query account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*AccountInfo, error)

	// Query max offer id for an account
	GetMaxOfferId(accountIndex uint32) (uint64, error)

	// return total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*Block, error)

	SendTx(txType uint32, txInfo string) (string, error)

	// GetAccountInfoByPubKey(accountPk string) (*AccountInfo, error)
	// GetBlockByBlockHeight(blockHeight int64)(*Block,error)
	// GetmempoolTxsByAccountName 直接返回全部交易
	// GetNextNonce
	// // Easily get/monitor l2 rollbacks
	// IfRollbacksOccurred() (blockHeight uint32, err error)
}

func NewZecrey() ZecreySDK {
	return &client{
		zecreyLegendURL: "http://127.0.0.1:8888",
	}
}
