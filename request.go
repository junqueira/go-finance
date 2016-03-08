package finance

import (
	"encoding/csv"
	"net/http"
	"net/url"
)

// requestCSV fetches a csv from a supplied URL.
func requestCSV(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// buildURL takes a base URL and parameters returns the full URL.
func buildURL(base string, params map[string]string) string {

	url, _ := url.ParseRequestURI(base)
	q := url.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	url.RawQuery = q.Encode()

	return url.String()
}
