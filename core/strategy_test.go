package core

import "testing"

type DummyStrategy struct {
}

func (s *DummyStrategy) Name() string {
	return "dummy"
}

func (s *DummyStrategy) RequiredCandlesNumber() int {
	return 5
}

func (s *DummyStrategy) HandleTick(_ *TickerContext) StrategyDecision {
	return StrategyDecision{}
}

func TestNewStrategyDecision(t *testing.T) {
	indicators := CalculatedIndicators{"adx": "12.34"}
	d := NewStrategyDecision(
		StrategyDecisionBuy,
		StrategyDecisionBehaviorMarket,
		101.22,
		indicators,
	)

	if d.Behavior() != StrategyDecisionBehaviorMarket {
		t.Errorf("Behavior does not match with " + StrategyDecisionBuy)
	}

	if d.ShouldBuy() != true {
		t.Errorf("Action must be buy")
	}

	if d.ShouldSell() != false {
		t.Errorf("Action could not be sell")
	}

	if len(d.GetIndicators()) != 1 {
		t.Errorf("Indicators are not set")
	}
}
