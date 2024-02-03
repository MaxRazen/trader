package core

type IStrategy interface {
	Name() string
	RequiredCandlesNumber() int
	HandleTick(ctx *TickerContext) StrategyDecision
}

const (
	StrategyDecisionBuy  = "buy"
	StrategyDecisionSell = "sell"
)

const (
	StrategyDecisionBehaviorMarket = "market"
	StrategyDecisionBehaviorLimit  = "limit"
)

type CalculatedIndicators map[string]string

type StrategyDecision struct {
	indicators CalculatedIndicators
	action     string
	behavior   string
	price      float64
}

func (sd *StrategyDecision) ShouldBuy() bool {
	return sd.action == StrategyDecisionBuy
}

func (sd *StrategyDecision) ShouldSell() bool {
	return sd.action == StrategyDecisionSell
}

func (sd *StrategyDecision) Behavior() string {
	return sd.behavior
}

func (sd *StrategyDecision) GetIndicators() CalculatedIndicators {
	return sd.indicators
}

func NewStrategyDecision(action, behavior string, price float64, indicators CalculatedIndicators) StrategyDecision {
	return StrategyDecision{
		indicators: indicators,
		action:     action,
		behavior:   behavior,
		price:      price,
	}
}
