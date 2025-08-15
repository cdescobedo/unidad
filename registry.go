package unidad

import (
	"fmt"
)

type UnitRegistry struct {
	units map[Symbol]Unit
}

var DefaultUnitRegistry = &UnitRegistry{
	units: defaultUnits,
}

var unitRegistry = DefaultUnitRegistry

func (r *UnitRegistry) GetUnit(unitSymbol Symbol) (Unit, bool) {
	unit, ok := r.units[unitSymbol]
	return unit, ok
}

func (r *UnitRegistry) Register(unit Unit) error {
	if _, ok := r.GetUnit(unit.Symbol); ok {
		return fmt.Errorf("unit already registered")
	}
	r.units[unit.Symbol] = unit
	return nil
}
