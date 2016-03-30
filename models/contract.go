package models

import "github.com/shopspring/decimal"

// Contract represents an instance of an option contract.
type Contract struct {
	ContractID    string
	Security      string
	Strike        decimal.Decimal
	Price         decimal.Decimal
	ChangeNominal decimal.Decimal
	ChangePercent decimal.Decimal
	Bid           decimal.Decimal
	Ask           decimal.Decimal
	Volume        int
	OpenInterest  int
}

// NewContract creates a new instance of an options contract.
func NewContract(option map[string]string) Contract {

	c := Contract{
		ContractID:    option["cid"],
		Security:      option["s"],
		Strike:        ToDecimal(option["strike"]),
		Price:         ToDecimal(option["p"]),
		ChangeNominal: ToDecimal(option["c"]),
		Bid:           ToDecimal(option["b"]),
		Ask:           ToDecimal(option["a"]),
		Volume:        ToInt(option["vol"]),
		OpenInterest:  ToInt(option["oi"]),
	}

	if c.Price.IntPart() != 0 {
		hundred, _ := decimal.NewFromString("100")
		c.ChangePercent = ((c.ChangeNominal).Div(c.Price.Sub(c.ChangeNominal))).Mul(hundred).Truncate(2)
	}

	return c
}

// NewContractSlice creates a new slice of contracts.
func NewContractSlice(options []map[string]string) (contracts []Contract) {

	for _, op := range options {
		contracts = append(contracts, NewContract(op))
	}

	return contracts
}
