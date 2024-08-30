package calculator

const CONSUMPTION_TAX float64 = 0.10

func TaxCalcuation(price int64) int64 {
	taxPrice := float64(price) * CONSUMPTION_TAX
	return int64(taxPrice)
}
