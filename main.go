package main

import (
	"fmt"
	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

var InfuraProvider = "https://mainnet.infura.io/"
var VanbexProvider = "https://geth-m1.etherparty.com/"
var RopstenProvider = "https://geth-r1.etherparty.com/"

func main() {

	fmt.Println("\n")
	fmt.Println("-- Starting program --")
	fmt.Println("\n")


	blockNumber := 7913731
	tokenAddress := "0xEA38eAa3C86c8F9B751533Ba2E562deb9acDED40"
	currentProvider := InfuraProvider

	//hashOfTransferEvent := ""

	fmt.Println("Creating a snapshot for token: " + tokenAddress);
	fmt.Printf("At block: %v\n", blockNumber);
	fmt.Printf("Using provider: %v\n\n", currentProvider);


	lib.BuildSnapshot(tokenAddress, currentProvider, int64(blockNumber));

	fmt.Println("\n")
	fmt.Println("-- Closing Program --")
	fmt.Println("\n")
}



//lib.GetEthLog(address)

// hash := "0x06d0eb420066385183109fdc851a944cd17dc4bff1a339f8f5f69412f9c14115"
// lib.GetTxByHash(hash)

// startBlock := uint64(5000000)
// numberOfBlocks := uint64(1000)

// for i := startBlock; i < startBlock+numberOfBlocks; i++ {
// 	lib.GetBlockByNumber(i)
// }