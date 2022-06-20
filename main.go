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
	fmt.Println(txs, err)
	// for _, tx := range txs {
	// 	txInfo, err := parser.ParseFullExitTxInfo(tx.TxInfo)
	// 	fmt.Println(txInfo, err)
	// }

	offerId, err := zecrey.GetMaxOfferId(1)
	fmt.Println(offerId, err)

	total, blocks, err := zecrey.GetBlocks(0, 1)
	fmt.Println(total, blocks, err)

	txId, err := zecrey.SendTx(parser.TxTypeTransfer, "")
	fmt.Println(txId, err)
}
