package client

import (
	"github.com/bnb-chain/zkbnb-eth-rpc/core"
	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/bnb-chain/zkbnb-go-sdk/accounts"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
)

type ZkBNBClient interface {
	ZkBNBQuerier
	ZkBNBTxSender
}

type getTxOption struct {
	Types    []int64
	FromHash string
}

type GetTxOptionFunc func(*getTxOption)

func GetTxWithTypes(txTypes []int64) GetTxOptionFunc {
	return func(o *getTxOption) {
		o.Types = txTypes
	}
}

// Get txs from the tx hash record
func GetTxWithFromHash(hash string) GetTxOptionFunc {
	return func(o *getTxOption) {
		o.FromHash = hash
	}
}

type ZkBNBQuerier interface {
	// GetCurrentHeight returns current block height
	GetCurrentHeight() (int64, error)

	// GetBlocks returns total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*types.Block, error)

	// GetBlockByHeight returns block by height
	GetBlockByHeight(blockHeight int64) (*types.Block, error)

	// GetBlockByCommitment returns block by commitment
	GetBlockByCommitment(blockCommitment string) (*types.Block, error)

	// GetTx returns tx by tx hash
	GetTx(hash string) (*types.EnrichedTx, error)

	// GetTxsByAccountPk returns txs by account public key
	GetTxsByAccountPk(accountPk string, offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error)

	// GetTxsByL1Address returns txs by account address
	GetTxsByL1Address(l1Address string, offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error)

	// GetTxs returns txs list
	GetTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetTxsByAccountIndex returns txs list by account index
	GetTxsByAccountIndex(accountIndex int64, offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error)

	// GetTxsByBlockHeight return txs in block
	GetTxsByBlockHeight(blockHeight uint32) ([]*types.Tx, error)

	// GetPendingTxs returns the pending txs
	GetPendingTxs(offset, limit uint32) (total uint32, txs []*types.Tx, err error)

	// GetPendingTxsByL1Address returns the pending txs by account address
	GetPendingTxsByL1Address(l1Address string, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error)

	// GetPendingTxs returns the executed txs
	GetExecutedTxs(offset, limit uint32, options ...GetTxOptionFunc) (total uint32, txs []*types.Tx, err error)

	// GetAccountByL1Address returns account (mainly pubkey) by account l1Address
	GetAccountByL1Address(l1Address string) (*types.Account, error)

	// GetAccounts returns accounts by query conditions
	GetAccounts(offset, limit uint32) (*types.Accounts, error)

	// GetAccountByPk returns account by public key
	GetAccountByPk(accountPk string) (*types.Account, error)

	// GetAccountByIndex returns account by account index
	GetAccountByIndex(accountIndex int64) (*types.Account, error)

	// GetNextNonce returns nonce of account
	GetNextNonce(accountIndex int64) (int64, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex int64) (uint64, error)

	// GetAssetById returns asset by asset id
	GetAssetById(id uint32) (*types.Asset, error)

	// GetAssetBySymbol returns asset by asset symbol
	GetAssetBySymbol(symbol string) (*types.Asset, error)

	// GetAssets returns asset list
	GetAssets(offset, limit uint32) (*types.Assets, error)

	// GetGasFeeAssets returns gas fee asset list
	GetGasFeeAssets() (*types.GasFeeAssets, error)

	// GetGasFee returns gas fee for asset
	GetGasFee(assetId int64, txType int) (*big.Int, error)

	// Search returns data type by queried keyword
	Search(keyword string) (*types.Search, error)

	// GetLayer2BasicInfo returns layer 2 basic info
	GetLayer2BasicInfo() (*types.Layer2BasicInfo, error)

	// GetGasAccount returns gas account of layer 2
	GetGasAccount() (*types.GasAccount, error)

	// GetNftsByAccountIndex returns nfts by account index
	GetNftsByAccountIndex(accountIndex, offset, limit int64) (*types.Nfts, error)

	// GetRollbacks returns tx rollback info
	GetRollbacks(fromBlockHeight, offset, limit int64) (total uint32, rollbacks []*types.Rollback, err error)

	// GetMaxCollectionId returns max collection id  by accountIndex
	GetMaxCollectionId(accountIndex int64) (*types.MaxCollectionId, error)

	// GetNftByTxHash returns nfts by txHash
	GetNftByTxHash(txHash string) (*types.NftIndex, error)

	// UpdateNftByIndex updates mutable attribute by NftIndex
	UpdateNftByIndex(nft *types.UpdateNftReq, signatureList ...string) (*types.Mutable, error)
}

type ZkBNBTxSender interface {

	// KeyManager returns the key manager for signing txs.
	KeyManager() accounts.KeyManager

	// SendRawTx sends signed raw transaction and returns tx hash
	SendRawTx(txType uint32, txInfo string) (string, error)

	// ChangePubKey will sign tx with key manager and send signed transaction
	ChangePubKey(tx *types.ChangePubKeyReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// GenerateSignBody generates the signature body for caller to calculate signature
	GenerateSignBody(txData interface{}, ops *types.TransactOpts) (string, error)

	// GenerateSignature generates the signature for l1 identifier validation
	GenerateSignature(privateKey string, txData interface{}) (string, error)

	// NOTE: You need to call SetKeyManager first before using following functions

	// MintNft will sign tx with key manager and send signed transaction
	MintNft(tx *types.MintNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// CreateCollection will sign tx with key manager and send signed transaction
	CreateCollection(tx *types.CreateCollectionTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// CancelOffer will sign tx with key manager and send signed transaction
	CancelOffer(tx *types.CancelOfferTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// AtomicMatch will sign tx with key manager and send signed transaction
	AtomicMatch(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) (string, error)

	// WithdrawNft will sign tx with key manager and send signed transaction
	WithdrawNft(tx *types.WithdrawNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// TransferNft will sign tx with key manager and send signed transaction
	TransferNft(tx *types.TransferNftTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// Transfer will sign tx with key manager and send signed transaction
	Transfer(tx *types.TransferTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)

	// Withdraw will sign tx with key manager and send signed transaction
	Withdraw(tx *types.WithdrawTxReq, ops *types.TransactOpts, signatureList ...string) (string, error)
}

type ZkBNBL1Client interface {
	// SetPrivateKey will set the private key of the l1 account
	SetPrivateKey(pk string) error

	// DepositBNB will deposit specific amount bnb to l2
	DepositBNB(l1Address string, amount *big.Int) (common.Hash, error)

	// DepositBEP20 will deposit specific amount of bep20 token to l2
	DepositBEP20(token common.Address, l1Address string, amount *big.Int) (common.Hash, error)

	// DepositNft will deposit specific nft to l2
	DepositNft(nftL1Address common.Address, l1Address string, nftL1TokenId *big.Int) (common.Hash, error)

	// RequestFullExit will request full exit from l2
	RequestFullExit(accountIndex uint32, asset common.Address) (common.Hash, error)

	// RequestFullExitNft will request full nft exit from l2
	RequestFullExitNft(accountIndex uint32, nftIndex uint32) (common.Hash, error)
}

func NewZkBNBClientWithSeed(url, seed string, chainId uint64) (ZkBNBClient, error) {
	keyManager, err := accounts.NewSeedKeyManager(seed)
	if err != nil {
		return nil, err
	}

	return &l2Client{
		endpoint:   url,
		privateKey: "",
		chainId:    chainId,
		l1Signer:   nil,
		keyManager: keyManager,
	}, nil
}

func NewZkBNBL1Client(provider, zkbnbContract string) (ZkBNBL1Client, error) {
	bscClient, err := rpc.NewClient(provider)
	if err != nil {
		panic(err)
	}
	zkbnbContractInstance, err := core.NewZkBNB(common.HexToAddress(zkbnbContract), bscClient)
	if err != nil {
		panic("new proxy contract error")
	}

	return &l1Client{
		bscClient:             bscClient,
		zkbnbContractInstance: zkbnbContractInstance,
	}, nil
}
