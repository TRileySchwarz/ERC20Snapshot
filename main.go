package main

import (
	"fmt"
	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

func main() {

	fmt.Println("\n")
	fmt.Println("-- Starting program --")
	fmt.Println("\n")


	fromBlockNumber := 7898966
	toblockNumber := 7901263
	tokenAddress := "0xEA38eAa3C86c8F9B751533Ba2E562deb9acDED40"

	//hashOfTransferEvent := ""

	fmt.Println("Creating a snapshot for token: " + tokenAddress);
	fmt.Printf("From blocks: %v to %v \n", fromBlockNumber, toblockNumber);
	fmt.Println("Using provider: " + lib.Provider);


	lib.BuildSnapshot(tokenAddress, uint64(fromBlockNumber), uint64(toblockNumber));

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