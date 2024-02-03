package core

import (
	"hexnet/trader/market"
)

const (
	OrderFulfillmentStrategyMarket = "market"
)

type OrderManager struct {
	client market.IClient
}

func (om *OrderManager) List() {

}

//func (om *OrderManager) CreateOrder() (market.Order, error) {
//}
//
//func (om *OrderManager) CancelOrder(orderId int64) error {
//	return nil
//}

func (om *OrderManager) NewOrder(pair string, price float64, strategy string) market.Order {
	return market.Order{}
}

func (om *OrderManager) HandleDecision(decision StrategyDecision) {
	// todo
}

func NewOrderManager(client market.IClient) *OrderManager {
	return &OrderManager{
		client: client,
	}
}
