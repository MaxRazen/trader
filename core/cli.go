package core

import (
	"log"
	"strings"
)

type CliArguments struct {
	Timeframe string
	Pair      string
	TimeRange TimeRange
	DryMode   bool
	NoCache   bool
}

func ResolveCliArguments(args []string) CliArguments {
	// --timeframe=1h
	// --pair=BTCUSDT
	// --timerange=2020/02/04-2022/02/02
	// --no-cache
	// --dry-mode

	cliArgs := CliArguments{
		Timeframe: "1h",
		Pair:      "BTCUSDT",
		TimeRange: GetLast30Days(),
		DryMode:   false,
		NoCache:   false,
	}

	for _, arg := range args {
		segments := strings.Split(arg, "=")
		argName := strings.TrimPrefix(segments[0], "--")
		argValue := ""
		if len(segments) > 1 {
			argValue = segments[1]
		}
		switch argName {
		case "pair":
			pairCode := strings.ReplaceAll(argValue, "-", "")
			cliArgs.Pair = strings.ToUpper(pairCode)
		case "timeframe":
			cliArgs.Timeframe = argValue
		case "timerange":
			dt := strings.Split(argValue, "-")
			if len(dt) != 2 {
				log.Fatalf("unsupported date format (use yyyy/mm/dd-yyyy/mm/dd)")
			}
			start, _ := ParseDate2Time(dt[0])
			end, _ := ParseDate2Time(dt[1])

			if start.Unix() > end.Unix() {
				start, end = end, start
			}
			cliArgs.TimeRange = TimeRange{
				Start: start,
				End:   end,
			}
		case "dry-mode":
			cliArgs.DryMode = true
		case "no-cache":
			cliArgs.NoCache = true
		}
	}

	return cliArgs
}
