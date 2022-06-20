package sdk

type ZecreySDK interface {
	// Wallet can sign/send different transactions

	// Easily get/monitor l2 rollbacks
	IfRollbacksOccurred() (blockHeight uint32, err error)

	// Query account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*AccountInfo, error)

	// Query max offer id for an account
	GetMaxOfferId(accountIndex uint32) (OfferId uint64, err error)

	// Send raw transaction
	GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error)
}

func NewZecrey() ZecreySDK {
	return &client{
		zecreyLegendURL: "http://127.0.0.1:8888",
	}
}
