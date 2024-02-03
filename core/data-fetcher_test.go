package core

import (
	"hexnet/trader/market"
	"os"
	"testing"
)

func TestNewHistoricalDataFetcher(t *testing.T) {
	tr := GetToday()
	c := market.NewMockedClient(market.GetMockCandleData(1))
	fetcher := NewHistoricalDataFetcher(c, tr, true, "/tmp")

	if fetcher == nil {
		t.Errorf("NewHistoricalDataFetcher could not be initialized")
	}
}

func TestHistoricalDataFetcher_Fetch(t *testing.T) {
	pair := "PAIR"
	timeframe := "1h"
	cachePath := "/tmp"
	tr := GetToday()
	data := market.GetMockCandleData(1)
	c := market.NewMockedClient(data)

	fetcher := NewHistoricalDataFetcher(c, tr, true, cachePath)
	candles := fetcher.Fetch(pair, timeframe, 0)

	if len(candles) != len(data) || candles[0] != data[0] {
		t.Errorf("data-fetcher: candles data is now equal")
	}

	filename := fetcher.getFilename(pair, timeframe)
	f, err := os.OpenFile(fetcher.resolveCachePath(filename), os.O_RDONLY, 644)
	if err != nil {
		t.Errorf("data-fetcher: cache file could be opened")
	}
	info, err := f.Stat()
	if err != nil {
		t.Errorf("data-fetcher: cache file data is unavailable")
	}
	if info.Size() < 1 {
		t.Errorf("data-fetcher: cache file is empty")
	}
	err = os.Remove(fetcher.resolveCachePath(filename))
	if err != nil {
		t.Errorf("data-fetcher: cache file could not be deleted")
	}
}
