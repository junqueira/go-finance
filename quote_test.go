package finance

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_NewQuote(t *testing.T) {

	// Given we have a csv with quote data.
	table := getFixtureAsTable("quote_fixture.csv")

	// When we create a new quote instance,
	quote := newQuote(table[0])

	// Then quote should have some equal fields-
	assert.Equal(t, "AAPL", quote.Symbol)
	assert.Equal(t, "Apple Inc.", quote.Name)
	assert.Equal(t, decimal.NewFromFloat(108.435), quote.LastTradePrice)
	assert.Equal(t, 26303703, quote.Volume)
	assert.Equal(t, "NMS", quote.Exchange)
	assert.Equal(t, "USD", quote.Currency)
	assert.Equal(t, "601.23B", quote.MarketCap)
	assert.Equal(t, 44780600, quote.AvgDailyVolume)
	assert.Equal(t, "82.79B", quote.EBITDA)
}

// map[g:108.121 d1:4/7/2016 x:NMS e7:9.060 r1:2/11/2016 a:108.460 b4:23.134 c1:-2.525 k:134.540 l1:108.435 n:Apple Inc. o:110.000 t1:3:44pm a2:44780600 q:2/4/2016 s7:1.470 j4:82.79B m4:107.594 y:1.870 e8:10.000 r:11.539 b:108.450 j:92.000 p:110.960 p2:-2.276% d:2.080 s:AAPL j1:601.23B v:26303703 e:9.397 e9:1.780 p5:2.618 p6:4.796 h:110.420 r5:1.050 t8:134.210 m3:102.801 c4:USD]
