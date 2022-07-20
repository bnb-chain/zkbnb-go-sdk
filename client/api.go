package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/zkbas-go-sdk/accounts"
	"github.com/bnb-chain/zkbas-go-sdk/client/abi"
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
	GetMaxOfferId(accountIndex int64) (uint64, error)

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
	GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (*big.Int, error)

	// GetGasFeeAssetList returns gas fee asset list
	GetGasFeeAssetList() (*types.RespGetGasFeeAssetList, error)

	// GetGasFee returns gas fee for asset
	GetGasFee(assetId int64) (*big.Int, error)

	// Search returns data type by queried info
	Search(info string) (*types.RespSearch, error)

	// GetLayer2BasicInfo returns layer 2 basic info
	GetLayer2BasicInfo() (*types.RespGetLayer2BasicInfo, error)
}

type ZkBASTxSender interface {
	// SetKeyManager sets the key manager for signing txs.
	SetKeyManager(keyManager accounts.KeyManager)

	// KeyManager returns the key manager for signing txs.
	KeyManager() accounts.KeyManager

	// SendRawTx sends signed raw transaction and returns tx id
	SendRawTx(txType uint32, txInfo string) (string, error)

	// SendRawMintNftTx sends signed raw mint nft transaction and returns nft id
	SendRawMintNftTx(txInfo string) (int64, error)

	// SendRawCreateCollectionTx sends signed raw create collection transaction and returns collection id
	SendRawCreateCollectionTx(txInfo string) (int64, error)

	// MintNft will sign tx with key manager and send signed transaction
	MintNft(tx *types.MintNftTxInfo, ops *types.TransactOpts) (int64, error)

	// CreateCollection will sign tx with key manager and send signed transaction
	CreateCollection(tx *types.CreateCollectionTxInfo, ops *types.TransactOpts) (int64, error)

	// CancelOffer will sign tx with key manager and send signed transaction
	CancelOffer(tx *types.CancelOfferTxInfo, ops *types.TransactOpts) (string, error)

	// AtomicMatch will sign tx with key manager and send signed transaction
	AtomicMatch(tx *types.AtomicMatchTxInfo, ops *types.TransactOpts) (string, error)

	// WithdrawNft will sign tx with key manager and send signed transaction
	WithdrawNft(tx *types.WithdrawNftTxInfo, ops *types.TransactOpts) (string, error)

	// SendTransferNft will sign tx with key manager and send signed transaction
	SendTransferNft(tx *types.TransferNftTxInfo, ops *types.TransactOpts) (string, error)

	// Transfer will sign tx with key manager and send signed transaction
	Transfer(tx *types.TransferTxInfo, ops *types.TransactOpts) (string, error)

	// Swap will sign tx with key manager and send signed transaction
	Swap(tx *types.SwapTxInfo, ops *types.TransactOpts) (string, error)

	// AddLiquidity will sign tx with key manager and send signed transaction
	AddLiquidity(tx *types.AddLiquidityTxInfo, ops *types.TransactOpts) (string, error)

	// RemoveLiquidity will sign tx with key manager and send signed transaction
	RemoveLiquidity(tx *types.RemoveLiquidityTxInfo, ops *types.TransactOpts) (string, error)

	// Withdraw will sign tx with key manager and send signed transaction
	Withdraw(tx *types.WithdrawTxInfo, ops *types.TransactOpts) (string, error)
}

type ZkBASL1Client interface {
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
	UpdatePairRate(pairInfo abi.ZkbasPairInfo) (common.Hash, error)
}

func NewZkBASClient(url string) ZkBASClient {
	return &l2Client{
		endpoint: url,
	}
}

func NewZkBASL1Client(provider, zkBasContract string) (ZkBASL1Client, error) {
	bscClient, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	zkBASContractInstance, err := abi.NewZkbas(common.HexToAddress(zkBasContract), bscClient)
	if err != nil {
		panic("new proxy contract error")
	}

	return &l1Client{
		bscClient:             bscClient,
		zkBASContractInstance: zkBASContractInstance,
	}, nil
}
