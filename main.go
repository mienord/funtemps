package main

import (
	"flag"
	"fmt"
	//fun "github.com/mienord/funtemps/funfacts"
	//con "github.com/mienord/funtemps/conv" 
	
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
	
// Check the out flag and convert temperature - "%.2f %s\n" sørger for at det kun printer 2 desimaler i output
	switch out {
	case "C":
		if fahr != 0 {
			celsius = (fahr - 32) * 5 / 9
			fmt.Printf("%.2f %s\n", celsius, "°C")
		} else if kelvin != 0 {
			celsius = kelvin - 273.15
			fmt.Printf("%.2f %s\n", celsius, "°C")
		} else {
			fmt.Printf("%.2f %s\n", celsius, "°C")
		}
	case "F":
		if celsius != 0 {
			fahr = (celsius * 9 / 5) + 32
			fmt.Printf("%.2f %s\n", fahr, "°F")
		} else if kelvin != 0 {
			fahr = (kelvin*9/5 - 459.67)
			fmt.Printf("%.2f %s\n", fahr, "°F")
		} else {
			fmt.Printf("%.2f %s\n", fahr, "°F")
		}
	case "K":
		if fahr != 0 {
			kelvin = (fahr + 459.67) * 5 / 9
			fmt.Printf("%.2f %s\n", kelvin, "K")
		} else if celsius != 0 {
			kelvin = celsius + 273.15
			fmt.Printf("%.2f %s\n", kelvin, "K")
		} else {
			fmt.Printf("%.2f %s\n", kelvin, "K")
		}
	default:
		fmt.Println("Invalid temperature scale: Enter a valid temperature unit (C, F or K)")
	}
	
}
	

/*
func  GetFunFacts(about string) {
	funFacts := fun.GetFunFacts(about)
	for i, fact := range funFacts {
		fmt.Println(i+1, fact)
	}
}
*/
