package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BaseURl = "https://api.etherscan.io/api?module=account&action=tokentx"

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Result  []Tx   `json:"result,omitempty"`
}

type Tx struct {
	BlockNumber     string `json:"blockNumber,omitempty"`
	TimeStamp       string `json:"timeStamp,omitempty"`
	Hash            string `json:"hash,omitempty"`
	From            string `json:"from,omitempty"`
	To              string `json:"to,omitempty"`
	Value           string `json:"value,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Input           string `json:"input,omitempty"`
	Type            string `json:"type,omitempty"`
	Gas             string `json:"gas,omitempty"`
	GasUsed         string `json:"gasUsed,omitempty"`
	TraceID         string `json:"traceId,omitempty"`
	IsError         string `json:"isError,omitempty"`
	ErrCode         string `json:"errCode,omitempty"`
}

func GetERC20Transactions(tokenAddress string, walletAddress string, fromBlock string, toBlock string) {

	url := BaseURl + "&contractaddress=" + tokenAddress + "&address=" + walletAddress
	// + "&page=" 1
	// + "&offset=" 100
	// + "&sort=" asc
	//+ "&apikey=" YourApiKeyToken

	println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	test := string(body)

	var response Response

	json.Unmarshal([]byte(test), &response)

	bs, _ := json.Marshal(response)

	fmt.Printf("%s", bs)
}
