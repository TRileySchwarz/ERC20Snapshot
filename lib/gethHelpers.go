package lib

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Snapshot struct {
	TokenAddress string          `json:"tokenAddress,omitempty"`
	StartBlock   string          `json:"startBlock,omitempty"`
	EndBlock     string          `json:"endBlock,omitempty"`
	Balances     []WalletAddress `json:"balances,omitempty"`
}

type WalletAddress struct {
	Address       string `json:"address,omitempty"`
	WalletDetails Wallet `json:"walletDetails,omitempty"`
}

type Wallet struct {
	Balance string `json:"balance,omitempty"`
}

type ERC20Ledger struct {
	Wallets map[string]Wallet `json:"wallets,omitempty"`
}

// Base URL used for querying Etherscan API token tx
var EtherscanBaseURl = "https://api.etherscan.io/api?module=account&action=tokentx"

// The max number of results returned by the Etherscan API
var maxResults = 5000

// Token ledger responsible for holding the values obtained via Etherscan API
var TokenLedger = make(map[string]string)

// Used to indicate the last block number we parsed
var lastBlockParsed = big.NewInt(0)

// Used to indicate the total amount of minted tokens ie total supply in most cases for valid ERC20
var totalMintedAmount = ""

// Stores the zero address wallet
var zeroAddress = "0x0000000000000000000000000000000000000000"

// Indicates whether to print verbose msgs
var Verbose = false

// Indicates how many addresses are contained in the current query
var numAddress = 0

// The number of geth response errors we allow before exiting the program
var maxError = 5

// Number of go routines we can have at any given point
// If having problems with rate limits, reduce this value
var concurrencyLimit = 500

// mutex used to lock values while handling global variables
var mu sync.Mutex

// Used to build a ERC20 token balance snapshot at a given token address, block number, using a specified provider
func BuildSnapshot(tokenAddress string, provider string, block int64) {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract to query contract balance
	token, err := NewERC20Token(common.HexToAddress(tokenAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Returns all wallets located in the token according to Etherscan
	arrayOfWallets := GetTokenWallets(tokenAddress, block)

	// Set global value to amount of addresses
	numAddress = len(arrayOfWallets)

	// Compare the Geth values to our Etherscan results
	CheckGethValues(&arrayOfWallets, block, token)

	// Writes the values to a local csv file
	WriteToCsv(arrayOfWallets)
}

// Responsible for going through a list of wallet addresses and checking their balance against a geth node
func CheckGethValues(arrayOfWallets *[]string, block int64, token *ERC20Token) {
	// Create a channel to initialize concurrency with the GETH provider
	channel := make(chan string, concurrencyLimit)
	var waitGroup sync.WaitGroup

	// Check all of the balances obtained during Etherscan parse compared to Geth query
	for _, address := range *arrayOfWallets {
		// Add a value to the channel
		channel <- address
		// Adds a value to the wait group
		waitGroup.Add(1)

		// Initiate go routine
		go GetBalanceAtBlock(address, block, token, &waitGroup, channel)
	}

	// Wait until all concurrency calls have been made
	waitGroup.Wait()
}

// Takes in an array of wallet addresses and writes the corresponding values into a csv
func WriteToCsv(arrayOfWallets []string) {
	PrintVerbose("Start of WriteToCSV \n")

	// Creates a new blank file for storing the resulting csv
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("error creating the initial csv file", err)
	}
	defer file.Close()

	// Create a 2d array of strings to push into the csv writer
	storedValues := make([][]string, len(arrayOfWallets))

	// Populate the data struct to push to csv
	for i := range storedValues {
		currentWallet := arrayOfWallets[i]

		storedValues[i] = []string{currentWallet, TokenLedger[currentWallet]}
	}

	// Create a new writer responsible for handling the write to file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Parse the 2d array and input values into writer
	for _, value := range storedValues {
		if err := writer.Write(value); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Ensure there are no errors during the writing process
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	PrintVerbose("End of WriteToCSV\n")
}

// Takes in a given wallet address, a block number, and a ERC20Token contract
// Then the Geth node provider set upon instantiation is used to query the Ethereum Blockchain at
// the specific block for the amount of tokens held by that wallet
func GetBalanceAtBlock(walletAddress string, block int64, token *ERC20Token, wg *sync.WaitGroup, channel chan string) {
	defer func(){
		<- channel
		wg.Done()
	}()

	// Create a new CallOpts instance specifying the block of interest
	ops := &bind.CallOpts{
		BlockNumber: big.NewInt(block),
	}

	// Converts address to hex values
	hexAddress := common.HexToAddress(walletAddress)

	var balance *big.Int
	var err error
	errorTicker := 0

	// Loops the geth call in case of errors
	for {
		// Determine the balance of the wallet
		balance, err = token.BalanceOf(ops, hexAddress)
		if err != nil {
			log.Printf("Failed to retrieve token balance: %v", err)
		} else {
			break
		}

		// If we reach the limit of errors from the provider then shut down the program
		errorTicker++
		if(errorTicker > maxError){
			log.Fatalf("Reached the limit of error retries with the following: %v", err)
		}
	}

	// Create values to very our results from Etherscan mixed with the Geth Results
	// If confident that the Etherscan is only returning values we care about, then we comment out this step
	actual := balance.String()
	expected := TokenLedger[walletAddress]

	if actual != expected {
		log.Printf("\n           !!!   Mismatched balances, expected wallet: %v to contain %v, instead it contains %v \n",
			walletAddress,
			actual,
			expected,
		)
	}

	// Locks the global numAddress to ensure there is no over writing in another go routine
	mu.Lock()
	numAddress--
	fmt.Printf("Token balance for address %v - %v ... %v left \n ", walletAddress, balance, numAddress)
	mu.Unlock()
}

// We will use Etherscan to build a list of all wallets holding tokens at a given block.
// We can verify these numbers as total all the holders should equal the total supply.
// By using two different sources, Etherscan and the chosen Geth Node Provider, you can ensure your data is credible
// Versus relying on one source for both pieces of information
func GetTokenWallets(tokenAddress string, endBlock int64) []string {

	// Initial page number to paginate response
	pageNumber := 1
	// This is needed to deal with tokens that contain more than 10000 transfer events
	currentNumResults := maxResults

	// As long as our returned results array is greater than the maximum returned amount by the api,
	// it indicates we haven't reached the last page yet
	for currentNumResults == maxResults {

		// Etherscan API url
		url := EtherscanBaseURl +
			"&contractaddress=" +
			tokenAddress + "&page=" +
			strconv.Itoa(pageNumber) +
			"&offset=" +
			strconv.Itoa(maxResults) +
			"&sort=dsc" +
			"&endblock=" +
			strconv.FormatInt(endBlock, 10)

		log.Printf("Querying Etherscan API with the following url: %v \n", url)

		// Create httpGet Response
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Cast the body as a string
		bodyString := string(body)

		// Marshall the json body to response
		var response GetTxResponse
		json.Unmarshal([]byte(bodyString), &response)

		// Parse the array of transactions to store the concurrent balances
		for _, tx := range response.Result {
			ProcessTransfer(tx.From, tx.To, tx.Value, tx.BlockNumber)
		}

		pageNumber++
		currentNumResults = len(response.Result)
	}

	// Resets the zero address value to actual balance, assumes that nobody burnt tokens by sending to zero address
	// by accident...
	totalMintedAmount = TokenLedger[zeroAddress]
	TokenLedger[zeroAddress] = "0"

	// Once the ledger has been updated return an array of all the token holder wallet addresses
	return GetKeys(TokenLedger)
}

// When Verbose is set to true this will print imbedded messages, good for turning off console debugging etc.
func PrintVerbose(msg string) {
	if(Verbose) {
		log.Print(msg)
	}
}

// Consume a mapping of addresses and returns just an array of strings representing an array of the key values
func GetKeys(mapping map[string]string) []string {
	arrayToReturn := make([]string, len(mapping))

	i := 0

	for key, _ := range mapping {
		arrayToReturn[i] = key
		i += 1
	}

	return arrayToReturn
}

// Take in a transaction and updates the ledger balances
// This assumes transactions Are parsed chronologically
func ProcessTransfer(fromAddress string, toAddress string, amount string, blockNumber string) {
	// Checks that we are parsing the transactions in order
	blockNumberInt := big.NewInt(0)
	blockNumberInt.SetString(blockNumber, 10)

	if blockNumberInt.Cmp(lastBlockParsed) == -1 {
		log.Fatal("A previous block has been parsed that shouldn't have been")
	} else {
		lastBlockParsed = blockNumberInt
	}

	transferAmountInt := big.NewInt(0)
	transferAmountInt.SetString(amount, 10)

	fromAmount := TokenLedger[fromAddress]
	fromAmountInt := big.NewInt(0)
	fromAmountInt.SetString(fromAmount, 10)
	fromAmountInt = fromAmountInt.Sub(fromAmountInt, transferAmountInt)

	TokenLedger[fromAddress] = fromAmountInt.Text(10)

	toAmount := TokenLedger[toAddress]
	toAmountInt := big.NewInt(0)
	toAmountInt.SetString(toAmount, 10)
	toAmountInt = toAmountInt.Add(toAmountInt, transferAmountInt)

	TokenLedger[toAddress] = toAmountInt.Text(10)
}

