package lib

type GetLogsResponse struct {
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

type GetTxResponse struct {
	Jsonrpc string     `json:"jsonrpc,omitempty"`
	ID      int        `json:"id,omitempty"`
	Result  TxResponse `json:"result,omitempty"`
}

type TxResponse struct {
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	From             string `json:"from,omitempty"`
	Gas              string `json:"gas,omitempty"`
	GasPrice         string `json:"gasPrice,omitempty"`
	Hash             string `json:"hash,omitempty"`
	Input            string `json:"input,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	To               string `json:"to,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value            string `json:"value,omitempty"`
	V                string `json:"v,omitempty"`
	R                string `json:"r,omitempty"`
	S                string `json:"s,omitempty"`
}

type BlockResponse struct {
	ID      int                 `json:"id,omitempty"`
	Jsonrpc string              `json:"jsonrpc,omitempty"`
	Result  BlockResponseResult `json:"result,omitempty"`
}

type BlockResponseResult struct {
	Number           string       `json:"number,omitempty"`
	Hash             string       `json:"hash,omitempty"`
	ParentHash       string       `json:"parentHash,omitempty"`
	Nonce            string       `json:"nonce,omitempty"`
	Sha3Uncles       string       `json:"sha3Uncles,omitempty"`
	LogsBloom        string       `json:"logsBloom,omitempty"`
	TransactionsRoot string       `json:"transactionsRoot,omitempty"`
	StateRoot        string       `json:"stateRoot,omitempty"`
	Miner            string       `json:"miner,omitempty"`
	Difficulty       string       `json:"difficulty,omitempty"`
	TotalDifficulty  string       `json:"totalDifficulty,omitempty"`
	ExtraData        string       `json:"extraData,omitempty"`
	Size             string       `json:"size,omitempty"`
	GasLimit         string       `json:"gasLimit,omitempty"`
	GasUsed          string       `json:"gasUsed,omitempty"`
	Timestamp        string       `json:"timestamp,omitempty"`
	Transactions     []TxResponse `json:"transactions,omitempty"`
	Uncles           []string     `json:"uncles,omitempty"`
}
