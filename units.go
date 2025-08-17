package unidad

import (
	"log"

	"github.com/shopspring/decimal"
)

type (
	Symbol   string
	UnitName string
	Quantity string
)

type Unit struct {
	Name     UnitName
	Symbol   Symbol
	Quantity Quantity
	ToBase   decimal.Decimal
}

const (
	// Common Metric Prefixes. Spaces are used in between the zeros in the comments for readability.
	Tera  = 10e12  // 1 000 000 000 000
	Giga  = 10e9   // 1 000 000 000
	Mega  = 10e6   // 1 000 000
	Kilo  = 10e3   // 1 000
	Hecto = 10e2   // 100
	Deka  = 10e1   // 10
	Deci  = 10e-1  // 0.1
	Milli = 10e-3  // 0.001
	Micro = 10e-6  // 0.000 001
	Nano  = 10e-9  // 0.000 000 001
	Pico  = 10e-12 // 0.000 000 000 001

	// Base Quantity Types
	Length          Quantity = "length"
	Mass            Quantity = "mass"
	Time            Quantity = "time"
	ElectricCurrent Quantity = "electric current"
	Temperature     Quantity = "temperature"
	// Derived Quantity Types
	Area Quantity = "area"
	// This is a unitless unit to represent constant of 1
	// SI Base Units
	Meter    UnitName = "meter"
	Kilogram UnitName = "kilogram"
	Second   UnitName = "second"
	Ampere   UnitName = "ampere"
	Kelvin   UnitName = "kelvin"
	// Length
	Millimeter UnitName = "millimeter"
	Centimeter UnitName = "centimeter"
	Kilometer  UnitName = "kilometer"
	// Derived Units
	// Area
	SquareMeter UnitName = "square meter"
)

var unitRegistry = map[UnitName]Unit{
	Millimeter:  {Millimeter, "mm", Length, newDecimalFromString("0.001")},
	Centimeter:  {Centimeter, "cm", Length, newDecimalFromString("0.01")},
	Meter:       {Meter, "m", Length, decimal.NewFromInt(1)},
	Kilometer:   {Kilometer, "km", Length, decimal.NewFromInt(1000)},
	Kilogram:    {Kilogram, "kg", Mass, decimal.NewFromInt(1)},
	Second:      {Second, "s", Time, decimal.NewFromInt(1)},
	Ampere:      {Ampere, "A", ElectricCurrent, decimal.NewFromInt(1)},
	Kelvin:      {Kelvin, "K", Temperature, decimal.NewFromInt(1)},
	SquareMeter: {SquareMeter, "m2", Area, decimal.NewFromInt(1)},
}

func GetUnit(unitName UnitName) (Unit, bool) {
	unit, ok := unitRegistry[unitName]
	return unit, ok
}

func newDecimalFromString(value string) decimal.Decimal {
	decVal, err := decimal.NewFromString(value)
	if err != nil {
		log.Fatal(err)
	}
	return decVal
}
