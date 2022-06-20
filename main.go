package main

import (
	"fmt"

	"github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"
	"github.com/zecrey-labs/zecrey-legend-go-sdk/sdk/parser"
)

func main() {
	zecrey := sdk.NewZecrey()
	account, err := zecrey.GetAccountInfoByAccountName("gas.legend")
	fmt.Println(account, err)

	txs, err := zecrey.GetTxsListByBlockHeight(1)
	for _, tx := range txs {
		fmt.Println(txs, err)
		txInfo, err := parser.ParseFullExitTxInfo(tx.TxInfo)
		fmt.Println(txInfo, err)
	}
}
