package cmd

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/rickiey/loggo"
	"net/http"
)

var LotusNodeAddr = "http://api.node.glif.io/rpc/v0"

var lapi api.FullNodeStruct
var ctx = context.Background()

func init() {
	//headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	headers := http.Header{
		"content-type": []string{"application/json"},
		//"Authorization": []string{"Bearer eyJh..............1gNY"},
	}
	closer, err := jsonrpc.NewMergeClient(context.Background(), LotusNodeAddr, "Filecoin", []interface{}{&lapi.Internal, &lapi.CommonStruct.Internal}, headers)
	if err != nil {
		loggo.Panicf("connecting with lotus failed: %s", err)
	}
	defer closer()

}
