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

	"github.com/bnb-chain/zkbnb-go-sdk/client/abi"
)

type l1Client struct {
	bscClient             *ethclient.Client
	zkbnbContractInstance *abi.ZkBNB
	privateKey            *ecdsa.PrivateKey
}

func (c *l1Client) SetPrivateKey(pk string) error {
	key, err := crypto.HexToECDSA(pk)
	if err != nil {
		return err
	}
	c.privateKey = key
	return nil
}

func (c *l1Client) DepositBNB(accountName string, amount *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	opts.Value = amount
	tx, err := c.zkbnbContractInstance.DepositBNB(opts, accountName)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) DepositBEP20(token common.Address, accountName string, amount *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.zkbnbContractInstance.DepositBEP20(opts, token, amount, accountName)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) DepositNft(nftL1Address common.Address, accountName string, nftL1TokenId *big.Int) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.zkbnbContractInstance.DepositNft(opts, accountName, nftL1Address, nftL1TokenId)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

// TODO: need query the charge fee
func (c *l1Client) RegisterZNS(name string, owner common.Address, value *big.Int, pubKeyX [32]byte, pubKeyY [32]byte) (common.Hash, error) {
	opts, err := c.getTransactor(value)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.zkbnbContractInstance.RegisterZNS(opts, name, owner, pubKeyX, pubKeyY)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) RequestFullExit(accountName string, asset common.Address) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.zkbnbContractInstance.RequestFullExit(opts, accountName, asset)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) RequestFullExitNft(accountName string, nftIndex uint32) (common.Hash, error) {
	opts, err := c.getTransactor(nil)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.zkbnbContractInstance.RequestFullExitNft(opts, accountName, nftIndex)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (c *l1Client) getTransactor(value *big.Int) (*bind.TransactOpts, error) {
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
