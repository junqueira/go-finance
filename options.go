package finance

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
		return nil, err
	}

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
	result, err := getOptionsData(url)
	if err != nil {
		return expirations, price, err
	}

	return parseExpirations(result.Expirations), toDecimal(result.Price), err
}

// parseExpirations returns valid dates from the malformed input text.
func parseExpirations(tm []timeMap) (expirations []time.Time) {
	for _, t := range tm {
		dString := t.Month + "/" + t.Day + "/" + t.Year
		expirations = append(expirations, parseDate(dString))
	}
	return expirations
}

// getOptionsData fetches data from the endpoint and returns an intermediate result.
func getOptionsData(url string) (fr *fetchResult, err error) {

	b, err := request(url)
	if err != nil {
		return nil, fmt.Errorf("options fetch error:  (error was: %s)\n", err.Error())
	}

	err = json.Unmarshal(b, &fr)
	if err != nil {
		return nil, fmt.Errorf("options format error:  (error was: %s)\n", err.Error())
	}

	return fr, nil
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

	result, err := getOptionsData(url)
	if err != nil {
		return err
	}

	chain.Calls = newContractSlice(result.Calls)
	chain.Puts = newContractSlice(result.Puts)

	return err
}
