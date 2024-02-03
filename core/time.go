package core

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const DateFormatLayout = "20060102"

type TimeRange struct {
	Start time.Time
	End   time.Time
}

// ParseDate2Time casts string date (YYYY-MM-DD, YYYY/MM/DD, YYYYMMDD) to time.Time
func ParseDate2Time(dt string) (time.Time, error) {
	layout := "2006-01-02"
	if strings.Contains(dt, "/") {
		layout = "2006/01/02"
	} else if !strings.Contains(dt, "-") {
		layout = DateFormatLayout
	}
	return time.Parse(layout, dt)
}

// GetTimeframeDuration returns count of seconds from the given string period
// (5m => 300)
// (1h => 3600)
func GetTimeframeDuration(timeframe string) int {
	re := regexp.MustCompile(`(\d+)([mhd])`)
	s := re.FindAllStringSubmatch(timeframe, -1)
	if len(s) <= 0 || len(s[0]) <= 2 {
		return 0
	}
	segments := s[0][1:]
	value, _ := strconv.Atoi(segments[0])
	switch segments[1] {
	case "m":
		return int(time.Minute * time.Duration(value) / time.Second)
	case "h":
		return int(time.Hour * time.Duration(value) / time.Second)
	case "d":
		return int(time.Hour * 24 * time.Duration(value) / time.Second)
	}
	return 0
}

// SplitDurationOnPeriods returns a slice of time ranges split by duration*limit seconds
func SplitDurationOnPeriods(tr TimeRange, duration, limit int) []TimeRange {
	start := tr.Start.Unix()
	end := tr.End.Unix()
	durations := int64(duration * limit)
	var periods []TimeRange

	periodsNum := int(math.Ceil(float64(end-start) / float64(durations)))
	if periodsNum <= 1 {
		periods = append(periods, tr)
		return periods
	}

	i := 0
	cursor := start
	for i < periodsNum {
		periodEnd := time.Unix(cursor+durations-1, 0).In(time.UTC)
		if periodEnd.Unix() > end {
			periodEnd = tr.End
		}
		periods = append(periods, TimeRange{
			Start: time.Unix(cursor, 0).In(time.UTC),
			End:   periodEnd,
		})
		cursor += durations
		i++
	}
	return periods
}

// GetNow returns current time with the UTC timezone
func GetNow() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second(), 0, time.UTC)
}

func GetToday() TimeRange {
	return TimeRange{
		Start: todayDate(),
		End:   GetNow(),
	}
}

func GetLast30Days() TimeRange {
	return TimeRange{
		Start: todayDate().Add(-time.Hour * 24 * 30),
		End:   GetNow(),
	}
}

func GetLast365Days() TimeRange {
	return TimeRange{
		Start: todayDate().Add(-time.Hour * 24 * 365),
		End:   GetNow(),
	}
}

func todayDate() time.Time {
	now := GetNow()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
