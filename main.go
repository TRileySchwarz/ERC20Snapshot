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

	blockHex := "0x06d0"
	lib.GetBlockByNumber(blockHex)

	fmt.Println("\n")
}
