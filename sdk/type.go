package sdk

type Asset struct {
	Id         uint32
	BalanceEnc string
}

type AccountInfo struct {
	Index     uint32   `json:"account_index"`
	Name      string   `json:"account_name"`
	AccountPk string   `json:"account_pk"`
	Assets    []*Asset `json:"assets"`
}

type RawTx struct {
	TxType uint32
	TxInfo string //globalrpc => sendAddliquidity.go
	TxHash string
}

type TxHash struct {
	TxHash    string `json:"tx_hash"`
	CreatedAt int64  `json:"created_at"`
}

type TxDetail struct {
	TxId            int64  `json:"tx_id"`
	AssetId         int64  `json:"asset_id"`
	AssetType       int64  `json:"asset_type"`
	AccountIndex    int64  `json:"account_index"`
	AccountName     string `json:"account_name"`
	Balance         string `json:"balance"`
	BalanceDelta    string `json:"balance_delta"`
	Order           int64  `json:"order"`
	AccountOrder    int64  `json:"account_order"`
	Nonce           int64  `json:"nonce"`
	CollectionNonce int64  `json:"collection_nonce"`
}

type Tx struct {
	TxHash        string      `json:"tx_hash"`
	TxType        int64       `json:"tx_type"`
	GasFee        string      `json:"gas_fee"`
	GasFeeAssetId int64       `json:"gas_fee_asset_id"`
	TxStatus      int64       `json:"tx_status"`
	BlockHeight   int64       `json:"block_height"`
	BlockId       int64       `json:"block_id"`
	StateRoot     string      `json:"state_root"`
	NftIndex      int64       `json:"nft_index"`
	PairIndex     int64       `json:"pair_index"`
	AssetId       int64       `json:"asset_id"`
	TxAmount      string      `json:"tx_amount"`
	NativeAddress string      `json:"native_address"`
	TxInfo        string      `json:"tx_info"`
	TxDetails     []*TxDetail `json:"tx_details"`
	ExtraInfo     string      `json:"extra_info"`
	Memo          string      `json:"memo"`
	AccountIndex  int64       `json:"account_index"`
	Nonce         int64       `json:"nonce"`
	ExpiredAt     int64       `json:"expired_at"`
}

type RespGetTxsListByBlockHeight struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type Block struct {
	BlockCommitment                 string `json:"block_commitment"`
	BlockHeight                     int64  `json:"block_height"`
	StateRoot                       string `json:"state_root"`
	PriorityOperations              int64  `json:"priority_operations"`
	PendingOnChainOperationsHash    string `json:"pending_on_chain_operations_hash"`
	PendingOnChainOperationsPubData string `json:"pending_on_chain_operations_hub_data"`
	CommittedTxHash                 string `json:"committed_tx_hash"`
	CommittedAt                     int64  `json:"committed_at"`
	VerifiedTxHash                  string `json:"verified_tx_hash"`
	VerifiedAt                      int64  `json:"verified_at"`
	Txs                             []*Tx  `json:"txs"`
	BlockStatus                     int64  `json:"block_status"`
}

type RespGetBlocks struct {
	Total  uint32   `json:"total"`
	Blocks []*Block `json:"blocks"`
}

type RespGetMaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
}

type RespSendTx struct {
	TxId string `json:"tx_id"`
}
