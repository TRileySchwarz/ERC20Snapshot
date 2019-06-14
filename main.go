package main

import (
	"fmt"
	"github.com/TRileySchwarz/ERC20Snapshot/lib"
)

func main() {
	fmt.Println("\n -- Starting program -- \n")

	// 1, Set this to the desired block number of the snapshot
	blockNumber := 7464576
	// 2, Set this to the desired contract address
	tokenAddress := "0xff56Cc6b1E6dEd347aA0B7676C85AB0B3D08B0FA"
	// 3, Set this to the desired Geth provider, eg...
	// currentProvider := "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"
	currentProvider := lib.InfuraProvider
	// 4, Set this to indicate whether you want more descriptive console log messages
	lib.Verbose = true
	// 5, Set this to true if you dont want the final csv to contain entries of
	// wallets with zero balances at the block
	lib.IgnoreZeroBalance = true

	fmt.Println("Creating a snapshot for token: " + tokenAddress)
	fmt.Printf("At block: %v\n", blockNumber)
	fmt.Printf("Using provider: %v\n\n", currentProvider)

	lib.BuildSnapshot(tokenAddress, currentProvider, int64(blockNumber))

	fmt.Printf("\nThe total ledger supply at block %v is %v \n", blockNumber, lib.TotalMintedAmount)
	fmt.Printf("The total supply minted from the 0x0 address is %v", lib.TotalSupply)

	fmt.Println("\n -- Closing Program --")
}
