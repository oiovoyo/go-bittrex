package bittrex

/*
{"Id":30776765,
"Amount":0.30443319,
"Currency":"BTC",
"Confirmations":2,
"LastUpdated":"2017-09-10T23:34:59.393",
"TxId":"d68b6c3b1ab88822db184f7162fffbc2091582feec4df45ffb9e8b1bb883cd18",
"CryptoAddress":"12VU5uv6tAV5cDo3q65uf9bVxZiWS8rC6P"}
*/

type Deposit struct {
	Id            int64   `json:"Id"`
	Amount        float64 `json:"Amount"`
	Currency      string  `json:"Currency"`
	Confirmations int     `json:"Confirmations"`
	LastUpdated   jTime   `json:"LastUpdated"`
	TxId          string  `json:"TxId"`
	CryptoAddress string  `json:"CryptoAddress"`
}
