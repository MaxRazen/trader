package core

import (
	"testing"
	"time"
)

func TestParseDate2Time(t *testing.T) {
	expectedTime := time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)

	args := []struct {
		input string
		time  time.Time
		isErr bool
	}{
		{
			input: "2006-01-02",
			time:  expectedTime,
			isErr: false,
		},
		{
			input: "20060102",
			time:  expectedTime,
			isErr: false,
		},
		{
			input: "2006/01/02",
			time:  expectedTime,
			isErr: false,
		},
		{
			input: "2s006-01-02 15:04:05",
			time:  time.Time{},
			isErr: true,
		},
	}

	for i, v := range args {
		ts, err := ParseDate2Time(v.input)
		if v.isErr && err == nil {
			t.Errorf("date parsing error must be returned")
			return
		}
		if ts != v.time {
			t.Errorf("date converting error %v != %v [case %v]", ts, v.time, i)
		}
	}
}

func TestGetTimeframeDuration(t *testing.T) {
	testCases := []struct {
		timeframe string
		expected  int
	}{
		// success cases
		{
			timeframe: "1m",
			expected:  60,
		},
		{
			timeframe: "2m",
			expected:  120,
		},
		{
			timeframe: "5m",
			expected:  300,
		},
		{
			timeframe: "12m",
			expected:  720,
		},
		{
			timeframe: "30m",
			expected:  1800,
		},
		{
			timeframe: "1h",
			expected:  3600,
		},
		{
			timeframe: "2h",
			expected:  7200,
		},
		{
			timeframe: "4h",
			expected:  14400,
		},
		{
			timeframe: "1d",
			expected:  86400,
		},
		{
			timeframe: "2d",
			expected:  86400 * 2,
		},
		// failed cases
		{
			timeframe: "",
			expected:  0,
		},
	}
	for i, testCase := range testCases {
		res := GetTimeframeDuration(testCase.timeframe)

		if res != testCase.expected {
			t.Errorf("duration is not match %v != %v [case %v]", res, testCase.expected, i)
		}
	}
}

func TestSplitDurationOnPeriods(t *testing.T) {
	testCases := []struct {
		timeRange      TimeRange
		duration       int
		limit          int
		expectedLength int
	}{
		{
			timeRange: TimeRange{
				Start: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			duration:       5 * 60,
			limit:          1000,
			expectedLength: 1,
		},
		{
			timeRange: TimeRange{
				Start: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 1, 9, 23, 59, 59, 0, time.UTC),
			},
			duration:       5 * 60,
			limit:          1000,
			expectedLength: 3,
		},
		{
			timeRange: TimeRange{
				Start: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2020, 5, 8, 15, 22, 12, 0, time.UTC),
			},
			duration:       60 * 60,
			limit:          1000,
			expectedLength: 4,
		},
	}

	for i, testCase := range testCases {
		res := SplitDurationOnPeriods(testCase.timeRange, testCase.duration, testCase.limit)

		if len(res) != testCase.expectedLength {
			t.Errorf("\t[case #%v] length is not match %v != %v", i, len(res), testCase.expectedLength)
		}
		if res[0].Start != testCase.timeRange.Start {
			t.Errorf("\t[case #%v] start period is not match with expected", i)
		}
		if res[len(res)-1].End != testCase.timeRange.End {
			t.Errorf("\t[case #%v] end period is not match with expected", i)
		}
	}
}
