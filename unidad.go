package unidad

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type (
	UnitSymbol string
	UnitName   string
	UnitType   string
)

type Unit struct {
	Symbol UnitSymbol
	Name   UnitName
	Type   UnitType
	ToBase decimal.Decimal
}

const (
	MassUnit     UnitType   = "mass"
	VolumeUnit   UnitType   = "volume"
	EnergyUnit   UnitType   = "energy"
	Gram         UnitSymbol = "g"
	Kilogram     UnitSymbol = "kg"
	Tonne        UnitSymbol = "t"
	Ounce        UnitSymbol = "oz"
	Pound        UnitSymbol = "lb"
	ShortTon     UnitSymbol = "T"
	LongTon      UnitSymbol = "lt"
	CubicMeter   UnitSymbol = "cm3"
	Liter        UnitSymbol = "L"
	Gallon       UnitSymbol = "gal"
	Joule        UnitSymbol = "J"
	KiloJoule    UnitSymbol = "kJ"
	MegaJoule    UnitSymbol = "MJ"
	Calorie      UnitSymbol = "cal"
	KiloCalorie  UnitSymbol = "kcal"
	WattHour     UnitSymbol = "Wh"
	KiloWattHour UnitSymbol = "kWh"
	Btu          UnitSymbol = "Btu"
	MMBtu        UnitSymbol = "MMBtu"
	Therm        UnitSymbol = "thm"
	DecaTherm    UnitSymbol = "dth"
)

var (
	// Mass Units
	gram = Unit{
		Symbol: Gram,
		Name:   "gram",
		Type:   MassUnit,
		ToBase: newDecimalFromString("0.001"),
	}
	kilogram = Unit{
		Symbol: Kilogram,
		Name:   "kilogram",
		Type:   MassUnit,
		ToBase: decimal.NewFromInt(1),
	}
	tonne = Unit{
		Symbol: Tonne,
		Name:   "tonne",
		Type:   MassUnit,
		ToBase: decimal.NewFromInt(1000),
	}
	ounce = Unit{
		Symbol: Ounce,
		Name:   "ounce",
		Type:   MassUnit,
		ToBase: newDecimalFromString("0.028349523125"),
	}
	pound = Unit{
		Symbol: Pound,
		Name:   "pound",
		Type:   MassUnit,
		ToBase: newDecimalFromString("0.45359237"),
	}
	shortTon = Unit{
		Symbol: ShortTon,
		Name:   "ton",
		Type:   MassUnit,
		ToBase: newDecimalFromString("907.18474"),
	}
	longTon = Unit{
		Symbol: LongTon,
		Name:   "long ton",
		Type:   MassUnit,
		ToBase: newDecimalFromString("1016.0469088"),
	}
	// Volume Units
	cubicMeter = Unit{
		Symbol: CubicMeter,
		Name:   "cubic meter",
		Type:   VolumeUnit,
		ToBase: decimal.NewFromInt(1),
	}
	liter = Unit{
		Symbol: Liter,
		Name:   "liter",
		Type:   VolumeUnit,
		ToBase: newDecimalFromString("0.001"),
	}
	gallon = Unit{
		Symbol: Gallon,
		Name:   "gallon",
		Type:   VolumeUnit,
		ToBase: newDecimalFromString("0.003785411784"),
	}
	// Energy Units
	joule = Unit{
		Symbol: Joule,
		Name:   "Joule",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(1),
	}
	kiloJoule = Unit{
		Symbol: KiloJoule,
		Name:   "kilojoule",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(1000),
	}
	megaJoule = Unit{
		Symbol: MegaJoule,
		Name:   "megajoule",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(1_000_000),
	}
	calorie = Unit{
		Symbol: Calorie,
		Name:   "calorie",
		Type:   EnergyUnit,
		ToBase: newDecimalFromString("4.184"),
	}
	kiloCalorie = Unit{
		Symbol: KiloCalorie,
		Name:   "kilocalorie",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(4184),
	}
	wattHour = Unit{
		Symbol: WattHour,
		Name:   "watt-hour",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(3600),
	}
	kiloWattHour = Unit{
		Symbol: KiloWattHour,
		Name:   "kilowatt-hour",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(3_600_000),
	}
	btu = Unit{
		Symbol: Btu,
		Name:   "British thermal unit",
		Type:   EnergyUnit,
		ToBase: newDecimalFromString("1055.056"),
	}
	mmbtu = Unit{
		Symbol: MMBtu,
		Name:   "million British thermal unit",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(1_055_056_000),
	}
	therm = Unit{
		Symbol: Therm,
		Name:   "therm",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(105_505_600),
	}
	decatherm = Unit{
		Symbol: DecaTherm,
		Name:   "decatherm",
		Type:   EnergyUnit,
		ToBase: decimal.NewFromInt(1_055_056_000),
	}
)

var unitRegistry = map[UnitSymbol]Unit{
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

func GetUnit(unitSymbol UnitSymbol) (Unit, bool) {
	unit, ok := unitRegistry[unitSymbol]
	return unit, ok
}

func newDecimalFromString(val string) decimal.Decimal {
	decimalVal, err := decimal.NewFromString(val)
	if err != nil {
		panic(err)
	}
	return decimalVal
}

type Measurement struct {
	Value decimal.Decimal
	Unit  UnitSymbol
}

func NewFromFloat(value float64, unit UnitSymbol) Measurement {
	decimalValue := decimal.NewFromFloat(value)
	return Measurement{decimalValue, unit}
}

func NewFromString(value string, unit UnitSymbol) (Measurement, error) {
	decimalValue, err := decimal.NewFromString(value)
	if err != nil {
		return Measurement{}, err
	}
	return Measurement{decimalValue, unit}, nil
}

func (m Measurement) String() string {
	return fmt.Sprintf("%v %s", m.Value, m.Unit)
}

func (m Measurement) ConvertTo(targetUnit UnitSymbol) (Measurement, error) {
	sourceUnit, ok := unitRegistry[m.Unit]
	if !ok {
		return Measurement{}, &UnitNotRegisteredError{UnitSymbol: m.Unit}
	}

	newUnit, ok := unitRegistry[targetUnit]
	if !ok {
		return Measurement{}, &UnitNotRegisteredError{UnitSymbol: targetUnit}
	}

	if sourceUnit.Type != newUnit.Type {
		return Measurement{}, &TypeMismatchError{SourceUnit: sourceUnit.Type, TargetUnit: newUnit.Type}
	}

	srcValueInBase := m.Value.Mul(sourceUnit.ToBase)
	newValue := srcValueInBase.Div(newUnit.ToBase)
	return Measurement{newValue, targetUnit}, nil
}
