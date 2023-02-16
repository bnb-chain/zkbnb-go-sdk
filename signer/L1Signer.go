package signer

import (
	"crypto/ecdsa"
	accounts2 "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type L1Signer interface {
	Sign(body string) (string, error)
	GetPublicKey() string
	GetAddress() string
}

type DefaultL1Singer struct {
	privateKey *ecdsa.PrivateKey
}

func NewL1Singer(privateKey string) (L1Signer, error) {
	privateKeyInEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return &DefaultL1Singer{
		privateKey: privateKeyInEcdsa,
	}, nil
}

func (singer *DefaultL1Singer) Sign(body string) (string, error) {
	bodyHash := accounts2.TextHash([]byte(body))
	signatureBytes, err := crypto.Sign(bodyHash, singer.privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	signatureString := hexutil.Encode(signatureBytes)
	return signatureString, nil
}

func (singer *DefaultL1Singer) GetPublicKey() string {
	pubKey := crypto.FromECDSAPub(&singer.privateKey.PublicKey)
	return hexutil.Encode(pubKey)
}

func (singer *DefaultL1Singer) GetAddress() string {
	address := crypto.PubkeyToAddress(singer.privateKey.PublicKey)
	return address.Hex()
}
