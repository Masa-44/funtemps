package main

import (
	"flag"
	"fmt"

	"github.com/Masa-44/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var celsius float64
var Kelvin float64

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
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&Kelvin, "K", 0.0, "temperatur i grader kelvin")

}

func main() {
	//conv.CelsiusToFahrenheit()
	//conv.CelsiusToKelvin()
	//conv.CelsiusToFahrenheit()
	//conv.FarhenheitToCelsius()
	//conv.CelsiusToKelvin()
	//conv.FahrenheitToKelvin()

	flag.Parse()

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
		conv.FarhenheitToCelsius(fahr)

		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println(conv.FarhenheitToCelsius(fahr))

	}
	if out == "F" && isFlagPassed("C") {
		conv.CelsiusToFahrenheit(celsius)
		fmt.Println(conv.CelsiusToFahrenheit(celsius))
	}

	if out == "K" && isFlagPassed("C") {
		conv.CelsiusToKelvin(celsius)
		fmt.Println(conv.CelsiusToKelvin(celsius))
	}
	if out == "K" && isFlagPassed("F") {
		conv.FahrenheitToKelvin(fahr)
		fmt.Println(conv.FahrenheitToKelvin(fahr))
	}
	if out == "C" && isFlagPassed("K") {
		conv.KelvinToCelsius(Kelvin)
		fmt.Println(conv.KelvinToCelsius(Kelvin))
	}
	if out == "F" && isFlagPassed("K") {
		conv.KelvinToCelsius(Kelvin)
		fmt.Println(conv.KelvinToFahrenheit(Kelvin))
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
