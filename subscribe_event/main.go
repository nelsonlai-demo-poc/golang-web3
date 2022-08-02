package main

import (
	"context"
	"fmt"
	"golang-web3/nft2309"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// https://oneclickdapp.com/carbon-field

func main() {
	key := os.Getenv("INFURA_KEY")
	fmt.Println("Key: ", key)

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/" + key)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x9CC3C79356C36C8A9D087849d171b24099cd329c")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(nft2309.Nft2309ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			fmt.Println("Log: ", vLog)

			i, err := contractAbi.Unpack("ConsecutiveTransfer", vLog.Data)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("FromTokenId: ", vLog.Topics[1].Big())
				fmt.Println("ToTokenId: ", i[0])
				fmt.Println("FromAddress: ", common.HexToAddress(vLog.Topics[2].Hex()))
				fmt.Println("ToAddress: ", common.HexToAddress(vLog.Topics[3].Hex()))
			}
		}
	}
}
