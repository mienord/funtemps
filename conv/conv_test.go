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


func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: -128.92, want: -89.4},
		{input: 16445.93, want: 9118.85},
		{input: 27000032, want: 15000000},
		{input: 9941.53, want: 5504.85},
		{input: -297.4, want: -183},
		{input: 221.8, want: 106},

	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.7, want: 329.82},
		{input: -89.4, want: 183.75},
		{input: 9118.85, want: 9392},
		{input: 15000000, want: 15000273.15},
		{input: 5504.85, want: 5778},
		{input: -183, want: 90.15},
		{input: 106, want: 379.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want: 56.7},
	    {input: 183.75, want: -89.4},
    	{input: 9392, want: 9118.85},
    	{input: 15000273.15, want: 15000000},
    	{input: 5778, want: 5504.85},
    	{input: 90.15, want: -183},
    	{input: 379.15, want: 106},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.7, want: 134},
		{input: -89.4, want: -128.92},
        {input: 9008.183, want: 16246.73},
        {input: 14962204, want: 26932000},
        {input: 5504.85, want: 9941.53},
        {input: -183, want: -297.4},
        {input: 106, want: 221.8},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 329.82},
		{input: -128.92, want: 183.75},
        {input: 16445.93, want: 9392},
        {input: 27000032, want: 15000273.15},
        {input: 9941.53, want: 5778},
        {input: -297.4, want: 90.15},
        {input: 221.8, want: 379.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
    }
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 329.82, want:134 },
		{input: 183.75, want: -128.92},
        {input: 9281.333, want: 16246.73},
        {input: 14962478, want: 26932000},
        {input: 5778, want: 9941.53},
        {input: 90.15, want: -297.4},
        {input: 379.15, want: 221.8},

	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(tc.want - got) > 1 {
			t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
		}
	}
}

/*

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
*/