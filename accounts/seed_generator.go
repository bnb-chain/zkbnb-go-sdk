package accounts

import (
	"fmt"
	"github.com/bnb-chain/zkbnb-go-sdk/signer"
	"strings"
)

const SEED_FORMAT = "Access zkbnb account.\n\nOnly sign this message for a trusted client!\nChain ID: %d."

func GenerateSeed(privateKey string, chainId uint64) (string, error) {
	signer, err := signer.NewL1Singer(privateKey)
	if err != nil {
		return "", err
	}

	messageText := fmt.Sprintf(SEED_FORMAT, chainId)
	seedString, err := signer.Sign(messageText)
	// if seedString starts with 0x as the prefix, directly trim the 0x prefix
	if strings.HasPrefix(seedString, "0x") {
		seedString = seedString[2:]
	}
	if err != nil {
		return "", err
	}
	return seedString, nil
}
