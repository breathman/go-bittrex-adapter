package bittrex

import (
	"encoding/json"
	"errors"
)

const (
	API_BASE = "https://bittrex.com/api"
	API_VERSION = "/v1.1"
)

type bittrex struct {
	Client *client
}

type jsonResponse struct {
	Success bool
	Message string
	Result json.RawMessage
}

var response jsonResponse

func New() *bittrex {
	return &bittrex{newClient(API_BASE + API_VERSION)}
}

func handleError(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

func handleResponse(body []byte) error {
	if err := json.Unmarshal(body, &response); err != nil {
		return errors.New(err.Error())
	}

	if err := handleError(response); err != nil {
		return err
	}

	return nil
}

func GetMarkets(c *client) ([]Market, error) {
	var (
		markets []Market
		err error
		body []byte
	)
	method := "/public/getmarkets"

	if body, err = c.do(c.Url + method); err != nil {
		return nil, err
	}

	if err = handleResponse(body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response.Result, &markets); err != nil {
		return nil, errors.New(err.Error())
	}

	return markets, nil
}

func GetCurrencies(c *client) ([]Currency, error) {
	var (
		currencies []Currency
		err error
		body []byte
	)
	method := "/public/getcurrencies"

	if body, err = c.do(c.Url + method); err != nil {
		return nil, err
	}

	if err = handleResponse(body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(response.Result, &currencies); err != nil {
		return nil, errors.New(err.Error())
	}

	return currencies, nil
}
