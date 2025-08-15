package unidad

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type measurement struct {
	value decimal.Decimal
	unit  Unit
}

func NewFromInt(value int64, unitSymbol Symbol) (measurement, error) {
	decimalValue := decimal.NewFromInt(value)
	unit, ok := unitRegistry.GetUnit(unitSymbol)
	if !ok {
		return measurement{}, &UnitNotRegisteredError{unitSymbol}
	}
	return measurement{decimalValue, unit}, nil
}

func NewFromFloat(value float64, unitSymbol Symbol) (measurement, error) {
	decimalValue := decimal.NewFromFloat(value)
	unit, ok := unitRegistry.GetUnit(unitSymbol)
	if !ok {
		return measurement{}, &UnitNotRegisteredError{Symbol: unitSymbol}
	}

	return measurement{value: decimalValue, unit: unit}, nil
}

func NewFromString(value string, unitSymbol Symbol) (measurement, error) {
	decimalValue, err := decimal.NewFromString(value)
	if err != nil {
		return measurement{}, err
	}

	unit, ok := DefaultUnitRegistry.GetUnit(unitSymbol)
	if !ok {
		return measurement{}, &UnitNotRegisteredError{Symbol: unitSymbol}
	}

	return measurement{value: decimalValue, unit: unit}, nil
}

func (m measurement) String() string {
	return fmt.Sprintf("%v %s", m.value, m.unit.Symbol)
}

func (m measurement) Add(other measurement) (measurement, error) {
	if !m.IsType(other.unit.Type) {
		return measurement{}, &TypeMismatchError{SourceUnit: m.unit.Type, TargetUnit: other.unit.Type}
	}

	otherNew, err := other.ConvertTo(m.unit.Symbol)
	if err != nil {
		return measurement{}, err
	}
	newValue := m.value.Add(otherNew.value)
	return measurement{value: newValue, unit: m.unit}, nil
}

func (m measurement) Sub(other measurement) (measurement, error) {
	if !m.IsType(other.unit.Type) {
		return measurement{}, &TypeMismatchError{SourceUnit: m.unit.Type, TargetUnit: other.unit.Type}
	}

	otherNew, err := other.ConvertTo(m.unit.Symbol)
	if err != nil {
		return measurement{}, err
	}
	newValue := m.value.Sub(otherNew.value)
	return measurement{value: newValue, unit: m.unit}, nil
}

func (m measurement) IsType(unitType UnitType) bool {
	if m.unit.Type != unitType {
		return false
	}
	return true
}

func (m measurement) ConvertTo(targetUnit Symbol) (measurement, error) {
	newUnit, ok := unitRegistry.GetUnit(targetUnit)
	if !ok {
		return measurement{}, &UnitNotRegisteredError{Symbol: targetUnit}
	}

	if !m.IsType(newUnit.Type) {
		return measurement{}, &TypeMismatchError{SourceUnit: m.unit.Type, TargetUnit: newUnit.Type}
	}

	srcValueInBase := m.value.Mul(m.unit.ToBase)
	newValue := srcValueInBase.Div(newUnit.ToBase)
	return measurement{value: newValue, unit: newUnit}, nil
}
