package market

import (
	"github.com/adshao/go-binance/v2"
	"math"
	"math/rand"
	"time"
)

type IClient interface {
	DataFetcher
	OrderMaker
}

type DataFetcher interface {
	Fetch(pair, timeframe string, limit int) []Candle
	FetchPeriod(pair, timeframe string, start int64, end int64) []Candle
}

type OrderMaker interface {
	ListOrders(pair string) []Order
	OpenMarketOrder(order Order) (Order, error)
}

type Candle struct {
	OpenTime  int64   `json:"openTime"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
	CloseTime int64   `json:"closeTime"`
}

type Order struct {
	OrderID     int64
	Price       float64
	Pair        string
	IsOpen      bool
	IsCanceled  bool
	IsCompleted bool
}

func NewClient(apiKey, secretKey string) IClient {
	return &BinanceClient{
		client: binance.NewClient(apiKey, secretKey),
	}
}

func RandomizeCandle(candle Candle) Candle {
	h := candle.High - candle.Low
	r := rand.Float64()

	ns := time.Now().Nanosecond()
	if ns <= 1 {
		ns = 3
	}
	candle.Close = candle.Low + h*r
	candle.Volume += candle.Volume * r / math.Log10(float64(ns))

	return candle
}
