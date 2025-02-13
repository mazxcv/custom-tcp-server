package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that fetches the price of a ticker
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher is a struct that implements the PriceFetcher interface
type priceFetcher struct {
}

func (p *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {

	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"AAPL": 100,
	"MSFT": 200,
	"GOOG": 300,
	"AMZN": 400,
	"FB":   500,
	"BTC":  6_000_000,
	"ETH":  4_020_000,
	"XRP":  2000,
	"DOGE": 1_000,
	"DOT":  200,
	"UNI":  1,
	"LINK": 400,
	"DAI":  3_000,
	"USDT": 2_000,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return 0, fmt.Errorf("ticker %s not found", ticker)
	}
	return price, nil
}
