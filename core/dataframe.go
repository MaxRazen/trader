package core

import (
	"hexnet/trader/market"
	"log"
	"time"
)

type IDataframe interface {
	GetCurrent() market.Candle
	GetDataset(limit int) market.Candle
	GetClosePrices(limit int) []float64
	GetVolumes(limit int) []float64
	loadDataset()
	getSafeLimit(int) int
}

type Dataframe struct {
	dataFetcher DataFetcher
	candles     []market.Candle
	pair        string
	timeframe   string
	minCandles  int
	maxCandles  int
}

func NewDataframe(pair, timeframe string, minCandles int, fetcher DataFetcher) *Dataframe {
	return &Dataframe{
		dataFetcher: fetcher,
		pair:        pair,
		timeframe:   timeframe,
		minCandles:  minCandles,
		maxCandles:  50,
	}
}

func (df *Dataframe) GetCurrent() market.Candle {
	_ = df.getSafeLimit(1)
	return df.candles[0]
}

func (df *Dataframe) GetDataset(limit int) []market.Candle {
	limit = df.getSafeLimit(limit)
	return df.candles[0:limit]
}

func (df *Dataframe) GetClosePrices(limit int) []float64 {
	var prices []float64
	limit = df.getSafeLimit(limit)
	for i := 0; i < limit; i++ {
		prices = append(prices, df.candles[i].Close)
	}
	return prices
}

func (df *Dataframe) GetVolumes(limit int) []float64 {
	var volumes []float64
	limit = df.getSafeLimit(limit)
	for i := 0; i < limit; i++ {
		volumes = append(volumes, df.candles[i].Volume)
	}
	return volumes
}

func (df *Dataframe) loadDataset() {
	df.candles = df.dataFetcher.Fetch(df.pair, df.timeframe, df.minCandles)
}

func (df *Dataframe) getSafeLimit(limit int) int {
	if len(df.candles) == 0 {
		log.Fatalf("dataframe: candles slice is empty")
		return 0
	}
	if limit <= 0 || len(df.candles) < limit {
		limit = len(df.candles)
	}
	return limit
}

type DataframeBacktesting struct {
	Dataframe
	timeCursor time.Time
}

func (df *DataframeBacktesting) GetDataset(limit int) []market.Candle {
	limit = df.getSafeLimit(limit)
	var candles []market.Candle
	candles = append(candles, df.candles[0])
	candles = append(candles, df.candles[1:limit]...)
	return candles
}

func (df *DataframeBacktesting) loadDataset() {
	allCandles := df.dataFetcher.Fetch(df.pair, df.timeframe, df.minCandles)
	ts := df.timeCursor.UnixMilli()
	var candles []market.Candle
	for _, c := range allCandles {
		if c.CloseTime <= ts {
			candles = append(candles, c)
		}
	}
	if len(candles) > 0 {
		candles[0] = market.RandomizeCandle(candles[0])
	}
	df.candles = candles
}

func (df *DataframeBacktesting) setTimeCursor(tc time.Time) {
	df.timeCursor = tc
}

func NewDataframeBacktesting(pair, timeframe string, minCandles int, fetcher DataFetcher) *DataframeBacktesting {
	return &DataframeBacktesting{
		Dataframe: Dataframe{
			dataFetcher: fetcher,
			pair:        pair,
			timeframe:   timeframe,
			minCandles:  minCandles,
			maxCandles:  50,
		},
		timeCursor: GetNow(),
	}
}
