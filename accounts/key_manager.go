package accounts

import (
	"hash"

	"github.com/bnb-chain/zkbnb-crypto/ecc/ztwistededwards/tebn254"
	"github.com/consensys/gnark-crypto/signature"
)

type Signer interface {
	Sign(message []byte, hFunc hash.Hash) ([]byte, error)
}

type KeyManager interface {
	Signer
	PubKey() signature.PublicKey
	PubKeyPoint() [2][32]byte
}

type seedKeyManager struct {
	key *tebn254.PrivateKey
}

func NewSeedKeyManager(seed string) (KeyManager, error) {
	key, err := tebn254.GenerateEddsaPrivateKey(seed)
	if err != nil {
		return nil, err
	}
	return &seedKeyManager{key: key}, nil
}

func (key *seedKeyManager) Sign(message []byte, hFunc hash.Hash) ([]byte, error) {
	return key.key.Sign(message, hFunc)
}

func (key *seedKeyManager) PubKey() signature.PublicKey {
	return key.key.Public()
}

func (key *seedKeyManager) PubKeyPoint() (res [2][32]byte) {
	copy(res[0][:], key.key.PublicKey.A.X.Marshal())
	copy(res[1][:], key.key.PublicKey.A.Y.Marshal())
	return res
}
