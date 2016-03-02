package finance

import (
	"database/sql"
	"fmt"

	// YQL Driver.
	_ "github.com/mattn/go-yql"
)

// Client manages yahoo developer authentication.
type Client struct {
	*sql.DB
}

// NewAuthenticatedClient trys to aunthenticate with yahoo and tries to establish a private YQL connection.
func NewAuthenticatedClient(key string, secret string) (*Client, error) {

	auth := key + "|" + secret
	db, err := sql.Open("yql", auth)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to YQL: %s", err)
	}

	return &Client{
		db,
	}, err
}

// NewClient trys to establish a public YQL connection.
func NewClient() (*Client, error) {

	db, err := sql.Open("yql", "||store://datatables.org/alltableswithkeys")
	if err != nil {
		return nil, fmt.Errorf("Error connecting to YQL: %s", err)
	}

	return &Client{
		db,
	}, err
}
