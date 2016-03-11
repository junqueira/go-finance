package finance

import (
	"fmt"
	"strings"

	"github.com/FlashBoys/go-finance/models"
)

// USDGBP -  Currency Pairs.
const (
	USDGBP = "USDGBP=X"
	USDEUR = "USDEUR=X"
	USDAUD = "USDAUD=X"
	USDCHF = "USDCHF=X"
	USDJPY = "USDJPY=X"
	USDCAD = "USDCAD=X"
	USDSGD = "USDSGD=X"
	USDNZD = "USDNZD=X"
	USDHKD = "USDHKD=X"

	GBPUSD = "GBPUSD=X"
	GBPEUR = "GBPEUR=X"
	GBPAUD = "GBPAUD=X"
	GBPCHF = "GBPCHF=X"
	GBPJPY = "GBPJPY=X"
	GBPCAD = "GBPCAD=X"
	GBPSGD = "GBPSGD=X"
	GBPNZD = "GBPNZD=X"
	GBPHKD = "GBPHKD=X"

	EURUSD = "EURUSD=X"
	EURGBP = "EURGBP=X"
	EURAUD = "EURAUD=X"
	EURCHF = "EURCHF=X"
	EURJPY = "EURJPY=X"
	EURCAD = "EURCAD=X"
	EURSGD = "EURSGD=X"
	EURNZD = "EURNZD=X"
	EURHKD = "EURHKD=X"

	AUDUSD = "AUDUSD=X"
	AUDGBP = "AUDGBP=X"
	AUDEUR = "AUDEUR=X"
	AUDCHF = "AUDCHF=X"
	AUDJPY = "AUDJPY=X"
	AUDCAD = "AUDCAD=X"
	AUDSGD = "AUDSGD=X"
	AUDNZD = "AUDNZD=X"
	AUDHKD = "AUDHKD=X"

	CHFGBP = "CHFGBP=X"
	CHFEUR = "CHFEUR=X"
	CHFAUD = "CHFAUD=X"
	CHFJPY = "CHFJPY=X"
	CHFCAD = "CHFCAD=X"
	CHFSGD = "CHFSGD=X"
	CHFNZD = "CHFNZD=X"
	CHFHKD = "CHFHKD=X"

	JPYUSD = "JPYUSD=X"
	JPYGBP = "JPYGBP=X"
	JPYEUR = "JPYEUR=X"
	JPYAUD = "JPYAUD=X"
	JPYCHF = "JPYCHF=X"
	JPYCAD = "JPYCAD=X"
	JPYSGD = "JPYSGD=X"
	JPYNZD = "JPYNZD=X"
	JPYHKD = "JPYHKD=X"

	CADUSD = "CADUSD=X"
	CADGBP = "CADGBP=X"
	CADEUR = "CADEUR=X"
	CADAUD = "CADAUD=X"
	CADCHF = "CADCHF=X"
	CADJPY = "CADJPY=X"
	CADSGD = "CADSGD=X"
	CADNZD = "CADNZD=X"
	CADHKD = "CADHKD=X"

	SGDUSD = "SGDUSD=X"
	SGDGBP = "SGDGBP=X"
	SGDEUR = "SGDEUR=X"
	SGDAUD = "SGDAUD=X"
	SGDCHF = "SGDCHF=X"
	SGDJPY = "SGDJPY=X"
	SGDCAD = "SGDCAD=X"
	SGDNZD = "SGDNZD=X"
	SGDHKD = "SGDHKD=X"

	NZDUSD = "NZDUSD=X"
	NZDGBP = "NZDGBP=X"
	NZDEUR = "NZDEUR=X"
	NZDAUD = "NZDAUD=X"
	NZDCHF = "NZDCHF=X"
	NZDJPY = "NZDJPY=X"
	NZDCAD = "NZDCAD=X"
	NZDSGD = "NZDSGD=X"
	NZDHKD = "NZDHKD=X"

	HKDUSD = "HKDUSD=X"
	HKDGBP = "HKDGBP=X"
	HKDEUR = "HKDEUR=X"
	HKDAUD = "HKDAUD=X"
	HKDCHF = "HKDCHF=X"
	HKDJPY = "HKDJPY=X"
	HKDCAD = "HKDCAD=X"
	HKDSGD = "HKDSGD=X"
	HKDNZD = "HKDNZD=X"
)

// GetCurrencyPairQuote fetches a single currency pair's quote from Yahoo Finance.
func GetCurrencyPairQuote(symbol string) models.FXPairQuote {

	params := map[string]string{
		"s": symbol,
		"f": strings.Join(models.QuoteFields[:], ""),
		"e": ".csv",
	}

	table, err := requestCSV(buildURL(quoteURL, params))
	if err != nil {
		fmt.Println("Error fetching pair: ", err)
		return models.FXPairQuote{}
	}

	return generatePairQuotes(table)[0]
}

func generatePairQuotes(table [][]string) (pairs []models.FXPairQuote) {

	for _, row := range table {
		pairs = append(pairs, models.NewFXPairQuote(row))
	}
	return pairs
}
