package tron

type Value struct {
	Amount       int64  `json:"amount"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
}

type Parameter struct {
	Value   Value  `json:"value"`
	TypeURL string `json:"type_url"`
}

type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}

type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Timestamp     int64      `json:"timestamp"`
}

type Transaction struct {
	Visible    bool    `json:"visible"`
	TxID       string  `json:"txID"`
	RawData    RawData `json:"raw_data"`
	RawDataHex string  `json:"raw_data_hex"`
}

type TransactionWithSign struct {
	Transaction
	Signature []string
}

type Broadcast struct {
	Result bool   `json:"result"`
	TxID   string `json:"txid"`
}

type TransResult struct {
	Total        int64       `json:"total"`
	Data         ResultData  `json:"data"`
	ContractMap  interface{} `json:"contractMap"`
	RangeTotal   int64       `json:"rangeTotal"`
	ContractInfo interface{} `json:"contractInfo"`
}

type ResultData struct {
	ID                  string `json:"id"`
	Block               int64  `json:"block"`
	TransactionHash     string `json:"transactionHash"`
	Timestamp           int64
	TransferFromAddress string `json:"transferFromAddress"`
	TransferToAddress   string `json:"transferToAddress"`
	Amount              int64  `json:"amount"`
	TokenName           string `json:"tokenName"`
	Confirmed           bool   `json:"confirmed"`
	Data                string `json:"data"`
	ContractRet         string `json:"contractRet"`
	Revert              bool   `json:"revert"`
	TokenInfo           TokenInfo
}

type TokenInfo struct {
	TokenID      string `json:"tokenId"`
	TokenAbbr    string `json:"tokenAbbr"`
	TokenName    string `json:"tokenName"`
	TokenDecimal int    `json:"tokenDecimal"`
	TokenCanShow int    `json:"tokenCanShow"`
	TokenType    string `json:"tokenType"`
	TokenLogo    string `json:"tokenLogo"`
	TokenLevel   string `json:"tokenLevel"`
}
