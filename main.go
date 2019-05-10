package main

import (
	//"fmt"

	"github.com/TRileySchwarz/EtherscanAPI/lib"
)

func main() {

	tokenAddress := "0xc7bba5b765581efb2cdd2679db5bea9ee79b201f"
	walletAddress := "0xf9c317bcfe4dae770dfb9639fd821d54ffdb3457"
	fromBlock := "5097900"
	toBlock := "5480750"

	lib.GetERC20Transactions(tokenAddress, walletAddress, fromBlock, toBlock)
}
