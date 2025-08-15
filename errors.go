package unidad

import "fmt"

type UnitNotRegisteredError struct {
	UnitSymbol UnitSymbol
}

func (e *UnitNotRegisteredError) Error() string {
	return fmt.Sprintf("unit %s is not registered", e.UnitSymbol)
}

type TypeMismatchError struct {
	SourceUnit UnitType
	TargetUnit UnitType
}

func (e *TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: cannot perform conversion between %s and %s", e.SourceUnit, e.TargetUnit)
}
