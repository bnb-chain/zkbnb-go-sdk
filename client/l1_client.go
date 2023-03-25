package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/bnb-chain/zkbnb-eth-rpc/core"
	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type L1Client struct {
	*rpc.ProviderClient
	ZkbnbContractInstance *core.ZkBNB
	PrivateKey            *ecdsa.PrivateKey
}

func (c *L1Client) SetPrivateKey(pk string) error {
	key, err := crypto.HexToECDSA(pk)
	if err != nil {
		return err
	}
	c.PrivateKey = key
	return nil
}

func (c *L1Client) DepositBNB(l1Address string, amount *big.Int) (common.Hash, error) {
	tx, err := c.DepositBNBWithTxReturn(l1Address, amount)
	return tx.Hash(), err
}

func (c *L1Client) DepositBNBWithTxReturn(l1Address string, amount *big.Int) (*types.Transaction, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return nil, err
	}

	opts.Value = amount
	tx, err := c.ZkbnbContractInstance.DepositBNB(opts, common.HexToAddress(l1Address))
	if err != nil {
		return tx, err
	}
	return tx, nil
}

func (c *L1Client) DepositBEP20(token common.Address, l1Address string, amount *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.ZkbnbContractInstance.DepositBEP20(opts, token, amount, common.HexToAddress(l1Address))
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *L1Client) DepositNft(nftL1Address common.Address, l1Address string, nftL1TokenId *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.ZkbnbContractInstance.DepositNft(opts, common.HexToAddress(l1Address), nftL1Address, nftL1TokenId)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *L1Client) RequestFullExit(accountIndex uint32, asset common.Address) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.ZkbnbContractInstance.RequestFullExit(opts, accountIndex, asset)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *L1Client) RequestFullExitNft(accountIndex uint32, nftIndex uint32) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}
	println(opts)

	//tx, err := c.ZkbnbContractInstance.RequestFullExitNft(opts, accountIndex, nftIndex,1)
	if err != nil {
		return common.Hash{}, err
	}
	//return tx.Hash(), nil
	return common.Hash{}, nil
}

func (c *L1Client) getTransactor(value *big.Int) (*bind.TransactOpts, error) {
	if c.PrivateKey == nil {
		return nil, fmt.Errorf("private key is not set")
	}

	nonce, err := c.PendingNonceAt(context.Background(), getAddressFromPrivateKey(c.PrivateKey))
	if err != nil {
		return nil, err
	}
	chainId, err := c.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value // in wei
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
