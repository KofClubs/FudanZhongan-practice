package transfer_controller

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	url            = "https://ropsten.infura.io/v3"
	projectIDPath  = "./Project-ID.txt"
	privKeyHexPath = "./Private-Key.txt"
	toAddrHex      = "0x717d6ac45c741a41aAC254b78A1609D74Ec654bf"
	tokenAddrHex   = "0xD65F818D787b343F8E9d08dE921E910D9E86B3b6"
	txSignStr      = "transfer(address,uint256)"
)

// Print 打印代币交易的响应
func Print() {
	projectID, err := ioutil.ReadFile(projectIDPath)
	if err != nil {
		fmt.Println(err)
	}
	client, err := ethclient.Dial(url + "/" + string(projectID))
	if err != nil {
		fmt.Println(err)
	}

	privKeyByte, err := ioutil.ReadFile(privKeyHexPath)
	if err != nil {
		fmt.Println(err)
	}
	kr, err := crypto.HexToECDSA(string(privKeyByte))
	if err != nil {
		fmt.Println(err)
	}

	ku := kr.Public()
	kuECDSA, ok := ku.(*ecdsa.PublicKey)
	if !ok {
		return
	}

	fromAddr := crypto.PubkeyToAddress(*kuECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		fmt.Println(err)
	}

	tokenAddr := common.HexToAddress(tokenAddrHex)

	value := big.NewInt(0) /* 代币转账不消耗以太币 */

	gasLimit := uint64(300000) /* 手续费上限 */

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	methodID := crypto.Keccak256([]byte(txSignStr))[:4]
	// methodID := sha3.NewKeccak256().Write([]byte(txSignStr)).Sum(nil)[:4]            /* 调用代币的转账方法 */
	paddedAddr := common.LeftPadBytes(common.HexToAddress(toAddrHex).Bytes(), 32)    /* 向左增广到32位的接收地址 */
	paddedAmount := common.LeftPadBytes(big.NewInt(1000000000000000000).Bytes(), 32) /* 代币数目，1个，向左增广到32位 */
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddr...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddr, value, gasLimit, gasPrice, data)

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(networkID), kr)
	if err != nil {
		fmt.Println(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Tx Hash:\t" + signedTx.Hash().Hex())
}
