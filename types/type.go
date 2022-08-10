package types

type Asset struct {
	AssetId                  uint32 `json:"asset_id"`
	Balance                  string `json:"balance"`
	LpAmount                 string `json:"lp_amount"`
	OfferCanceledOrFinalized string `json:"offer_canceled_or_finalized"`
}

type AccountInfo struct {
	Index     int64    `json:"account_index"`
	Nonce     int64    `json:"nonce"`
	AccountPk string   `json:"account_pk"`
	Assets    []*Asset `json:"assets"`
}

type RawTx struct {
	TxType uint32
	TxInfo string
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

type RespGetBlockByBlockHeight struct {
	Block *Block `json:"block"`
}

type RespGetMaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
}

type RespSendTx struct {
	TxId string `json:"tx_id"`
}

type RespGetNextNonce struct {
	Nonce int64 `json:"nonce"`
}

type RespGetmempoolTxsByAccountName struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"mempool_txs"`
}

type RespGetMempoolTxs struct {
	Total      uint32 `json:"total"`
	MempoolTxs []*Tx  `json:"mempool_txs"`
}

type RespGetTxByHash struct {
	Tx          Tx    `json:"result"`
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
	AssetAId    int64 `json:"asset_a_id"`
	AssetBId    int64 `json:"asset_b_id"`
}

type RespGetAccountStatusByAccountPk struct {
	AccountStatus int64  `json:"account_status"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `json:"account_name"`
}

type RespGetAccountStatusByAccountName struct {
	AccountStatus uint32 `json:"account_status"`
	AccountIndex  uint32 `json:"account_index"`
	AccountPk     string `json:"account_pk"`
}

type RespGetBalanceInfoByAssetIdAndAccountName struct {
	Balance string `json:"balance_enc"`
}

type RespGetBlockByCommitment struct {
	Block Block `json:"block"`
}

type RespGetLayer2BasicInfo struct {
	BlockCommitted             int64    `json:"block_committed"`
	BlockVerified              int64    `json:"block_verified"`
	TotalTransactions          int64    `json:"total_transactions"`
	TransactionsCountYesterday int64    `json:"transactions_count_yesterday"`
	TransactionsCountToday     int64    `json:"transactions_count_today"`
	DauYesterday               int64    `json:"dau_yesterday"`
	DauToday                   int64    `json:"dau_today"`
	ContractAddresses          []string `json:"contract_addresses"`
}

type RespGetAssetsList struct {
	Assets []*AssetInfo `json:"assets"`
}

type AssetInfo struct {
	AssetId       uint32 `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	AssetDecimals uint32 `json:"asset_decimals"`
	AssetSymbol   string `json:"asset_symbol"`
	AssetAddress  string `json:"asset_address"`
	IsGasAsset    uint32 `json:"is_gas_asset"`
}

type RespGetCurrencyPriceBySymbol struct {
	AssetId uint32 `json:"assetId"`
	Price   uint64 `json:"price"`
}

type RespGetCurrencyPrices struct {
	Data []*DataCurrencyPrices `json:"data"`
}

type DataCurrencyPrices struct {
	Pair    string `json:"pair"`
	AssetId uint32 `json:"assetId"`
	Price   uint64 `json:"price"`
}

type RespGetGasFee struct {
	GasFee string `json:"gas_fee"`
}

type RespGetGasFeeAssetList struct {
	Assets []AssetInfo `json:"assets"`
}

type RespGetAccounts struct {
	Total    uint32      `json:"total"`
	Accounts []*Accounts `json:"accounts"`
}

type Accounts struct {
	AccountIndex uint32 `json:"account_index"`
	AccountName  string `json:"account_name"`
	PublicKey    string `json:"public_key"`
}

type RespSearch struct {
	DataType int32 `json:"data_type"`
}

type RespGetTxsList struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsListByAccountIndex struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByAccountIndexAndTxType struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByAccountName struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByPubKey struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetAccountInfoByPubKey struct {
	AccountStatus uint32          `json:"account_status"`
	AccountName   string          `json:"account_name"`
	AccountIndex  int64           `json:"account_index"`
	Assets        []*AccountAsset `json:"assets"`
}

type AccountAsset struct {
	AssetId                  uint32 `json:"asset_id"`
	Balance                  string `json:"balance"`
	LpAmount                 string `json:"lp_amount"`
	OfferCanceledOrFinalized string `json:"offer_canceled_or_finalized"`
}

type RespGetAccountInfoByAccountIndex struct {
	AccountStatus uint32          `json:"account_status"`
	AccountName   string          `json:"account_name"`
	AccountPk     string          `json:"account_pk"`
	Assets        []*AccountAsset `json:"assets"`
}

type RespSendCreateCollectionTx struct {
	CollectionId int64 `json:"collection_id"`
}

type RespSendMintNftTx struct {
	NftIndex int64 `json:"nft_index"`
}

type ReqGetSwapAmount struct {
	PairIndex   uint32 `form:"pair_index"`
	AssetId     uint32 `form:"asset_id"`
	AssetAmount string `form:"asset_amount"`
	IsFrom      bool   `form:"is_from"`
}

type RespGetSwapAmount struct {
	ResAssetAmount string `json:"res_asset_amount"`
	ResAssetId     uint32 `json:"res_asset_id"`
}

type RespGetAvailablePairs struct {
	Pairs []*Pair `json:"result"`
}

type Pair struct {
	PairIndex    uint32 `json:"pair_index"`
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetAAmount string `json:"asset_a_amount"`
	AssetBId     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	AssetBAmount string `json:"asset_b_amount"`
	FeeRate      int64  `json:"fee_Rate"`
	TreasuryRate int64  `json:"treasury_rate"`
}

type RespGetLPValue struct {
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetAAmount string `json:"asset_a_amount"`
	AssetBid     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	AssetBAmount string `json:"asset_b_amount"`
}

type RespGetPairInfo struct {
	AssetAId      uint32 `json:"asset_a_id"`
	AssetAAmount  string `json:"asset_a_amount"`
	AssetBId      uint32 `json:"asset_b_id"`
	AssetBAmount  string `json:"asset_b_amount"`
	TotalLpAmount string `json:"total_lp_amount"`
}

type RespGetGasAccount struct {
	AccountStatus int64  `json:"account_status"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `json:"account_name"`
}

type Nft struct {
	NftIndex            int64  `json:"nft_index"`
	CreatorAccountIndex int64  `json:"creator_account_index"`
	OwnerAccountIndex   int64  `json:"owner_account_index"`
	NftContentHash      string `json:"nft_content_hash"`
	NftL1Address        string `json:"nft_l1_address"`
	NftL1TokenId        string `json:"nft_l1_token_id"`
	CreatorTreasuryRate int64  `json:"creator_treasury_rate"`
	CollectionId        int64  `json:"collection_id"`
}

type RespGetAccountNftList struct {
	Total int64  `json:"total"`
	Nfts  []*Nft `json:"nfts"`
}
