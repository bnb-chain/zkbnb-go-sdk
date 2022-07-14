package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/zkbas-go-sdk/client/abi"
)

type l1Client struct {
	provider string

	proxyContractAddress string
	bscClient            *ethclient.Client
	proxyContract        *abi.Zkbas

	privateKey *ecdsa.PrivateKey
}

func NewL1Client(provider, proxyContractAddress string) *l1Client {
	bscClient, err := ethclient.Dial(provider)
	if err != nil {
		panic("new bsc client error")
	}

	proxyContract, err := abi.NewZkbas(common.HexToAddress(proxyContractAddress), bscClient)
	if err != nil {
		panic("new proxy contract error")
	}

	return &l1Client{
		provider:             provider,
		proxyContractAddress: proxyContractAddress,
		bscClient:            bscClient,
		proxyContract:        proxyContract,
	}
}

func (c *l1Client) SetPrivateKey(key *ecdsa.PrivateKey) {
	c.privateKey = key
}

func (c *l1Client) DepositBNB(accountName string, amount *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	opts.Value = amount
	tx, err := c.proxyContract.DepositBNB(opts, accountName)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) DepositBEP20(token common.Address, accountName string, amount *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.DepositBEP20(opts, token, amount, accountName)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) DepositNft(nftL1Address common.Address, accountName string, nftL1TokenId *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.DepositNft(opts, accountName, nftL1Address, nftL1TokenId)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) RegisterZNS(name string, owner common.Address, pubKeyX [32]byte, pubKeyY [32]byte) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.RegisterZNS(opts, name, owner, pubKeyX, pubKeyY)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) CreatePair(tokenA common.Address, tokenB common.Address) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.CreatePair(opts, tokenA, tokenB)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) RequestFullExit(accountName string, asset common.Address) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.RequestFullExit(opts, accountName, asset)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) RequestFullExitNft(accountName string, nftIndex uint32) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.RequestFullExitNft(opts, accountName, nftIndex)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) UpdatePairRate(pairInfo abi.ZkbasPairInfo) (common.Hash, error) {
	opts, err := c.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.proxyContract.UpdatePairRate(opts, pairInfo)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) getTransactor() (*bind.TransactOpts, error) {
	if c.privateKey == nil {
		return nil, fmt.Errorf("private key is not set")
	}

	nonce, err := c.bscClient.PendingNonceAt(context.Background(), getAddressFromPrivateKey(c.privateKey))
	if err != nil {
		return nil, err
	}
	chainId, err := c.bscClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, chainId)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.bscClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func getAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("get public key error")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
