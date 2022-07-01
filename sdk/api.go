package sdk

type ZkbasSDK interface {
	// SetKeyManager sets the key manager for txs to sign
	SetKeyManager(keyManager KeyManager)

	// GetTxsListByBlockHeight return txs in block
	GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error)

	// GetAccountInfoByAccountName returns account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*AccountInfo, error)

	// GetNextNonce returns nonce of account
	GetNextNonce(accountIdx int64) (int64, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex uint32) (uint64, error)

	// GetBlocks returns total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*Block, error)

	// GetBlockByBlockHeight returns block by height
	GetBlockByBlockHeight(blockHeight int64) (*Block, error)

	// GetMempoolTxs returns the mempool txs
	GetMempoolTxs(offset, limit uint32) (total uint32, txs []*Tx, err error)

	// GetMempoolTxsByAccountName returns the mempool txs by account name
	GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*Tx, err error)

	// GetBalanceByAssetIdAndAccountName returns the balance by asset id and account name
	GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error)

	// GetAccountStatusByAccountName returns account status by account name
	GetAccountStatusByAccountName(accountName string) (*RespGetAccountStatusByAccountName, error)

	// GetAccountStatusByAccountPk returns account status by account public key
	GetAccountStatusByAccountPk(accountPk string) (*RespGetAccountStatusByAccountPk, error)

	// GetCurrencyPriceBySymbol returns currency price by symbol
	GetCurrencyPriceBySymbol(symbol string) (*RespGetCurrencyPriceBySymbol, error)

	// GetCurrencyPrices returns all currency prices
	GetCurrencyPrices() (*RespGetCurrencyPrices, error)

	// GetSwapAmount returns swap amount by request
	GetSwapAmount(req ReqGetSwapAmount) (*RespGetSwapAmount, error)

	// GetAvailablePairs returns available pairs
	GetAvailablePairs() (*RespGetAvailablePairs, error)

	// GetLPValue returns lp value
	GetLPValue(pairIndex uint32, lpAmount string) (*RespGetLPValue, error)

	// GetPairInfo returns pair info by pair index
	GetPairInfo(pairIndex uint32) (*RespGetPairInfo, error)

	// GetTxByHash returns tx by tx hash
	GetTxByHash(txHash string) (*RespGetTxByHash, error)

	// GetBlockByCommitment returns block by commitment
	GetBlockByCommitment(blockCommitment string) (*Block, error)

	// GetTxsByPubKey returns txs by public key
	GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*Tx, err error)

	// GetTxsByAccountName returns txs by account name
	GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*Tx, err error)

	// GetTxsByAccountIndexAndTxType returns txs by account index and tx type
	GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*Tx, err error)

	// GetTxsListByAccountIndex returns txs list by account index
	GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*Tx, err error)

	// Search returns data type by queried info
	Search(info string) (*RespSearch, error)

	// GetAccounts returns accounts by query conditions
	GetAccounts(offset, limit uint32) (*RespGetAccounts, error)

	// GetGasFeeAssetList returns gas fee asset list
	GetGasFeeAssetList() (*RespGetGasFeeAssetList, error)

	// GetWithdrawGasFee returns withdraw gas fee
	GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error)

	// GetGasFee returns gas fee for asset
	GetGasFee(assetId uint32) (int64, error)

	// GetAssetsList returns asset list
	GetAssetsList() (*RespGetAssetsList, error)

	// GetLayer2BasicInfo returns layer 2 basic info
	GetLayer2BasicInfo() (*RespGetLayer2BasicInfo, error)

	// GetAccountInfoByPubKey returns account info by public key
	GetAccountInfoByPubKey(accountPk string) (*RespGetAccountInfoByPubKey, error)

	// GetAccountInfoByAccountIndex returns account info by account index
	GetAccountInfoByAccountIndex(accountIndex int64) (*RespGetAccountInfoByAccountIndex, error)

	// SendTx sends signed raw transaction and returns tx id
	SendTx(txType uint32, txInfo string) (string, error)

	// SendMintNftTx sends signed raw mint nft transaction and returns nft id
	SendMintNftTx(txInfo string) (int64, error)

	// SendCreateCollectionTx sends signed raw create collection transaction and returns collection id
	SendCreateCollectionTx(txInfo string) (int64, error)

	// SignAndSendMintNftTx will sign tx with key manager and send signed transaction
	SignAndSendMintNftTx(tx *MintNftTxInfo) (int64, error)

	// SignAndSendCreateCollectionTx will sign tx with key manager and send signed transaction
	SignAndSendCreateCollectionTx(tx *CreateCollectionTxInfo) (int64, error)

	// SignAndSendCancelOfferTx will sign tx with key manager and send signed transaction
	SignAndSendCancelOfferTx(tx *CancelOfferTxInfo) (string, error)

	// SignAndSendAtomicMatchTx will sign tx with key manager and send signed transaction
	SignAndSendAtomicMatchTx(tx *AtomicMatchTxInfo) (string, error)

	// SignAndSendOfferTx will sign tx with key manager and send signed transaction
	SignAndSendOfferTx(tx *OfferTxInfo) (string, error)

	// SignAndSendWithdrawNftTx will sign tx with key manager and send signed transaction
	SignAndSendWithdrawNftTx(tx *WithdrawNftTxInfo) (string, error)

	// SignAndSendTransferNftTx will sign tx with key manager and send signed transaction
	SignAndSendTransferNftTx(tx *TransferNftTxInfo) (string, error)
}

func NewZkbasSDK(url string) ZkbasSDK {
	return &client{
		zkbasURL: url,
	}
}
