package zecrey

type Asset struct {
	Id         uint32
	BalanceEnc string
}

type AccountInfo struct {
	Index     uint32
	Name      string
	AccountPk string
	Assets    []*Asset
}

type RawTx struct {
	TxType uint32
	TxInfo string //globalrpc => sendAddliquidity.go
}

type Block struct {
	BlockCommitment              string
	BlockHeight                  int64
	StateRoot                    string
	PriorityOperations           int64
	PendingOnchainOperationsHash string
	CommittedTxHash              string
	CommittedAt                  int64
	VerifiedTxHash               string
	VerifiedAt                   int64
	RawTxs                       []*RawTx
	BlockStatus                  int64
}
