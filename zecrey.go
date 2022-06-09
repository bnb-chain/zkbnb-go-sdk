package zecrey

type zecrey struct{}

func (z *zecrey) IfRollbacksOccurred() (blockHeight uint32, err error)
func (z *zecrey) GetAccountInfoByAccountName(accountName string) (AccountInfo, error)
func (z *zecrey) GetMaxOfferId(accountIndex uint32) (OfferId uint64, err error)
func (z *zecrey) GetBlockByBlockHeight(blockHeight uint32) ([]*Block, error)
