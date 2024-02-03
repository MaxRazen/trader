package core

import (
	"encoding/json"
	"fmt"
	"hexnet/trader/market"
	"log"
	"os"
)

type DataFetcher interface {
	Fetch(pair, timeframe string, limit int) []market.Candle
}

type HistoricalDataFetcher struct {
	client    market.IClient
	timeRange TimeRange
	cachePath string
	nocache   bool
}

func (f *HistoricalDataFetcher) Fetch(pair, timeframe string, _ int) []market.Candle {
	filename := f.getFilename(pair, timeframe)
	var candles []market.Candle

	if !f.nocache {
		candles = f.loadFromFile(filename)
	}

	if len(candles) > 0 {
		return candles
	}

	duration := GetTimeframeDuration(timeframe)
	periods := SplitDurationOnPeriods(f.timeRange, duration, 1000)
	for _, period := range periods {
		cs := f.client.FetchPeriod(pair, timeframe, period.Start.UnixMilli(), period.End.UnixMilli())
		for _, c := range cs {
			candles = append(candles, c)
		}
	}
	if len(candles) <= 0 {
		log.Fatalf("DataFetcher: data has not been fetched")
	}

	err := f.saveToFile(filename, candles)
	if err != nil {
		log.Fatalf("DataFetcher: can't save historical data to file: " + err.Error())
	}

	return candles
}

func (f *HistoricalDataFetcher) getFilename(pair, timeframe string) string {
	return fmt.Sprintf(
		"%s-%s-%s-%s.json",
		pair,
		timeframe,
		f.timeRange.Start.Format(DateFormatLayout),
		f.timeRange.End.Format(DateFormatLayout),
	)
}

func (f *HistoricalDataFetcher) loadFromFile(filename string) []market.Candle {
	data, err := os.ReadFile(f.resolveCachePath(filename))
	var candles []market.Candle
	if err != nil {
		return candles
	}

	err = json.Unmarshal(data, &candles)
	if err != nil {
		return candles
	}
	return candles
}

func (f *HistoricalDataFetcher) saveToFile(filename string, candles []market.Candle) error {
	data, err := json.Marshal(candles)
	if err != nil {
		return err
	}
	return os.WriteFile(f.resolveCachePath(filename), data, 0664)
}

func (f *HistoricalDataFetcher) resolveCachePath(filename string) string {
	return fmt.Sprintf("%s/%s", f.cachePath, filename)
}

type LiveDataFetcher struct {
	client market.IClient
}

func (f *LiveDataFetcher) Fetch(pair, timeframe string, limit int) []market.Candle {
	return f.client.Fetch(pair, timeframe, limit)
}

func NewHistoricalDataFetcher(
	client market.IClient,
	tr TimeRange,
	nocache bool,
	cachePath string,
) *HistoricalDataFetcher {
	if cachePath == "" { // todo move to the config
		cachePath = "../.data/historical"
	}
	return &HistoricalDataFetcher{
		client:    client,
		timeRange: tr,
		cachePath: cachePath,
		nocache:   nocache,
	}
}

func NewLiveDataFetcher(client market.IClient) *LiveDataFetcher {
	return &LiveDataFetcher{
		client: client,
	}
}
