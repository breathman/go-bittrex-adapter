package bittrex

type Currency struct {
	Currency string
	CurrencyLong string
	MinConfirmation int
	TxFee float64
	IsActive bool
	CoinType string
	BaseAddress string
	Notice interface{}
}
