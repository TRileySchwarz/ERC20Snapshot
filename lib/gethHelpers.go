package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"net/http"
)

type GetLogPayload struct {
	Jsonrpc string         `json:"jsonrpc,omitempty"`
	Method  string         `json:"method,omitempty"`
	Params  []GetLogParams `json:"params,omitempty"`
	ID      int            `json:"id,omitempty"`
}

type GetLogParams struct {
	FromBlock string `json:"fromBlock,omitempty"`
	ToBlock string `json:"toBlock,omitempty"`
	Address   string `json:"address,omitempty"`
}

type GetTxPayload struct {
	Jsonrpc string   `json:"jsonrpc,omitempty"`
	Method  string   `json:"method,omitempty"`
	Params  []string `json:"params,omitempty"`
	ID      int      `json:"id,omitempty"`
}

type GetBlockByNumberPayload struct {
	Jsonrpc string        `json:"jsonrpc,omitempty"`
	Method  string        `json:"method,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id,omitempty"`
}

// var Provider = "https://mainnet.infura.io/"
var Provider = "https://geth-m1.etherparty.com/"
var RopstenProvider = "https://geth-r1.etherparty.com/"

func BuildSnapshot(tokenAddress string, fromBlock uint64, toBlock uint64) {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(Provider)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	token, err := NewERC20Token(common.HexToAddress(tokenAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)
	
	//GetEthLog(tokenAddress, fromBlock, toBlock)
}

func GetEthLog(address string, fromBlock uint64, toBlock uint64) {

	fromHexString := hexutil.EncodeUint64(fromBlock)
	toHexString := hexutil.EncodeUint64(toBlock)

	currentPayload := GetLogPayload{
		Jsonrpc: "2.0",
		Method:  "eth_getLogs",
		Params: []GetLogParams{GetLogParams{
			FromBlock: fromHexString,
			ToBlock: toHexString,
			Address:   address,
		}},
		ID: 74,
	}
	fmt.Printf("\n Using Payload: %+v \n", currentPayload)

	body, err := json.Marshal(currentPayload)
	if err != nil {
		fmt.Println(err)
	}

	var readerBody = bytes.NewReader(body)
	responseString := getPostRequest(readerBody)

	printGetLogs(responseString)
}

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


func getPostRequest(body *bytes.Reader) string {
	req, err := http.NewRequest("POST", Provider, body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(responseBody)
}

//func printRequestByteResult(bytesArray []byte) {
//
//	readerBody := string(bytesArray)
//
//	printGetTx(readerBody)
//}

func printGetLogs(bytesString string) {
	var response GetLogsResponse

	json.Unmarshal([]byte(bytesString), &response)
	bs, _ := json.Marshal(response)

	if(response.Error.Message != "") {
		fmt.Println("\n Oops, there was an error in the request!")
		fmt.Printf("%s", bs)
	} else {
		fmt.Printf("%s", bs)
	}
}

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
