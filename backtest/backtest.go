package main

import (
	"hexnet/trader/core"
	"hexnet/trader/market"
	"hexnet/trader/strategies"
	"log"
	"os"
)

var (
	config       *core.Config
	marketClient market.IClient
)

func initApplication() {
	conf := core.LoadConfig("../.env")
	config = &conf
	marketClient = market.NewClient(conf.Env.BinanceApiKey, conf.Env.BinanceSecretKey)
}

func loadStrategy() core.IStrategy {
	registeredStrategies := []core.IStrategy{
		&strategies.Scalpel{},
	}
	for _, s := range registeredStrategies {
		if s.Name() == config.Env.StrategyName {
			return s
		}
	}
	log.Fatalf("strategy is not found: %s", config.Env.StrategyName)
	return nil
}

func main() {
	println("backtest")

	initApplication()
	cliArgs := core.ResolveCliArguments(os.Args[1:])

	strategy := loadStrategy()

	dataFetcher := core.NewHistoricalDataFetcher(
		marketClient,
		cliArgs.TimeRange,
		cliArgs.NoCache,
		config.Env.HistoricalDataPath,
	)

	dataframe := core.NewDataframeBacktesting(
		cliArgs.Pair,
		cliArgs.Timeframe,
		strategy.RequiredCandlesNumber(),
		dataFetcher,
	)

	orderManager := core.NewOrderManager(marketClient)
	tickerContext := core.NewTickerContext(config, dataframe, orderManager, strategy)
	ticker := core.NewTicker(tickerContext)
	ticker.SetTimeRange(cliArgs.TimeRange)
	ticker.SetIsBacktesting(true)
	ticker.Run()

	/*
		timeCursor = timerange.Start
		for timeCursor < timerange.End {
			dataframe.SetCurrentTime(timeCursor)
			orderManager.HandleStrategyDecision(strategy.HandleTick(context))

			timeCursor += config.Env.TickerTimeout // 60sec
		}
	*/

	/**
	- order manager
	- ticker
	- dump orders
	- graceful shutdown
	*/

	//fmt.Printf("%+v\n", dataframe.GetCurrentCandle())
	/*

		data := marketClient.Fetch("ETHUSDT", "15m", 2)
		strategy := loadStrategy()

		if strategy == nil {
			log.Fatalf("strategy must be defined: %v", config.Env.StrategyName)
		}
		if strategy.ShouldBuy(data) {
			println("should buy")
		}
		if strategy.ShouldSell(data) {
			println("should sell")
		}
		//fmt.Printf("len: %v\n: %+v\n", len(data), data)

		fmt.Printf("%+v\n", core.ResolveCliArguments(os.Args))
	*/
	println("backtest done")
}
