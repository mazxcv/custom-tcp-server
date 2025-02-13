package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/mazxcv/custom-tcp-server/price-fetcher/client"
)

func main() {

	client := client.New("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "USD")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("price: %f\n", price)
	return
	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()
	svc := NewLoggingService(&priceFetcher{})

	server := NewJsonAPIServer(*listenAddr, svc)

	server.Run()

}
