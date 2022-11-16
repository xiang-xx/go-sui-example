package main

import (
	"context"
	"go-sui-example/common"
	"net/http"
	"time"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const TestnetRpcUrl = "https://fullnode.testnet.sui.io"

func main() {
	c, err := client.DialWithClient(TestnetRpcUrl, &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    3,
			IdleConnTimeout: 30 * time.Second,
		},
		Timeout: 90 * time.Second,
	})
	common.PanicIfError(err)

	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	common.PanicIfError(err)
	objs, err := c.GetObjectsOwnedByAddress(context.Background(), *address)
	common.PanicIfError(err)
	for _, obj := range objs {
		println(obj.Type)
		detail, err := c.GetObject(context.Background(), *obj.ObjectId)
		common.PanicIfError(err)
		println(detail.Status)
	}

	// transfer
	txn, err := c.SplitCoinEqual(context.Background(), *address, *objs[0].ObjectId, 10, objs[1].ObjectId, 100000)
	common.PanicIfError(err)
	acc := common.GetAccount()
	signedTxn := txn.SignWith(acc.PrivateKey)
	resp, err := c.ExecuteTransaction(context.Background(), *signedTxn, types.TxnRequestTypeWaitForLocalExecution)
	common.PanicIfError(err)
	println(resp.TxCert)
}
