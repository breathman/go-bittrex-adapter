package bittrex

import (
	"net/http"
	"errors"
	"io/ioutil"
)

type client struct {
	Url string
}

func newClient(url string) (*client) {
	return &client{url}
}

func (c *client) do(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return body, nil
}
