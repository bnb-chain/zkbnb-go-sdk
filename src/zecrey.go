package src

type zecrey struct{}

func (z *zecrey) IfRollbacksOccurred() (blockHeight uint32, err error){
	// todo implement
	return blockHeight, err
}
func (z *zecrey) GetAccountInfoByAccountName(accountName string) (account AccountInfo, err error){
	// todo implement
	return account, err
}
func (z *zecrey) GetMaxOfferId(accountIndex uint32) (offerId uint64, err error){
	// todo implement
	return offerId, err
}
func (z *zecrey) GetBlockByBlockHeight(blockHeight uint32) (block *Block, err error){
	// todo implement
	return block, err
}
