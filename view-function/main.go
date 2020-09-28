package main

import (
	"fmt"
	"io/ioutil"
	"practice/view-function/view"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	url           = "https://ropsten.infura.io/v3"
	contractHex   = "0x8aC8dA2B230f891101Cb4d30a55A5abBcCBabE6d"
	projectIDPath = "../Project-ID.txt"
)

func main() {
	content, err := ioutil.ReadFile(projectIDPath)
	if err != nil {
		fmt.Println(err)
	}
	client, err := ethclient.Dial(url + "/" + string(content))
	if err != nil {
		fmt.Println(err)
	}
	v, err := view.NewView(common.HexToAddress(contractHex), client)
	if err != nil {
		fmt.Println(err)
	}
	str, err := v.GetStr(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
