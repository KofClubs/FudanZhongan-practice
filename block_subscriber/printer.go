package block_subscriber

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	url           = "wss://ropsten.infura.io/ws/v3"
	projectIDPath = "./Project-ID.txt"
)

// Print 打印最新区块哈希和高度
func Print() {
	content, err := ioutil.ReadFile(projectIDPath)
	if err != nil {
		fmt.Println(err)
	}
	client, err := ethclient.Dial(url + "/" + string(content))
	if err != nil {
		fmt.Println(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				fmt.Println()
			}
			fmt.Println(block.NumberU64())
		}
	}
}
