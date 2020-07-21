package server

import (
	"context"
	"log"
	"time"

	"github.com/chutified/currencies/data"
	"github.com/chutified/currencies/protos/currency"
)

// Currency is a server which serves the gRPC currency calls.
type Currency struct {
	log *log.Logger
	ds  *data.Service
}

// New is a constructor of the currency server.
func New(log *log.Logger, ds *data.Service) *Currency {
	c := &Currency{
		log: log,
		ds:  ds,
	}

	// monitoring
	go func() {
		c.handleUpdates()
	}()

	return c
}

// GetCurrency handles GetCurrency gRPC request calls.
func (c *Currency) GetCurrency(ctx context.Context, req *currency.GetCurrencyRequest) (*currency.GetCurrencyResponse, error) {

	// handle
	resp, err := c.handleGetCurrencyRequest(req)
	if err != nil {
		return nil, err
	}
	// success
	return resp, nil
}

// GetRate handles GetRate gRPC request calls.
func (c *Currency) GetRate(ctx context.Context, req *currency.GetRateRequest) (*currency.GetRateResponse, error) {

	// handle
	resp, err := c.handleGetRateRequest(req)
	if err != nil {
		return nil, err
	}
	// success
	return resp, nil
}

func (c *Currency) handleUpdates() {

	updates, errs := c.ds.MonitorData(15 * time.Second)

	// log errors
	go func() {
		for err := range errs {
			c.log.Printf("[error] monitor data: %s", err)
		}
	}()

	// update
	for range updates {

		// TODO
		// data got updated
	}
}
