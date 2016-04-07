package finance

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

const (
	optionsURL = "http://www.google.com/finance/option_chain?"
)

type timeMap struct {
	Month string `json:"m"`
	Day   string `json:"d"`
	Year  string `json:"y"`
}

type fetchResult struct {
	Expiry      json.RawMessage     `json:"expiry"`
	Expirations []timeMap           `json:"expirations"`
	Underlying  json.RawMessage     `json:"underlying_id"`
	Price       string              `json:"underlying_price"`
	Calls       []map[string]string `json:"calls,omitempty"`
	Puts        []map[string]string `json:"puts,omitempty"`
}

// NewOptionsChain creates a new OptionChain instance.
func NewOptionsChain(symbol string) (oc *OptionChain, err error) {

	oc = &OptionChain{Symbol: symbol}
	oc.Expirations, oc.UnderlyingPrice, err = getExpirations(symbol)

	return oc, err
}

func getExpirations(symbol string) (expirations []time.Time, price decimal.Decimal, err error) {

	params := map[string]string{
		"q":      symbol,
		"expd":   "4",
		"expm":   "4",
		"expy":   "2014",
		"output": "json",
	}

	url := buildURL(optionsURL, params)
	b, err := request(url)
	if err != nil {
		return expirations, price, err
	}

	var fr fetchResult
	err = json.Unmarshal(b, &fr)
	if err != nil {
		return expirations, price, err
	}

	for _, t := range fr.Expirations {
		dString := t.Month + "/" + t.Day + "/" + t.Year
		expirations = append(expirations, parseDate(dString))
	}

	price = toDecimal(fr.Price)

	return expirations, price, err
}

// GetOptionsExpiringNext fetches calls and puts with the shortest expiration date.
func (chain *OptionChain) GetOptionsExpiringNext() (err error) {

	expiry := chain.Expirations[0]

	params := map[string]string{
		"q":      chain.Symbol,
		"expd":   strconv.Itoa(expiry.Day()),
		"expm":   strconv.Itoa(int(expiry.Month())),
		"expy":   strconv.Itoa(expiry.Year()),
		"output": "json",
	}

	url := buildURL(optionsURL, params)

	b, err := request(url)
	if err != nil {
		return err
	}

	var fr fetchResult
	err = json.Unmarshal(b, &fr)
	if err != nil {
		return err
	}

	chain.Calls = NewContractSlice(fr.Calls)
	chain.Puts = NewContractSlice(fr.Puts)

	return err
}
