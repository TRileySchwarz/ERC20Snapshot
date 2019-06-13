package main

import (
	"fmt"
	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

var InfuraProvider = "https://mainnet.infura.io/v3/"
var VanbexProvider = "https://geth-m1.etherparty.com/"
var RopstenProvider = "https://geth-r1.etherparty.com/"

func main() {

	fmt.Println("\n")
	fmt.Println("-- Starting program --")
	fmt.Println("\n")

	blockNumber := 4995790
	// !!! The token address must be in this hex format, not all lowercase
	tokenAddress := "0xEA38eAa3C86c8F9B751533Ba2E562deb9acDED40"
	currentProvider := InfuraProvider + lib.ProjectID
	lib.Verbose = true

	fmt.Println("Creating a snapshot for token: " + tokenAddress)
	fmt.Printf("At block: %v\n", blockNumber)
	fmt.Printf("Using provider: %v\n\n", currentProvider)

	lib.BuildSnapshot(tokenAddress, currentProvider, int64(blockNumber))

	fmt.Println("\n")
	fmt.Println("-- Closing Program --")
	fmt.Println("\n")
}
