package conv

import (
	"math"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/

func TestFarenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 32, want: 0},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}
}
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}
}
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -273.15, want: 0},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}
}
func TestCelsiusToFarenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 273.2},
		{input: -32, want: -25.6},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}
}
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
		{input: 373.15, want: 212},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
		{input: 373.15, want: 212},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
		//		if !reflect.DeepEqual(tc.want, got) {
		//			t.Errorf("expected: %v, got: %v", tc.want, got)
		//		}
	}

}

// De andre testfunksjonene implementeres her
// ...
