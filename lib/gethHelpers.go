package lib

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

//type GetLogPayload struct {
//	Jsonrpc string         `json:"jsonrpc,omitempty"`
//	Method  string         `json:"method,omitempty"`
//	Params  []GetLogParams `json:"params,omitempty"`
//	ID      int            `json:"id,omitempty"`
//}
//
//type GetLogParams struct {
//	FromBlock string `json:"fromBlock,omitempty"`
//	ToBlock string `json:"toBlock,omitempty"`
//	Address   string `json:"address,omitempty"`
//}
//
//type GetTxPayload struct {
//	Jsonrpc string   `json:"jsonrpc,omitempty"`
//	Method  string   `json:"method,omitempty"`
//	Params  []string `json:"params,omitempty"`
//	ID      int      `json:"id,omitempty"`
//}
//
//type GetBlockByNumberPayload struct {
//	Jsonrpc string        `json:"jsonrpc,omitempty"`
//	Method  string        `json:"method,omitempty"`
//	Params  []interface{} `json:"params,omitempty"`
//	ID      int           `json:"id,omitempty"`
//}

type Snapshot struct {
	TokenAddress string `json:"tokenAddress,omitempty"`
	StartBlock string `json:"startBlock,omitempty"`
	EndBlock string `json:"endBlock,omitempty"`
	Balances []WalletAddress `json:"balances,omitempty"`
}

type WalletAddress struct {
	Address string `json:"address,omitempty"`
	WalletDetails Wallet `json:"walletDetails,omitempty"`
}

type Wallet struct {
	Balance string `json:"balance,omitempty"`
}

type ERC20Ledger struct {
	Wallets map[string]Wallet `json:"wallets,omitempty"`
}

var EtherscanBaseURl = "https://api.etherscan.io/api?module=account&action=tokentx"

var TokenLedger = make(map[string]string)


func BuildSnapshot(tokenAddress string, provider string, block int64) {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	token, err := NewERC20Token(common.HexToAddress(tokenAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Returns all wallets located in the token
	arrayOfWallets := GetTokenWallets(tokenAddress, block)

	for _, element := range arrayOfWallets{
		GetBalanceAtBlock(element, block, token)
	}

}

func GetBalanceAtBlock(walletAddress string, block int64, token *ERC20Token) {
	ops := &bind.CallOpts{
		BlockNumber: big.NewInt(block),
	}

	hexAddress := common.HexToAddress(walletAddress)

	balance, err := token.BalanceOf(ops, hexAddress)
	if err != nil {
		log.Fatalf("Failed to retrieve token balance: %v", err)
	}

	actual := balance.String()
	expected := TokenLedger[walletAddress]

	if(actual != expected){
		fmt.Printf("\n UHOH expected %v to be == to %v \n", actual, expected)
	}

	fmt.Printf("Token balance for address %v - %v \n", walletAddress, balance)
}

// We will use Etherescan to build a list of all wallets holding tokens at a given block.
// We can verify these numbers as totalling all the holders should equal the total supply.
// By using two different sources, Etherscan and the chosen Geth Node, you can ensure your data is credible
// Versus relying on one source for both pieces of information
func GetTokenWallets(tokenAddress string, endBlock int64) ([]string) {
	pageNumber := 1
	maxResults := 1000

	for pageNumber < 2 {
		url := EtherscanBaseURl + "&contractaddress=" + tokenAddress + "&page=" + strconv.Itoa(pageNumber) + "&offset=" + strconv.Itoa(maxResults) + "&sort=asc" + "&endblock=" + strconv.FormatInt(endBlock, 10)

		println(url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			//return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			//return
		}

		test := string(body)

		var response GetTxResponse

		json.Unmarshal([]byte(test), &response)

		for _, tx := range response.Result {
			ProcessTransfer(tx.From, tx.To, tx.Value)
		}

		//bs, _ := json.Marshal(response)

		//fmt.Printf("%s", bs)

		pageNumber++
	}

	//fmt.Print(TokenLedger)

	//arrayToReturn := []string{"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", "0x0D0707963952f2fBA59dD06f2b425ace40b492Fe"}

	return GetKeys(TokenLedger)
}

func GetKeys(mapping map[string]string) ([]string){
	arrayToReturn := make([]string, len(mapping))

	i := 0

	for key, _ := range mapping {
		arrayToReturn[i] = key
		i += 1
	}

	return arrayToReturn
}

func ProcessTransfer(fromAddress string, toAddress string, amount string) {
	//fmt.Print(TokenLedger)

	fromAmount := TokenLedger[fromAddress];
	toAmount := TokenLedger[toAddress];

	fromAmountInt := big.NewInt(0)
	fromAmountInt.SetString(fromAmount, 10)

	toAmountInt := big.NewInt(0)
	toAmountInt.SetString(toAmount, 10)

	transferAmountInt := big.NewInt(0)
	transferAmountInt.SetString(amount, 10)

	fromAmountInt = fromAmountInt.Sub(fromAmountInt, transferAmountInt)
	toAmountInt = toAmountInt.Add(toAmountInt, transferAmountInt)

	TokenLedger[fromAddress] = fromAmountInt.Text(10)
	TokenLedger[toAddress] = toAmountInt.Text(10)

	//fmt.Print(TokenLedger)
}

//func GetEthLog(address string, fromBlock uint64, toBlock uint64) {
//
//	fromHexString := hexutil.EncodeUint64(fromBlock)
//	toHexString := hexutil.EncodeUint64(toBlock)
//
//	currentPayload := GetLogPayload{
//		Jsonrpc: "2.0",
//		Method:  "eth_getLogs",
//		Params: []GetLogParams{GetLogParams{
//			FromBlock: fromHexString,
//			ToBlock: toHexString,
//			Address:   address,
//		}},
//		ID: 74,
//	}
//	fmt.Printf("\n Using Payload: %+v \n", currentPayload)
//
//	body, err := json.Marshal(currentPayload)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	var readerBody = bytes.NewReader(body)
//	responseString := getPostRequest(readerBody)
//
//	printGetLogs(responseString)
//}

//func GetBlockByNumber(blockNumber uint64) {
//	hexString := hexutil.EncodeUint64(blockNumber)
//
//	currentPayload := GetBlockByNumberPayload{
//		Jsonrpc: "2.0",
//		Method:  "eth_getBlockByNumber",
//		Params: []interface{}{
//			hexString,
//			true,
//		},
//		ID: 1,
//	}
//	//fmt.Printf("\n The payload created is as follows: %+v \n", currentPayload)
//
//	body, err := json.Marshal(currentPayload)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	var readerBody = bytes.NewReader(body)
//	responseString := getPostRequest(readerBody)
//
//	printGetBlock(responseString)
//}


//func getPostRequest(body *bytes.Reader) string {
//	req, err := http.NewRequest("POST", Provider, body)
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//	defer resp.Body.Close()
//
//	responseBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//
//	return string(responseBody)
//}

//func printRequestByteResult(bytesArray []byte) {
//
//	readerBody := string(bytesArray)
//
//	printGetTx(readerBody)
//}

//func printGetLogs(bytesString string) {
//	var response GetLogsResponse
//
//	json.Unmarshal([]byte(bytesString), &response)
//	bs, _ := json.Marshal(response)
//
//	if(response.Error.Message != "") {
//		fmt.Println("\n Oops, there was an error in the request!")
//		fmt.Printf("%s", bs)
//	} else {
//		fmt.Printf("%s", bs)
//	}
//}

//func printGetTx(bytesString string) {
//	var response GetTxResponse
//
//	json.Unmarshal([]byte(bytesString), &response)
//	bs, _ := json.Marshal(response)
//
//	fmt.Printf("%s", bs)
//}

//func printGetBlock(bytesString string) {
//	var response BlockResponse
//
//	json.Unmarshal([]byte(bytesString), &response)
//
//	//bs, _ := json.Marshal(response)
//	//fmt.Printf("%s", bs)
//
//	//printTxValues(response.Result.Transactions)
//
//	//i, err := hexutil.DecodeUint64(response.Result.Number)
//	////i, err:= strconv.ParseInt("558913", 16, 64)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//fmt.Printf("There are %v transactions inside of block %v", len(response.Result.Transactions), i)
//}

//func printTxValues(arrayOfTransactions []TxResponse) {
//
//	fmt.Println(len(arrayOfTransactions))
//
//	//for _, tx := range arrayOfTransactions {
//	//	fmt.Println(tx.Hash)
//	//}
//}
