package market

type MockedClient struct {
	data []Candle
}

func (c *MockedClient) Fetch(_, _ string, _ int) []Candle {
	return c.data
}

func (c *MockedClient) FetchPeriod(_, _ string, _ int64, _ int64) []Candle {
	return c.data
}

func (c *MockedClient) ListOrders(pair string) []Order {
	return []Order{
		{
			OrderID:     1,
			Price:       1.25,
			Pair:        pair,
			IsOpen:      false,
			IsCanceled:  false,
			IsCompleted: true,
		},
		{
			OrderID:     2,
			Price:       4.99,
			Pair:        "BTCUSDT",
			IsOpen:      true,
			IsCanceled:  false,
			IsCompleted: false,
		},
	}
}

func (c *MockedClient) OpenMarketOrder(_ Order) (Order, error) {
	return Order{
		OrderID:     2,
		Price:       4.99,
		Pair:        "BTCUSDT",
		IsOpen:      true,
		IsCanceled:  false,
		IsCompleted: false,
	}, nil
}

func NewMockedClient(data []Candle) IClient {
	return &MockedClient{
		data: data,
	}
}

func GetMockCandleData(limit int) []Candle {
	c := []Candle{
		{
			OpenTime:  1668729600000,
			Open:      16692.56000000,
			High:      17011.00000000,
			Low:       16546.04000000,
			Close:     16574.92000000,
			Volume:    182284.38651000,
			CloseTime: 1668815999999,
		},
		{
			OpenTime:  1668643200000,
			Open:      16661.61000000,
			High:      16751.00000000,
			Low:       16410.74000000,
			Close:     16692.56000000,
			Volume:    228038.97873000,
			CloseTime: 1668729599999,
		},
		{
			OpenTime:  1668556800000,
			Open:      16900.57000000,
			High:      17015.92000000,
			Low:       16378.61000000,
			Close:     16662.76000000,
			Volume:    261493.40809000,
			CloseTime: 1668643199999,
		},
		{
			OpenTime:  1668470400000,
			Open:      16617.72000000,
			High:      17134.69000000,
			Low:       16527.72000000,
			Close:     16900.57000000,
			Volume:    282461.84391000,
			CloseTime: 1668556799999,
		},
		{
			OpenTime:  1668384000000,
			Open:      16331.78000000,
			High:      17190.00000000,
			Low:       15815.21000000,
			Close:     16619.46000000,
			Volume:    380210.77750000,
			CloseTime: 1668470399999,
		},
		{
			OpenTime:  1668297600000,
			Open:      16813.16000000,
			High:      16954.28000000,
			Low:       16229.00000000,
			Close:     16329.85000000,
			Volume:    184960.78846000,
			CloseTime: 1668383999999,
		},
		{
			OpenTime:  1668211200000,
			Open:      17069.98000000,
			High:      17119.10000000,
			Low:       16631.39000000,
			Close:     16812.08000000,
			Volume:    167819.96035000,
			CloseTime: 1668297599999,
		},
		{
			OpenTime:  1668124800000,
			Open:      17602.45000000,
			High:      17695.00000000,
			Low:       16361.60000000,
			Close:     17070.31000000,
			Volume:    393552.86492000,
			CloseTime: 1668211199999,
		},
		{
			OpenTime:  1668038400000,
			Open:      15922.68000000,
			High:      18199.00000000,
			Low:       15754.26000000,
			Close:     17601.15000000,
			Volume:    608448.36432000,
			CloseTime: 1668124799999,
		},
		{
			OpenTime:  1667952000000,
			Open:      18545.38000000,
			High:      18587.76000000,
			Low:       15588.00000000,
			Close:     15922.81000000,
			Volume:    731926.92972900,
			CloseTime: 1668038399999,
		},
	}
	if limit > len(c) {
		limit = len(c)
	}
	return c[len(c)-limit:]
}
