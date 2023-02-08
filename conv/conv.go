package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres.
*/

	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende

// Konverterer Celcius til Kelvin 
func CelsiusToKelvin(value float64) float64 {
  return value + 273.15
}

// Konverterer Kelvin til Celcius 
func KelvinToCelsius(value float64) float64 {
  return value - 273.15
}

// Konverterer Celcius til Farhenhrit
func CelsiusToFarhenheit(value float64) float64 {
  return value*(9/5) + 32
}

// Konverterer Farhenheit til Celcius
func FarhenheitToCelsius(value float64) float64 {
  return (value - 32)*(5/9)
}

// Konverterer Farhenheit til Kelvin 
func FarhenheitToKelvin(value float64) float64 {
  return (value - 32)*(5/9) + 273.15
}

// Konverterer Kelvin til Ferhenheit
func KelvinToFarhenheit(value float64) float64 {
  return (value - 273.15)*(9/5) + 32
}


