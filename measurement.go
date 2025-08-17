package unidad

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Measurement struct {
	Value decimal.Decimal
	Unit  Unit
}

func NewFromString(value string, unitName UnitName) (Measurement, error) {
	unit, ok := GetUnit(unitName)
	if !ok {
		return Measurement{}, fmt.Errorf("unit not registered: %s is not in unit registry", unitName)
	}

	decVal, err := decimal.NewFromString(value)
	if err != nil {
		return Measurement{}, err
	}
	return Measurement{decVal, unit}, nil
}

func NewFromInt(value int64, unitName UnitName) (Measurement, error) {
	unit, ok := GetUnit(unitName)
	if !ok {
		return Measurement{}, fmt.Errorf("unit not registered: %s is not in unit registry", unitName)
	}

	decValue := decimal.NewFromInt(value)
	return Measurement{decValue, unit}, nil
}

func (m Measurement) String() string {
	return fmt.Sprintf("%s %s", m.Value, m.Unit.Symbol)
}

func (m Measurement) ConvertTo(unitName UnitName) (Measurement, error) {
	// TODO: implement rounding procedures 4.4.1.1 and 4.4.1.2
	targetUnit, ok := GetUnit(unitName)
	if !ok {
		return Measurement{}, fmt.Errorf("unit not registered: %s is not in unit registry", unitName)
	}

	if targetUnit.Quantity != m.Unit.Quantity {
		return Measurement{}, fmt.Errorf("quantity mismatch: source quantity %s cannot be converted to target quantity %s", m.Unit.Quantity, targetUnit.Quantity)
	}
	srcValueInBase := m.Value.Mul(m.Unit.ToBase)
	newValue := srcValueInBase.Div(targetUnit.ToBase)
	newMeasurement := Measurement{
		Value: newValue,
		Unit:  targetUnit,
	}
	return newMeasurement, nil
}

func (m Measurement) Add(other Measurement) (Measurement, error) {
	if m.Unit.Quantity != other.Unit.Quantity {
		return Measurement{}, fmt.Errorf("quantity mismatch: quantity %s cannot be added to quantity %s", other.Unit.Quantity, m.Unit.Quantity)
	}

	mBaseValue := m.Value.Mul(m.Unit.ToBase)
	otherBaseValue := other.Value.Mul(other.Unit.ToBase)

	sumInBase := mBaseValue.Add(otherBaseValue)
	newValue := sumInBase.Div(m.Unit.ToBase)

	return Measurement{Value: newValue, Unit: m.Unit}, nil
}

func (m Measurement) Sub(other Measurement) (Measurement, error) {
	if m.Unit.Quantity != other.Unit.Quantity {
		return Measurement{}, fmt.Errorf("quantity mismatch: quantity %s cannot be added to quantity %s", other.Unit.Quantity, m.Unit.Quantity)
	}

	mBaseValue := m.Value.Mul(m.Unit.ToBase)
	otherBaseValue := m.Value.Mul(other.Unit.ToBase)

	sumInBase := mBaseValue.Sub(otherBaseValue)
	newValue := sumInBase.Div(m.Unit.ToBase)

	return Measurement{newValue, m.Unit}, nil
}

func (m Measurement) Mul(other Measurement) (Measurement, error) {
	return Measurement{}, nil
}

func (m Measurement) ScalarMul(scalar decimal.Decimal) Measurement {
	newValue := m.Value.Mul(scalar)
	return Measurement{newValue, m.Unit}
}

func (m Measurement) Div(other Measurement) (Measurement, error) {
	return Measurement{}, nil
}

func (m Measurement) ScalarDiv(scalar decimal.Decimal) (Measurement, error) {
	if scalar.IsZero() {
		return Measurement{}, fmt.Errorf("division by zero")
	}
	newValue := m.Value.Div(scalar)
	return Measurement{newValue, m.Unit}, nil
}
