package finance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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
