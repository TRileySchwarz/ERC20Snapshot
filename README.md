# ERC20 Snapshot
A Go implementation of an ERC20 Snapshot tool. 

There is not currently an easy way to gather a series of Ethereum wallet addresses holding a specific token. In addition, 
there is also not a trivial way of determining a snapshot of token holders for a given block number. 
This tool consumes a token address and a block number and outputs a .csv file that contains a list of all 
addresses that held or currently holds at that point in time. 

The simplest way of achieving this is by parsing Etherscan transfer events emitted and building up an exhaustive 
list of addresses associated with those events. This means that for the tool to work properly the token must use events
correctly to ensure an accurate result. This means for every single token that changes possession or is minted needs to
be recorded. In this case minting means transferring from the "Zero Address". While we were parsing the addresses 
originally via Etherscan, we might as well take advantage of the fact we can aggregate the tokens balances to later 
verify against the Geth results. 

Once we have achieved a series of addresses that are holding or held tokens at one point, the Geth node is queried to 
see what the balance of that account was at a specific point in the chain(blocknumber). We pair this up against the balances
attained in the Etherscan parse, and in theory everything should match up. The code will output console logs 
indicating any discrepancies found. 

## Concurrency
When verifying the Etherscan results against the Geth node, we must make an individual call per address. In most
cases the token can have thousands of addresses we need to lookup. To make this process faster most Geth providers
have rather generous rate limits so multiple calls can be made without getting rate limited. There is a global value
inside of the Snapshot.go file that allows you to set the limit on the amount of go routines that can be called. If 
you are having issues with rate limits, reduce this number. 

TODO:
- Write tests for all functions
- Implement command line flags/parameters
- Write Medium article
- Get peer review
- Create better way of logging errors or mismatches in results for larger datasets
- Example use cases
- Add creative(gif) of program running

