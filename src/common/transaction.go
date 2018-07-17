package common

type Transaction struct {
	BlockNumber        int64   `bson:"blockNumber" json:"blockNumber"`                    //"blockNumber":"5176578",
	TimeStamp          int64   `bson:"timeStamp" json:"timeStamp,string"`                 //"timeStamp":"1519898336",
	Hash               string  `json:"hash" bson:"_id,omitempty"`                         //"hash":"0x05993680c24fc81909daab444d224c04ddad4ffbd514d5b24bc004ee801b1c6d",
	Nonce              int64   `bson:"nonce" json:"nonce,string"`                         //"nonce":"1",
	BlockHash          string  `bson:"blockHash" json:"blockHash"`                        //"blockHash":"0xa5a1dc164508738eff07ad5bb00dcae1f3f2d304a70cffb7869bca70c37b67ee",
	TransactionIndex   int64   `bson:"transactionIndex" json:"transactionIndex,string"`   //"transactionIndex":"8",
	From               string  `bson:"from" json:"from"`                                  //"from":"0x95343e65c188952ad41a56869b7cb6d89df8dd25",
	To                 string  `to:"from" json:"to"`                                      //"to":"0xfedae5642668f8636a11987ff386bfd215f942ee",
	Value              float64 `bson:"value" json:"value,string"`                         //"value":"0",
	Gas                int64   `bson:"gas" json:"gas,string"`                             //"gas":"80137",
	GasPrice           float64 `bson:"gasPrice" json:"gasPrice,string"`                   //"gasPrice":"41000000000",
	IsError            int64   `bson:"isError" json:"isError,string"`                     //"isError":"0",
	TxReceiptStatus    int64   `bson:"txreceipt_status" json:"txreceipt_status,string"`   //"txreceipt_status":"1",
	Input              string  `bson:"input" json:"input"`                                //"input":"0xa9059cbb000000000000000000000000ab0f6b8c486fec656b270ff2b53ae09e454e12e10000000000000000000000000000000000000000000034511860c551606c0000",
	ContractAddress    string  `bson:"contractAddress" json:"contractAddress"`            //"contractAddress":"",
	CommulativeGasUsed float64 `bson:"cumulativeGasUsed" json:"cumulativeGasUsed,string"` //"cumulativeGasUsed":"672102",
	GasUsed            float64 `bson:"gasUsed" json:"gasUsed"`                            //"gasUsed":"53425",
	Confirmations      int64   `bson:"confirmations" json:"confirmations"`                //"confirmations":"585807"
}
