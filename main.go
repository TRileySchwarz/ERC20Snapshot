package main

import (
	"fmt"
	"github.com/TRileySchwarz/ERC20Snapshot/lib"
)

func main() {
	fmt.Println("\n -- Starting program -- \n")

	blockNumber := 4653195
	tokenAddress := "0x419c4dB4B9e25d6Db2AD9691ccb832C8D9fDA05E"
	currentProvider := lib.InfuraProvider
	lib.Verbose = true
	lib.IgnoreZeroBalance = true

	fmt.Println("Creating a snapshot for token: " + tokenAddress)
	fmt.Printf("At block: %v\n", blockNumber)
	fmt.Printf("Using provider: %v\n\n", currentProvider)

	lib.BuildSnapshot(tokenAddress, currentProvider, int64(blockNumber))

	fmt.Printf("\nThe total ledger supply at block %v is %v \n", blockNumber, lib.TotalMintedAmount)
	fmt.Printf("The total supply minted from the 0x0 address is %v", lib.TotalSupply)

	fmt.Println("\n -- Closing Program --")
}
