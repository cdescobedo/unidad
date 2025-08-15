package unidad

import "github.com/shopspring/decimal"

type (
	Symbol   string
	UnitName string
	UnitType string
)

const (
	// Unit Types
	Length      UnitType = "length"
	Mass        UnitType = "mass"
	Time        UnitType = "time"
	Temperature UnitType = "temperature"
	Volume      UnitType = "volume"
	Energy      UnitType = "energy"
	// Cannonical Base Units
	Meter    Symbol = "m"
	Kilogram Symbol = "kg"
	Second   Symbol = "s"
	Kelvin   Symbol = "K"
	Joule    Symbol = "J"
	// Mass
	Gram     Symbol = "g"
	Tonne    Symbol = "t"
	Ounce    Symbol = "oz"
	Pound    Symbol = "lb"
	ShortTon Symbol = "T"
	LongTon  Symbol = "lt"
	// Time
	Minute Symbol = "min"
	Hour   Symbol = "hr"
	// Volume
	CubicMeter Symbol = "m3"
	Liter      Symbol = "L"
	Gallon     Symbol = "gal"
	// Energy
	KiloJoule    Symbol = "kJ"
	MegaJoule    Symbol = "MJ"
	Calorie      Symbol = "cal"
	KiloCalorie  Symbol = "kcal"
	WattHour     Symbol = "Wh"
	KiloWattHour Symbol = "kWh"
	Btu          Symbol = "Btu"
	MMBtu        Symbol = "MMBtu"
	Therm        Symbol = "thm"
	DecaTherm    Symbol = "dth"
)

var (
	// Mass Units
	gram = Unit{
		Symbol: Gram,
		Name:   "gram",
		Type:   Mass,
		ToBase: newDecimalFromString("0.001"),
	}
	kilogram = Unit{
		Symbol: Kilogram,
		Name:   "kilogram",
		Type:   Mass,
		ToBase: decimal.NewFromInt(1),
	}
	tonne = Unit{
		Symbol: Tonne,
		Name:   "tonne",
		Type:   Mass,
		ToBase: decimal.NewFromInt(1000),
	}
	ounce = Unit{
		Symbol: Ounce,
		Name:   "ounce",
		Type:   Mass,
		ToBase: newDecimalFromString("0.028349523125"),
	}
	pound = Unit{
		Symbol: Pound,
		Name:   "pound",
		Type:   Mass,
		ToBase: newDecimalFromString("0.45359237"),
	}
	shortTon = Unit{
		Symbol: ShortTon,
		Name:   "ton",
		Type:   Mass,
		ToBase: newDecimalFromString("907.18474"),
	}
	longTon = Unit{
		Symbol: LongTon,
		Name:   "long ton",
		Type:   Mass,
		ToBase: newDecimalFromString("1016.0469088"),
	}
	// Volume Units
	cubicMeter = Unit{
		Symbol: CubicMeter,
		Name:   "cubic meter",
		Type:   Volume,
		ToBase: decimal.NewFromInt(1),
	}
	liter = Unit{
		Symbol: Liter,
		Name:   "liter",
		Type:   Volume,
		ToBase: newDecimalFromString("0.001"),
	}
	gallon = Unit{
		Symbol: Gallon,
		Name:   "gallon",
		Type:   Volume,
		ToBase: newDecimalFromString("0.003785411784"),
	}
	// Energy Units
	joule = Unit{
		Symbol: Joule,
		Name:   "Joule",
		Type:   Energy,
		ToBase: decimal.NewFromInt(1),
	}
	kiloJoule = Unit{
		Symbol: KiloJoule,
		Name:   "kilojoule",
		Type:   Energy,
		ToBase: decimal.NewFromInt(1000),
	}
	megaJoule = Unit{
		Symbol: MegaJoule,
		Name:   "megajoule",
		Type:   Energy,
		ToBase: decimal.NewFromInt(1_000_000),
	}
	calorie = Unit{
		Symbol: Calorie,
		Name:   "calorie",
		Type:   Energy,
		ToBase: newDecimalFromString("4.184"),
	}
	kiloCalorie = Unit{
		Symbol: KiloCalorie,
		Name:   "kilocalorie",
		Type:   Energy,
		ToBase: decimal.NewFromInt(4184),
	}
	wattHour = Unit{
		Symbol: WattHour,
		Name:   "watt-hour",
		Type:   Energy,
		ToBase: decimal.NewFromInt(3600),
	}
	kiloWattHour = Unit{
		Symbol: KiloWattHour,
		Name:   "kilowatt-hour",
		Type:   Energy,
		ToBase: decimal.NewFromInt(3_600_000),
	}
	btu = Unit{
		Symbol: Btu,
		Name:   "British thermal unit",
		Type:   Energy,
		ToBase: newDecimalFromString("1055.056"),
	}
	mmbtu = Unit{
		Symbol: MMBtu,
		Name:   "million British thermal unit",
		Type:   Energy,
		ToBase: decimal.NewFromInt(1_055_056_000),
	}
	therm = Unit{
		Symbol: Therm,
		Name:   "therm",
		Type:   Energy,
		ToBase: decimal.NewFromInt(105_505_600),
	}
	decatherm = Unit{
		Symbol: DecaTherm,
		Name:   "decatherm",
		Type:   Energy,
		ToBase: decimal.NewFromInt(1_055_056_000),
	}
)

var defaultUnits = map[Symbol]Unit{
	// Length

	// Mass
	Gram:     gram,
	Kilogram: kilogram,
	Tonne:    tonne,
	Ounce:    ounce,
	Pound:    pound,
	ShortTon: shortTon,
	LongTon:  longTon,
	// Volume
	CubicMeter: cubicMeter,
	Liter:      liter,
	Gallon:     gallon,
	// Energy
	Joule:        joule,
	KiloJoule:    kiloJoule,
	MegaJoule:    megaJoule,
	Calorie:      calorie,
	KiloCalorie:  kiloCalorie,
	WattHour:     wattHour,
	KiloWattHour: kiloWattHour,
	Btu:          btu,
	MMBtu:        mmbtu,
	Therm:        therm,
	DecaTherm:    decatherm,
}

type Unit struct {
	Name   UnitName
	Symbol Symbol
	Type   UnitType
	ToBase decimal.Decimal
}

func NewUnit(name UnitName, unitSymbol Symbol, unitType UnitType, toBase decimal.Decimal) Unit {
	return Unit{
		Name:   name,
		Symbol: unitSymbol,
		Type:   unitType,
		ToBase: toBase,
	}
}
