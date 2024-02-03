package strategies

import (
	"hexnet/trader/core"
	"hexnet/trader/market"
)

type Scalpel struct {
}

func (s *Scalpel) Name() string {
	return "scalpel"
}

func (s *Scalpel) RequiredCandlesNumber() int {
	return 5
}

func (s *Scalpel) HandleTick(ctx *core.TickerContext) core.StrategyDecision {
	indicators := s.populateIndicators(ctx.GetDataframe().GetDataset(50))
	var action, behavior string

	return core.NewStrategyDecision(
		action,
		behavior,
		0,
		indicators,
	)
}

func (s *Scalpel) populateIndicators(_ []market.Candle) core.CalculatedIndicators {
	indicatorsMap := make(core.CalculatedIndicators)
	// todo
	return indicatorsMap
}

func (s *Scalpel) shouldBuy(_ []market.Candle) bool {
	return false
}

func (s *Scalpel) shouldSell(_ []market.Candle) bool {
	return false
}
