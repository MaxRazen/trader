package core

import (
	"testing"
	"time"
)

type CliTestCase struct {
	Args        []string
	Expectation CliArguments
}

var CliTestCases = []CliTestCase{
	// default values must be applied when test
	{
		Args: []string{},
		Expectation: CliArguments{
			Timeframe: "1h",
			Pair:      "BTCUSDT",
			TimeRange: GetLast30Days(),
			DryMode:   false,
		},
	},
	// all the arguments
	{
		Args: []string{"--timeframe=5m", "--dry-mode", "--timerange=2021/01/01-2021/12/15", "--pair=ethusdt"},
		Expectation: CliArguments{
			Timeframe: "5m",
			Pair:      "ETHUSDT",
			TimeRange: TimeRange{
				Start: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2021, time.December, 15, 0, 0, 0, 0, time.UTC),
			},
			DryMode: true,
		},
	},
	// only some arguments
	{
		Args: []string{"--pair=eth-usdt", "--timeframe=1d"},
		Expectation: CliArguments{
			Timeframe: "1d",
			Pair:      "ETHUSDT",
			TimeRange: GetLast30Days(),
			DryMode:   false,
		},
	},
}

func TestResolveCliArguments(t *testing.T) {
	for i, testCase := range CliTestCases {
		cliArgs := ResolveCliArguments(testCase.Args)
		if cliArgs != testCase.Expectation {
			t.Errorf("result does not match expectation in case #%v\n", i)
			t.Errorf("[case: #%v]\n\t%+v\n\t%+v\n", i, cliArgs, testCase.Expectation)
		}
	}
}
