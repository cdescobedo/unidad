package unidad

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

func TestSuccessfulConvertion(t *testing.T) {
	tests := []struct {
		name       string
		measurment Measurement
		targetUnit Symbol
		want       decimal.Decimal
	}{
		{
			"grams to kilograms",
			Measurement{
				decimal.NewFromInt(1000),
				Gram,
			},
			Kilogram,
			decimal.NewFromInt(1),
		},
		{
			"Btu to Joule",
			Measurement{
				decimal.NewFromInt(1),
				Btu,
			},
			Joule,
			newDecimalFromString("1055.056"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newMeasurement, err := tt.measurment.ConvertTo(tt.targetUnit)
			if err != nil {
				t.Fatal(err)
			}
			if !newMeasurement.Value.Equal(tt.want) {
				t.Errorf("Convert(%v, %s) = %v %s, want: %v %s", tt.measurment, tt.targetUnit, newMeasurement.Value, newMeasurement.Unit, tt.want, tt.targetUnit)
			}
		})
	}
}

func TestConvertUnitNotRegisteredError(t *testing.T) {
	tests := []struct {
		name       string
		measurment Measurement
		targetUnit Symbol
		want       decimal.Decimal
	}{
		{
			"stones to grams",
			Measurement{
				decimal.NewFromInt(1),
				"st",
			},
			Gram,
			decimal.NewFromInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var unitNotRegisteredError *UnitNotRegisteredError
			_, err := tt.measurment.ConvertTo(tt.targetUnit)
			if err == nil {
				t.Fatalf("expected: %T, actual: nil", UnitNotRegisteredError{})
			}
			if !errors.As(err, &unitNotRegisteredError) {
				t.Fatalf("expected error of type(%T) but got error of type(%T)", UnitNotRegisteredError{}, err)
			}
		})
	}
}

func TestConvertQuantityMismatchError(t *testing.T) {
	tests := []struct {
		name       string
		measurment Measurement
		targetUnit Symbol
		want       decimal.Decimal
	}{
		{
			"mass to energy",
			Measurement{
				decimal.NewFromInt(1),
				Gram,
			},
			Joule,
			decimal.NewFromInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var typeMismatchError *QuantityMismatchError
			_, err := tt.measurment.ConvertTo(tt.targetUnit)
			if err == nil {
				t.Fatalf("expected: %T, actual: nil", QuantityMismatchError{})
			}
			if !errors.As(err, &typeMismatchError) {
				t.Fatalf("expected error of type(%T) but got error of type(%T)", QuantityMismatchError{}, err)
			}
		})
	}
}
