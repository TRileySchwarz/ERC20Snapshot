# ERC20 Snapshot
A Go implementation of a ERC20 Snapshot tool. 

There is not currently an easy way to gather a series of Ethereum wallet addresses holding a specific token. In addition, 
there is also not a trivial way of determining a snapshot of token holders for a given block number. 
This tool consumes a token address and a block number and outputs a .csv file that contains a list of all addresses that held or currently
holds at that point in time. 

# Things to explain
- Can use any geth provider
- Should specify a project id if using infura to avoid rate limit
- explain concurrency method and reasoning
- Explain checks along the way and why it uses Etherscan to verify the balances of both
- Why you might have mismatched balances... Invalid ERC20 implementations
- Rate limits
- How to adjust concurrency limits

- Example use cases
- Create gif of the program running
- Explain bindings/link the tool / tutorial
- Why we paginate results

TODO:
- Include option to store addresses with zero balances
- Finalize log statements
- Check with jackson the best way to store private variables
- Implement a double check for column totals of total supply etc. 
- Write tests for all functions
- Allow for variable flags in the command line running of this program
- Specify providor and error handling for bad providers
- Organize project into proper class names and folder organizations
- Pass off Jack and Sons to verify I didnt do something dumb
- Write article explaining this for Medium
- Set license create public repo
- Go over grammar in coding comments
- Check with path variables for loading to public repo with jacknson


