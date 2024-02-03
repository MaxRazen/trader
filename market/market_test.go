package market

import (
	"fmt"
	"testing"
)

func TestRandomizeCandle(t *testing.T) {
	testCases := []Candle{
		{
			OpenTime:  0,
			CloseTime: 0,
			Open:      1324.0,
			Close:     1330.0,
			High:      1342.0,
			Low:       1321.0,
			Volume:    54024.24,
		},
		{
			OpenTime:  1668755700000,
			CloseTime: 1668756599999,
			Open:      2.26410000,
			Close:     2.27200000,
			High:      2.27430000,
			Low:       2.26240000,
			Volume:    19515.00000000,
		},
		{
			OpenTime:  1668756600000,
			CloseTime: 1668757499999,
			Open:      0.08552000,
			Close:     0.08537000,
			High:      0.08559000,
			Low:       0.08516000,
			Volume:    3373325.00000000,
		},
	}

	for i, testCase := range testCases {
		res := RandomizeCandle(testCase)

		if res.Volume < testCase.Volume {
			t.Errorf("[case #%v] volume can not be less %v expected %v", i, res.Volume, testCase.Volume)
		}
		if res.Low < testCase.Low {
			t.Errorf("[case #%v] low price %v can not be less than %v", i, res.Low, testCase.Low)
		}
		if res.High > testCase.High {
			t.Errorf("[case #%v] high price %v can not be great than %v", i, res.High, testCase.High)
		}
		if res.Open != testCase.Open {
			t.Errorf("[case #%v] open price %v must not be changed, expected %v", i, res.Open, testCase.Open)
		}
		if res.Close == testCase.Close {
			t.Logf("[case #%v] warning: close price has not been changed, input %v, result %v", i, testCase.Close, res.Close)
		}
	}
}

func TestGetMockCandleData(t *testing.T) {
	data := GetMockCandleData(1)
	errMsg := "mocked-client: received data is not matched with expectations (%v)"
	if len(data) != 1 || data[0].OpenTime != 1667952000000 {
		t.Fatalf(fmt.Sprintf(errMsg, 1))
	}
	data = GetMockCandleData(2)
	if len(data) != 2 || data[0].OpenTime != 1668038400000 || data[1].OpenTime != 1667952000000 {
		t.Fatalf(fmt.Sprintf(errMsg, 2))
	}
	data = GetMockCandleData(100)
	if len(data) != 10 || data[0].OpenTime != 1668729600000 || data[9].OpenTime != 1667952000000 {
		t.Fatalf(fmt.Sprintf(errMsg, 10))
	}
}
