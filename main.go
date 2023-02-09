package main

import (
	"flag"
	"fmt"
	"github.com/mienord/funtemps/funfacts"
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
			celsius = (fahr - 32) * 5 / 9
			fmt.Println(celsius, "°C")
		} else if kelvin != 0 {
			celsius = kelvin - 273.15
			fmt.Println(celsius, "°C")
		} else {
			fmt.Println(celsius, "°C")
		}
	case "F":
		if celsius != 0 {
			fahr = (celsius * 9 / 5) + 32
			fmt.Println(fahr, "°F")
		} else if kelvin != 0 {
			fahr = (kelvin*9/5 - 459.67)
			fmt.Println(fahr, "°F")
		} else {
			fmt.Println(fahr, "°F")
		}
	case "K":
		if fahr != 0 {
			kelvin = (fahr + 459.67) * 5 / 9
			fmt.Println(kelvin, "K")
		} else if celsius != 0 {
			kelvin = celsius + 273.15
			fmt.Println(kelvin, "K")
		} else {
			fmt.Println(kelvin, "K")
		}
	default:
		fmt.Println("Invalid temperature scale: Enter a valid temperature unit (C, F or K)")
	}
		
			
				// Tests that makes sure these combinations are fulfilled:
				// -F, -C, -K can not be used similtaniously
				if (fahr != 0) && (celsius != 0 || kelvin != 0) {
					fmt.Println("Error: Cannot use -F flag with -C or -K flags.")
					return
				}
				if (celsius != 0) && (fahr != 0 || kelvin != 0) {
					fmt.Println("Error: Cannot use -C flag with -F or -K flags.")
					return
				}
				if (kelvin != 0) && (fahr != 0 || celsius != 0) {
					fmt.Println("Error: Cannot use -K flag with -F or -C flags.")
					return
				}
			
				// -F, -C, -K can be used with -out but not with -funfacts
				if (fahr != 0 || celsius != 0 || kelvin != 0) && funfacts != "sun" {
					fmt.Println("Error: Cannot use -F, -C, or -K flags with -funfacts flag.")
					return
				}
			
				// funfacts can only be used with -t
				if (funfacts != "sun") && tempSkala == "C" {
					fmt.Println("Error: Must use -t flag with -funfacts flag.")
					return
				}
			
				fmt.Println("len(flag.Args())", len(flag.Args()))
				fmt.Println("flag.NFlag()", flag.NFlag())
			
				// Check the out flag and convert temperature
				switch out {
				case "C":
					if fahr != 0 {
						celsius = (fahr - 32) * 5 / 9
						fmt.Println(celsius, "°C")
					} else if kelvin != 0 {
						celsius = kelvin - 273.15
						fmt.Println(celsius, "°C")
					} else {
						fmt.Println(celsius, "°C")
					}
				case "F":
					if celsius != 0 {
						fahr = (celsius * 9 / 5) + 32
						fmt.Println(fahr, "°F")
					} else if kelvin != 0 {
						fahr = (kelvin*9/5 - 459.67)
						fmt.Println(fahr, "°F")
					} else {
						fmt.Println(fahr, "°F")
					}
				case "K":
					if fahr != 0 {
						kelvin = (fahr + 459.67) * 5 / 9
						fmt.Println(kelvin, "K")
					} else if celsius != 0 {
						kelvin = celsius + 273.15
						fmt.Println(kelvin, "K")
					} else {
						fmt.Println(kelvin, "K")
					}
				default:
					fmt.Println("Invalid temperature scale")
				}
			}
			
