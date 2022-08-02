package main

import (
	"fmt"
	"golang-web3/nft2309"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	key := os.Getenv("INFURA_KEY")
	fmt.Println("Key: ", key)

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/" + key)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xc00c68d19F49E30caD09465D3688b1BEe6883782")
	nft, err := nft2309.NewNft2309(address, client)
	if err != nil {
		log.Fatal(err)
	}

	b, err := nft.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total supply: ", b)
}
