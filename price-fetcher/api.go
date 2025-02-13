package main

import (
	"context"
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJsonAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.IntN(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(context.Background(), w, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}
func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return writeJSON(w, http.StatusOK, &priceResponse)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
