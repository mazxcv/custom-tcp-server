package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (l *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
			"ticker":    ticker,
			"requestID": ctx.Value("requestID"),
		}).Infof("Fetching...")
	}(time.Now())

	return l.next.FetchPrice(ctx, ticker)
}
