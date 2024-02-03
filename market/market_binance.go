package market

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"log"
	"strconv"
)

type BinanceClient struct {
	client *binance.Client
}

func (c *BinanceClient) Fetch(pair, timeframe string, limit int) []Candle {
	klines, err := c.client.NewKlinesService().
		Symbol(pair).
		Interval(timeframe).
		Limit(limit).
		Do(context.TODO())

	candles := make([]Candle, 0, len(klines))

	if err != nil {
		log.Printf("market: error on fetching candles data (%s)", err.Error())
		return candles
	}

	for i := len(klines) - 1; i >= 0; i-- {
		candles = append(candles, kline2Candle(klines[i]))
	}

	return candles
}

func (c *BinanceClient) FetchPeriod(pair, timeframe string, start int64, end int64) []Candle {
	klines, err := c.client.NewKlinesService().
		Symbol(pair).
		Interval(timeframe).
		StartTime(start).
		EndTime(end).
		Do(context.TODO())

	candles := make([]Candle, 0, len(klines))

	if err != nil {
		log.Printf("market: error on fetching candles data: %s", err.Error())
		return candles
	}

	for i := len(klines) - 1; i >= 0; i-- {
		candles = append(candles, kline2Candle(klines[i]))
	}

	return candles
}

func (c *BinanceClient) ListOrders(pair string) []Order {
	binanceOrders, err := c.client.NewListOrdersService().Symbol(pair).Do(context.TODO())
	if err != nil {
		log.Printf("market: error on fetching orders: %s", err.Error())
	}
	var orders []Order
	for _, binanceOrder := range binanceOrders {
		orders = append(orders, binanceOrder2Order(binanceOrder))
	}
	return orders
}

func (c *BinanceClient) OpenMarketOrder(o Order) (Order, error) {
	res, err := c.client.NewCreateOrderService().
		Symbol(o.Pair).
		Type(binance.OrderTypeMarket).
		Do(context.TODO())

	if err != nil {
		return o, err
	}
	return createOrderResponse2Order(res), err
}

func kline2Candle(k *binance.Kline) Candle {
	openPrice, _ := strconv.ParseFloat(k.Open, 64)
	closePrice, _ := strconv.ParseFloat(k.Close, 64)
	highPrice, _ := strconv.ParseFloat(k.High, 64)
	lowPrice, _ := strconv.ParseFloat(k.Low, 64)
	volume, _ := strconv.ParseFloat(k.Volume, 64)

	return Candle{
		OpenTime:  k.OpenTime,  // idx in response: 0
		Open:      openPrice,   // idx in response: 1
		High:      highPrice,   // idx in response: 2
		Low:       lowPrice,    // idx in response: 3
		Close:     closePrice,  // idx in response: 4
		Volume:    volume,      // idx in response: 5
		CloseTime: k.CloseTime, // idx in response: 6
	}
}

func binanceOrder2Order(order *binance.Order) Order {
	price, _ := strconv.ParseFloat(order.Price, 64)
	isOpen, isCanceled, isCompleted := resolveOrderStatusFlags(order.Status)
	return Order{
		OrderID:     order.OrderID,
		Price:       price,
		Pair:        order.Symbol,
		IsOpen:      isOpen,
		IsCanceled:  isCanceled,
		IsCompleted: isCompleted,
	}
}

func createOrderResponse2Order(res *binance.CreateOrderResponse) Order {
	price, _ := strconv.ParseFloat(res.Price, 64)
	isOpen, isCanceled, isCompleted := resolveOrderStatusFlags(res.Status)
	return Order{
		OrderID:     res.OrderID,
		Price:       price,
		Pair:        res.Symbol,
		IsOpen:      isOpen,
		IsCanceled:  isCanceled,
		IsCompleted: isCompleted,
	}
}

func resolveOrderStatusFlags(status binance.OrderStatusType) (bool, bool, bool) {
	isOpen, isCanceled, isCompleted := false, false, false
	if status == binance.OrderStatusTypeNew || status == binance.OrderStatusTypePartiallyFilled {
		isOpen = true
	}
	if status == binance.OrderStatusTypeCanceled ||
		status == binance.OrderStatusTypeExpired ||
		status == binance.OrderStatusTypeRejected {
		isCanceled = true
	}
	if status == binance.OrderStatusTypeFilled {
		isCompleted = true
	}
	return isOpen, isCanceled, isCompleted
}
