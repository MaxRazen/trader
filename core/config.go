package core

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	EnvDbUrl              = "DB_URL"
	EnvBacktestDbUrl      = "BACKTEST_DB_URL"
	EnvStrategyName       = "STRATEGY_NAME"
	EnvBinanceApiKey      = "BINANCE_API_KEY"
	EnvBinanceSecretKey   = "BINANCE_SECRET_KEY"
	EnvTickerTimeout      = "TICKER_TIMEOUT"
	EnvHistoricalDataPath = "HISTORICAL_DATA_PATH"
)

var config Config

type AppEnv struct {
	DbUrl              string
	BacktestDbUrl      string
	StrategyName       string
	BinanceApiKey      string
	BinanceSecretKey   string
	HistoricalDataPath string
	TickerTimeout      int
}

type Config struct {
	Env AppEnv
}

func LoadConfig(envPath string) Config {
	if envPath == "" {
		envPath = ".env"
	}

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("environment configuration file not found in %s", envPath)
		return Config{}
	}

	config = Config{
		Env: AppEnv{
			StrategyName:       os.Getenv(EnvStrategyName),
			DbUrl:              resolveStr(os.Getenv(EnvDbUrl), ".data/db.sql"),
			TickerTimeout:      resolveInt(os.Getenv(EnvTickerTimeout), 120),
			BacktestDbUrl:      resolveStr(os.Getenv(EnvBacktestDbUrl), ".data/db-backtest.sql"),
			BinanceApiKey:      os.Getenv(EnvBinanceApiKey),
			BinanceSecretKey:   os.Getenv(EnvBinanceSecretKey),
			HistoricalDataPath: resolveStr(os.Getenv(EnvHistoricalDataPath), "../.data/historical"),
		},
	}

	return config
}

func GetConfig() Config {
	return config
}

func resolveStr(str, def string) string {
	if str == "" {
		return def
	}
	return str
}

func resolveInt(str string, def int) int {
	if str == "" {
		return def
	}
	v, e := strconv.Atoi(str)
	if e != nil {
		return def
	}
	return v
}
