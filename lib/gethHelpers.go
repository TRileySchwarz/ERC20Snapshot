package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Address   string `json:"address,omitempty"`
}

type GetTxPayload struct {
	Jsonrpc string   `json:"jsonrpc,omitempty"`
	Method  string   `json:"method,omitempty"`
	Params  []string `json:"params,omitempty"`
	ID      int      `json:"id,omitempty"`
}

type GetBlockByNumberPayload struct {
	Jsonrpc string                   `json:"jsonrpc,omitempty"`
	Method  string                   `json:"method,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int                      `json:"id,omitempty"`
}

func GetEthLog(address string) {
	currentPayload := GetLogPayload{
		Jsonrpc: "2.0",
		Method:  "eth_getLogs",
		Params: []GetLogParams{GetLogParams{
			FromBlock: "0x539755",
			Address:   "0xE6AF66a8345a8A3E80D27740B72d682272aA11dB",
		}},
		ID: 74,
	}
	fmt.Printf("\n The payload created is as follows: %+v \n", currentPayload)

	body, err := json.Marshal(currentPayload)
	if err != nil {
		fmt.Println(err)
	}

	var readerBody = bytes.NewReader(body)
	responseString := getPostRequest(readerBody)

	printGetLogs(responseString)
}


func GetBlockByNumber(blockNumberHex string) {
	currentPayload := GetBlockByNumberPayload{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params: []interface{}{
			"0x558913",
			true,
		},
		ID: 1,
	}
	fmt.Printf("\n The payload created is as follows: %+v \n", currentPayload)

	body, err := json.Marshal(currentPayload)
	if err != nil {
		fmt.Println(err)
	}

	var readerBody = bytes.NewReader(body)
	responseString := getPostRequest(readerBody)

	printGetBlock(responseString)
}

func GetTxByHash(txHash string) {
	currentPayload := GetLogPayload{
		Jsonrpc: "2.0",
		Method:  "eth_getLogs",
		Params: []GetLogParams{GetLogParams{
			FromBlock: "0x539755",
			Address:   "0xE6AF66a8345a8A3E80D27740B72d682272aA11dB",
		}},
		ID: 74,
	}
	fmt.Printf("\n The payload created is as follows: %+v \n", currentPayload)

	body, err := json.Marshal(currentPayload)
	if err != nil {
		fmt.Println(err)
	}

	var readerBody = bytes.NewReader(body)
	responseString := getPostRequest(readerBody)

	printGetTx(responseString)
}

func getPostRequest(body *bytes.Reader) string {
	req, err := http.NewRequest("POST", "https://geth-r1.etherparty.com/", body)
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

func printRequestByteResult(bytesArray []byte) {

	readerBody := string(bytesArray)

	printGetTx(readerBody)
}

func printGetLogs(bytesString string) {
	var response GetLogsResponse

	json.Unmarshal([]byte(bytesString), &response)
	bs, _ := json.Marshal(response)

	fmt.Printf("%s", bs)
}

func printGetTx(bytesString string) {
	var response GetTxResponse

	json.Unmarshal([]byte(bytesString), &response)
	bs, _ := json.Marshal(response)

	fmt.Printf("%s", bs)
}

func printGetBlock(bytesString string) {
	var response BlockResponse

	json.Unmarshal([]byte(bytesString), &response)
	
	//bs, _ := json.Marshal(response)
	//fmt.Printf("%s", bs)

	printTxValues(response.Result.Transactions)
}

func printTxValues(arrayOfTransactions []TxResponse) {

	for _, tx := range arrayOfTransactions {
		fmt.Println(tx.Hash)
	}
}
