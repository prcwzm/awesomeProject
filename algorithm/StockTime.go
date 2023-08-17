package algorithm

import "math"

func StockProfit(price []float64) (earn float64) {
	if price == nil || len(price) == 0 {
		return 0
	}
	priceChange := make([]float64, len(price), len(price))
	for i := 1; i < len(price); i++ {
		priceChange[i] = price[i] - price[i-1]
	}
	earn = praseflow(priceChange)
	return
}

func praseflow(priceChange []float64) (res float64) {
	res = 0
	for i := 1; i < len(priceChange); i++ {
		priceChange[i] += math.Max(priceChange[i-1], 0)
		res = math.Max(res, priceChange[i])
	}
	return
}
