package client

import (
	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/types"
)

type ZkBASClient interface {
	ZkBASQuerier
	ZkBASTxSender
}

type ZkBASQuerier interface {
	// GetBlocks returns total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*types.Block, error)

	// GetBlockByHeight returns block by height
	GetBlockByHeight(blockHeight int64) (*types.Block, error)

	// GetBlockByCommitment returns block by commitment
	GetBlockByCommitment(blockCommitment string) (*types.Block, error)

	// GetTxByHash returns tx by tx hash
	GetTxByHash(txHash string) (*types.RespGetTxByHash, error)

	// GetTxsByPubKey returns txs by public key
	GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByAccountName returns txs by account name
	GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByAccountIndexAndTxType returns txs by account index and tx type
	GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsListByAccountIndex returns txs list by account index
	GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsListByBlockHeight return txs in block
	GetTxsListByBlockHeight(blockHeight uint32) ([]*types.Tx, error)

	// GetMempoolTxs returns the mempool txs
	GetMempoolTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetMempoolTxsByAccountName returns the mempool txs by account name
	GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*types.Tx, err error)

	// GetBalanceByAssetIdAndAccountName returns the balance by asset id and account name
	GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error)

	// GetAccountStatusByAccountName returns account status by account name
	GetAccountStatusByAccountName(accountName string) (*types.RespGetAccountStatusByAccountName, error)

	// GetAccountStatusByAccountPk returns account status by account public key
	GetAccountStatusByAccountPk(accountPk string) (*types.RespGetAccountStatusByAccountPk, error)

	// GetAccountInfoByAccountName returns account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*types.AccountInfo, error)

	// GetAccounts returns accounts by query conditions
	GetAccounts(offset, limit uint32) (*types.RespGetAccounts, error)

	// GetAccountInfoByPubKey returns account info by public key
	GetAccountInfoByPubKey(accountPk string) (*types.RespGetAccountInfoByPubKey, error)

	// GetAccountInfoByAccountIndex returns account info by account index
	GetAccountInfoByAccountIndex(accountIndex int64) (*types.RespGetAccountInfoByAccountIndex, error)

	// GetNextNonce returns nonce of account
	GetNextNonce(accountIdx int64) (int64, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex uint32) (uint64, error)

	// GetCurrencyPriceBySymbol returns currency price by symbol
	GetCurrencyPriceBySymbol(symbol string) (*types.RespGetCurrencyPriceBySymbol, error)

	// GetCurrencyPrices returns all currency prices
	GetCurrencyPrices() (*types.RespGetCurrencyPrices, error)

	// GetSwapAmount returns swap amount by request
	GetSwapAmount(req *types.ReqGetSwapAmount) (*types.RespGetSwapAmount, error)

	// GetAvailablePairs returns available pairs
	GetAvailablePairs() (*types.RespGetAvailablePairs, error)

	// GetLPValue returns lp value
	GetLPValue(pairIndex uint32, lpAmount string) (*types.RespGetLPValue, error)

	// GetPairInfo returns pair info by pair index
	GetPairInfo(pairIndex uint32) (*types.RespGetPairInfo, error)

	// GetAssetsList returns asset list
	GetAssetsList() (*types.RespGetAssetsList, error)

	// GetWithdrawGasFee returns withdraw gas fee
	GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error)

	// GetGasFeeAssetList returns gas fee asset list
	GetGasFeeAssetList() (*types.RespGetGasFeeAssetList, error)

	// GetGasFee returns gas fee for asset
	GetGasFee(assetId uint32) (int64, error)

	// Search returns data type by queried info
	Search(info string) (*types.RespSearch, error)

	// GetLayer2BasicInfo returns layer 2 basic info
	GetLayer2BasicInfo() (*types.RespGetLayer2BasicInfo, error)
}

type ZkBASTxSender interface {
	// SetKeyManager sets the key manager for signing txs.
	SetKeyManager(keyManager accounts.KeyManager)

	// SendRawTx sends signed raw transaction and returns tx id
	SendRawTx(txType uint32, txInfo string) (string, error)

	// SendRawMintNftTx sends signed raw mint nft transaction and returns nft id
	SendRawMintNftTx(txInfo string) (int64, error)

	// SendRawCreateCollectionTx sends signed raw create collection transaction and returns collection id
	SendRawCreateCollectionTx(txInfo string) (int64, error)

	// MintNft will sign tx with key manager and send signed transaction
	MintNft(tx *types.MintNftTxInfo) (int64, error)

	// CreateCollection will sign tx with key manager and send signed transaction
	CreateCollection(tx *types.CreateCollectionTxInfo) (int64, error)

	// CancelOffer will sign tx with key manager and send signed transaction
	CancelOffer(tx *types.CancelOfferTxInfo) (string, error)

	// AtomicMatch will sign tx with key manager and send signed transaction
	AtomicMatch(tx *types.AtomicMatchTxInfo) (string, error)

	// Offer will sign tx with key manager and send signed transaction
	Offer(tx *types.OfferTxInfo) (string, error)

	// WithdrawNft will sign tx with key manager and send signed transaction
	WithdrawNft(tx *types.WithdrawNftTxInfo) (string, error)

	// SendTransferNft will sign tx with key manager and send signed transaction
	SendTransferNft(tx *types.TransferNftTxInfo) (string, error)

	// Transfer will sign tx with key manager and send signed transaction
	Transfer(tx *types.TransferTxInfo) (string, error)

	// Swap will sign tx with key manager and send signed transaction
	Swap(tx *types.SwapTxInfo) (string, error)

	// AddLiquidity will sign tx with key manager and send signed transaction
	AddLiquidity(tx *types.AddLiquidityTxInfo) (string, error)

	// RemoveLiquidity will sign tx with key manager and send signed transaction
	RemoveLiquidity(tx *types.RemoveLiquidityTxInfo) (string, error)

	// Withdraw will sign tx with key manager and send signed transaction
	Withdraw(tx *types.WithdrawTxInfo) (string, error)
}

func NewZkBASClient(url string) ZkBASClient {
	return &client{
		endpoint: url,
	}
}
