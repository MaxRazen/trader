package core

import (
	"errors"
	"time"
)

type ITicker interface {
	Run()
	SetTimeCursor(time.Time)
	SetTimeRange(TimeRange)
	SetIsBacktesting(bool)
}

type ITickerContext interface {
}

type TickerContext struct {
	config       *Config
	dataframe    *DataframeBacktesting
	orderManager *OrderManager
	strategy     IStrategy
	// config, dataframe, strategy, order manager
}

type TickerMetadata struct {
	Iterations int
}

func (tc *TickerContext) GetDataframe() *DataframeBacktesting {
	return tc.dataframe
}

func NewTickerContext(
	config *Config,
	dataframe *DataframeBacktesting,
	orderManager *OrderManager,
	strategy IStrategy,
) *TickerContext {
	return &TickerContext{
		config:       config,
		dataframe:    dataframe,
		orderManager: orderManager,
		strategy:     strategy,
	}
}

func NewTicker(ctx *TickerContext) ITicker {
	return &Ticker{
		context:    ctx,
		timeRange:  nil,
		timeCursor: GetNow(),
		duration:   GetTimeframeDuration(ctx.dataframe.timeframe),
		metadata:   TickerMetadata{},
	}
}

type Ticker struct {
	context       *TickerContext
	timeRange     *TimeRange
	timeCursor    time.Time
	duration      int
	isBacktesting bool
	metadata      TickerMetadata
}

func (t *Ticker) SetTimeCursor(tc time.Time) {
	t.timeCursor = tc
}

func (t *Ticker) SetTimeRange(tr TimeRange) {
	t.timeRange = &tr
}

func (t *Ticker) SetIsBacktesting(enabled bool) {
	t.isBacktesting = enabled
}

func (t *Ticker) getMetadata() TickerMetadata {
	return t.metadata
}

func (t *Ticker) shiftTimeCursor() {
	t.timeCursor = t.timeCursor.Add(t.getTickerTimeout())
}

func (t *Ticker) getTickerTimeout() time.Duration {
	return time.Second * time.Duration(t.context.config.Env.TickerTimeout)
}

func (t *Ticker) hasEnoughData() bool {
	return len(t.context.dataframe.candles) >= t.context.strategy.RequiredCandlesNumber()
}

func (t *Ticker) await() error {
	if !t.isBacktesting {
		time.Sleep(t.getTickerTimeout())
		return nil
	}
	t.shiftTimeCursor()
	if t.timeCursor.Unix() >= t.timeRange.End.Unix() {
		return errors.New("ticker ")
	}
	return nil
}

func (t *Ticker) Run() {
	for {
		t.context.dataframe.loadDataset()

		if len(t.context.dataframe.candles) < 1 {
			break
		}

		if t.hasEnoughData() {
			decision := t.context.strategy.HandleTick(t.context)
			t.context.orderManager.HandleDecision(decision)
		}

		t.metadata.Iterations++

		if err := t.await(); err != nil {
			return
		}
	}

	// check ticker condition
	// 		- not enough data
	// 		- end period reached
	// get [n]-candles, n - should be defined in the strategy, or take 5 by default
	// feed [n]-candles / dataframe to Strategy
}
