package signer

type SignParam struct {
	generalSignature   string
	buyOfferSignature  string
	sellOfferSignature string
}

func NewSignParamForGeneral(generalSignature string) *SignParam {
	return &SignParam{
		generalSignature: generalSignature,
	}
}

func NewSignParamForAtomicMatch(buySignature, sellSignature string) *SignParam {
	return &SignParam{
		buyOfferSignature:  buySignature,
		sellOfferSignature: sellSignature,
	}
}
