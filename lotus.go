package main

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"net/http"
	"time"
)

var (
	lotusNode *node
)

type node struct {
	node   api.FullNode
	closer jsonrpc.ClientCloser
}

func Setup(url string) error {
	var err error
	lotusNode, err = getNode(url)
	if err != nil {
		return err
	}
	return err
}
func Node() api.FullNode {
	return lotusNode.node
}
func getNode(url string) (*node, error) {
	var headers http.Header
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	addr := "ws://" + url + "/rpc/v1"
	fullNode, closer, err := client.NewFullNodeRPCV1(ctx, addr, headers)
	if err != nil {
		fmt.Printf("[lotus] get lotus client from node[%s] err: %s", url, err.Error())
		if closer != nil {
			closer()
		}

		return nil, err
	}

	return &node{
		node:   fullNode,
		closer: closer,
	}, nil

}
