package unidad

import (
	"github.com/shopspring/decimal"
)

func newDecimalFromString(val string) decimal.Decimal {
	decimalVal, err := decimal.NewFromString(val)
	if err != nil {
		panic(err)
	}
	return decimalVal
}
