package main

import (
	"flag"
	"fmt"
)

// Definerer flag-variablene i hoved-"scope"
fahr float64
celsius float64
kelvin float64
out string
funfacts string
tempSkala string


// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&tempSkala, "t", "C", "the temperature scale to use when showing funfacts")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()
//denne koden sjekker om F, C og K blir brukt samtidig, fordi det er ikke lov
	if (fahr != 0 && celsius != 0) || (fahr != 0 && kelvin != 0) || (celsius != 0 && kelvin != 0) {
		fmt.Println("Error: Kun en av: -F, -C, og -K kan brukes samtidig")
		//hvis de blir brukt samtidig slutter den programmet med denne koden
		os.Exit(1)
	}
	//sjekker om F, C eller K blir brukt med funfacts, os.Exit(1) - fordi det er ikke lov
	if fahr != 0 && funfacts != "" {
		fmt.Println("Error: -F kan ikke brukes med -funfacts")
		os.Exit(1)
	}
	if celsius != 0 && funfacts != "" {
		fmt.Println("Error: -C kan ikke brukes med -funfacts")
		os.Exit(1)
	}
	if kelvin != 0 && funfacts != "" {
		fmt.Println("Error: -K kan ikke brukes med -funfacts")
		os.Exit(1)
	}
	//sjekker om -t er brukt med funfacts, noe den skal! Så hvis det ikke skjer så slutter den også programmet
	if funfacts != "" && tempSkala == "" {
		fmt.Println("Error: -funfacts kan kun brukes med -t")
		os.Exit(1)
	}
//til slutt, hvis alt er greit og betingelsene er møtt vil den se hvilken temp som har verdi og gjøre en av fire ting
	if fahr != 0 {
		// konverterer fra farhenheit til ønskelig temp 
	} else if celsius != 0 {
		// konverterer fra celsius til ønskelig temp 
	} else if kelvin != 0 {
		// konverterer fra kelvin til ønskelig temp 
	} else if funfacts != "" {
		// Viser funfacts om den verdien man har valgt 
	}
}



	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
