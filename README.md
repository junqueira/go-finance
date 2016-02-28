# go-finance [![Build Status](https://travis-ci.org/FlashBoys/go-finance.svg?branch=master)](https://travis-ci.org/FlashBoys/go-finance)

`go-finance` is a no-nonsense library for accessing the Yahoo! Finance API.

To install go-finance, use the following command:

```
go get gopkg.in/go-finance.v1
```

`go-finance` aims to provide a clean, flexible, and thorough implementation of the full suite of Yahoo! Finance APIs. After much consideration and exploration of the labyrinthian endpoint system in the Yahoo Finance ecosystem, I've done my best to select and document proper consumption of the seemingly most stable endpoints. I've taken leaves out of the books of the following projects, chosen for their stability, wide-spread usage, completeness and accuracy of the financial data we all know to be available, as well as how much I like using them in other projects given their own concise syntax inherent in their design.

  * [pandas](https://github.com/pydata/pandas) (Python)
  * [yahoofinance-ap](https://github.com/sstrickx/yahoofinance-api) (Java)



Given Yahoo! Finance's own perpetuation of a rabbit warren -like system of financial data endpoints in varying degrees of deprecation, conflation/realtime, exchange availability, protocol access, and overall lack of consistent usage guidelines/documentation, I advise users of this library to be aware that you should not depend on it returning data to you 100% of the time. Build fail-safes and back-up plans into your own systems tasked with handling these cases as they arise. You should also probably complain to Yahoo to build better financial engineering tools since so many of us depend on them.

## Status

  * As of writing, coverage is hovering around 95%, but more thorough testing
    is always useful and desirable.


## Features

```go

```
