package main

import (
	"flag"
	"fmt"
	"strings"
	//fun "github.com/mienord/funtemps/funfacts"
	"github.com/mienord/funtemps/conv" 
	
)

var (
	fahr    float64
	celsius float64
	kelvin  float64
	out     string
	funfact string
	tempSkala string
)

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperature in Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&out, "out", "C", "output temperature in C (Celsius), F (Fahrenheit), K (Kelvin)")
	flag.StringVar(&funfact, "funfact", "sun", "\"fun-facts\" about sun, moon and earth")
	flag.StringVar(&tempSkala, "t", "C", "temperature scale to use when providing funfact")
}


func main() {
	flag.Parse()

	// Check if multiple temperature flags are not used simultaneously
	if (fahr != 0 && (celsius != 0 || kelvin != 0)) ||
		(celsius != 0 && (fahr != 0 || kelvin != 0)) ||
		(kelvin != 0 && (fahr != 0 || celsius != 0)) {
		fmt.Println("Error: Cannot use multiple temperature flags together.")
		return
	}

	// Check if temperature flags and funfact flag are used together
	if (fahr != 0 || celsius != 0 || kelvin != 0) && funfact != "sun" {
		fmt.Println("Error: Cannot use temperature flags with funfact flag.")
		return
	}

	// Check if funfact flag is used with temperature scale flag
	if funfact != "sun" && tempSkala == "C" {
		fmt.Println("Error: Must use temperature scale flag with funfact flag.")
		return
	}
	

// Check the out flag and convert temperature
switch out {
case "C":
	if fahr != 0 {
		celsius = (conv.FahrenheitToCelsius(fahr))
		fmt.Printf("%s %s er %s %s\n", formatNumber(fahr, 2), "°F", formatNumber(celsius, 2), "°C")
	} else if kelvin != 0 {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Printf("%s %s er %s %s\n", formatNumber(kelvin, 2), "K", formatNumber(celsius, 2), "°C")
	} else {
		fmt.Printf("%s %s er %s %s\n", formatNumber(celsius, 2), "°C", formatNumber(celsius, 2), "°C")
	}
case "F":
	if celsius != 0 {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%s %s er %s %s\n", formatNumber(celsius, 2), "°C", formatNumber(fahr, 2), "°F")
	} else if kelvin != 0 {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Printf("%s %s er %s %s\n", formatNumber(kelvin, 2), "K", formatNumber(fahr, 2), "°F")
	} else {
		fmt.Printf("%s %s er %s %s\n", formatNumber(fahr, 2), "°F", formatNumber(fahr, 2), "°F")
	}
case "K":
	if fahr != 0 {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%s %s er %s %s\n", formatNumber(fahr, 2), "°F", formatNumber(kelvin, 2), "K")
	} else if celsius != 0 {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%s %s er %s %s\n", formatNumber(celsius, 2), "°C", formatNumber(kelvin, 2), "K")
	} else {
		fmt.Printf("%s %s er %s %s\n", formatNumber(kelvin, 2), "K", formatNumber(kelvin, 2), "K")
	}
}


}

func formatNumber(num float64, decimals int) string {
    // Convert float to string with specified number of decimal places
    // and split into whole and fractional parts
    parts := strings.Split(fmt.Sprintf(fmt.Sprintf("%%.%df", decimals), num), ".")

    // If there's only a whole number part, format it with thousands separator
    if len(parts) == 1 {
        whole := parts[0]
        formatted := ""
        for i, c := range whole {
            if i > 0 && (len(whole)-i)%3 == 0 {
                formatted += " " // Add space as thousands separator
            }
            formatted += string(c)
        }
        return formatted
    }

    // If there's a fractional part, format the whole and fractional parts separately
    whole := parts[0]
    frac := parts[1]
    formatted := ""
    for i, c := range whole {
        if i > 0 && (len(whole)-i)%3 == 0 {
            formatted += " " // Add space as thousands separator
        }
        formatted += string(c)
    }

    // Trim trailing zeros from fractional part and return formatted number
    frac = strings.TrimRight(frac, "0")
    if frac == "" {
        return formatted
    }
    return fmt.Sprintf("%s.%s", formatted, frac)
}



/*
func  GetFunFacts(about string) {
	funFacts := fun.GetFunFacts(about)
	for i, fact := range funFacts {
		fmt.Println(i+1, fact)
	}
}
*/
