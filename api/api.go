package api

import (
	"encoding"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, usually 200
	Code int

	// Account Balance
	Balance int64
}

type Error struct {
	// Error Code
	Code int

	// Error Message
	Message string
}
