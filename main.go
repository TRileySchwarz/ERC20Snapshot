package main

import (
	"fmt"
	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

func main() {

	// tokenAddress := "0xc7bba5b765581efb2cdd2679db5bea9ee79b201f"
	// walletAddress := "0xf9c317bcfe4dae770dfb9639fd821d54ffdb3457"
	// fromBlock := "5097900"
	// toBlock := "5480750"


	fmt.Println("\n")

	// address := "0xE6AF66a8345a8A3E80D27740B72d682272aA11dB"
	// lib.GetEthLog(address)

	// hash := "0x06d0eb420066385183109fdc851a944cd17dc4bff1a339f8f5f69412f9c14115"
	// lib.GetTxByHash(hash)

	blockHex := "0x06d0"
	lib.GetBlockByNumber(blockHex)

	fmt.Println("\n")
}
