package conv

import (
	"testing"
	"math"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/

//testfunskjon for farhenheit til celsius
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

//testfunksjon for celsius til kelvin
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.7, want: 329.82},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

//testfunksjon for kelvin til celsius
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.7},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

//testfunksjon for celsius til farhenheit
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.7, want: 134},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

//testfunksjon for farhenheit til kelvin
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
}
}

//testfunksjon for kelvin til farhenheit
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 134},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
}
}

