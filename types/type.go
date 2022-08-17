package types

type Status struct {
	Status    uint32 `json:"status"`
	NetworkId uint32 `json:"network_id"`
}

type AccountAsset struct {
	AssetId                  uint32 `json:"asset_id"`
	AssetName                string `json:"asset_name"`
	Balance                  string `json:"balance"`
	LpAmount                 string `json:"lp_amount"`
	OfferCanceledOrFinalized string `json:"offer_canceled_or_finalized"`
}

type Account struct {
	AccountStatus uint32          `json:"account_status"`
	AccountIndex  int64           `json:"account_index"`
	AccountName   string          `json:"account_name"`
	AccountPk     string          `json:"account_pk"`
	Nonce         int64           `json:"nonce"`
	Assets        []*AccountAsset `json:"assets"`
}

type SimpleAccount struct {
	AccountIndex int64  `json:"account_index"`
	AccountName  string `json:"account_name"`
	AccountPk    string `json:"account_pk"`
}

type Accounts struct {
	Total    uint32           `json:"total"`
	Accounts []*SimpleAccount `json:"accounts"`
}

type Asset struct {
	AssetId       uint32 `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	AssetDecimals uint32 `json:"asset_decimals"`
	AssetSymbol   string `json:"asset_symbol"`
	AssetAddress  string `json:"asset_address"`
	IsGasAsset    uint32 `json:"is_gas_asset"`
}

type Assets struct {
	Total  uint32   `json:"total"`
	Assets []*Asset `json:"assets"`
}

type Block struct {
	BlockCommitment                 string `json:"block_commitment"`
	BlockHeight                     int64  `json:"block_height"`
	StateRoot                       string `json:"state_root"`
	PriorityOperations              int64  `json:"priority_operations"`
	PendingOnChainOperationsHash    string `json:"pending_on_chain_operations_hash"`
	PendingOnChainOperationsPubData string `json:"pending_on_chain_operations_pub_data"`
	CommittedTxHash                 string `json:"committed_tx_hash"`
	CommittedAt                     int64  `json:"committed_at"`
	VerifiedTxHash                  string `json:"verified_tx_hash"`
	VerifiedAt                      int64  `json:"verified_at"`
	Txs                             []*Tx  `json:"txs"`
	BlockStatus                     int64  `json:"block_status"`
}

type Blocks struct {
	Total  uint32   `json:"total"`
	Blocks []*Block `json:"blocks"`
}

type CurrentHeight struct {
	Height int64 `json:"height"`
}

type ContractAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Layer2BasicInfo struct {
	BlockCommitted             int64             `json:"block_committed"`
	BlockVerified              int64             `json:"block_verified"`
	TotalTransactions          int64             `json:"total_transactions"`
	TransactionsCountYesterday int64             `json:"transactions_count_yesterday"`
	TransactionsCountToday     int64             `json:"transactions_count_today"`
	DauYesterday               int64             `json:"dau_yesterday"`
	DauToday                   int64             `json:"dau_today"`
	ContractAddresses          []ContractAddress `json:"contract_addresses"`
}

type CurrencyPrice struct {
	Pair    string `json:"pair"`
	AssetId uint32 `json:"asset_id"`
	Price   string `json:"price"`
}

type CurrencyPrices struct {
	Total          uint32           `json:"total"`
	CurrencyPrices []*CurrencyPrice `json:"currency_prices"`
}

type GasFee struct {
	GasFee string `json:"gas_fee"`
}

type GasAccount struct {
	AccountStatus int64  `json:"account_status"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `json:"account_name"`
}

type GasFeeAssets struct {
	Assets []Asset `json:"assets"`
}

type Search struct {
	DataType int32 `json:"data_type"`
}

type SwapAmount struct {
	AssetId     uint32 `json:"asset_id"`
	AssetName   string `json:"asset_name"`
	AssetAmount string `json:"asset_amount"`
}

type Pair struct {
	PairIndex     uint32 `json:"pair_index"`
	AssetAId      uint32 `json:"asset_a_id"`
	AssetAName    string `json:"asset_a_name"`
	AssetAAmount  string `json:"asset_a_amount"`
	AssetBId      uint32 `json:"asset_b_id"`
	AssetBName    string `json:"asset_b_name"`
	AssetBAmount  string `json:"asset_b_amount"`
	FeeRate       int64  `json:"fee_rate"`
	TreasuryRate  int64  `json:"treasury_rate"`
	TotalLpAmount string `json:"total_lp_amount"`
}

type Pairs struct {
	Pairs []*Pair `json:"pairs"`
}

type LpValue struct {
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetAAmount string `json:"asset_a_amount"`
	AssetBId     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	AssetBAmount string `json:"asset_b_amount"`
}

type Tx struct {
	TxHash        string `json:"tx_hash"`
	TxType        int64  `json:"tx_type,range=[1:64]"`
	TxAmount      string `json:"tx_amount"`
	TxInfo        string `json:"tx_info"`
	TxStatus      int64  `json:"tx_status"`
	GasFeeAssetId int64  `json:"gas_fee_asset_id"`
	GasFee        string `json:"gas_fee"`
	NftIndex      int64  `json:"nft_index"`
	PairIndex     int64  `json:"pair_index"`
	AssetId       int64  `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	NativeAddress string `json:"native_adress"`
	ExtraInfo     string `json:"extra_info"`
	Memo          string `json:"memo"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `json:"account_name"`
	Nonce         int64  `json:"nonce"`
	ExpiredAt     int64  `json:"expire_at"`
	Status        int64  `json:"status,options=0|1|2"`
	BlockId       int64  `json:"block_id"`
	BlockHeight   int64  `json:"block_height"`
	CreatedAt     int64  `json:"created_at"`
	StateRoot     string `json:"state_root"`
}

type Txs struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type MempoolTxs struct {
	Total      uint32 `json:"total"`
	MempoolTxs []*Tx  `json:"mempool_txs"`
}

type TxHash struct {
	TxHash string `json:"tx_hash"`
}

type NextNonce struct {
	Nonce uint64 `json:"nonce"`
}

type EnrichedTx struct {
	Tx          Tx    `json:"tx"`
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
	AssetAId    int64 `json:"asset_a_id"`
	AssetBId    int64 `json:"asset_b_id"`
}

type MaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
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

type Nfts struct {
	Total int64  `json:"total"`
	Nfts  []*Nft `json:"nfts"`
}
