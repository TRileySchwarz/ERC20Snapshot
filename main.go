package main

import (
	"fmt"
	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

func main() {

	fmt.Println("\n")

	// address := "0xE6AF66a8345a8A3E80D27740B72d682272aA11dB"
	// lib.GetEthLog(address)

	// hash := "0x06d0eb420066385183109fdc851a944cd17dc4bff1a339f8f5f69412f9c14115"
	// lib.GetTxByHash(hash)

	startBlock := uint64(5000000)
	numberOfBlocks := uint64(1000)

	for i := startBlock; i < startBlock+numberOfBlocks; i++ {
		lib.GetBlockByNumber(i)
	}

	fmt.Println("Done")
	fmt.Println("\n")
}
