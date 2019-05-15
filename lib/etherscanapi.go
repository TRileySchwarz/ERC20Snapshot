package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BaseURl = "https://api.etherscan.io/api?module=account&action=tokentx"

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
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	test := string(body)

	var response GetLogsResponse

	json.Unmarshal([]byte(test), &response)

	bs, _ := json.Marshal(response)

	fmt.Printf("%s", bs)
}
