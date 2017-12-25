package bittrex

type Market struct {
	MarketCurrency string
	BaseCurrency string
	MarketCurrencyLong string
	BaseCurrencyLong string
	MinTradeSize float64
	MarketName string
	IsActive bool
	Created string
	Notice interface{}
	IsSponsored interface{}
	LogoUrl string
}
