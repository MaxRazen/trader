package core

import (
	"hexnet/trader/market"
	"testing"
	"time"
)

func TestNewDataframeBacktesting(t *testing.T) {
	tr := GetToday()
	data := market.GetMockCandleData(1)
	c := market.NewMockedClient(data)
	fetcher := NewHistoricalDataFetcher(c, tr, false, "/tmp")
	dataframe := NewDataframeBacktesting("PAIR", "5m", 2, fetcher)
	if dataframe == nil {
		t.Errorf("dataframe: instance could not be initialized")
	}
	if dataframe.pair != "PAIR" {
		t.Errorf("dataframe: pair not is not match")
	}
}

func TestBacktesting_loadDataset(t *testing.T) {
	dataframe := initDataframeBacktesting(5, 2)
	if len(dataframe.candles) > 0 {
		t.Errorf("dataframe: candles have already had values")
	}
	dataframe.loadDataset()
	if len(dataframe.candles) < 0 {
		t.Errorf("dataframe: data could not been received from fetcher")
	}

	data := market.GetMockCandleData(2)
	timeCursor := time.UnixMilli(data[0].CloseTime)
	dataframe.setTimeCursor(timeCursor)
	dataframe.loadDataset()
	if len(dataframe.candles) != 2 {
		t.Errorf("dataframe: candles count %v is not matched with expected count 2", len(dataframe.candles))
	}
	timeCursor = time.UnixMilli(timeCursor.UnixMilli() + 86400*1000)
	dataframe.setTimeCursor(timeCursor)
	dataframe.loadDataset()
	if len(dataframe.candles) != 3 {
		t.Errorf("dataframe: candles count %v is not matched with expected count 3", len(dataframe.candles))
	}
}

func TestDataframeBacktesting_getSafeLimit(t *testing.T) {
	maxLimit := 5
	dataframe := initDataframeBacktesting(maxLimit, 2)
	dataframe.loadDataset()
	l := dataframe.getSafeLimit(0)
	if l != maxLimit {
		t.Errorf("dataframe: limit must have max available value for 0")
	}
	l = dataframe.getSafeLimit(10)
	if l != maxLimit {
		t.Errorf("dataframe: limit must have max available value for 10")
	}
	l = dataframe.getSafeLimit(3)
	if l != 3 {
		t.Errorf("dataframe: limit must equal to 3")
	}
}

func TestBacktesting_GetCurrent(t *testing.T) {
	data := market.GetMockCandleData(2)
	dataframe := initDataframeBacktesting(5, 2)
	timeCursor := time.UnixMilli(data[1].CloseTime + 30*1000)

	dataframe.setTimeCursor(timeCursor)
	dataframe.loadDataset()
	currentCandle := dataframe.GetCurrent()

	currentCandle2 := dataframe.GetCurrent()
	if currentCandle.CloseTime != currentCandle2.CloseTime {
		t.Errorf("dataframe: current candle time is changed")
	}
	if currentCandle.Close != currentCandle2.Close {
		t.Errorf("dataframe: current candle close price should not changed")
	}
	if currentCandle.Volume != currentCandle2.Volume {
		t.Errorf("dataframe: current candle volume should not changed")
	}
}

func TestDataframeBacktesting_GetDataset(t *testing.T) {
	dataframe := initDataframeBacktesting(10, 2)

	for i, c := range [...]int{2, 3, 5, 7, 10} {
		expected := market.GetMockCandleData(c)
		length := len(expected)
		timeCursor := time.UnixMilli(expected[0].CloseTime + 86399*1000)
		dataframe.setTimeCursor(timeCursor)
		dataframe.loadDataset()
		candles := dataframe.GetDataset(c)

		if len(candles) != length {
			t.Fatalf("datafarame: [case #%v] expected length is not match", i)
		}
		if candles[0].CloseTime != expected[0].CloseTime || candles[length-1].OpenTime != expected[length-1].OpenTime {
			t.Errorf("dataframe: [case #%v] timerange is wrong", i)
		}
	}
}

func initDataframeBacktesting(limit, minCandles int) *DataframeBacktesting {
	tr := GetToday()
	data := market.GetMockCandleData(limit)
	client := market.NewMockedClient(data)
	fetcher := NewHistoricalDataFetcher(client, tr, true, "/tmp")
	return NewDataframeBacktesting("PAIR", "5m", minCandles, fetcher)
}
