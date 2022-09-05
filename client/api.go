package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
	"github.com/bnb-chain/zkbnb-go-sdk/client/abi"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
)

type ZkBNBClient interface {
	ZkBNBQuerier
	ZkBNBTxSender
}

type ZkBNBQuerier interface {
	// GetBlocks returns total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*types.Block, error)

	// GetBlockByHeight returns block by height
	GetBlockByHeight(blockHeight int64) (*types.Block, error)

	// GetBlockByCommitment returns block by commitment
	GetBlockByCommitment(blockCommitment string) (*types.Block, error)

	// GetTx returns tx by tx hash
	GetTx(hash string) (*types.EnrichedTx, error)

	// GetTxsByAccountPk returns txs by account public key
	GetTxsByAccountPk(accountPk string, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByAccountName returns txs by account name
	GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxs returns txs list
	GetTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByAccountIndex returns txs list by account index
	GetTxsByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByBlockHeight return txs in block
	GetTxsByBlockHeight(blockHeight uint32) ([]*types.Tx, error)

	// GetMempoolTxs returns the mempool txs
	GetMempoolTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetMempoolTxsByAccountName returns the mempool txs by account name
	GetMempoolTxsByAccountName(accountName string) (total uint32, txs []*types.Tx, err error)

	// GetAccountByName returns account (mainly pubkey) by using account_name
	GetAccountByName(accountName string) (*types.Account, error)

	// GetAccounts returns accounts by query conditions
	GetAccounts(offset, limit uint32) (*types.Accounts, error)

	// GetAccountByPk returns account info by public key
	GetAccountByPk(accountPk string) (*types.Account, error)

	// GetAccountByIndex returns account info by account index
	GetAccountByIndex(accountIndex int64) (*types.Account, error)

	// GetNextNonce returns nonce of account
	GetNextNonce(accountIdx int64) (int64, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex int64) (uint64, error)

	// GetCurrencyPrice returns currency price by symbol
	GetCurrencyPrice(symbol string) (*types.CurrencyPrice, error)

	// GetCurrencyPrices returns all currency prices
	GetCurrencyPrices(offset, limit uint32) (*types.CurrencyPrices, error)

	// GetSwapAmount returns swap amount by request
	GetSwapAmount(pairIndex, assetId int64, assetAmount string, isFrom bool) (*types.SwapAmount, error)

	// GetPairs returns available pairs
	GetPairs(offset, limit uint32) (*types.Pairs, error)

	// GetLpValue returns lp value
	GetLpValue(pairIndex uint32, lpAmount string) (*types.LpValue, error)

	// GetPair returns pair by pair index
	GetPair(index uint32) (*types.Pair, error)

	// GetAssets returns asset list
	GetAssets(offset, limit uint32) (*types.Assets, error)

	// GetWithdrawGasFee returns withdraw gas fee
	GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (*big.Int, error)

	// GetGasFeeAssets returns gas fee asset list
	GetGasFeeAssets() (*types.GasFeeAssets, error)

	// GetGasFee returns gas fee for asset
	GetGasFee(assetId int64) (*big.Int, error)

	// Search returns data type by queried info
	Search(keyword string) (*types.Search, error)

	// GetLayer2BasicInfo returns layer 2 basic info
	GetLayer2BasicInfo() (*types.Layer2BasicInfo, error)

	// GetGasAccount returns gas account of layer 2
	GetGasAccount() (*types.GasAccount, error)

	// GetNftsByAccountIndex returns nfts by account index
	GetNftsByAccountIndex(accountIndex, offset, limit int64) (*types.Nfts, error)
}

type ZkBNBTxSender interface {
	// SetKeyManager sets the key manager for signing txs.
	SetKeyManager(keyManager accounts.KeyManager)

	// KeyManager returns the key manager for signing txs.
	KeyManager() accounts.KeyManager

	// SendRawTx sends signed raw transaction and returns tx id
	SendRawTx(txType uint32, txInfo string) (string, error)

	// MintNft will sign tx with key manager and send signed transaction
	MintNft(tx *types.MintNftTxReq, ops *types.TransactOpts) (string, error)

	// CreateCollection will sign tx with key manager and send signed transaction
	CreateCollection(tx *types.CreateCollectionReq, ops *types.TransactOpts) (string, error)

	// CancelOffer will sign tx with key manager and send signed transaction
	CancelOffer(tx *types.CancelOfferReq, ops *types.TransactOpts) (string, error)

	// AtomicMatch will sign tx with key manager and send signed transaction
	AtomicMatch(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (string, error)

	// WithdrawNft will sign tx with key manager and send signed transaction
	WithdrawNft(tx *types.WithdrawNftTxReq, ops *types.TransactOpts) (string, error)

	// TransferNft will sign tx with key manager and send signed transaction
	TransferNft(tx *types.TransferNftTxReq, ops *types.TransactOpts) (string, error)

	// Transfer will sign tx with key manager and send signed transaction
	Transfer(tx *types.TransferTxReq, ops *types.TransactOpts) (string, error)

	// Swap will sign tx with key manager and send signed transaction
	Swap(tx *types.SwapTxReq, ops *types.TransactOpts) (string, error)

	// AddLiquidity will sign tx with key manager and send signed transaction
	AddLiquidity(tx *types.AddLiquidityReq, ops *types.TransactOpts) (string, error)

	// RemoveLiquidity will sign tx with key manager and send signed transaction
	RemoveLiquidity(tx *types.RemoveLiquidityReq, ops *types.TransactOpts) (string, error)

	// Withdraw will sign tx with key manager and send signed transaction
	Withdraw(tx *types.WithdrawReq, ops *types.TransactOpts) (string, error)
}

type ZkBNBL1Client interface {
	// SetPrivateKey will set the private key of the l1 account
	SetPrivateKey(pk string) error

	// DepositBNB will deposit specific amount bnb to l2
	DepositBNB(accountName string, amount *big.Int) (common.Hash, error)

	// DepositBEP20 will deposit specific amount of bep20 token to l2
	DepositBEP20(token common.Address, accountName string, amount *big.Int) (common.Hash, error)

	// DepositNft will deposit specific nft to l2
	DepositNft(nftL1Address common.Address, accountName string, nftL1TokenId *big.Int) (common.Hash, error)

	// RegisterZNS will register account in l2
	RegisterZNS(name string, owner common.Address, value *big.Int, pubKeyX [32]byte, pubKeyY [32]byte) (common.Hash, error)

	// CreatePair will create swap pair in l2
	CreatePair(tokenA common.Address, tokenB common.Address) (common.Hash, error)

	// RequestFullExit will request full exit from l2
	RequestFullExit(accountName string, asset common.Address) (common.Hash, error)

	// RequestFullExitNft will request full nft exit from l2
	RequestFullExitNft(accountName string, nftIndex uint32) (common.Hash, error)

	// UpdatePairRate will update pair info in l2
	UpdatePairRate(pairInfo abi.ZkBNBPairInfo) (common.Hash, error)
}

func NewZkBNBClient(url string) ZkBNBClient {
	return &l2Client{
		endpoint: url,
	}
}

func NewZkBNBL1Client(provider, zkBasContract string) (ZkBNBL1Client, error) {
	bscClient, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	zkBASContractInstance, err := abi.NewZkBNB(common.HexToAddress(zkBasContract), bscClient)
	if err != nil {
		panic("new proxy contract error")
	}

	return &l1Client{
		bscClient:             bscClient,
		zkBASContractInstance: zkBASContractInstance,
	}, nil
}
