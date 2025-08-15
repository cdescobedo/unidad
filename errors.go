package unidad

import "fmt"

type UnitNotRegisteredError struct {
	Symbol Symbol
}

func (e *UnitNotRegisteredError) Error() string {
	return fmt.Sprintf("unit %s is not registered", e.Symbol)
}

type TypeMismatchError struct {
	SourceUnit UnitType
	TargetUnit UnitType
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: cannot perform conversion between %s and %s", e.SourceUnit, e.TargetUnit)
}
