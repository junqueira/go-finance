package finance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getFixtureAsTable(filename string) [][]string {

	path := "./fixtures/" + filename
	f, err := os.Open(path)
	r := csv.NewReader(bufio.NewReader(f))
	r.FieldsPerRecord = -1
	table, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return table
}

func getFixtureAsString(filename string) string {

	path := "./fixtures/" + filename
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

func startTestServer(fixtureFile string) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getFixtureAsString(fixtureFile))
	}))
}

func Test_GetQuotesTable(t *testing.T) {

	ts := startTestServer("quotes_fixture.csv")
	defer ts.Close()

	quotes, err := getQuotesTable(ts.URL)
	assert.Nil(t, err)

	// Then the returned slice of rows should have a length of 2.
	assert.Len(t, quotes, 2)

	aapl := quotes[0][0]
	twtr := quotes[1][0]

	// And the first row's symbol should be AAPL.
	assert.Equal(t, "AAPL", aapl)
	// And the second row's symbol should be TWTR.
	assert.Equal(t, "TWTR", twtr)

}

func Test_GenerateQuotes(t *testing.T) {

	// Given we have a multi-quote csv.
	table := getFixtureAsTable("quotes_fixture.csv")

	// When we generate quotes,
	quotes := generateQuotes(table)

	// Then the returned slice of quote pointers should have a length of 2.
	assert.Len(t, quotes, 2)

	aapl := quotes[0]
	twtr := quotes[1]

	// And the first quote symbol should be AAPL.
	assert.Equal(t, "AAPL", aapl.Symbol)
	// And the second quote symbol should be TWTR.
	assert.Equal(t, "TWTR", twtr.Symbol)

}
