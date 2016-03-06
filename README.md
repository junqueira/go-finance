# go-finance [![Build Status](https://travis-ci.org/FlashBoys/go-finance.svg?branch=master)](https://travis-ci.org/FlashBoys/go-finance)

`go-finance` is a no-nonsense library for accessing the Yahoo! Finance API. It's not completely done yet - check Status section.

To install go-finance, use the following command:

```
go get github.com/FlashBoys/go-finance
```

`go-finance` aims to provide a clean, flexible, and thorough implementation of the full suite of Yahoo! Finance APIs. After much consideration and exploration of the labyrinthian endpoint system in the Yahoo Finance ecosystem, I've done my best to select and document proper consumption of the seemingly most stable endpoints. I've taken leaves out of the books of the following projects, chosen for their stability, wide-spread usage, completeness and accuracy of the financial data we all know to be available, as well as how much I like using them in other projects given their own concise syntax inherent in their design:

  * [pandas datareader](https://github.com/pydata/pandas-datareader) (Python) wide-spread use in academia.
  * [yahoofinance-api](https://github.com/sstrickx/yahoofinance-api) (Java) most popular java library for this purpose.



Given Yahoo! Finance's own perpetuation of a rabbit warren -like system of financial data YQL tables in varying degrees of deprecation, conflation/realtime, exchange availability, protocol access, and overall lack of consistent usage guidelines/documentation, I advise users of this library to be aware that you should not depend on it returning data to you 100% of the time. Build fail-safes and back-up plans into your own systems tasked with handling these cases as they arise. You should also probably complain to Yahoo to build better financial engineering tools since so many of us depend on them.

While dataframes (tabular data structures used for analytical operations atypical of what you see in the beaten track of web programming) are popular in the financial development community, those concepts are not the focus of this project. A good place to start there would be [Gota](https://github.com/kniren/gota). In the future, tabular data will be the focus! I just have to determine the easiest way to integrate it in my current projects. Anyways!- The primary technical tenants of this project are:

  * Make financial data easy and fun to work with in Go-lang.
  * Abstract the burden of non-sexy model serialization away from the user.


## Status

All of these are available through Yahoo finance:

- [x] Single security quotes
- [x] Multiple securities quotes
- [x] Single security quote history
- [ ] Single security dividend/split history
- [ ] Security look-up by symbol/name
- [ ] Option chains
- [ ] Currency pairs quotes
- [ ] International securities quotes
- [ ] Sector/Industry compositions
- [ ] Key Stats
- [ ] Test Coverage w/ mocks


## Features


### Single security quotes

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

  // 15-min delayed full quote for Apple.
	q := finance.GetQuote("AAPL")
	fmt.Println(q)

}
```

### Multiple securities quotes

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

  // 15-min delayed full quotes for Apple, Twitter, and Facebook.
  symbols := []string{"AAPL", "TWTR", "FB"}
	quotes := finance.GetQuotes(symbols)
	fmt.Println(quotes)

}
```

### Quote history

```go
package main

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {

  // Set time bounds to 1 month starting Jan. 1.
	start, _ := time.Parse(time.RFC3339, "2016-01-01T16:00:00+00:00")
	end := start.AddDate(0, 1, 0)

	// Request daily history for TWTR.
	// IntervalDaily OR IntervalWeekly OR IntervalMonthly are supported.
	bars := finance.GetQuoteHistory("TWTR", start, end, finance.IntervalDaily)
	fmt.Println(bars)

}
```
