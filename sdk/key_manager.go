package sdk

import (
	"hash"

	"github.com/consensys/gnark-crypto/signature"
	"github.com/zecrey-labs/zecrey-crypto/ecc/ztwistededwards/tebn254"
)

type KeyManager interface {
	Sign(message []byte, hFunc hash.Hash) ([]byte, error)
	Public() signature.PublicKey
}

type SeedKeyManager struct {
	privateKey *tebn254.PrivateKey
}

func NewSeedKeyManager(seed string) (KeyManager, error) {
	key, err := tebn254.GenerateEddsaPrivateKey(seed)
	if err != nil {
		return nil, err
	}

	return &SeedKeyManager{privateKey: key}, nil
}

func (key *SeedKeyManager) Sign(message []byte, hFunc hash.Hash) ([]byte, error) {
	return key.privateKey.Sign(message, hFunc)
}

func (key *SeedKeyManager) Public() signature.PublicKey {
	return key.privateKey.Public()
}
